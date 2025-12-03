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

func (r *gormInstrumentoRepositorio) BuscarInstrumentoPorID(tx *gorm.DB, instrumentoID uint) (*dominio.Instrumento, error) {
	var instrumento dominio.Instrumento
	if err := tx.First(&instrumento, instrumentoID).Error; err != nil {
		return nil, err
	}
	return &instrumento, nil
}

func (r *gormInstrumentoRepositorio) CriarAtribuicao(tx *gorm.DB, atribuicao *dominio.Atribuicao) error {
	return tx.Create(atribuicao).Error
}
