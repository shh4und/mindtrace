package postgres

import (
	"mindtrace/backend/interno/dominio"

	"gorm.io/gorm"
)

type RegistroHumorRepositorio interface {
	CriarRegistroHumor(tx *gorm.DB, registro *dominio.RegistroHumor) error
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
