package postgres

import (
	"mindtrace/backend/internal/domain"

	"gorm.io/gorm"
)

type UsuarioRepository interface {
	Salvar(usuario *domain.Usuario) error
}

type gormUsuarioRepository struct {
	db *gorm.DB
}

func NewGormUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &gormUsuarioRepository{db: db}
}

func (r *gormUsuarioRepository) Salvar(usuario *domain.Usuario) error {
	result := r.db.Create(usuario)
	return result.Error
}
