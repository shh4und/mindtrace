package postgres

import (
	"mindtrace/backend/interno/dominio"
	"time"

	"gorm.io/gorm"
)

type RegistroHumorRepositorio interface {
	CriarRegistroHumor(tx *gorm.DB, registro *dominio.RegistroHumor) error
	BuscarPorPacienteEPeriodo(pacienteID uint, inicio, fim time.Time) ([]*dominio.RegistroHumor, error)
}

type gormRegistroHumorRepositorio struct {
	db *gorm.DB
}

func NovoGormRegistroHumorRepositorio(db *gorm.DB) RegistroHumorRepositorio {
	return &gormRegistroHumorRepositorio{db: db}
}

func (r *gormRegistroHumorRepositorio) CriarRegistroHumor(tx *gorm.DB, registro *dominio.RegistroHumor) error {
	return tx.Create(registro).Error
}

func (r *gormRegistroHumorRepositorio) BuscarPorPacienteEPeriodo(pacienteID uint, inicio, fim time.Time) ([]*dominio.RegistroHumor, error) {
	var registros []*dominio.RegistroHumor
	err := r.db.Where("paciente_id = ? AND data_hora_registro BETWEEN ? AND ?", pacienteID, inicio, fim).Find(&registros).Error
	return registros, err
}
