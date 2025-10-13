package servicos

import (
	"errors"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegistrarProfissionalDTO representa os dados para registrar um profissional
type RegistrarProfissionalDTO struct {
	Nome                 string
	Email                string
	Senha                string
	DataNascimento       time.Time
	Especialidade        string
	RegistroProfissional string
	CPF                  string
	Contato              string
}

// RegistrarPacienteDTO representa os dados para registrar um paciente
type RegistrarPacienteDTO struct {
	Nome                 string
	Email                string
	Senha                string
	Dependente           bool
	DataNascimento       time.Time
	DataInicioTratamento *time.Time
	HistoricoSaude       string
	CPF                  string
	NomeResponsavel      string
	ContatoResponsavel   string
	Contato              string
}

// AtualizarPerfilDTO representa os dados para atualizar o perfil do usuario
type AtualizarPerfilDTO struct {
	Nome    string `json:"nome" binding:"required"`
	Contato string `json:"contato"`
	Bio     string `json:"bio"`
	// Campos para Profissional
	Especialidade        string `json:"especialidade,omitempty"`
	RegistroProfissional string `json:"registro_profissional,omitempty"`
	// IdadeProfissional    *int8  `json:"idade_profissional,omitempty"` // Se aplicavel
	// Campos para Paciente
	DataNascimento     *time.Time `json:"data_nascimento,omitempty"`
	Dependente         *bool      `json:"dependente,omitempty"`
	NomeResponsavel    string     `json:"nome_responsavel,omitempty"`
	ContatoResponsavel string     `json:"contato_responsavel,omitempty"`
}

// AlterarSenhaDTO representa os dados para alterar a senha
type AlterarSenhaDTO struct {
	SenhaAtual  string `json:"senha_atual" binding:"required"`
	NovaSenha   string `json:"nova_senha" binding:"required,min=8"`
	NovaSenhaRe string `json:"nova_senha_re" binding:"required,min=8"`
}

var ErrEmailJaCadastrado = errors.New("e-mail existente")
var ErrCrendenciaisInvalidas = errors.New("credenciais invalidas")
var ErrUsuarioNaoEncontrado = errors.New("usuario nao encontrado")
var ErrSenhaNaoConfere = errors.New("a nova senha e a senha de confirmacao nao conferem")

// UsuarioServico define os metodos para gerenciamento de usuarios
type UsuarioServico interface {
	RegistrarProfissional(dto RegistrarProfissionalDTO) (*dominio.Profissional, error)
	RegistrarPaciente(dto RegistrarPacienteDTO) (*dominio.Paciente, error)
	Login(email, senha string) (string, error)
	BuscarUsuarioPorID(id uint) (*dominio.Usuario, error)
	ProprioPerfilPaciente(id uint) (*dominio.Paciente, error)
	ProprioPerfilProfissional(id uint) (*dominio.Profissional, error)
	ListarPacientesDoProfissional(userID uint) ([]dominio.Paciente, error)
	AtualizarPerfil(userID uint, dto AtualizarPerfilDTO) (*dominio.Usuario, error)
	AlterarSenha(userID uint, dto AlterarSenhaDTO) error
	DeletarPerfil(userID uint) error
}

// usuarioServico implementa a interface UsuarioServico
type usuarioServico struct {
	db          *gorm.DB
	repositorio repositorios.UsuarioRepositorio
}

// NovoUsuarioServico cria uma nova instancia de UsuarioServico
func NovoUsuarioServico(db *gorm.DB, repo repositorios.UsuarioRepositorio) UsuarioServico {
	return &usuarioServico{db: db, repositorio: repo}
}

// RegistrarProfissional registra um novo profissional no sistema
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
			Contato:     dto.Contato,
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
			Contato:     dto.Contato,
		}
		if err := s.repositorio.CriarUsuario(tx, novoUsuario); err != nil {
			return err
		}

		novoPaciente := &dominio.Paciente{
			UsuarioID:          novoUsuario.ID,
			DataNascimento:     dto.DataNascimento,
			Dependente:         dto.Dependente,
			NomeResponsavel:    dto.NomeResponsavel,
			ContatoResponsavel: dto.ContatoResponsavel,
		}
		if err := s.repositorio.CriarPaciente(tx, novoPaciente); err != nil {
			return err
		}

		// Preparar o objeto de retorno completo
		novoPaciente.Usuario = *novoUsuario
		pacienteCompleto = novoPaciente

		return nil // Sucesso na transacao
	})

	return pacienteCompleto, err
}

// Login autentica o usuario e retorna um token JWT
func (s *usuarioServico) Login(email, senha string) (string, error) {
	// Busca usuario pelo e-mail
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

	// Gera o token JWT
	claims := jwt.MapClaims{
		"sub":  usuario.ID,                           // Subject com o ID do usuario
		"role": usuario.TipoUsuario,                  // Adiciona o tipo de usuario como role
		"iat":  time.Now().Unix(),                    // Issued At indica quando o token foi criado
		"exp":  time.Now().Add(time.Hour * 1).Unix(), // Define expiracao do token em uma hora
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// BuscarUsuarioPorID busca um usuario pelo ID
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

// ProprioPerfilPaciente busca o perfil proprio do paciente
func (s *usuarioServico) ProprioPerfilPaciente(id uint) (*dominio.Paciente, error) {
	var pacienteEncontado *dominio.Paciente

	err := s.db.Transaction(func(tx *gorm.DB) error {

		paciente, err := s.repositorio.BuscarPacientePorUsuarioID(tx, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUsuarioNaoEncontrado
			}
			return err
		}
		pacienteEncontado = paciente
		return nil
	})
	return pacienteEncontado, err
}

func (s *usuarioServico) ProprioPerfilProfissional(id uint) (*dominio.Profissional, error) {
	var profissionalEncontrado *dominio.Profissional

	err := s.db.Transaction(func(tx *gorm.DB) error {

		profissional, err := s.repositorio.BuscarProfissionalPorUsuarioID(tx, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUsuarioNaoEncontrado
			}
			return err
		}
		profissionalEncontrado = profissional
		return nil
	})
	return profissionalEncontrado, err
}

// AtualizarPerfil atualiza o perfil do usuario
func (s *usuarioServico) AtualizarPerfil(userID uint, dto AtualizarPerfilDTO) (*dominio.Usuario, error) {
	var usuarioAtualizado *dominio.Usuario
	err := s.db.Transaction(func(tx *gorm.DB) error {
		usuario, err := s.repositorio.BuscarUsuarioPorID(userID)
		if err != nil {
			return err
		}
		usuario.Nome = dto.Nome
		usuario.Contato = dto.Contato
		usuario.Bio = dto.Bio
		if err := s.repositorio.Atualizar(tx, usuario); err != nil {
			return err
		}
		// Atualiza dados especificos conforme o tipo do usuario
		switch usuario.TipoUsuario {
		case "profissional":
			profissional, err := s.repositorio.BuscarProfissionalPorUsuarioID(tx, userID)
			if err != nil {
				return err
			}
			if dto.Especialidade != "" {
				profissional.Especialidade = dto.Especialidade
			}
			if dto.RegistroProfissional != "" {
				profissional.RegistroProfissional = dto.RegistroProfissional
			}

			if err := s.repositorio.AtualizarProfissional(tx, profissional); err != nil {
				return err
			}
		case "paciente":
			paciente, err := s.repositorio.BuscarPacientePorUsuarioID(tx, userID)
			if err != nil {
				return err
			}
			if dto.DataNascimento != nil {
				paciente.DataNascimento = *dto.DataNascimento
			}
			if dto.Dependente != nil {
				paciente.Dependente = *dto.Dependente
			}
			if dto.NomeResponsavel != "" {
				paciente.NomeResponsavel = dto.NomeResponsavel
			}
			if dto.ContatoResponsavel != "" {
				paciente.ContatoResponsavel = dto.ContatoResponsavel
			}
			if err := s.repositorio.AtualizarPaciente(tx, paciente); err != nil {
				return err
			}
		}
		usuarioAtualizado = usuario
		return nil
	})
	return usuarioAtualizado, err
}

// AlterarSenha altera a senha do usuario
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

func (s *usuarioServico) ListarPacientesDoProfissional(userID uint) ([]dominio.Paciente, error) {
	var pacientes []dominio.Paciente
	err := s.db.Transaction(func(tx *gorm.DB) error {
		profissional, err := s.repositorio.BuscarProfissionalPorUsuarioID(tx, userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUsuarioNaoEncontrado
			}
			return err
		}
		pacientes, err = s.repositorio.BuscarPacientesDoProfissional(tx, profissional.ID)
		return err
	})
	return pacientes, err
}

// DeletarPerfil deleta o perfil do usuario
func (s *usuarioServico) DeletarPerfil(userID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		_, err := s.repositorio.BuscarUsuarioPorID(userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUsuarioNaoEncontrado
			}
			return err
		}
		return s.repositorio.DeletarUsuario(tx, userID)
	})
}
