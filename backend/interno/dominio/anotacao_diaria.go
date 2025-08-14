package dominio

import (
	"time"
)

// AnotacaoDiaria armazena as entradas do diário do paciente.
type AnotacaoDiaria struct {
	ID               uint      `gorm:"primaryKey"`
	PacienteID       uint      `json:"-" gorm:"not null"`
	Paciente         Paciente  `json:"paciente" gorm:"foreignKey:PacienteID"`
	Conteudo         string    `json:"conteudo" gorm:"type:text;not null"`
	DataHoraAnotacao time.Time `json:"data_hora_anotacao" gorm:"not null;default:now()"`
	CreatedAt        time.Time `json:"created_at"`
}


func (AnotacaoDiaria) TableName() string {
  return "anotacoes_diarias"
}