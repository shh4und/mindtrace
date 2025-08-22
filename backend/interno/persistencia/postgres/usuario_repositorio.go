package postgres

import (
	"mindtrace/backend/interno/dominio"

	"gorm.io/gorm"
)

type UsuarioRepositorio interface {
	CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error
	CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error
	CriarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error
	CriarResponsavelLegal(tx *gorm.DB, paciente *dominio.ResponsavelLegal) error
	BuscarPorEmail(email string) (*dominio.Usuario, error)
	BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error)
	BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error)
	BuscarUsuarioPorID(id uint) (*dominio.Usuario, error)
	Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error
}

type gormUsuarioRepositorio struct {
	db *gorm.DB
}

func NovoGormUsuarioRepositorio(db *gorm.DB) UsuarioRepositorio {
	return &gormUsuarioRepositorio{db: db}
}

func (r *gormUsuarioRepositorio) CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error {
	return tx.Create(usuario).Error
}

func (r *gormUsuarioRepositorio) CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	return tx.Create(profissional).Error
}

func (r *gormUsuarioRepositorio) CriarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	return tx.Create(paciente).Error
}

func (r *gormUsuarioRepositorio) CriarResponsavelLegal(tx *gorm.DB, responsavel *dominio.ResponsavelLegal) error {
	return tx.Create(responsavel).Error
}

func (r *gormUsuarioRepositorio) BuscarPorEmail(email string) (*dominio.Usuario, error) {
	var usuario dominio.Usuario
	if err := r.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *gormUsuarioRepositorio) BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error) {
	var profissional dominio.Profissional
	if err := tx.First(&profissional, id).Error; err != nil {
		return nil, err
	}
	return &profissional, nil
}

func (r *gormUsuarioRepositorio) BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	var paciente dominio.Paciente
	if err := tx.First(&paciente, id).Error; err != nil {
		return nil, err
	}
	return &paciente, nil
}

func (r *gormUsuarioRepositorio) BuscarUsuarioPorID(id uint) (*dominio.Usuario, error) {
	var usuario dominio.Usuario
	if err := r.db.First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *gormUsuarioRepositorio) Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error {
	return tx.Save(usuario).Error
}
