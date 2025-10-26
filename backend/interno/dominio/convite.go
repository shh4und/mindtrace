package dominio

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Convite struct {
	gorm.Model
	ProfissionalID uint         `gorm:"not null"`
	Profissional   Profissional `gorm:"foreignKey:ProfissionalID;constraint:OnDelete:CASCADE"`
	Token          string       `gorm:"unique;not null"`
	DataExpiracao  time.Time    `gorm:"not null"`
	Usado          bool         `gorm:"default:false"`
	PacienteID     *uint        // Ponteiro para permitir nulo, indica qual paciente usou o convite
	Paciente       Paciente     `gorm:"foreignKey:PacienteID;constraint:OnDelete:CASCADE"`
}

// Metodos de validacao - LOGICA DE NEGOCIO (Convite)
func (c *Convite) ValidarToken() error {
	if c.Token == "" {
		return errors.New("token do convite nao pode estar vazio")
	}
	if len(c.Token) < 10 {
		return errors.New("token deve ter no minimo 10 caracteres")
	}
	return nil
}

func (c *Convite) ValidarDataExpiracao() error {
	if c.DataExpiracao.IsZero() {
		return errors.New("data de expiracao e obrigatoria")
	}
	if c.DataExpiracao.Before(time.Now()) {
		return errors.New("data de expiracao nao pode ser no passado")
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
func (c *Convite) UtilizarConvite(pacienteID uint) error {
	if !c.EstaValido() {
		if c.EstaExpirado() {
			return errors.New("convite expirado")
		}
		if c.JaFoiUtilizado() {
			return errors.New("convite ja foi utilizado")
		}
	}
	c.Usado = true
	c.PacienteID = &pacienteID
	return nil
}
