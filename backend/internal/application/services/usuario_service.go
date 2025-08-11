package services

import (
	"mindtrace/backend/internal/domain"
	"mindtrace/backend/internal/persistence/postgres"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterProfissionalDTO struct {
	Name                 string
	Email                string
	Password             string
	Specialty            string
	ProfessionalRegistry string
}

type UsuarioService interface {
	RegisterProfissional(dto RegisterProfissionalDTO) (*domain.Profissional, error)
}

type usuarioService struct {
	db   *gorm.DB
	repo postgres.UsuarioRepository
}

func NewUsuarioService(db *gorm.DB, repo postgres.UsuarioRepository) UsuarioService {
	return &usuarioService{db: db, repo: repo}
}

func (s *usuarioService) RegisterProfissional(dto RegisterProfissionalDTO) (*domain.Profissional, error) {
	var registeredProfissional *domain.Profissional

	err := s.db.Transaction(func(tx *gorm.DB) error {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		newUsuario := &domain.Usuario{
			Name:     dto.Name,
			Email:    dto.Email,
			Password: string(hashedPassword),
		}

		if err := s.repo.CreateUsuario(tx, newUsuario); err != nil {
			return err
		}

		newProfissional := &domain.Profissional{
			Specialty:            dto.Specialty,
			ProfessionalRegistry: dto.ProfessionalRegistry,
			UsuarioID:            newUsuario.ID,
		}

		if err := s.repo.CreateProfissional(tx, newProfissional); err != nil {
			return err
		}

		newProfissional.Usuario = *newUsuario
		registeredProfissional = newProfissional

		return nil
	})

	return registeredProfissional, err
}
