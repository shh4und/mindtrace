package dominio

import (
	"time"
)

// RegistroHumor armazena as entradas de humor do paciente.
type RegistroHumor struct {
	ID               uint      `gorm:"primaryKey"`
	PacienteID       uint      `json:"-" gorm:"not null"`
	Paciente         Paciente  `json:"paciente" gorm:"foreignKey:PacienteID"`
	Nivel            int16     `json:"nivel" gorm:"not null;check:nivel >= 1 AND nivel <= 5"`
	Observacoes      string    `json:"observacoes" gorm:"type:text"`
	DataHoraRegistro time.Time `json:"data_hora_registro" gorm:"not null;default:now()"`
	CreatedAt        time.Time `json:"created_at"`
}

func (RegistroHumor) TableName() string {
	return "registros_humors"
}
