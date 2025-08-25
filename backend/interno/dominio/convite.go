package dominio

import (
	"time"

	"gorm.io/gorm"
)

type Convite struct {
	gorm.Model
	ProfissionalID uint         `gorm:"not null"`
	Profissional   Profissional `gorm:"foreignKey:ProfissionalID"`
	Token          string       `gorm:"unique;not null"`
	DataExpiracao  time.Time    `gorm:"not null"`
	Usado          bool         `gorm:"default:false"`
	PacienteID     *uint        // Ponteiro para permitir nulo, indica qual paciente usou o convite
	Paciente       Paciente     `gorm:"foreignKey:PacienteID"`
}
