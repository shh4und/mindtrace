package servicos

import (
	"errors"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/mappers"
	"mindtrace/backend/interno/persistencia/repositorios"

	"gorm.io/gorm"
)

// ResumoServico define os metodos para gerar resumos
type ResumoServico interface {
	GerarResumoPaciente(userID uint) (*dtos.ResumoPacienteDTOOut, error)
	// GerarResumoPacienteDoProfissional(pacienteID uint, filtroPeriodo int64) (*dtos.ResumoPacienteDTOOut, error)
}

// resumoServico implementa a interface ResumoServico
type resumoServico struct {
	db                       *gorm.DB
	registroHumorRepositorio repositorios.RegistroHumorRepositorio
	usuarioRepositorio       repositorios.UsuarioRepositorio
}

// NovoResumoServico cria uma nova instancia de ResumoServico
func NovoResumoServico(db *gorm.DB, registroHumorRepo repositorios.RegistroHumorRepositorio, usuarioRepo repositorios.UsuarioRepositorio) ResumoServico {
	return &resumoServico{
		db:                       db,
		registroHumorRepositorio: registroHumorRepo,
		usuarioRepositorio:       usuarioRepo,
	}
}

// GerarResumoPaciente gera um resumo para o paciente autenticado
func (rs *resumoServico) GerarResumoPaciente(userID uint) (*dtos.ResumoPacienteDTOOut, error) {
	// Validar userID
	if userID == 0 {
		return nil, errors.New("id do usuario invalido")
	}

	resumoPacienteFeito := &dtos.ResumoPacienteDTOOut{}

	paciente, err := rs.usuarioRepositorio.BuscarPacientePorUsuarioID(rs.db, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("paciente nao encontrado")
		}
		return nil, err
	}

	registroHumor, err := rs.registroHumorRepositorio.BuscarUltimoRegistroDePaciente(paciente.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Sem registros e ok, retorna resumo vazio
			return resumoPacienteFeito, nil
		}
		return nil, err
	}

	// Validar o registro antes de mapear
	if err := registroHumor.Validar(); err != nil {
		return nil, errors.New("dados do registro invalidos: " + err.Error())
	}

	return mappers.ResumoPacienteParaDTOOut(registroHumor), nil
}
