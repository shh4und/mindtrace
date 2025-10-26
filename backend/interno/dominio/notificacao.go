package dominio

import (
	"time"
)

// Notificacao representa uma notificacao enviada a um usuario.
type Notificacao struct {
	ID        uint      `gorm:"primaryKey"`
	UsuarioID uint      `gorm:"not null"`
	Usuario   Usuario   `gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE"`
	AlertaID  *uint     // Ponteiro para permitir valor NULL
	Conteudo  string    `gorm:"type:text;not null"`
	Status    string    `gorm:"type:varchar(50);not null;default:'Não lida'"`
	DataEnvio time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (Notificacao) TableName() string {
	return "notificacoes"
}
