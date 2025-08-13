package dominio

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome  string `json:"nome" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Senha string `json:"-" gorm:"not null"`
}

type Profissional struct {
	gorm.Model
	UsuarioID            uint    `json:"-"`
	Usuario              Usuario `gorm:"foreignKey:UsuarioID"`
	Especialidade        string  `json:"especialidade"`
	RegistroProfissional string  `json:"registro_profissional" gorm:"unique;not null"`
}

type Paciente struct {
	gorm.Model
	UsuarioID    uint         `json:"-"`
	Usuario      Usuario      `gorm:"foreignKey:UsuarioID"`
	Idade        int          `json:"idade" gorm:"not null"`
	EhDependente bool         `json:"eh_dependente" gorm:"not null"`
	Profissional Profissional `gorm:"foreignKey:ProfissionalID"`
}
