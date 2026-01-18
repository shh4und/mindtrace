package postgres

import (
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"

	"gorm.io/gorm"
)

type gormUsuarioRepositorio struct {
	db *gorm.DB
}

// NovoGormUsuarioRepositorio cria uma nova instancia do repositorio de usuario com GORM
func NovoGormUsuarioRepositorio(db *gorm.DB) repositorios.UsuarioRepositorio {
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

func (r *gormUsuarioRepositorio) BuscarPorEmail(email string) (*dominio.Usuario, error) {
	var usuario dominio.Usuario
	if err := r.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *gormUsuarioRepositorio) BuscarUsuarioPorID(id uint) (*dominio.Usuario, error) {
	var usuario dominio.Usuario
	if err := r.db.First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *gormUsuarioRepositorio) BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error) {
	var profissional dominio.Profissional
	if err := tx.Preload("Usuario").First(&profissional, id).Error; err != nil {
		return nil, err
	}
	return &profissional, nil
}

func (r *gormUsuarioRepositorio) BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	var paciente dominio.Paciente
	if err := tx.Preload("Usuario").First(&paciente, id).Error; err != nil {
		return nil, err
	}
	return &paciente, nil
}

func (r *gormUsuarioRepositorio) BuscarProfissionalPorUsuarioID(tx *gorm.DB, id uint) (*dominio.Profissional, error) {
	var profissional dominio.Profissional
	err := tx.Preload("Usuario").Where("usuario_id = ?", id).First(&profissional).Error
	if err != nil {
		return nil, err
	}
	return &profissional, nil
}

func (r *gormUsuarioRepositorio) BuscarPacientePorUsuarioID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	var paciente dominio.Paciente
	err := tx.Preload("Usuario").Where("usuario_id = ?", id).First(&paciente).Error
	if err != nil {
		return nil, err
	}
	return &paciente, nil
}

func (r *gormUsuarioRepositorio) BuscarPacientesDoProfissional(tx *gorm.DB, profissionalID uint) ([]dominio.Paciente, error) {
	var profissional dominio.Profissional
	err := tx.Preload("Pacientes").Preload("Pacientes.Usuario").Where("id = ?", profissionalID).First(&profissional).Error
	if err != nil {
		return nil, err
	}
	return profissional.Pacientes, nil
}

func (r *gormUsuarioRepositorio) Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error {
	return tx.Save(usuario).Error
}

func (r *gormUsuarioRepositorio) AtualizarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	return tx.Save(profissional).Error
}
func (r *gormUsuarioRepositorio) AtualizarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	return tx.Save(paciente).Error
}

func (r *gormUsuarioRepositorio) DeletarUsuario(tx *gorm.DB, id uint) error {
	// Deleta o usuario com hard delete ignorando soft delete
	return tx.Unscoped().Delete(&dominio.Usuario{}, id).Error
}

func (r *gormUsuarioRepositorio) BuscarUsuarioPorTokenHash(tokenHash string) (*dominio.Usuario, error) {
	var usuario dominio.Usuario
	if err := r.db.Where("email_verif_hash = ?", tokenHash).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}
