package sqlite

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
	// Busca o usuario para determinar o tipo
	usuario, err := r.BuscarUsuarioPorID(id)
	if err != nil {
		return err
	}

	// Deleta registros especificos conforme o tipo
	switch usuario.TipoUsuario {
	case 2: // tipo de usuario 2 = profissional
		// Busca o profissional para limpar associacoes
		profissional, err := r.BuscarProfissionalPorUsuarioID(tx, id)
		if err != nil {
			return err
		}
		// Limpa associacoes many-to-many removendo entradas da tabela profissional_paciente
		if err := tx.Model(profissional).Association("Pacientes").Clear(); err != nil {
			return err
		}
		// Deleta o profissional
		if err := tx.Where("usuario_id = ?", id).Delete(&dominio.Profissional{}).Error; err != nil {
			return err
		}
	case 3: // tipo de usuario 3 = paciente
		// Busca o paciente para limpar associacoes
		paciente, err := r.BuscarPacientePorUsuarioID(tx, id)
		if err != nil {
			return err
		}
		// Limpa associacoes many-to-many removendo entradas da tabela profissional_paciente
		if err := tx.Model(paciente).Association("Profissionais").Clear(); err != nil {
			return err
		}
		// Deleta o paciente
		if err := tx.Where("usuario_id = ?", id).Delete(&dominio.Paciente{}).Error; err != nil {
			return err
		}
	}

	// Deleta o usuario com hard delete ignorando soft delete
	return tx.Unscoped().Delete(&dominio.Usuario{}, id).Error
}
