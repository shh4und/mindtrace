package dominio

import (
	"time"

	"gorm.io/gorm"
)

// Usuario é a base para todos os tipos de usuários.
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

// Profissional tem seus próprios dados e uma referência ao Usuario.
type Profissional struct {
	ID                   uint       `gorm:"primaryKey"`
	UsuarioID            uint       `json:"-" gorm:"unique;not null"`
	Usuario              Usuario    `json:"usuario" gorm:"foreignKey:UsuarioID"`
	DataNascimento       time.Time  `json:"data_nascimento" gorm:"not null"`
	Especialidade        string     `json:"especialidade" gorm:"type:varchar(255)"`
	RegistroProfissional string     `json:"registro_profissional" gorm:"type:varchar(12);unique;not null"`
	Pacientes            []Paciente `json:"pacientes" gorm:"many2many:profissional_paciente;"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

func (Profissional) TableName() string {
	return "profissionais"
}

// Paciente tem seus próprios dados e uma referência ao Usuario.
type Paciente struct {
	ID                   uint           `json:"id" gorm:"primaryKey"`
	UsuarioID            uint           `json:"-" gorm:"unique;not null"`
	Usuario              Usuario        `json:"usuario" gorm:"foreignKey:UsuarioID"`
	DataNascimento       time.Time      `json:"data_nascimento" gorm:"not null"`
	Dependente           bool           `json:"dependente" gorm:"not null"`
	NomeResponsavel      string         `json:"nome_responsavel,omitempty" gorm:"type:varchar(255)"`
	ContatoResponsavel   string         `json:"contato_responsavel,omitempty" gorm:"type:varchar(100)"`
	DataInicioTratamento *time.Time     `json:"data_inicio_tratamento"`
	Profissionais        []Profissional `json:"profissionais" gorm:"many2many:profissional_paciente;"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

func (Paciente) TableName() string {
	return "pacientes"
}
