package postgres

import (
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"

	"gorm.io/gorm"
)

type gormInstrumentoRepositorio struct {
	db *gorm.DB
}

func NovoGormInstrumentoRepositorio(db *gorm.DB) repositorios.InstrumentoRepositorio {
	return &gormInstrumentoRepositorio{db: db}
}

func (r *gormInstrumentoRepositorio) BuscarTodosAtivos(tx *gorm.DB) ([]*dominio.Instrumento, error) {
	var instrumentos []*dominio.Instrumento
	err := tx.Where("esta_ativo = TRUE").Find(&instrumentos).Error
	return instrumentos, err
}
