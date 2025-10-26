package dominio

import (
	"errors"
	"regexp"
	"time"

	"gorm.io/gorm"
)

var (
	ErrEmailJaCadastrado     = errors.New("e-mail existente")
	ErrCrendenciaisInvalidas = errors.New("credenciais invalidas")
	ErrUsuarioNaoEncontrado  = errors.New("usuario nao encontrado")
	ErrSenhaNaoConfere       = errors.New("a nova senha e a senha de confirmacao nao conferem")
	ErrEmailInvalido         = errors.New("email invalido")
	ErrSenhaFraca            = errors.New("senha deve ter no minimo 8 caracteres")
	ErrNomeVazio             = errors.New("nome nao pode estar vazio")
)

// Usuario e a base para todos os tipos de usuarios.
type Usuario struct {
	ID          uint   `gorm:"primaryKey"`
	TipoUsuario uint8  `gorm:"type:smallint;not null;check:tipo_usuario >= 1"`
	Nome        string `gorm:"type:varchar(255);not null"`
	Email       string `gorm:"type:varchar(255);unique;not null"`
	Senha       string `gorm:"type:text;not null"`
	Contato     string `gorm:"type:varchar(11)"`
	Bio         string `gorm:"type:text"`
	CPF         string `gorm:"type:varchar(11);unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

// Métodos de validação - LÓGICA DE NEGÓCIO
func (u *Usuario) ValidarEmail() error {
	regex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !regex.MatchString(u.Email) {
		return ErrEmailInvalido
	}
	return nil
}

func (u *Usuario) ValidarSenha(senhaPlana string) error {
	if len(senhaPlana) < 8 {
		return ErrSenhaFraca
	}
	return nil
}

func (u *Usuario) ValidarNome() error {
	if u.Nome == "" {
		return ErrNomeVazio
	}
	return nil
}

// Validação completa
func (u *Usuario) Validar() error {
	if err := u.ValidarEmail(); err != nil {
		return err
	}
	if err := u.ValidarNome(); err != nil {
		return err
	}
	return nil
}

// Profissional tem seus proprios dados e uma referencia ao Usuario.
type Profissional struct {
	ID                   uint    `gorm:"primaryKey"`
	UsuarioID            uint    `gorm:"unique;not null"`
	Usuario              Usuario `gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE"`
	DataNascimento       time.Time
	Especialidade        string
	RegistroProfissional string     `gorm:"type:varchar(12);unique;not null"`
	Pacientes            []Paciente `gorm:"many2many:profissional_paciente;constraint:OnDelete:CASCADE;"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func (Profissional) TableName() string {
	return "profissionais"
}

// Paciente tem seus proprios dados e uma referencia ao Usuario.
type Paciente struct {
	ID                   uint    `gorm:"primaryKey"`
	UsuarioID            uint    `gorm:"unique;not null"`
	Usuario              Usuario `gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE"`
	DataNascimento       time.Time
	Dependente           bool
	NomeResponsavel      string `gorm:"type:varchar(255)"`
	ContatoResponsavel   string `gorm:"type:varchar(11)"`
	DataInicioTratamento *time.Time
	Profissionais        []Profissional `gorm:"many2many:profissional_paciente;constraint:OnDelete:CASCADE;"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func (Paciente) TableName() string {
	return "pacientes"
}
