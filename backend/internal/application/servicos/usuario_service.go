package services

import (
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/postgres"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegistrarProfissionalDTO struct {
	Nome                 string
	Email                string
	Senha                string
	Especialidade        string
	RegistroProfissional string
}

type UsuarioServico interface {
	RegistrarProfissional(dto RegistrarProfissionalDTO) (*dominio.Profissional, error)
}

type usuarioServico struct {
	db          *gorm.DB
	repositorio postgres.UsuarioRepositorio
}

func NewUsuarioServico(db *gorm.DB, repo postgres.UsuarioRepositorio) UsuarioServico {
	return &usuarioServico{db: db, repositorio: repo}
}

func (s *usuarioServico) RegistrarProfissional(dto RegistrarProfissionalDTO) (*dominio.Profissional, error) {
	var profissionalRegistrado *dominio.Profissional

	err := s.db.Transaction(func(tx *gorm.DB) error {
		senhaComHash, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		novoUsuario := &dominio.Usuario{
			Nome:  dto.Nome,
			Email: dto.Email,
			Senha: string(senhaComHash),
		}

		if err := s.repositorio.CriarUsuario(tx, novoUsuario); err != nil {
			return err
		}

		novoProfissional := &dominio.Profissional{
			Especialidade:        dto.Especialidade,
			RegistroProfissional: dto.RegistroProfissional,
			UsuarioID:            novoUsuario.ID,
		}

		if err := s.repositorio.CriarProfissional(tx, novoProfissional); err != nil {
			return err
		}

		novoProfissional.Usuario = *novoUsuario
		profissionalRegistrado = novoProfissional

		return nil
	})

	return profissionalRegistrado, err
}
