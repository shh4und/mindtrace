package servicos

import (
	"errors"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/postgres"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegistrarProfissionalDTO struct {
	Nome                 string
	Email                string
	Senha                string
	Idade                int8
	Especialidade        string
	RegistroProfissional string
}

type RegistrarPacienteDTO struct {
	Nome                 string
	Email                string
	Senha                string
	EhDependente         bool
	Idade                int8
	DataInicioTratamento *time.Time
	HistoricoSaude       string
}

type RegistrarResponsavelDTO struct {
	Nome             string
	Email            string
	Senha            string
	Parentesco       string
	ContatoPrincipal string
}

type UsuarioServico interface {
	RegistrarProfissional(dto RegistrarProfissionalDTO) (*dominio.Profissional, error)
	RegistrarPaciente(dto RegistrarPacienteDTO) (*dominio.Paciente, error)
	Login(email, senha string) (string, error)
}

type usuarioServico struct {
	db          *gorm.DB
	repositorio postgres.UsuarioRepositorio
}

func NovoUsuarioServico(db *gorm.DB, repo postgres.UsuarioRepositorio) UsuarioServico {
	return &usuarioServico{db: db, repositorio: repo}
}

func (s *usuarioServico) RegistrarProfissional(dto RegistrarProfissionalDTO) (*dominio.Profissional, error) {
	var profissionalRegistrado *dominio.Profissional

	err := s.db.Transaction(func(tx *gorm.DB) error {

		_, err := s.repositorio.BuscarPorEmail(dto.Email)
		if err == nil {
			return errors.New("e-mail já cadastrado")
		}

		hashSenha, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		novoUsuario := &dominio.Usuario{
			Nome:  dto.Nome,
			Email: dto.Email,
			Senha: string(hashSenha),
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

func (s *usuarioServico) RegistrarPaciente(dto RegistrarPacienteDTO) (*dominio.Paciente, error) {
	var pacienteCompleto *dominio.Paciente

	err := s.db.Transaction(func(tx *gorm.DB) error {

		_, err := s.repositorio.BuscarPorEmail(dto.Email)
		if err == nil {
			return errors.New("e-mail já cadastrado")
		}

		hashSenha, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		novoUsuario := &dominio.Usuario{
			Nome:  dto.Nome,
			Email: dto.Email,
			Senha: string(hashSenha),
		}
		if err := s.repositorio.CriarUsuario(tx, novoUsuario); err != nil {
			return err
		}

		novoPaciente := &dominio.Paciente{
			UsuarioID:    novoUsuario.ID,
			Idade:        dto.Idade,
			EhDependente: dto.EhDependente,
		}
		if err := s.repositorio.CriarPaciente(tx, novoPaciente); err != nil {
			return err
		}

		// Preparar o objeto de retorno completo
		novoPaciente.Usuario = *novoUsuario
		pacienteCompleto = novoPaciente

		return nil // Sucesso na transação
	})

	return pacienteCompleto, err
}

func (s *usuarioServico) Login(email, senha string) (string, error) {
	// Buscar usuário pelo e-mail
	usuario, err := s.repositorio.BuscarPorEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("credenciais inválidas")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	// Gerar o token JWT
	claims := jwt.MapClaims{
		"sub": usuario.ID,                            // "Subject", o ID do usuário
		"iat": time.Now().Unix(),                     // "Issued At", quando o token foi criado
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Data de expiração (24 horas)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
