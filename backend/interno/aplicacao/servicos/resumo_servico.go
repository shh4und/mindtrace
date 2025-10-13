package servicos

import (
	"errors"
	"mindtrace/backend/interno/persistencia/repositorios"
	"time"

	"gorm.io/gorm"
)

// PontoDeDadosDTO representa um ponto de dados para graficos

// ResumoPacienteDTO representa o resumo de um paciente <=> ultimo registro
type ResumoPacienteDTO struct {
	Data     time.Time `json:"data"`
	Humor    int16     `json:"humor"`
	Anotacao string    `json:"anotacao,omitempty"`
}

// ResumoServico define os metodos para gerar resumos
type ResumoServico interface {
	GerarResumoPaciente(userID uint) (*ResumoPacienteDTO, error)
	// GerarResumoPacienteDoProfissional(pacienteID uint, filtroPeriodo int64) (*ResumoPacienteDTO, error)
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
func (rs *resumoServico) GerarResumoPaciente(userID uint) (*ResumoPacienteDTO, error) {
	resumoPacienteFeito := &ResumoPacienteDTO{}

	paciente, err := rs.usuarioRepositorio.BuscarPacientePorUsuarioID(rs.db, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resumoPacienteFeito, nil
		}
		return nil, err
	}

	registroHumor, err := rs.registroHumorRepositorio.BuscarUltimoRegistroDePaciente(paciente.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resumoPacienteFeito, nil
		}
		return nil, err
	}

	resumoPacienteFeito.Data = registroHumor.DataHoraRegistro
	resumoPacienteFeito.Humor = registroHumor.NivelHumor
	resumoPacienteFeito.Anotacao = registroHumor.Observacoes

	return resumoPacienteFeito, nil
}
