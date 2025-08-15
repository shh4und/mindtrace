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

		_, err := s.repositorio.BuscarPorEmail(dto.Email)
		if err == nil {
			return errors.New("e-mail já cadastrado")
		}

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
			UsuarioID: novoUsuario.ID,
		}
		if err := s.repositorio.CriarPaciente(tx, novoPaciente); err != nil {
			return err
		}

		// // --- Passo 5: Criar a associação (A NOVA FORMA) ---

		// // 5a. Primeiro, encontramos a instância do profissional que está fazendo o cadastro.
		// profissional, err := s.repositorio.BuscarProfissionalPorID(tx, dto.ProfissionalID)
		// if err != nil {
		//     return errors.New("profissional não encontrado")
		// }

		// // 5b. Agora, usamos o método Association do GORM para adicionar o novo paciente.
		// // O GORM vai cuidar de inserir a linha na tabela 'profissional_paciente' para nós.
		// if err := tx.Model(profissional).Association("Pacientes").Append(novoPaciente); err != nil {
		//     return err
		// }

		// Preparar o objeto de retorno completo
		novoPaciente.Usuario = *novoUsuario
		pacienteCompleto = novoPaciente

		return nil // Sucesso na transação
	})

	return pacienteCompleto, err
}
