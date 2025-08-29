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
	CPF                  string
}

type RegistrarPacienteDTO struct {
	Nome                 string
	Email                string
	Senha                string
	EhDependente         bool
	Idade                int8
	DataInicioTratamento *time.Time
	HistoricoSaude       string
	CPF                  string
	NomeResponsavel      string
	ContatoResponsavel   string
}

type AtualizarPerfilDTO struct {
	Nome    string `json:"nome" binding:"required"`
	Contato string `json:"contato"`
	Bio     string `json:"bio"`
}

type AlterarSenhaDTO struct {
	SenhaAtual  string `json:"senha_atual" binding:"required"`
	NovaSenha   string `json:"nova_senha" binding:"required,min=8"`
	NovaSenhaRe string `json:"nova_senha_re" binding:"required,min=8"`
}

var ErrEmailJaCadastrado = errors.New("e-mail existente")
var ErrCrendenciaisInvalidas = errors.New("credenciais invalidas")
var ErrUsuarioNaoEncontrado = errors.New("usuario nao encontrado")
var ErrSenhaNaoConfere = errors.New("a nova senha e a senha de confirmação não conferem")

type UsuarioServico interface {
	RegistrarProfissional(dto RegistrarProfissionalDTO) (*dominio.Profissional, error)
	RegistrarPaciente(dto RegistrarPacienteDTO) (*dominio.Paciente, error)
	Login(email, senha string) (string, error)
	BuscarUsuarioPorID(id uint) (*dominio.Usuario, error)
	BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error)
	AtualizarPerfil(userID uint, dto AtualizarPerfilDTO) (*dominio.Usuario, error)
	AlterarSenha(userID uint, dto AlterarSenhaDTO) error
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
			return ErrEmailJaCadastrado
		}

		hashSenha, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		novoUsuario := &dominio.Usuario{
			Nome:        dto.Nome,
			Email:       dto.Email,
			Senha:       string(hashSenha),
			TipoUsuario: "profissional",
			CPF:         dto.CPF,
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
			return ErrEmailJaCadastrado
		}

		hashSenha, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		novoUsuario := &dominio.Usuario{
			Nome:        dto.Nome,
			Email:       dto.Email,
			Senha:       string(hashSenha),
			TipoUsuario: "paciente",
			CPF:         dto.CPF,
		}
		if err := s.repositorio.CriarUsuario(tx, novoUsuario); err != nil {
			return err
		}

		novoPaciente := &dominio.Paciente{
			UsuarioID:          novoUsuario.ID,
			Idade:              dto.Idade,
			EhDependente:       dto.EhDependente,
			NomeResponsavel:    dto.NomeResponsavel,
			ContatoResponsavel: dto.ContatoResponsavel,
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
			return "", ErrUsuarioNaoEncontrado
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))
	if err != nil {
		return "", ErrCrendenciaisInvalidas
	}

	// Gerar o token JWT
	claims := jwt.MapClaims{
		"sub":  usuario.ID,                           // "Subject", o ID do usuário
		"role": usuario.TipoUsuario,                  // Adiciona o tipo de usuário (role)
		"iat":  time.Now().Unix(),                    // "Issued At", quando o token foi criado
		"exp":  time.Now().Add(time.Hour * 1).Unix(), // Data de expiração (1 hora)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *usuarioServico) BuscarUsuarioPorID(id uint) (*dominio.Usuario, error) {
	usuario, err := s.repositorio.BuscarUsuarioPorID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUsuarioNaoEncontrado
		}
		return nil, err
	}
	return usuario, nil
}

func (s *usuarioServico) BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	paciente, err := s.repositorio.BuscarPacientePorID(tx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPacienteNaoEncontrado
		}
		return nil, err
	}
	return paciente, nil
}

func (s *usuarioServico) AtualizarPerfil(userID uint, dto AtualizarPerfilDTO) (*dominio.Usuario, error) {
	var usuarioAtualizado *dominio.Usuario
	err := s.db.Transaction(func(tx *gorm.DB) error {
		usuario, err := s.repositorio.BuscarUsuarioPorID(userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUsuarioNaoEncontrado
			}
			return err
		}

		usuario.Nome = dto.Nome
		usuario.Contato = dto.Contato
		usuario.Bio = dto.Bio

		if err := s.repositorio.Atualizar(tx, usuario); err != nil {
			return err
		}
		usuarioAtualizado = usuario
		return nil
	})

	return usuarioAtualizado, err
}

func (s *usuarioServico) AlterarSenha(userID uint, dto AlterarSenhaDTO) error {

	if dto.NovaSenha != dto.NovaSenhaRe {
		return ErrSenhaNaoConfere
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		usuario, err := s.repositorio.BuscarUsuarioPorID(userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUsuarioNaoEncontrado
			}
			return err
		}

		if err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(dto.SenhaAtual)); err != nil {
			return ErrCrendenciaisInvalidas
		}

		novaSenhaHash, err := bcrypt.GenerateFromPassword([]byte(dto.NovaSenha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		usuario.Senha = string(novaSenhaHash)

		return s.repositorio.Atualizar(tx, usuario)
	})

	return err
}
