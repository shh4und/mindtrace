package dominio

import (
	"time"
)

// Notificacao representa uma notificacao enviada a um usuario.
type Notificacao struct {
	ID        uint      `gorm:"primaryKey"`
	UsuarioID uint      `json:"-" gorm:"not null"`
	Usuario   Usuario   `json:"usuario" gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE"`
	AlertaID  *uint     `json:"-"` // Ponteiro para permitir valor NULL
	Conteudo  string    `json:"conteudo" gorm:"type:text;not null"`
	Status    string    `json:"status" gorm:"type:varchar(50);not null;default:'Não lida'"`
	DataEnvio time.Time `json:"data_envio" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (Notificacao) TableName() string {
	return "notificacoes"
}
