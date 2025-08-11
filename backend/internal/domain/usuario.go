package domain

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
}

type Profissional struct {
	gorm.Model
	UsuarioID            uint    `json:"-"`
	Usuario              Usuario `gorm:"foreignKey:UsuarioID"`
	Specialty            string  `json:"specialty"`
	ProfessionalRegistry string  `json:"professional_registry" gorm:"unique;not null"`
}

type Paciente struct {
	gorm.Model
	UsuarioID    uint         `json:"-"`
	Usuario      Usuario      `gorm:"foreignKey:UsuarioID"`
	Age          int          `json:"age" gorm:"not null"`
	IsDependent  bool         `json:"is_dependent" gorm:"not null"`
	Profissional Profissional `gorm:"foreignKey:ProfissionalID"`
}