package servicos

import (
	"encoding/json"
	"errors"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"
	"time"

	"gorm.io/gorm"
)

type CriarRegistroHumorDTO struct {
	UsuarioID        uint
	NivelHumor       int16
	HorasSono        int16
	NivelStress      int16
	NivelEnergia     int16
	AutoCuidado      []string
	Observacoes      string
	DataHoraRegistro time.Time
}

var ErrPacienteNaoEncontrado = errors.New("usuario nao encontrado")

type RegistroHumorServico interface {
	CriarRegistroHumor(dto CriarRegistroHumorDTO) (*dominio.RegistroHumor, error)
}

type registroHumorServico struct {
	db                 *gorm.DB
	repositorio        repositorios.RegistroHumorRepositorio
	usuarioRepositorio repositorios.UsuarioRepositorio
}

func NovoRegistroHumorServico(db *gorm.DB, repo repositorios.RegistroHumorRepositorio, userRepo repositorios.UsuarioRepositorio) *registroHumorServico {
	return &registroHumorServico{db: db, repositorio: repo, usuarioRepositorio: userRepo}
}

func (rhs *registroHumorServico) CriarRegistroHumor(dto CriarRegistroHumorDTO) (*dominio.RegistroHumor, error) {
	var registroHumoRealizado *dominio.RegistroHumor

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

		registroHumoRealizado = novoRegistroHumor
		return nil
	})

	return registroHumoRealizado, err
}
