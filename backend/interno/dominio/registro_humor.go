package dominio

import (
	"errors"
	"time"
)

// RegistroHumor armazena as entradas de humor do paciente.
type RegistroHumor struct {
	ID               uint      `gorm:"primaryKey"`
	PacienteID       uint      `gorm:"not null"`
	Paciente         Paciente  `gorm:"foreignKey:PacienteID;constraint:OnDelete:CASCADE"`
	NivelHumor       int16     `gorm:"not null;check:nivel_humor >= 1 AND nivel_humor <= 5"`
	HorasSono        int16     `gorm:"not null;check:horas_sono >= 0 AND horas_sono <= 12"`
	NivelEnergia     int16     `gorm:"not null;check:nivel_energia >= 1 and nivel_energia <= 10"`
	NivelStress      int16     `gorm:"not null;check:nivel_stress >= 1 and nivel_stress <= 10"`
	AutoCuidado      string    `gorm:"type:text;not null"`
	Observacoes      string    `gorm:"type:text"`
	DataHoraRegistro time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedAt        time.Time
}

func (RegistroHumor) TableName() string {
	return "registros_humor"
}

// Metodos de validacao - LOGICA DE NEGOCIO (RegistroHumor)
func (rh *RegistroHumor) ValidarNivelHumor() error {
	if rh.NivelHumor < 1 || rh.NivelHumor > 5 {
		return errors.New("nivel de humor deve estar entre 1 e 5")
	}
	return nil
}

func (rh *RegistroHumor) ValidarHorasSono() error {
	if rh.HorasSono < 0 || rh.HorasSono > 12 {
		return errors.New("horas de sono deve estar entre 0 e 12")
	}
	return nil
}

func (rh *RegistroHumor) ValidarNivelEnergia() error {
	if rh.NivelEnergia < 1 || rh.NivelEnergia > 10 {
		return errors.New("nivel de energia deve estar entre 1 e 10")
	}
	return nil
}

func (rh *RegistroHumor) ValidarNivelStress() error {
	if rh.NivelStress < 1 || rh.NivelStress > 10 {
		return errors.New("nivel de stress deve estar entre 1 e 10")
	}
	return nil
}

func (rh *RegistroHumor) ValidarAutoCuidado() error {
	if rh.AutoCuidado == "" {
		return errors.New("auto cuidado nao pode estar vazio")
	}
	if len(rh.AutoCuidado) < 3 {
		return errors.New("auto cuidado deve ter no minimo 3 caracteres")
	}
	return nil
}

func (rh *RegistroHumor) ValidarDataHoraRegistro() error {
	if rh.DataHoraRegistro.IsZero() {
		return errors.New("data e hora do registro e obrigatoria")
	}
	if rh.DataHoraRegistro.After(time.Now()) {
		return errors.New("data e hora do registro nao pode ser no futuro")
	}
	return nil
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
