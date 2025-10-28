package dominio

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Erros de validacao - Convite
var (
	ErrTokenConviteVazio      = errors.New("token do convite nao pode estar vazio")
	ErrTokenConviteInvalido   = errors.New("token deve ter no minimo 10 caracteres")
	ErrDataExpiracaoVazia     = errors.New("data de expiracao e obrigatoria")
	ErrDataExpiracaoNoPassado = errors.New("data de expiracao nao pode ser no passado")
	ErrConviteExpirado        = errors.New("convite expirado")
	ErrConviteJaUtilizado     = errors.New("convite ja foi utilizado")
)

type Convite struct {
	ID             uint         `gorm:"primarykey"`
	ProfissionalID uint         `gorm:"not null"`
	Profissional   Profissional `gorm:"foreignKey:ProfissionalID;constraint:OnDelete:CASCADE"`
	Token          string       `gorm:"unique;not null"`
	DataExpiracao  time.Time    `gorm:"not null"`
	Usado          bool         `gorm:"default:false"`
	PacienteID     *uint        // Ponteiro para permitir nulo, indica qual paciente usou o convite
	Paciente       Paciente     `gorm:"foreignKey:PacienteID;constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// Metodos de validacao - LOGICA DE NEGOCIO (Convite)
func (c *Convite) ValidarToken() error {
	if c.Token == "" {
		return ErrTokenConviteVazio
	}
	if len(c.Token) < 10 {
		return ErrTokenConviteInvalido
	}
	return nil
}

func (c *Convite) ValidarDataExpiracao() error {
	if c.DataExpiracao.IsZero() {
		return ErrDataExpiracaoVazia
	}
	if c.DataExpiracao.Before(time.Now()) {
		return ErrDataExpiracaoNoPassado
	}
	return nil
}

// Validacao completa do Convite
func (c *Convite) Validar() error {
	if err := c.ValidarToken(); err != nil {
		return err
	}
	if err := c.ValidarDataExpiracao(); err != nil {
		return err
	}
	return nil
}

// EstaValido verifica se o convite ainda e valido
func (c *Convite) EstaValido() bool {
	return !c.Usado && c.DataExpiracao.After(time.Now())
}

// EstaExpirado verifica se o convite expirou
func (c *Convite) EstaExpirado() bool {
	return c.DataExpiracao.Before(time.Now())
}

// JaFoiUtilizado verifica se o convite ja foi utilizado
func (c *Convite) JaFoiUtilizado() bool {
	return c.Usado
}

// UtilizarConvite marca o convite como utilizado por um paciente
func (c *Convite) UtilizarConvite(pacienteID uint) {
	c.Usado = true
	c.PacienteID = &pacienteID
}
