package servicos

import (
	"errors"
	"log"
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
	email       EmailServico
}

// NovoUsuarioServico cria uma nova instancia de UsuarioServico
func NovoUsuarioServico(db *gorm.DB, repo repositorios.UsuarioRepositorio, emailSvc EmailServico) UsuarioServico {
	return &usuarioServico{db: db, repositorio: repo, email: emailSvc}
}

/*
TODO:
  - INSERIR USUARIO COM EMAIL HASH NO DB
  - USAR FLAG EstaAtivo PARA PERMITIR LOGIN
  - ADICIONAR TIMEOUT += 48H PARA INVALIDAR TOKEN
  - NOVO CONTROLADOR/ROTA PARA ATIVACAO DE EMAIL
  - testar.
*/
// RegistrarProfissional registra um novo profissional no sistema
func (s *usuarioServico) RegistrarProfissional(dtoIn *dtos.RegistrarProfissionalDTOIn) (*dtos.ProfissionalDTOOut, error) {
	var profissionalRegistrado *dominio.Profissional

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Verifica se o e-mail ja esta cadastrado
		_, err := s.repositorio.BuscarPorEmail(dtoIn.Email)
		if err == nil {
			return dominio.ErrEmailJaCadastrado
		}

		// retorna err se diferente de "record not found"
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// Usa o mapper para criar as entidades a partir do DTOIn
		novoUsuario, novoProfissional := mappers.RegistrarProfissionalDTOInParaEntidade(dtoIn)

		// Validar usuario
		if err := novoUsuario.Validar(); err != nil {
			return err
		}

		// Validar senha
		if err := novoUsuario.ValidarSenha(dtoIn.Senha); err != nil {
			return err
		}

		// Validar profissional
		novoProfissional.Usuario = *novoUsuario
		if err := novoProfissional.Validar(); err != nil {
			return err
		}

		// Hash da senha
		hashSenha, err := bcrypt.GenerateFromPassword([]byte(novoUsuario.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// Gera hash token para posterior verificacao/ativacao de conta
		tokenHash, err := GenerateSecureToken()
		if err != nil {
			return err
		}

		novoUsuario.Senha = string(hashSenha)
		novoUsuario.TipoUsuario = dominio.TipoUsuarioProfissional
		novoUsuario.EmailVerifHash = &tokenHash

		// Cria o usuario
		if err := s.repositorio.CriarUsuario(tx, novoUsuario); err != nil {
			return err
		}

		// Cria o profissional
		novoProfissional.Usuario = *novoUsuario
		if err := s.repositorio.CriarProfissional(tx, novoProfissional); err != nil {
			return err
		}

		// Prepara o objeto de retorno completo
		profissionalRegistrado = novoProfissional
		go func() { // Rodar em goroutine para não travar a resposta HTTP
			if err := s.email.EnviarEmailAtivacao(dtoIn.Email, tokenHash); err != nil {
				log.Printf("Error at sending activation email in go routine: %v", err)
			}
		}()

		return nil
	})

	return mappers.ProfissionalParaDTOOut(profissionalRegistrado), err
}

/*
TODO:
  - INSERIR USUARIO COM EMAIL HASH NO DB
  - USAR FLAG EstaAtivo PARA PERMITIR LOGIN
  - ADICIONAR TIMEOUT += 48H PARA INVALIDAR TOKEN
  - NOVO CONTROLADOR/ROTA PARA ATIVACAO DE EMAIL
  - testar.
*/
func (s *usuarioServico) RegistrarPaciente(dtoIn *dtos.RegistrarPacienteDTOIn) (*dtos.PacienteDTOOut, error) {
	var pacienteCompleto *dominio.Paciente

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Verifica se o e-mail ja esta cadastrado
		_, err := s.repositorio.BuscarPorEmail(dtoIn.Email)
		if err == nil {
			return dominio.ErrEmailJaCadastrado
		}
		// retorna err se diferente de "record not found"
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		// Usa o mapper para criar as entidades a partir do DTOIn
		novoUsuario, novoPaciente := mappers.RegistrarPacienteDTOInParaEntidade(dtoIn)

		// Validar usuario
		if err := novoUsuario.Validar(); err != nil {
			return err
		}

		// Validar senha
		if err := novoUsuario.ValidarSenha(dtoIn.Senha); err != nil {
			return err
		}

		// Validar paciente
		novoPaciente.Usuario = *novoUsuario
		if err := novoPaciente.Validar(); err != nil {
			return err
		}

		// Hash da senha
		hashSenha, err := bcrypt.GenerateFromPassword([]byte(novoUsuario.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		// Gera hash token para posterior verificacao/ativacao de conta
		tokenHash, err := GenerateSecureToken()
		if err != nil {
			return err
		}

		novoUsuario.Senha = string(hashSenha)
		novoUsuario.TipoUsuario = dominio.TipoUsuarioProfissional
		novoUsuario.EmailVerifHash = &tokenHash
		// Cria o usuario
		if err := s.repositorio.CriarUsuario(tx, novoUsuario); err != nil {
			return err
		}

		// Cria o paciente
		novoPaciente.Usuario = *novoUsuario
		if err := s.repositorio.CriarPaciente(tx, novoPaciente); err != nil {
			return err
		}

		// Prepara o objeto de retorno completo
		pacienteCompleto = novoPaciente
		go func() { // Rodar em goroutine para não travar a resposta HTTP
			if err := s.email.EnviarEmailAtivacao(dtoIn.Email, tokenHash); err != nil {
				log.Printf("Error at sending activation email in go routine: %v", err)
			}
		}()
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
			return "", dominio.ErrUsuarioNaoEncontrado
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))
	if err != nil {
		return "", dominio.ErrCrendenciaisInvalidas
	}

	// Gera o token JWT
	claims := jwt.MapClaims{
		"sub":  usuario.ID,                                         // Subject com o ID do usuario
		"role": dominio.TipoUsuarioParaString(usuario.TipoUsuario), // Adiciona o tipo de usuario como role (string)
		"iat":  time.Now().Unix(),                                  // Issued At indica quando o token foi criado
		"exp":  time.Now().Add(time.Hour * 1).Unix(),               // Define expiracao do token em uma hora
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
			return nil, dominio.ErrUsuarioNaoEncontrado
		}
		return nil, err
	}
	return mappers.UsuarioParaDTOOut(usuario), nil
}

// ProprioPerfilPaciente busca o perfil proprio do paciente
func (s *usuarioServico) ProprioPerfilPaciente(pacID uint) (*dtos.PacienteDTOOut, error) {
	var pacienteEncontrado *dominio.Paciente

	err := s.db.Transaction(func(tx *gorm.DB) error {

		paciente, err := s.repositorio.BuscarPacientePorUsuarioID(tx, pacID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return dominio.ErrUsuarioNaoEncontrado
			}
			return err
		}
		pacienteEncontrado = paciente
		return nil
	})
	return mappers.PacienteParaDTOOut(pacienteEncontrado), err
}

func (s *usuarioServico) ProprioPerfilProfissional(profID uint) (*dtos.ProfissionalDTOOut, error) {
	var profissionalEncontrado *dominio.Profissional

	err := s.db.Transaction(func(tx *gorm.DB) error {

		profissional, err := s.repositorio.BuscarProfissionalPorUsuarioID(tx, profID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return dominio.ErrUsuarioNaoEncontrado
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

		// Validar nome atualizado
		if err := usuario.ValidarNome(); err != nil {
			return err
		}

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

			// Validar dados do profissional
			profissional.Usuario = *usuario
			if err := profissional.Validar(); err != nil {
				return err
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

			// Validar dados do paciente
			paciente.Usuario = *usuario
			if err := paciente.Validar(); err != nil {
				return err
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
		return dominio.ErrSenhaNaoConfere
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		usuario, err := s.repositorio.BuscarUsuarioPorID(userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return dominio.ErrUsuarioNaoEncontrado
			}
			return err
		}

		if err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(dtoIn.SenhaAtual)); err != nil {
			return dominio.ErrCrendenciaisInvalidas
		}

		// Validar nova senha
		if err := usuario.ValidarSenha(dtoIn.NovaSenha); err != nil {
			return err
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
				return dominio.ErrUsuarioNaoEncontrado
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
				return dominio.ErrUsuarioNaoEncontrado
			}
			return err
		}
		return s.repositorio.DeletarUsuario(tx, userID)
	})
}
