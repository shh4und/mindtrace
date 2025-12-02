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
