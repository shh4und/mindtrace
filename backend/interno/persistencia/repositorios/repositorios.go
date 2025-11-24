package repositorios

import (
	"mindtrace/backend/interno/dominio"
	"time"

	"gorm.io/gorm"
)

type ConviteRepositorio interface {
	CriarConvite(tx *gorm.DB, convite *dominio.Convite) error
	BuscarConvitePorToken(tx *gorm.DB, token string) (*dominio.Convite, error)
	MarcarConviteComoUsado(tx *gorm.DB, convite *dominio.Convite) error
}

type RegistroHumorRepositorio interface {
	CriarRegistroHumor(tx *gorm.DB, registro *dominio.RegistroHumor) error
	BuscarPorPacienteEPeriodo(pacienteID uint, inicio, fim time.Time) ([]*dominio.RegistroHumor, error)
	BuscarUltimoRegistroDePaciente(pacienteID uint) (*dominio.RegistroHumor, error)
	BuscarPorNUltimosRegistros(pacienteID uint, numLimite int) ([]*dominio.RegistroHumor, error)
}

type UsuarioRepositorio interface {
	CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error
	CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error
	CriarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error
	BuscarPorEmail(email string) (*dominio.Usuario, error)
	BuscarUsuarioPorID(id uint) (*dominio.Usuario, error)
	BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error)
	BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error)
	BuscarProfissionalPorUsuarioID(tx *gorm.DB, usuarioID uint) (*dominio.Profissional, error)
	BuscarPacientePorUsuarioID(tx *gorm.DB, usuarioID uint) (*dominio.Paciente, error)
	BuscarPacientesDoProfissional(tx *gorm.DB, profissionalID uint) ([]dominio.Paciente, error)
	Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error
	AtualizarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error
	AtualizarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error
	DeletarUsuario(tx *gorm.DB, id uint) error
}

type InstrumentoRepositorio interface {
	BuscarTodosAtivos(tx *gorm.DB) ([]*dominio.Instrumento, error)
}
