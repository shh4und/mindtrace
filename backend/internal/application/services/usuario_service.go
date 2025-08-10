package services

import (
	"mindtrace/backend/internal/domain"
	"mindtrace/backend/internal/persistence/postgres"

	"golang.org/x/crypto/bcrypt"
)

type UsuarioService interface {
	Register(nome, email, senha string) (*domain.Usuario, error)
}

type usuarioService struct {
	repo postgres.UsuarioRepository
}

func NewUsuarioService(repo postgres.UsuarioRepository) UsuarioService {
	return &usuarioService{repo: repo}
}

func (s *usuarioService) Register(nome, email, senha string) (*domain.Usuario, error) {
	hashSenha, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	newUsuario := &domain.Usuario{
		Nome:  nome,
		Email: email,
		Senha: string(hashSenha),
	}

	err = s.repo.Save(newUsuario)

	if err != nil {
		return nil, err
	}

	return newUsuario, nil
}
