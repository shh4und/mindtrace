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

func (r *gormInstrumentoRepositorio) CriarReposta(tx *gorm.DB, resposta *dominio.Resposta, atribuicaoId uint) error {

	if err := tx.Model(&dominio.Atribuicao{}).Where("id = ? AND data_resposta is NULL", resposta.AtribuicaoID).Updates(map[string]interface{}{
		"status":        "RESPONDIDO",
		"data_resposta": resposta.DataResposta,
	}).Error; err != nil {
		return err
	}

	return tx.Create(resposta).Error
}
func (r *gormInstrumentoRepositorio) BuscarRespostaPorAtribuicaoID(tx *gorm.DB, atribuicaoID uint) (*dominio.Resposta, error) {
	var resposta *dominio.Resposta

	if err := tx.Where("atribuicao_id = ?", atribuicaoID).First(&resposta).Error; err != nil {
		return nil, err
	}

	return resposta, nil
}

func (r *gormInstrumentoRepositorio) BuscarRespostaCompletaPorAtribuicaoID(tx *gorm.DB, atribuicaoID uint) (*dominio.Resposta, error) {
	var resposta *dominio.Resposta

	if err := tx.
		Preload("Atribuicao").
		Preload("Atribuicao.Instrumento.Perguntas").
		Preload("Atribuicao.Instrumento.OpcoesEscala").
		Preload("Atribuicao.Paciente.Usuario").
		Preload("Atribuicao.Profissional.Usuario").
		Where("atribuicao_id = ?", atribuicaoID).
		First(&resposta).Error; err != nil {
		return nil, err
	}

	return resposta, nil
}
