package dominio

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Status da Atribuição
const (
	StatusPendente   = "PENDENTE"
	StatusRespondido = "RESPONDIDO"
	StatusExpirado   = "EXPIRADO"
)

var (
	ErrAtribuicaoSemPaciente    = errors.New("atribuicao deve ter um paciente")
	ErrAtribuicaoSemInstrumento = errors.New("atribuicao deve ter um instrumento")
)

// Atribuicao representa o envio de um questionário para um paciente
type Atribuicao struct {
	ID            uint        `gorm:"primaryKey"`
	PacienteID    uint        `gorm:"not null;index;column:paciente_id"`
	Paciente      Paciente    `gorm:"foreignKey:PacienteID"`
	InstrumentoID uint        `gorm:"not null;index;column:instrumento_id"`
	Instrumento   Instrumento `gorm:"foreignKey:InstrumentoID"`

	Status         string     `gorm:"default:'PENDENTE';index;column:status"`
	DataAtribuicao time.Time  `gorm:"autoCreateTime;column:data_atribuicao"`
	DataResposta   *time.Time `gorm:"column:data_resposta"`

	// Relacionamento inverso: Uma atribuição pode ter uma resposta
	Resposta *Resposta `gorm:"foreignKey:AtribuicaoID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Atribuicao) TableName() string {
	return "atribuicoes"
}

func (a *Atribuicao) Validar() error {
	if a.PacienteID == 0 {
		return ErrAtribuicaoSemPaciente
	}
	if a.InstrumentoID == 0 {
		return ErrAtribuicaoSemInstrumento
	}
	return nil
}
