package servicos

import (
	"testing"
	"time"

	"mindtrace/backend/interno/dominio"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// fakeUsuarioRepo implements the UsuarioRepositorio interface for tests.
type fakeUsuarioRepo struct {
	userToReturn    *dominio.Usuario
	patientToReturn *dominio.Paciente
	errToReturn     error
}

func (r *fakeUsuarioRepo) CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error {
	usuario.ID = 1
	usuario.CreatedAt = time.Now()
	return nil
}

func (r *fakeUsuarioRepo) CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	profissional.ID = 2
	profissional.CreatedAt = time.Now()
	return nil
}

func (r *fakeUsuarioRepo) CriarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	paciente.ID = 3
	paciente.CreatedAt = time.Now()
	return nil
}

func (r *fakeUsuarioRepo) CriarResponsavelLegal(tx *gorm.DB, responsavel *dominio.ResponsavelLegal) error {
	responsavel.ID = 4
	responsavel.CreatedAt = time.Now()
	return nil
}

func (r *fakeUsuarioRepo) BuscarPorEmail(email string) (*dominio.Usuario, error) {
	if r.errToReturn != nil {
		return nil, r.errToReturn
	}
	if r.userToReturn != nil && r.userToReturn.Email == email {
		return r.userToReturn, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (r *fakeUsuarioRepo) BuscarUsuarioPorID(id uint) (*dominio.Usuario, error) {
	if r.errToReturn != nil {
		return nil, r.errToReturn
	}
	if r.userToReturn != nil && r.userToReturn.ID == id {
		return r.userToReturn, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (r *fakeUsuarioRepo) Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error {
	if r.errToReturn != nil {
		return r.errToReturn
	}
	r.userToReturn = usuario
	return nil
}

func (r *fakeUsuarioRepo) BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error) {
	return nil, gorm.ErrRecordNotFound // Not implemented for this test suite
}

// CORREÇÃO: Implementação do método faltante.
func (r *fakeUsuarioRepo) BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	if r.errToReturn != nil {
		return nil, r.errToReturn
	}
	if r.patientToReturn != nil && r.patientToReturn.UsuarioID == id {
		return r.patientToReturn, nil
	}
	return nil, gorm.ErrRecordNotFound
}

// TestRegistrarProfissional_HappyPath verifies that a new professional is created when the email
// is not already registered.
func TestRegistrarProfissional_HappyPath(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	repo := &fakeUsuarioRepo{userToReturn: nil}
	svc := NovoUsuarioServico(db, repo)

	dto := RegistrarProfissionalDTO{
		Nome:                 "Alice Test",
		Email:                "alice@example.com",
		Senha:                "s3cr3t",
		Idade:                30,
		Especialidade:        "Psiquiatria",
		RegistroProfissional: "CRP-12345",
	}

	prof, err := svc.RegistrarProfissional(dto)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if prof == nil {
		t.Fatalf("expected profissional, got nil")
	}
	if prof.Usuario.ID == 0 {
		t.Errorf("expected usuario ID to be set, got 0")
	}
}
