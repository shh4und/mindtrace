package dominio

import (
	"time"
)

// RegistroHumor armazena as entradas de humor do paciente.
type RegistroHumor struct {
	ID               uint      `gorm:"primaryKey"`
	PacienteID       uint      `gorm:"not null"`
	Paciente         Paciente  `gorm:"foreignKey:PacienteID;constraint:OnDelete:CASCADE"`
	NivelHumor       int16     `gorm:"not null;check:nivel_humor >= 1 AND nivel_humor <= 5"`
	HorasSono        int16     `gorm:"not null;check:horas_sono >= 0 AND horas_sono <= 12"`
	NivelEnergia     int16     `gorm:"not null;check:nivel_energia >= 1 and nivel_energia <= 10"`
	NivelStress      int16     `gorm:"not null;check:nivel_stress >= 1 and nivel_stress <= 10"`
	AutoCuidado      string    `gorm:"type:text;not null"`
	Observacoes      string    `gorm:"type:text"`
	DataHoraRegistro time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedAt        time.Time
}

func (RegistroHumor) TableName() string {
	return "registros_humor"
}
