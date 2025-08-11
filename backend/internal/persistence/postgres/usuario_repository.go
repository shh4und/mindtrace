package postgres

import (
	"mindtrace/backend/internal/domain"

	"gorm.io/gorm"
)

type UsuarioRepository interface {
	CreateUsuario(tx *gorm.DB, usuario *domain.Usuario) error
	CreateProfissional(tx *gorm.DB, profissional *domain.Profissional) error
	FindByEmail(email string) (*domain.Usuario, error)
}

type gormUsuarioRepository struct {
	db *gorm.DB
}

func NewGormUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &gormUsuarioRepository{db: db}
}

func (r *gormUsuarioRepository) CreateUsuario(tx *gorm.DB, usuario *domain.Usuario) error {
	return tx.Create(usuario).Error
}

func (r *gormUsuarioRepository) CreateProfissional(tx *gorm.DB, profissional *domain.Profissional) error {
	return tx.Create(profissional).Error
}

func (r *gormUsuarioRepository) FindByEmail(email string) (*domain.Usuario, error) {
	var usuario domain.Usuario
	if err := r.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}
