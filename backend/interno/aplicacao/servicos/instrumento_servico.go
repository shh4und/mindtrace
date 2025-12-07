package servicos

import (
	"errors"
	"fmt"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/mappers"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"

	"gorm.io/gorm"
)

type InstrumentoServico interface {
	ListarInstrumentos(userID uint) ([]*dtos.InstrumentoDTOOut, error)
	CriarAtribuicao(userID, pacienteID, instrumentoID uint, instrumentoCodigo string) error
	ListarAtribuicoesProfissional(profId uint) ([]*dtos.AtribuicaoDTOOut, error)
	ListarAtribuicoesPaciente(pacId uint) ([]*dtos.AtribuicaoDTOOut, error)
	ListarPerguntasAtribuicao(usuarioId, atribuicaoId uint) (*dtos.AtribuicaoDTOOut, error)
}
type instrumentoServico struct {
	db              *gorm.DB
	instrumentoRepo repositorios.InstrumentoRepositorio
	usuarioRepo     repositorios.UsuarioRepositorio
}

func NovoInstrumentoServico(db *gorm.DB, instrumentoRepo repositorios.InstrumentoRepositorio, usuarioRepo repositorios.UsuarioRepositorio) InstrumentoServico {
	return &instrumentoServico{
		db:              db,
		instrumentoRepo: instrumentoRepo,
		usuarioRepo:     usuarioRepo,
	}
}

func (is *instrumentoServico) ListarInstrumentos(userID uint) ([]*dtos.InstrumentoDTOOut, error) {
	var instrumentos []*dominio.Instrumento
	// Checar a existencia do profissional
	_, err := is.usuarioRepo.BuscarProfissionalPorUsuarioID(is.db, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dominio.ErrUsuarioNaoEncontrado
		}
		return nil, err
	}

	// Buscar instrumentos
	instrumentos, err = is.instrumentoRepo.BuscarTodosAtivos(is.db)
	if err != nil {
		return nil, err
	}

	if instrumentos == nil {
		return nil, fmt.Errorf("slice instrumentos invalido: slice nulo")
	}
	for _, inst := range instrumentos {
		if err := inst.ValidarCodigo(); err != nil {

			return nil, fmt.Errorf("instrumento %s invalido: %w", inst.Codigo, err)
		}
		if err := inst.ValidarNome(); err != nil {

			return nil, fmt.Errorf("instrumento %s invalido: %w", inst.Nome, err)
		}
	}

	return mappers.InstrumentosParaDTOOut(instrumentos), nil
}

func (is *instrumentoServico) CriarAtribuicao(userID, pacienteID, instrumentoID uint, instrumentoCodigo string) error {
	err := is.db.Transaction(func(tx *gorm.DB) error {

		profissional, err := is.usuarioRepo.BuscarProfissionalPorUsuarioID(tx, userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return dominio.ErrUsuarioNaoEncontrado
			}
			return err
		}

		paciente, err := is.usuarioRepo.BuscarPacientePorID(tx, pacienteID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return dominio.ErrUsuarioNaoEncontrado
			}
			return err
		}

		instrumento, err := is.instrumentoRepo.BuscarInstrumentoPorID(tx, instrumentoID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("instrumento nao encontrado: %v", err)
			}
			return err
		}

		atribuicao := &dominio.Atribuicao{
			ProfissionalID: profissional.ID,
			Profissional:   *profissional,
			PacienteID:     paciente.ID,
			Paciente:       *paciente,
			InstrumentoID:  instrumento.ID,
			Instrumento:    *instrumento,
		}

		if err = is.instrumentoRepo.CriarAtribuicao(tx, atribuicao); err != nil {
			return err
		}

		if err = atribuicao.Validar(); err != nil {
			return err
		}
		return nil

	})

	return err
}

func (is *instrumentoServico) ListarAtribuicoesPaciente(usuarioId uint) ([]*dtos.AtribuicaoDTOOut, error) {
	var atribuicoes []*dominio.Atribuicao

	paciente, err := is.usuarioRepo.BuscarPacientePorUsuarioID(is.db, usuarioId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dominio.ErrUsuarioNaoEncontrado
		}
		return nil, err
	}

	atribuicoes, err = is.instrumentoRepo.BuscarAtribuicoesPaciente(is.db, paciente.ID)
	if err != nil {
		return nil, err
	}

	for _, atribc := range atribuicoes {
		if err = atribc.Validar(); err != nil {
			return nil, err
		}
	}

	return mappers.AtribuicoesParaDTOOutPaciente(atribuicoes), nil
}

func (is *instrumentoServico) ListarAtribuicoesProfissional(usuarioId uint) ([]*dtos.AtribuicaoDTOOut, error) {
	var atribuicoes []*dominio.Atribuicao
	profissional, err := is.usuarioRepo.BuscarProfissionalPorUsuarioID(is.db, usuarioId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dominio.ErrUsuarioNaoEncontrado
		}
		return nil, err
	}
	atribuicoes, err = is.instrumentoRepo.BuscarAtribuicoesProfissional(is.db, profissional.ID)
	if err != nil {
		return nil, err
	}

	for _, atribc := range atribuicoes {
		if err = atribc.Validar(); err != nil {
			return nil, err
		}
	}

	return mappers.AtribuicoesParaDTOOutProfissional(atribuicoes), nil
}

func (is *instrumentoServico) ListarPerguntasAtribuicao(usuarioId, atribuicaoId uint) (*dtos.AtribuicaoDTOOut, error) {
	var atribuicao *dominio.Atribuicao
	_, err := is.usuarioRepo.BuscarPacientePorUsuarioID(is.db, usuarioId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dominio.ErrUsuarioNaoEncontrado
		}
		return nil, err
	}
	atribuicao, err = is.instrumentoRepo.BuscarAtribuicaoPorID(is.db, atribuicaoId)
	if err != nil {
		return nil, err
	}

	for _, pergunta := range atribuicao.Instrumento.Perguntas {
		if err = pergunta.Validar(atribuicao.Instrumento.Codigo); err != nil {
			return nil, err
		}
	}

	return mappers.AtribuicaoComPerguntasDTOOut(atribuicao), nil
}
