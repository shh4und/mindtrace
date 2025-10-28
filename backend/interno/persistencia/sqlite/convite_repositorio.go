package sqlite

import (
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"

	"gorm.io/gorm"
)

type gormConviteRepositorio struct{ db *gorm.DB }

func NovoGormConviteRepositorio(db *gorm.DB) repositorios.ConviteRepositorio {
	return &gormConviteRepositorio{db: db}
}

func (r *gormConviteRepositorio) CriarConvite(tx *gorm.DB, convite *dominio.Convite) error {
	return tx.Create(convite).Error
}

func (r *gormConviteRepositorio) BuscarConvitePorToken(tx *gorm.DB, token string) (*dominio.Convite, error) {
	var convite dominio.Convite
	if err := tx.Where("token = ? AND usado = ?", token, false).First(&convite).Error; err != nil {
		return nil, err
	}
	return &convite, nil
}

func (r *gormConviteRepositorio) MarcarConviteComoUsado(tx *gorm.DB, convite *dominio.Convite) error {
	return tx.Model(&dominio.Convite{}).Where("id = ?", convite.ID).Updates(map[string]interface{}{
		"usado":       convite.Usado,
		"paciente_id": convite.PacienteID,
	}).Error
}
