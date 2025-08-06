package domain

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome  string `json:"nome" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
	Senha string `json:"senha" gorm:"not null"`
}

type Profissional struct {
	gorm.Model
	UsuarioID            uint    `gorm:"unique;not null"`
	Usuario              Usuario `gorm:"foreignKey:UsuarioID"`
	Especialidade        string  `json:"especialidade"`
	RegistroProfissional string  `json:"registro_profissional" gorm:"unique;not null"`
}

type Paciente struct {
	gorm.Model
	UsuarioID    uint         `gorm:"unique;not null"`
	Usuario      Usuario      `gorm:"foreignKey:UsuarioID"`
	Idade        int          `json:"idade" gorm:"not null"`
	Depdendente  bool         `json:"dependente" gorm:"not null"`
	Profissional Profissional `gorm:"foreignKey:ProfissionalID"`
}
