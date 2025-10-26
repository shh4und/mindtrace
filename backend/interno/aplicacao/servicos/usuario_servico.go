package servicos

import (
	"errors"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/mappers"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrEmailJaCadastrado = errors.New("e-mail existente")
var ErrCrendenciaisInvalidas = errors.New("credenciais invalidas")
var ErrUsuarioNaoEncontrado = errors.New("usuario nao encontrado")
var ErrSenhaNaoConfere = errors.New("a nova senha e a senha de confirmacao nao conferem")

// UsuarioServico define os metodos para gerenciamento de usuarios
type UsuarioServico interface {
	RegistrarProfissional(dtoIn *dtos.RegistrarProfissionalDTOIn) (*dtos.ProfissionalDTOOut, error)
	RegistrarPaciente(dtoIn *dtos.RegistrarPacienteDTOIn) (*dtos.PacienteDTOOut, error)
	Login(email, senha string) (string, error)
	BuscarUsuarioPorID(userID uint) (*dtos.UsuarioDTOOut, error)
	ProprioPerfilPaciente(pacID uint) (*dtos.PacienteDTOOut, error)
	ProprioPerfilProfissional(profID uint) (*dtos.ProfissionalDTOOut, error)
	ListarPacientesDoProfissional(userID uint) ([]dtos.PacienteDTOOut, error)
	AtualizarPerfil(userID uint, dtoIn *dtos.AtualizarPerfilDTOIn) error
	AlterarSenha(userID uint, dtoIn *dtos.AlterarSenhaDTOIn) error
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
func (s *usuarioServico) RegistrarProfissional(dtoIn *dtos.RegistrarProfissionalDTOIn) (*dtos.ProfissionalDTOOut, error) {
	var profissionalRegistrado *dominio.Profissional

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Verifica se o e-mail já está cadastrado
		_, err := s.repositorio.BuscarPorEmail(dtoIn.Email)
		if err == nil {
			return ErrEmailJaCadastrado
		}

		// Usa o mapper para criar as entidades a partir do DTOIn
		novoUsuario, novoProfissional := mappers.RegistrarProfissionalDTOInParaEntidade(dtoIn)

		// Hash da senha
		hashSenha, err := bcrypt.GenerateFromPassword([]byte(novoUsuario.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		novoUsuario.Senha = string(hashSenha)

		// Cria o usuário
		if err := s.repositorio.CriarUsuario(tx, novoUsuario); err != nil {
			return err
		}

		// Define o UsuarioID do profissional
		novoProfissional.UsuarioID = novoUsuario.ID

		// Cria o profissional
		if err := s.repositorio.CriarProfissional(tx, novoProfissional); err != nil {
			return err
		}

		// Prepara o objeto de retorno completo
		novoProfissional.Usuario = *novoUsuario
		profissionalRegistrado = novoProfissional

		return nil
	})

	return mappers.ProfissionalParaDTOOut(profissionalRegistrado), err
}

func (s *usuarioServico) RegistrarPaciente(dtoIn *dtos.RegistrarPacienteDTOIn) (*dtos.PacienteDTOOut, error) {
	var pacienteCompleto *dominio.Paciente

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Verifica se o e-mail já está cadastrado
		_, err := s.repositorio.BuscarPorEmail(dtoIn.Email)
		if err == nil {
			return ErrEmailJaCadastrado
		}

		// Usa o mapper para criar as entidades a partir do DTOIn
		novoUsuario, novoPaciente := mappers.RegistrarPacienteDTOInParaEntidade(dtoIn)

		// Hash da senha
		hashSenha, err := bcrypt.GenerateFromPassword([]byte(novoUsuario.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		novoUsuario.Senha = string(hashSenha)

		// Cria o usuário
		if err := s.repositorio.CriarUsuario(tx, novoUsuario); err != nil {
			return err
		}

		// Define o UsuarioID do paciente
		novoPaciente.UsuarioID = novoUsuario.ID

		// Cria o paciente
		if err := s.repositorio.CriarPaciente(tx, novoPaciente); err != nil {
			return err
		}

		// Prepara o objeto de retorno completo
		novoPaciente.Usuario = *novoUsuario
		pacienteCompleto = novoPaciente

		return nil
	})

	return mappers.PacienteParaDTOOut(pacienteCompleto), err
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
func (s *usuarioServico) BuscarUsuarioPorID(userID uint) (*dtos.UsuarioDTOOut, error) {
	usuario, err := s.repositorio.BuscarUsuarioPorID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUsuarioNaoEncontrado
		}
		return nil, err
	}
	return mappers.UsuarioParaDTOOut(usuario), nil
}

// ProprioPerfilPaciente busca o perfil proprio do paciente
func (s *usuarioServico) ProprioPerfilPaciente(pacID uint) (*dtos.PacienteDTOOut, error) {
	var pacienteEncontado *dominio.Paciente

	err := s.db.Transaction(func(tx *gorm.DB) error {

		paciente, err := s.repositorio.BuscarPacientePorUsuarioID(tx, pacID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUsuarioNaoEncontrado
			}
			return err
		}
		pacienteEncontado = paciente
		return nil
	})
	return mappers.PacienteParaDTOOut(pacienteEncontado), err
}

func (s *usuarioServico) ProprioPerfilProfissional(profID uint) (*dtos.ProfissionalDTOOut, error) {
	var profissionalEncontrado *dominio.Profissional

	err := s.db.Transaction(func(tx *gorm.DB) error {

		profissional, err := s.repositorio.BuscarProfissionalPorUsuarioID(tx, profID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUsuarioNaoEncontrado
			}
			return err
		}
		profissionalEncontrado = profissional
		return nil
	})
	return mappers.ProfissionalParaDTOOut(profissionalEncontrado), err
}

// AtualizarPerfil atualiza o perfil do usuario
func (s *usuarioServico) AtualizarPerfil(userID uint, dtoIn *dtos.AtualizarPerfilDTOIn) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		usuario, err := s.repositorio.BuscarUsuarioPorID(userID)
		if err != nil {
			return err
		}
		usuario.Nome = dtoIn.Nome
		usuario.Contato = dtoIn.Contato
		usuario.Bio = dtoIn.Bio
		if err := s.repositorio.Atualizar(tx, usuario); err != nil {
			return err
		}
		// Atualiza dados especificos conforme o tipo do usuario
		switch usuario.TipoUsuario {
		case 2: // tipo de usuario 2 = profissional
			profissional, err := s.repositorio.BuscarProfissionalPorUsuarioID(tx, userID)
			if err != nil {
				return err
			}
			if dtoIn.Especialidade != "" {
				profissional.Especialidade = dtoIn.Especialidade
			}
			if dtoIn.RegistroProfissional != "" {
				profissional.RegistroProfissional = dtoIn.RegistroProfissional
			}

			if err := s.repositorio.AtualizarProfissional(tx, profissional); err != nil {
				return err
			}
		case 3: // tipo de usuario 3 = paciente
			paciente, err := s.repositorio.BuscarPacientePorUsuarioID(tx, userID)
			if err != nil {
				return err
			}
			if dtoIn.DataNascimento != nil {
				paciente.DataNascimento = *dtoIn.DataNascimento
			}
			if dtoIn.Dependente != nil {
				paciente.Dependente = *dtoIn.Dependente
			}
			if dtoIn.NomeResponsavel != "" {
				paciente.NomeResponsavel = dtoIn.NomeResponsavel
			}
			if dtoIn.ContatoResponsavel != "" {
				paciente.ContatoResponsavel = dtoIn.ContatoResponsavel
			}
			if err := s.repositorio.AtualizarPaciente(tx, paciente); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// AlterarSenha altera a senha do usuario
func (s *usuarioServico) AlterarSenha(userID uint, dtoIn *dtos.AlterarSenhaDTOIn) error {

	if dtoIn.NovaSenha != dtoIn.NovaSenhaRe {
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

		if err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(dtoIn.SenhaAtual)); err != nil {
			return ErrCrendenciaisInvalidas
		}

		novaSenhaHash, err := bcrypt.GenerateFromPassword([]byte(dtoIn.NovaSenha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		usuario.Senha = string(novaSenhaHash)

		return s.repositorio.Atualizar(tx, usuario)
	})

	return err
}

func (s *usuarioServico) ListarPacientesDoProfissional(userID uint) ([]dtos.PacienteDTOOut, error) {
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

	return mappers.PacientesParaDTOOut(pacientes), err
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
