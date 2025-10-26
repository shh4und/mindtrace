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
	ID          uint           `gorm:"primaryKey"`
	TipoUsuario string         `json:"tipo_usuario" gorm:"type:varchar(50);not null"`
	Nome        string         `json:"nome" gorm:"type:varchar(255);not null"`
	Email       string         `json:"email" gorm:"type:varchar(255);unique;not null"`
	Senha       string         `json:"-" gorm:"type:text;not null"`
	Contato     string         `json:"contato,omitempty" gorm:"type:varchar(100)"`
	Bio         string         `json:"bio" gorm:"type:text"`
	CPF         string         `json:"cpf" gorm:"type:varchar(20);unique"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
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
	ID                   uint       `gorm:"primaryKey"`
	UsuarioID            uint       `json:"-" gorm:"unique;not null"`
	Usuario              Usuario    `json:"usuario" gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE"`
	DataNascimento       time.Time  `json:"data_nascimento" gorm:"type:date"`
	Especialidade        string     `json:"especialidade" gorm:"type:varchar(255)"`
	RegistroProfissional string     `json:"registro_profissional" gorm:"type:varchar(12);unique;not null"`
	Pacientes            []Paciente `json:"pacientes" gorm:"many2many:profissional_paciente;constraint:OnDelete:CASCADE;"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

func (Profissional) TableName() string {
	return "profissionais"
}

// Paciente tem seus proprios dados e uma referencia ao Usuario.
type Paciente struct {
	ID                   uint           `json:"id" gorm:"primaryKey"`
	UsuarioID            uint           `json:"-" gorm:"unique;not null"`
	Usuario              Usuario        `json:"usuario" gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE"`
	DataNascimento       time.Time      `json:"data_nascimento" gorm:"not null"`
	Dependente           bool           `json:"dependente" gorm:"not null"`
	NomeResponsavel      string         `json:"nome_responsavel,omitempty" gorm:"type:varchar(255)"`
	ContatoResponsavel   string         `json:"contato_responsavel,omitempty" gorm:"type:varchar(100)"`
	DataInicioTratamento *time.Time     `json:"data_inicio_tratamento"`
	Profissionais        []Profissional `json:"profissionais" gorm:"many2many:profissional_paciente;constraint:OnDelete:CASCADE;"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

func (Paciente) TableName() string {
	return "pacientes"
}
