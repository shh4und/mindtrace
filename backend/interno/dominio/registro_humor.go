package dominio

import (
	"errors"
	"math"
	"time"
)

// Erros de validacao - RegistroHumor
var (
	ErrNivelHumorInvalido        = errors.New("nivel de humor deve estar entre 1 e 5")
	ErrHorasSonoInvalido         = errors.New("horas de sono deve estar entre 0 e 12")
	ErrNivelEnergiaInvalido      = errors.New("nivel de energia deve estar entre 1 e 10")
	ErrNivelStressInvalido       = errors.New("nivel de stress deve estar entre 1 e 10")
	ErrAutoCuidadoVazio          = errors.New("auto cuidado nao pode estar vazio")
	ErrAutoCuidadoInvalido       = errors.New("auto cuidado deve ter no minimo 3 caracteres")
	ErrDataHoraRegistroVazia     = errors.New("data e hora do registro e obrigatoria")
	ErrDataHoraRegistroNoFuturo  = errors.New("data e hora do registro nao pode ser no futuro")
	ErrRegistroHumorMuitoRecente = errors.New("seu registro diário já foi realizado")
)

// RegistroHumor armazena as entradas de humor do paciente.
type RegistroHumor struct {
	ID                  uint      `gorm:"primaryKey"`
	PacienteID          uint      `gorm:"not null;index:idx_paciente_data"`
	Paciente            Paciente  `gorm:"foreignKey:PacienteID;constraint:OnDelete:CASCADE"`
	NivelHumor          int16     `gorm:"not null;check:nivel_humor >= 1 AND nivel_humor <= 5;"`
	HorasSono           int16     `gorm:"not null;check:horas_sono >= 0 AND horas_sono <= 12;"`
	NivelEnergia        int16     `gorm:"not null;check:nivel_energia >= 1 and nivel_energia <= 10;"`
	NivelStress         int16     `gorm:"not null;check:nivel_stress >= 1 and nivel_stress <= 10;"`
	IndiceBemEstarGeral float64   `gorm:"not null;default:0;check:indice_bem_estar_geral >= 0 and indice_bem_estar_geral <= 1"`
	AutoCuidado         string    `gorm:"type:jsonb;default:'[]';not null;"`
	Observacoes         string    `gorm:"type:text;"`
	DataHoraRegistro    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;index:idx_paciente_data"`
	CreatedAt           time.Time
}

func (RegistroHumor) TableName() string {
	return "registros_humor"
}

// Metodos de validacao - LOGICA DE NEGOCIO (RegistroHumor)
func (rh *RegistroHumor) ValidarNivelHumor() error {
	if rh.NivelHumor < 1 || rh.NivelHumor > 5 {
		return ErrNivelHumorInvalido
	}
	return nil
}

func (rh *RegistroHumor) ValidarHorasSono() error {
	if rh.HorasSono < 0 || rh.HorasSono > 12 {
		return ErrHorasSonoInvalido
	}
	return nil
}

func (rh *RegistroHumor) ValidarNivelEnergia() error {
	if rh.NivelEnergia < 1 || rh.NivelEnergia > 10 {
		return ErrNivelEnergiaInvalido
	}
	return nil
}

func (rh *RegistroHumor) ValidarNivelStress() error {
	if rh.NivelStress < 1 || rh.NivelStress > 10 {
		return ErrNivelStressInvalido
	}
	return nil
}

func (rh *RegistroHumor) ValidarAutoCuidado() error {
	if rh.AutoCuidado == "" || rh.AutoCuidado == "[]" || rh.AutoCuidado == "null" {
		return ErrAutoCuidadoVazio
	}
	return nil
}

func (rh *RegistroHumor) ValidarDataHoraRegistro() error {
	if rh.DataHoraRegistro.IsZero() {
		return ErrDataHoraRegistroVazia
	}
	if rh.DataHoraRegistro.After(time.Now()) {
		return ErrDataHoraRegistroNoFuturo
	}
	return nil
}

// Pesos do IBG
const (
	PesoSono    = 0.4
	PesoHumor   = 0.2
	PesoStress  = 0.2
	PesoEnergia = 0.2
)

// CalcularIBG calcula o Indice de Bem-Estar Geral a partir dos campos do registro
func (rh *RegistroHumor) CalcularIBG() float64 {
	normHumor := (float64(rh.NivelHumor) - 1.0) / 4.0
	if normHumor < 0 {
		normHumor = 0
	}
	normEnergia := (float64(rh.NivelEnergia) - 1.0) / 9.0
	normStress := 1.0 - ((float64(rh.NivelStress) - 1.0) / 9.0)
	distancia := math.Abs(float64(rh.HorasSono) - 8.0)
	normSono := 1.0 - (distancia / 4.0)
	if normSono < 0 {
		normSono = 0
	}
	return (normHumor * PesoHumor) +
		(normStress * PesoStress) +
		(normSono * PesoSono) +
		(normEnergia * PesoEnergia)
}

// Validacao completa do RegistroHumor
func (rh *RegistroHumor) Validar() error {
	if err := rh.ValidarNivelHumor(); err != nil {
		return err
	}
	if err := rh.ValidarHorasSono(); err != nil {
		return err
	}
	if err := rh.ValidarNivelEnergia(); err != nil {
		return err
	}
	if err := rh.ValidarNivelStress(); err != nil {
		return err
	}
	if err := rh.ValidarAutoCuidado(); err != nil {
		return err
	}
	if err := rh.ValidarDataHoraRegistro(); err != nil {
		return err
	}
	return nil
}
