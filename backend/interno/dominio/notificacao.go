package dominio

import (
	"time"
)

// Notificacao representa uma notificação enviada a um usuário.
type Notificacao struct {
	ID        uint      `gorm:"primaryKey"`
	UsuarioID uint      `json:"-" gorm:"not null"`
	Usuario   Usuario   `json:"usuario" gorm:"foreignKey:UsuarioID"`
	AlertaID  *uint     `json:"-"` // Ponteiro para permitir valor NULL
	Alerta    *Alerta   `json:"alerta,omitempty" gorm:"foreignKey:AlertaID"`
	Conteudo  string    `json:"conteudo" gorm:"type:text;not null"`
	Status    string    `json:"status" gorm:"type:varchar(50);not null;default:'Não lida'"`
	DataEnvio time.Time `json:"data_envio" gorm:"not null;default:now()"`
}

func (Notificacao) TableName() string {
  return "notificacoes"
}
