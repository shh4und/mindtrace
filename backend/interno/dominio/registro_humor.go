package dominio

import (
	"time"
)

type AutoCuidadoUnit struct {
	autocuidado string
}

// RegistroHumor armazena as entradas de humor do paciente.
type RegistroHumor struct {
	ID               uint             `gorm:"primaryKey"`
	PacienteID       uint             `json:"-" gorm:"not null"`
	Paciente         Paciente         `json:"paciente" gorm:"foreignKey:PacienteID"`
	NivelHumor       int16            `json:"nivel_humor" gorm:"not null;check:nivel_humor >= 1 AND nivel_humor <= 5"`
	HorasSono        int16            `json:"horas_sono" gorm:"not null;check:horas_sono >= 0 AND horas_sono <= 12"`
	NivelEnergia     int16            `json:"nivel_energia" gorm:"not null;check:nivel_energia >= 1 and nivel_energia <= 10"`
	NivelStress      int16            `json:"nivel_stress" gorm:"not null;check:nivel_stress >= 1 and nivel_stress <= 10"`
	AutoCuidado      string           `json:"auto_cuidado" gorm:"type:jsonb;not null"`
	Observacoes      string           `json:"observacoes" gorm:"type:text"`
	DataHoraRegistro time.Time        `json:"data_hora_registro" gorm:"not null;default:now()"`
	CreatedAt        time.Time        `json:"created_at"`
}

func (RegistroHumor) TableName() string {
	return "registros_humor"
}
