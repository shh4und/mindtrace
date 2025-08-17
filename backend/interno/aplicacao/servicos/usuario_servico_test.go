package servicos

import (
	"errors"
	"testing"
	"time"

	"mindtrace/backend/interno/dominio"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// fakeRepo implements the UsuarioRepositorio interface for tests.
type fakeRepo struct {
	userToReturn *dominio.Usuario
	errToReturn  error
}

func (r *fakeRepo) CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error {
	// Simulate that the DB assigns an ID
	usuario.ID = 1
	usuario.CreatedAt = time.Now()
	return nil
}

func (r *fakeRepo) CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	profissional.ID = 2
	profissional.CreatedAt = time.Now()
	return nil
}

func (r *fakeRepo) CriarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	paciente.ID = 3
	paciente.CreatedAt = time.Now()
	return nil
}

func (r *fakeRepo) CriarResponsavelLegal(tx *gorm.DB, responsavel *dominio.ResponsavelLegal) error {
	responsavel.ID = 4
	responsavel.CreatedAt = time.Now()
	return nil
}

func (r *fakeRepo) BuscarPorEmail(email string) (*dominio.Usuario, error) {
	if r.errToReturn != nil {
		return nil, r.errToReturn
	}
	if r.userToReturn != nil && r.userToReturn.Email == email {
		return r.userToReturn, nil
	}

	return nil, gorm.ErrRecordNotFound
}

func (r *fakeRepo) BuscarUsuarioPorID(id uint) (*dominio.Usuario, error) {
	if r.errToReturn != nil {
		return nil, r.errToReturn
	}
	if r.userToReturn != nil && r.userToReturn.ID == id {
		return r.userToReturn, nil
	}

	return nil, gorm.ErrRecordNotFound
}

func (r *fakeRepo) Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error {
	if r.errToReturn != nil {
		return r.errToReturn
	}
	r.userToReturn = usuario
	return nil
}

func (r *fakeRepo) BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error) {

	return nil, gorm.ErrRecordNotFound
}

// TestRegistrarProfissional_HappyPath verifies that a new professional is created when the email
// is not already registered.
func TestRegistrarProfissional_HappyPath(t *testing.T) {
	// Use in-memory sqlite to satisfy s.db.Transaction calls; repo is a fake that doesn't use DB.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	// Configurado para não  encontrar usuário

	repo := &fakeRepo{userToReturn: nil}
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

// TestRegistrarProfissional_EmailAlreadyRegistered verifies that the service returns
// ErrEmailJaCadastrado when BuscarPorEmail finds an existing user.
func TestRegistrarProfissional_EmailAlreadyRegistered(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	// Configura o repo para retornar um usuário, simulando que o e-mail já existe

	repo := &fakeRepo{userToReturn: &dominio.Usuario{Email: "bob@example.com"}}
	svc := NovoUsuarioServico(db, repo)

	dto := RegistrarProfissionalDTO{
		Nome:                 "Bob Test",
		Email:                "bob@example.com",
		Senha:                "password",
		Idade:                40,
		Especialidade:        "Neurologia",
		RegistroProfissional: "CRN-98765",
	}

	_, err = svc.RegistrarProfissional(dto)

	if !errors.Is(err, ErrEmailJaCadastrado) {
		t.Fatalf("expected ErrEmailJaCadastrado, got %v", err)
	}
}

func TestRegistrarPaciente_HappyPath(t *testing.T) {
	// Use in-memory sqlite to satisfy s.db.Transaction calls; repo is a fake that doesn't use DB.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	repo := &fakeRepo{userToReturn: nil}
	svc := NovoUsuarioServico(db, repo)

	dto := RegistrarPacienteDTO{
		Nome:         "Patty Test",
		Email:        "patty@example.com",
		Senha:        "s3cr3t234",
		Idade:        16,
		EhDependente: true,
	}

	pat, err := svc.RegistrarPaciente(dto)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if pat == nil {
		t.Fatalf("expected profissional, got nil")
	}
	if pat.Usuario.ID == 0 {
		t.Errorf("expected usuario ID to be set, got 0")
	}

}

func TestRegistrarPaciente_EmailAlreadyRegistered(t *testing.T) {
	// Use in-memory sqlite to satisfy s.db.Transaction calls; repo is a fake that doesn't use DB.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	repo := &fakeRepo{userToReturn: &dominio.Usuario{Email: "patty@example.com"}}
	svc := NovoUsuarioServico(db, repo)

	dto := RegistrarPacienteDTO{
		Nome:         "Patty Test",
		Email:        "patty@example.com",
		Senha:        "s3cr3t234",
		Idade:        16,
		EhDependente: true,
	}

	_, err = svc.RegistrarPaciente(dto)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, ErrEmailJaCadastrado) {
		t.Fatalf("expected ErrEmailJaCadastrado, got %v", err)
	}

}

func TestLogin_HappyPath(t *testing.T) {
	// Use in-memory sqlite to satisfy s.db.Transaction calls; repo is a fake that doesn't use DB.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	senhaQualquer := "senhaSuperSegura123"
	senhaHashed, err := bcrypt.GenerateFromPassword([]byte(senhaQualquer), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("falha ao gerar hash da senha: %v", err)
	}

	email := "dummy@email.com"
	user := &dominio.Usuario{
		ID:    10,
		Email: email,
		Senha: string(senhaHashed),
	}

	repo := &fakeRepo{userToReturn: user}
	svc := NovoUsuarioServico(db, repo)

	t.Setenv("JWT_SECRET", "minha-assinatura-de-teste")
	token, err := svc.Login(email, senhaQualquer)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token == "" {
		t.Fatalf("expected token, got empty string")
	}

}

func TestLogin_InvalidCredentials(t *testing.T) {
	// Use in-memory sqlite to satisfy s.db.Transaction calls; repo is a fake that doesn't use DB.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	senhaQualquer := "senhaSuperSegura123"
	senhaErrada := "senhaErrada"
	senhaHashed, err := bcrypt.GenerateFromPassword([]byte(senhaQualquer), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("falha ao gerar hash da senha: %v", err)
	}

	email := "dummy@email.com"
	user := &dominio.Usuario{
		ID:    10,
		Email: email,
		Senha: string(senhaHashed),
	}
	repo := &fakeRepo{userToReturn: user}
	svc := NovoUsuarioServico(db, repo)

	t.Setenv("JWT_SECRET", "minha-assinatura-de-teste")

	_, err = svc.Login(email, senhaErrada)
	if !errors.Is(err, ErrCrendenciaisInvalidas) {
		t.Fatalf("expected ErrCrendenciaisInvalidas, got %v", err)
	}

}

func TestLogin_UserNotFound(t *testing.T) {
	// Use in-memory sqlite to satisfy s.db.Transaction calls; repo is a fake that doesn't use DB.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	repo := &fakeRepo{userToReturn: nil}
	svc := NovoUsuarioServico(db, repo)

	t.Setenv("JWT_SECRET", "minha-assinatura-de-teste")

	_, err = svc.Login("alice@email.com", "s3cr3t")

	if !errors.Is(err, ErrUsuarioNaoEncontrado) {
		t.Fatalf("expected ErrUsuarioNaoEncontrado, got %v", err)
	}

}

func TestBuscarUsuarioPorID_HappyPath(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	expectedUser := &dominio.Usuario{ID: 1, Nome: "Test User"}
	repo := &fakeRepo{userToReturn: expectedUser}
	svc := NovoUsuarioServico(db, repo)

	user, err := svc.BuscarUsuarioPorID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.ID != expectedUser.ID {
		t.Errorf("expected user ID %d, got %d", expectedUser.ID, user.ID)
	}
}

func TestAtualizarPerfil_HappyPath(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	originalUser := &dominio.Usuario{ID: 1, Nome: "Original Name"}
	repo := &fakeRepo{userToReturn: originalUser}
	svc := NovoUsuarioServico(db, repo)

	dto := AtualizarPerfilDTO{
		Nome:    "New Name",
		Contato: "12345",
		Bio:     "New Bio",
	}

	updatedUser, err := svc.AtualizarPerfil(1, dto)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updatedUser.Nome != dto.Nome {
		t.Errorf("expected name to be '%s', got '%s'", dto.Nome, updatedUser.Nome)
	}
	if updatedUser.Contato != dto.Contato {
		t.Errorf("expected contact to be '%s', got '%s'", dto.Contato, updatedUser.Contato)
	}
	if updatedUser.Bio != dto.Bio {
		t.Errorf("expected bio to be '%s', got '%s'", dto.Bio, updatedUser.Bio)
	}
}

func TestAlterarSenha_HappyPath(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	currentPassword := "oldPassword123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(currentPassword), bcrypt.DefaultCost)
	user := &dominio.Usuario{ID: 1, Senha: string(hashedPassword)}
	repo := &fakeRepo{userToReturn: user}
	svc := NovoUsuarioServico(db, repo)

	dto := AlterarSenhaDTO{
		SenhaAtual:  currentPassword,
		NovaSenha:   "newPassword456",
		NovaSenhaRe: "newPassword456",
	}

	err = svc.AlterarSenha(1, dto)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify the password was changed
	err = bcrypt.CompareHashAndPassword([]byte(repo.userToReturn.Senha), []byte(dto.NovaSenha))
	if err != nil {
		t.Errorf("new password does not match hash: %v", err)
	}
}

func TestAlterarSenha_InvalidCurrentPassword(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	currentPassword := "oldPassword123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(currentPassword), bcrypt.DefaultCost)
	user := &dominio.Usuario{ID: 1, Senha: string(hashedPassword)}
	repo := &fakeRepo{userToReturn: user}
	svc := NovoUsuarioServico(db, repo)

	dto := AlterarSenhaDTO{
		SenhaAtual:  "wrongPassword",
		NovaSenha:   "newPassword456",
		NovaSenhaRe: "newPassword456",
	}

	err = svc.AlterarSenha(1, dto)
	if !errors.Is(err, ErrCrendenciaisInvalidas) {
		t.Fatalf("expected ErrCrendenciaisInvalidas, got %v", err)
	}
}

func TestAlterarSenha_PasswordMismatch(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	repo := &fakeRepo{}
	svc := NovoUsuarioServico(db, repo)

	dto := AlterarSenhaDTO{
		SenhaAtual:  "anyPassword",
		NovaSenha:   "newPassword456",
		NovaSenhaRe: "newPassword789", // Mismatch
	}

	err = svc.AlterarSenha(1, dto)
	if !errors.Is(err, ErrSenhaNaoConfere) {
		t.Fatalf("expected ErrSenhaNaoConfere, got %v", err)
	}
}