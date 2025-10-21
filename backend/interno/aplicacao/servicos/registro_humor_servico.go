package servicos

import (
	"encoding/json"
	"errors"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"

	"gorm.io/gorm"
)

var ErrPacienteNaoEncontrado = errors.New("usuario nao encontrado")

// RegistroHumorServico define os metodos para gerenciamento de registros de humor
type RegistroHumorServico interface {
	CriarRegistroHumor(dto dtos.CriarRegistroHumorDTOin) (*dominio.RegistroHumor, error)
}

// registroHumorServico implementa a interface RegistroHumorServico
type registroHumorServico struct {
	db                 *gorm.DB
	repositorio        repositorios.RegistroHumorRepositorio
	usuarioRepositorio repositorios.UsuarioRepositorio
}

// NovoRegistroHumorServico cria uma nova instancia de registroHumorServico
func NovoRegistroHumorServico(db *gorm.DB, repo repositorios.RegistroHumorRepositorio, userRepo repositorios.UsuarioRepositorio) *registroHumorServico {
	return &registroHumorServico{db: db, repositorio: repo, usuarioRepositorio: userRepo}
}

// CriarRegistroHumor cria um novo registro de humor para o paciente
func (rhs *registroHumorServico) CriarRegistroHumor(dto dtos.CriarRegistroHumorDTOin) (*dominio.RegistroHumor, error) {
	var registroHumorRealizado *dominio.RegistroHumor

	err := rhs.db.Transaction(func(tx *gorm.DB) error {
		paciente, err := rhs.usuarioRepositorio.BuscarPacientePorUsuarioID(tx, dto.UsuarioID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrPacienteNaoEncontrado
			}
			return err
		}

		autoCuidadoJSON, err := json.Marshal(dto.AutoCuidado)
		if err != nil {
			return err // Retorna erro de marshaling
		}

		novoRegistroHumor := &dominio.RegistroHumor{
			PacienteID:       paciente.ID,
			Paciente:         *paciente,
			NivelHumor:       dto.NivelHumor,
			HorasSono:        dto.HorasSono,
			NivelEnergia:     dto.NivelEnergia,
			NivelStress:      dto.NivelStress,
			AutoCuidado:      string(autoCuidadoJSON),
			Observacoes:      dto.Observacoes,
			DataHoraRegistro: dto.DataHoraRegistro,
		}

		if err := rhs.repositorio.CriarRegistroHumor(tx, novoRegistroHumor); err != nil {
			return err
		}

		registroHumorRealizado = novoRegistroHumor
		return nil
	})

	return registroHumorRealizado, err
}
