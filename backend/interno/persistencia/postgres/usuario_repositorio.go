package postgres

import (
	"mindtrace/backend/interno/dominio"

	"gorm.io/gorm"
)

type UsuarioRepositorio interface {
	CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error
	CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error
	BuscarPorEmail(email string) (*dominio.Usuario, error)
}

type gormUsuarioRepositorio struct {
	db *gorm.DB
}

func NewGormUsuarioRepositorio(db *gorm.DB) UsuarioRepositorio {
	return &gormUsuarioRepositorio{db: db}
}

func (r *gormUsuarioRepositorio) CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error {
	return tx.Create(usuario).Error
}

func (r *gormUsuarioRepositorio) CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	return tx.Create(profissional).Error
}

func (r *gormUsuarioRepositorio) BuscarPorEmail(email string) (*dominio.Usuario, error) {
	var usuario dominio.Usuario
	if err := r.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}
