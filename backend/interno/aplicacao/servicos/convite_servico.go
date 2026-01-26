package servicos

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/mappers"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"
	"time"

	"gorm.io/gorm"
)

// ConviteServico define os metodos para gerenciamento de convites
type ConviteServico interface {
	GerarConvite(userID uint, emailPac string) (*dtos.ConviteDTOOut, error)
	VincularPaciente(userID uint, token string) error
}

// conviteServico implementa a interface ConviteServico
type conviteServico struct {
	db                 *gorm.DB
	conviteRepositorio repositorios.ConviteRepositorio
	usuarioRepositorio repositorios.UsuarioRepositorio
	emailServico       EmailServico
}

// NovoConviteServico cria uma nova instancia de ConviteServico
func NovoConviteServico(db *gorm.DB, cr repositorios.ConviteRepositorio, ur repositorios.UsuarioRepositorio, es EmailServico) ConviteServico {
	return &conviteServico{
		db:                 db,
		conviteRepositorio: cr,
		usuarioRepositorio: ur,
		emailServico:       es,
	}
}

// GerarConvite gera um novo convite para o profissional
func (s *conviteServico) GerarConvite(userID uint, emailPac string) (*dtos.ConviteDTOOut, error) {
	emailValido := false
	var conviteGerado *dominio.Convite
	err := s.db.Transaction(func(tx *gorm.DB) error {
		profissional, err := s.usuarioRepositorio.BuscarProfissionalPorUsuarioID(tx, userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return dominio.ErrUsuarioNaoEncontrado
			}
			return err
		}

		if emailPac != "" {
			usuarioPaciente, err := s.usuarioRepositorio.BuscarPorEmail(emailPac)
			if err != nil {
				emailValido = false
				if errors.Is(err, gorm.ErrRecordNotFound) {

					return dominio.ErrUsuarioNaoEncontrado
				}
				return err
			}
			_, err = s.usuarioRepositorio.BuscarPacientePorUsuarioID(tx, usuarioPaciente.ID)
			if err != nil {
				emailValido = false
				if errors.Is(err, gorm.ErrRecordNotFound) {

					return dominio.ErrUsuarioNaoEncontrado
				}
				return err
			}

			emailValido = true
		}
		// token criptografado aleatoriamente
		tokenBytes := make([]byte, 16)
		if _, err := rand.Read(tokenBytes); err != nil {
			return err
		}
		// de bytes para string
		token := hex.EncodeToString(tokenBytes)

		convite := &dominio.Convite{
			ProfissionalID: profissional.ID,
			Token:          token,
			DataExpiracao:  time.Now().Add(24 * time.Hour),
			Usado:          false,
		}

		// Validar convite antes de criar
		if err := convite.Validar(); err != nil {
			return err
		}

		if err := s.conviteRepositorio.CriarConvite(tx, convite); err != nil {
			return err
		}

		conviteGerado = convite
		return nil

	})

	/*
		TODO:
		- terminar funcionalidade de envio de email de convite
		- testar
	*/
	if emailValido {
		go func() {
			s.emailServico.EnviarEmail([]string{emailPac}, "Convite de Vínculo", "link ativacao")
		}()
	}

	return mappers.ConviteParaDTOOut(conviteGerado), err // err = nil
}

// VincularPaciente vincula um paciente a um profissional usando um token de convite
func (s *conviteServico) VincularPaciente(userID uint, token string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Buscar convite pelo token recebido
		convite, err := s.conviteRepositorio.BuscarConvitePorToken(tx, token)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return dominio.ErrTokenConviteInvalido
			}
			return err
		}
		// Validar se o convite recebido esta valido com os metodos de dominio
		if !convite.EstaValido() {
			if convite.EstaExpirado() {
				return dominio.ErrConviteExpirado
			}
			if convite.JaFoiUtilizado() {
				return dominio.ErrConviteJaUtilizado
			}
		}
		// Busca paciente
		paciente, err := s.usuarioRepositorio.BuscarPacientePorUsuarioID(tx, userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return dominio.ErrUsuarioNaoEncontrado
			}
			return err
		}

		// Vincular paciente ao profissional (insere na tabela de associacao profissionais_pacientes)
		profissionalShell := &dominio.Profissional{ID: convite.ProfissionalID}
		if err := tx.Model(profissionalShell).Association("Pacientes").Append(paciente); err != nil {
			return err
		}

		convite.UtilizarConvite(paciente.ID)

		return s.conviteRepositorio.MarcarConviteComoUsado(tx, convite)
	})
}
