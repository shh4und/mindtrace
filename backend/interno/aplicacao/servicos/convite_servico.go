package servicos

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"
	"time"

	"gorm.io/gorm"
)

var (
	ErrConviteExpirado      = errors.New("o convite expirou")
	ErrConviteNaoEncontrado = errors.New("convite não encontrado ou já utilizado")
	ErrPerfilNaoEncontrado  = errors.New("perfil de profissional ou paciente não encontrado")
)

// ConviteServico define os metodos para gerenciamento de convites
type ConviteServico interface {
	GerarConvite(userID uint) (*dominio.Convite, error)
	VincularPaciente(userID uint, token string) error
}

// conviteServico implementa a interface ConviteServico
type conviteServico struct {
	db                 *gorm.DB
	conviteRepositorio repositorios.ConviteRepositorio
	usuarioRepositorio repositorios.UsuarioRepositorio
}

// NovoConviteServico cria uma nova instancia de ConviteServico
func NovoConviteServico(db *gorm.DB, cr repositorios.ConviteRepositorio, ur repositorios.UsuarioRepositorio) ConviteServico {
	return &conviteServico{
		db:                 db,
		conviteRepositorio: cr,
		usuarioRepositorio: ur,
	}
}

// GerarConvite gera um novo convite para o profissional
func (s *conviteServico) GerarConvite(userID uint) (*dominio.Convite, error) {
	profissional, err := s.usuarioRepositorio.BuscarProfissionalPorUsuarioID(s.db, userID)
	if err != nil {
		return nil, ErrPerfilNaoEncontrado
	}

	tokenBytes := make([]byte, 16)
	if _, err := rand.Read(tokenBytes); err != nil {
		return nil, err
	}
	token := hex.EncodeToString(tokenBytes)

	convite := &dominio.Convite{
		ProfissionalID: profissional.ID,
		Token:          token,
		DataExpiracao:  time.Now().Add(24 * time.Hour),
		Usado:          false,
	}

	if err := s.conviteRepositorio.CriarConvite(s.db, convite); err != nil {
		return nil, err
	}

	return convite, nil
}

// VincularPaciente vincula um paciente a um profissional usando um token de convite
func (s *conviteServico) VincularPaciente(userID uint, token string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		convite, err := s.conviteRepositorio.BuscarConvitePorToken(tx, token)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrConviteNaoEncontrado
			}
			return err
		}

		if time.Now().After(convite.DataExpiracao) {
			return ErrConviteExpirado
		}

		paciente, err := s.usuarioRepositorio.BuscarPacientePorUsuarioID(tx, userID)
		if err != nil {
			return ErrPerfilNaoEncontrado
		}

		profissionalShell := &dominio.Profissional{ID: convite.ProfissionalID}

		if err := tx.Model(profissionalShell).Association("Pacientes").Append(paciente); err != nil {
			return err
		}

		return s.conviteRepositorio.MarcarConviteComoUsado(tx, convite.ID, paciente.ID)
	})
}
