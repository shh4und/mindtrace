package postgres

import (
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"
	"time"

	"gorm.io/gorm"
)

type gormRegistroHumorRepositorio struct {
	db *gorm.DB
}

func NovoGormRegistroHumorRepositorio(db *gorm.DB) repositorios.RegistroHumorRepositorio {
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

func (r *gormRegistroHumorRepositorio) BuscarUltimoRegistroDePaciente(pacienteID uint) (*dominio.RegistroHumor, error) {
	var registro *dominio.RegistroHumor
	err := r.db.Where("paciente_id = ?", pacienteID).Last(&registro).Error
	return registro, err
}

func (r *gormRegistroHumorRepositorio) BuscarPorNUltimosRegistros(pacienteID uint, numLimite int) ([]*dominio.RegistroHumor, error) {
	var registros []*dominio.RegistroHumor
	err := r.db.Where("paciente_id = ?", pacienteID).Order("created_at DESC").Limit(numLimite).Find(&registros).Error
	return registros, err
}
