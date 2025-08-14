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
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Profissional tem seus próprios dados e uma referência ao Usuario.
type Profissional struct {
	ID                   uint       `gorm:"primaryKey"`
	UsuarioID            uint       `json:"-" gorm:"unique;not null"`
	Usuario              Usuario    `json:"usuario" gorm:"foreignKey:UsuarioID"`
	Especialidade        string     `json:"especialidade" gorm:"type:varchar(255)"`
	RegistroProfissional string     `json:"registro_profissional" gorm:"type:varchar(100);unique;not null"`
	Pacientes            []Paciente `json:"pacientes" gorm:"many2many:profissional_paciente;"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

// Paciente tem seus próprios dados e uma referência ao Usuario.
type Paciente struct {
	ID                   uint           `gorm:"primaryKey"`
	UsuarioID            uint           `json:"-" gorm:"unique;not null"`
	Usuario              Usuario        `json:"usuario" gorm:"foreignKey:UsuarioID"`
	DataInicioTratamento *time.Time     `json:"data_inicio_tratamento"`
	HistoricoSaude       string         `json:"historico_saude" gorm:"type:text"`
	Profissionais        []Profissional `json:"profissionais" gorm:"many2many:profissional_paciente;"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
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

// ProfissionalPaciente é a tabela de junção para o relacionamento N-N.
type ProfissionalPaciente struct {
	ProfissionalID uint `gorm:"primaryKey"`
	PacienteID     uint `gorm:"primaryKey"`
}
