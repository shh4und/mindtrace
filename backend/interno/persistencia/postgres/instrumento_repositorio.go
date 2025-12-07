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
	var instrumento *dominio.Instrumento
	if err := tx.Preload("Perguntas").Preload("OpcoesEscala").First(&instrumento, instrumentoID).Error; err != nil {
		return nil, err
	}
	return instrumento, nil
}

func (r *gormInstrumentoRepositorio) CriarAtribuicao(tx *gorm.DB, atribuicao *dominio.Atribuicao) error {
	return tx.Create(atribuicao).Error
}

func (r *gormInstrumentoRepositorio) BuscarAtribuicaoPorID(tx *gorm.DB, atribuicaoID uint) (*dominio.Atribuicao, error) {
	var atribuicao *dominio.Atribuicao

	if err := tx.
		Preload("Instrumento.Perguntas").
		Preload("Instrumento.OpcoesEscala").
		Preload("Profissional.Usuario").
		Preload("Paciente.Usuario").
		Find(&atribuicao, atribuicaoID).Error; err != nil {
		return nil, err
	}
	return atribuicao, nil
}

func (r *gormInstrumentoRepositorio) BuscarAtribuicoesPaciente(tx *gorm.DB, pacId uint) ([]*dominio.Atribuicao, error) {
	var atribuicoes []*dominio.Atribuicao

	if err := tx.
		Preload("Instrumento.Perguntas").
		Preload("Profissional.Usuario").
		Preload("Paciente.Usuario").
		Where("paciente_id = ?", pacId).
		Find(&atribuicoes).Error; err != nil {
		return nil, err
	}
	return atribuicoes, nil
}

func (r *gormInstrumentoRepositorio) BuscarAtribuicoesProfissional(tx *gorm.DB, profId uint) ([]*dominio.Atribuicao, error) {
	var atribuicoes []*dominio.Atribuicao

	if err := tx.
		Preload("Instrumento.Perguntas").
		Preload("Profissional.Usuario").
		Preload("Paciente.Usuario").
		Where("profissional_id = ?", profId).
		Find(&atribuicoes).Error; err != nil {
		return nil, err
	}
	return atribuicoes, nil
}
