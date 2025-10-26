package dominio

import (
	"errors"
	"time"
)

// Constantes para status de notificacao
const (
	NotificacaoNaoLida   = "Nao lida"
	NotificacaoLida      = "Lida"
	NotificacaoArquivada = "Arquivada"
)

// Erros de validacao - Notificacao
var (
	ErrConteudoNotificacaoVazio      = errors.New("conteudo da notificacao nao pode estar vazio")
	ErrConteudoNotificacaoMuitoCurto = errors.New("conteudo deve ter no minimo 3 caracteres")
	ErrConteudoNotificacaoMuitoLongo = errors.New("conteudo nao pode exceder 5000 caracteres")
	ErrStatusNotificacaoInvalido     = errors.New("status de notificacao invalido")
	ErrDataEnvioVazia                = errors.New("data de envio e obrigatoria")
	ErrDataEnvioNoFuturo             = errors.New("data de envio nao pode ser no futuro")
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

// Metodos de validacao - LOGICA DE NEGOCIO (Notificacao)
func (n *Notificacao) ValidarConteudo() error {
	if n.Conteudo == "" {
		return ErrConteudoNotificacaoVazio
	}
	if len(n.Conteudo) < 3 {
		return ErrConteudoNotificacaoMuitoCurto
	}
	if len(n.Conteudo) > 5000 {
		return ErrConteudoNotificacaoMuitoLongo
	}
	return nil
}

func (n *Notificacao) ValidarStatus() error {
	statusValidos := map[string]bool{
		NotificacaoNaoLida:   true,
		NotificacaoLida:      true,
		NotificacaoArquivada: true,
	}
	if !statusValidos[n.Status] {
		return ErrStatusNotificacaoInvalido
	}
	return nil
}

func (n *Notificacao) ValidarDataEnvio() error {
	if n.DataEnvio.IsZero() {
		return ErrDataEnvioVazia
	}
	if n.DataEnvio.After(time.Now()) {
		return ErrDataEnvioNoFuturo
	}
	return nil
}

// Validacao completa da Notificacao
func (n *Notificacao) Validar() error {
	if err := n.ValidarConteudo(); err != nil {
		return err
	}
	if err := n.ValidarStatus(); err != nil {
		return err
	}
	if err := n.ValidarDataEnvio(); err != nil {
		return err
	}
	return nil
}

// MarcarComoLida marca a notificacao como lida
func (n *Notificacao) MarcarComoLida() {
	n.Status = NotificacaoLida
}

// MarcarComoArquivada marca a notificacao como arquivada
func (n *Notificacao) MarcarComoArquivada() {
	n.Status = NotificacaoArquivada
}

// EstaLida verifica se a notificacao foi lida
func (n *Notificacao) EstaLida() bool {
	return n.Status == NotificacaoLida
}

// EstaArquivada verifica se a notificacao esta arquivada
func (n *Notificacao) EstaArquivada() bool {
	return n.Status == NotificacaoArquivada
}
