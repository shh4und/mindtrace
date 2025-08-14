package dominio

import (
	"time"
)

// Alerta representa um alerta gerado pelo sistema.
type Alerta struct {
	ID                 uint         `gorm:"primaryKey"`
	PacienteID         uint         `json:"-" gorm:"not null"`
	Paciente           Paciente     `json:"paciente" gorm:"foreignKey:PacienteID"`
	ProfissionalID     uint         `json:"-" gorm:"not null"`
	Profissional       Profissional `json:"profissional" gorm:"foreignKey:ProfissionalID"`
	Titulo             string       `json:"titulo" gorm:"type:varchar(255);not null"`
	DescricaoDetalhada string       `json:"descricao_detalhada" gorm:"type:text"`
	NivelUrgencia      string       `json:"nivel_urgencia" gorm:"type:varchar(50);not null"`
	Status             string       `json:"status" gorm:"type:varchar(50);not null;default:'Ativo'"`
	DataGeracao        time.Time    `json:"data_geracao" gorm:"not null;default:now()"`
}
