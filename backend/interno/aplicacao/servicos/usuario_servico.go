package servicos

import (
	"errors"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/postgres"
	"time"

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

type RegistrarPacienteDTO struct {
	Nome                 string
	Email                string
	Senha                string
	ProfissionalID       uint
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

func (s *usuarioServico) RegistrarPaciente(dto RegistrarPacienteDTO) (*dominio.Paciente, error) {
	var pacienteCompleto *dominio.Paciente

	// Inicia a transação
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// 1. Verificar se o e-mail já existe
		_, err := s.repositorio.BuscarPorEmail(dto.Email)
		if err == nil { // Se err for nil, significa que o usuário FOI encontrado
			return errors.New("e-mail já cadastrado")
		}

		// 2. Criptografar a senha
		hashSenha, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// 3. Criar a entidade Usuario
		novoUsuario := &dominio.Usuario{
			Nome:  dto.Nome,
			Email: dto.Email,
			Senha: string(hashSenha),
		}
		if err := s.repositorio.CriarUsuario(tx, novoUsuario); err != nil {
			return err
		}

		// 4. Criar a entidade Paciente, ligando ao Usuario recém-criado
		novoPaciente := &dominio.Paciente{
			UsuarioID: novoUsuario.ID,
		}
		if err := s.repositorio.CriarPaciente(tx, novoPaciente); err != nil {
			return err
		}

		// 5. Criar a associação entre o Profissional e o novo Paciente (PASSO CRÍTICO)
		if err := s.repositorio.CriarAssociacao(tx, dto.ProfissionalID, novoPaciente.ID); err != nil {
			return err
		}

		// Preparar o objeto de retorno completo
		novoPaciente.Usuario = *novoUsuario
		pacienteCompleto = novoPaciente

		return nil // Sucesso na transação
	})

	return pacienteCompleto, err
}
