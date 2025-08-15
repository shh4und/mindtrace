package dominio

import (
	"time"
)

// RegistroHumor armazena as entradas de humor do paciente.
type RegistroHumor struct {
	ID               uint      `gorm:"primaryKey"`
	PacienteID       uint      `json:"-" gorm:"not null"`
	Paciente         Paciente  `json:"paciente" gorm:"foreignKey:PacienteID"`
	NivelHumor       int16     `json:"nivel_humor" gorm:"not null;check:nivel >= 1 AND nivel <= 7"`
	HorasSono        int16     `json:"horas_sono" gorm:"not null;check:horas_sono >= 0 and horas_sono <= 24"`
	NivelStress      int16     `json:"nivel_stress" gorm:"not null;check:horas_sono >= 0 and horas_sono <= 10"`
	NivelEnergia     int16     `json:"nivel_energia" gorm:"not null;check:horas_sono >= 0 and horas_sono <= 10"`
	AtivdadeFisica   bool      `json:"atividade_fisica" gorm:"not null"`
	AutoCuidado      bool      `json:"auto_cuidado" gorm:"not null"`
	Observacoes      string    `json:"observacoes" gorm:"type:text"`
	DataHoraRegistro time.Time `json:"data_hora_registro" gorm:"not null;default:now()"`
	CreatedAt        time.Time `json:"created_at"`
}

func (RegistroHumor) TableName() string {
	return "registros_humor"
}
