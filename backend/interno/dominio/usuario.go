package dominio

import (
	"time"

	"gorm.io/gorm"
)

// Usuario é a base para todos os tipos de usuários.
type Usuario struct {
	ID        uint           `gorm:"primaryKey"`
	Nome      string         `json:"nome" gorm:"type:varchar(255);not null"`
	Email     string         `json:"email" gorm:"type:varchar(255);unique;not null"`
	Senha     string         `json:"-" gorm:"type:text;not null"`
	Contato   string         `json:"contato" gorm:"not null"`
	Bio       string         `json:"bio" gorm:"type:text"`
	CPF       string         `json:"cpf" gorm:"type:varchar(20);unique"`
	CNPJ      string         `json:"cnpj" gorm:"type:varchar(20);unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

// Profissional tem seus próprios dados e uma referência ao Usuario.
type Profissional struct {
	ID                   uint       `gorm:"primaryKey"`
	UsuarioID            uint       `json:"-" gorm:"unique;not null"`
	Usuario              Usuario    `json:"usuario" gorm:"foreignKey:UsuarioID"`
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
	ID                   uint           `gorm:"primaryKey"`
	UsuarioID            uint           `json:"-" gorm:"unique;not null"`
	Usuario              Usuario        `json:"usuario" gorm:"foreignKey:UsuarioID"`
	Idade                int8           `json:"idade" gorm:"not null"`
	EhDependente         bool           `json:"eh_dependente" gorm:"not null"`
	DataInicioTratamento *time.Time     `json:"data_inicio_tratamento"`
	HistoricoSaude       string         `json:"historico_saude" gorm:"type:text"`
	Profissionais        []Profissional `json:"profissionais" gorm:"many2many:profissional_paciente;"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

func (Paciente) TableName() string {
	return "pacientes"
}

// ResponsavelLegal representa a entidade de um responsável por um paciente.
type ResponsavelLegal struct {
	ID               uint      `gorm:"primaryKey"`
	UsuarioID        uint      `json:"-" gorm:"unique;not null"`
	Usuario          Usuario   `json:"usuario" gorm:"foreignKey:UsuarioID"`
	PacienteID       uint      `json:"-" gorm:"unique;not null"`
	Paciente         Paciente  `json:"paciente" gorm:"foreignKey:PacienteID"`
	Parentesco       string    `json:"parentesco" gorm:"type:varchar(100)"`
	ContatoPrincipal string    `json:"contato_principal" gorm:"type:varchar(100)"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (ResponsavelLegal) TableName() string {
	return "responsaveis_legais"
}
