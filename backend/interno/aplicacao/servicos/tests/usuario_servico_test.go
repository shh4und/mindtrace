package tests

import (
	"errors"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ========== Helper para criar DB de teste ==========

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	return db
}

// ========== Mock do Repositorio ==========

type MockUsuarioRepositorio struct {
	mock.Mock
}

func (m *MockUsuarioRepositorio) BuscarPorEmail(email string) (*dominio.Usuario, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Usuario), args.Error(1)
}

func (m *MockUsuarioRepositorio) CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error {
	args := m.Called(tx, usuario)
	return args.Error(0)
}

func (m *MockUsuarioRepositorio) CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	args := m.Called(tx, profissional)
	return args.Error(0)
}

func (m *MockUsuarioRepositorio) CriarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	args := m.Called(tx, paciente)
	return args.Error(0)
}

func (m *MockUsuarioRepositorio) BuscarUsuarioPorID(userID uint) (*dominio.Usuario, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Usuario), args.Error(1)
}

func (m *MockUsuarioRepositorio) BuscarPacientePorUsuarioID(tx *gorm.DB, userID uint) (*dominio.Paciente, error) {
	args := m.Called(tx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Paciente), args.Error(1)
}

func (m *MockUsuarioRepositorio) BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error) {
	args := m.Called(tx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Profissional), args.Error(1)
}

func (m *MockUsuarioRepositorio) BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	args := m.Called(tx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Paciente), args.Error(1)
}

func (m *MockUsuarioRepositorio) BuscarProfissionalPorUsuarioID(tx *gorm.DB, userID uint) (*dominio.Profissional, error) {
	args := m.Called(tx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Profissional), args.Error(1)
}

func (m *MockUsuarioRepositorio) Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error {
	args := m.Called(tx, usuario)
	return args.Error(0)
}

func (m *MockUsuarioRepositorio) AtualizarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	args := m.Called(tx, profissional)
	return args.Error(0)
}

func (m *MockUsuarioRepositorio) AtualizarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	args := m.Called(tx, paciente)
	return args.Error(0)
}

func (m *MockUsuarioRepositorio) BuscarPacientesDoProfissional(tx *gorm.DB, profissionalID uint) ([]dominio.Paciente, error) {
	args := m.Called(tx, profissionalID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dominio.Paciente), args.Error(1)
}

func (m *MockUsuarioRepositorio) DeletarUsuario(tx *gorm.DB, userID uint) error {
	args := m.Called(tx, userID)
	return args.Error(0)
}

// ========== Testes para RegistrarProfissional ==========

func TestUsuarioServico_RegistrarProfissional_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dtoIn := &dtos.RegistrarProfissionalDTOIn{
		Nome:                 "Dr. João Silva",
		Email:                "joao@example.com",
		Senha:                "Senha1234",
		DataNascimento:       time.Now().AddDate(-30, 0, 0),
		Especialidade:        "Psicologia Clínica",
		RegistroProfissional: "CRP12345",
	}

	// Mock: email nao cadastrado
	mockRepo.On("BuscarPorEmail", dtoIn.Email).Return(nil, gorm.ErrRecordNotFound)
	mockRepo.On("CriarUsuario", mock.Anything, mock.AnythingOfType("*dominio.Usuario")).Run(func(args mock.Arguments) {
		// Simula o banco atribuindo um ID
		user := args.Get(1).(*dominio.Usuario)
		user.ID = 1
	}).Return(nil)
	mockRepo.On("CriarProfissional", mock.Anything, mock.AnythingOfType("*dominio.Profissional")).Run(func(args mock.Arguments) {
		// Simula o banco atribuindo um ID
		prof := args.Get(1).(*dominio.Profissional)
		prof.ID = 1
	}).Return(nil)

	result, err := servico.RegistrarProfissional(dtoIn)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dtoIn.Nome, result.Usuario.Nome)
	assert.Equal(t, dtoIn.Email, result.Usuario.Email)
	assert.Equal(t, dtoIn.Especialidade, result.Especialidade)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_RegistrarProfissional_EmailJaCadastrado(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dtoIn := &dtos.RegistrarProfissionalDTOIn{
		Nome:                 "Dr. João Silva",
		Email:                "joao@example.com",
		Senha:                "Senha123!",
		DataNascimento:       time.Now().AddDate(-30, 0, 0),
		Especialidade:        "Psicologia Clínica",
		RegistroProfissional: "CRP12345",
	}

	usuarioExistente := &dominio.Usuario{ID: 1, Email: dtoIn.Email}
	mockRepo.On("BuscarPorEmail", dtoIn.Email).Return(usuarioExistente, nil)

	result, err := servico.RegistrarProfissional(dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrEmailJaCadastrado, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_RegistrarProfissional_EmailInvalido(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dtoIn := &dtos.RegistrarProfissionalDTOIn{
		Nome:                 "Dr. João Silva",
		Email:                "email-invalido",
		Senha:                "Senha123!",
		DataNascimento:       time.Now().AddDate(-30, 0, 0),
		Especialidade:        "Psicologia Clínica",
		RegistroProfissional: "CRP12345",
	}

	mockRepo.On("BuscarPorEmail", dtoIn.Email).Return(nil, gorm.ErrRecordNotFound)

	result, err := servico.RegistrarProfissional(dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrEmailInvalido, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_RegistrarProfissional_SenhaFraca(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dtoIn := &dtos.RegistrarProfissionalDTOIn{
		Nome:                 "Dr. João Silva",
		Email:                "joao@example.com",
		Senha:                "123",
		DataNascimento:       time.Now().AddDate(-30, 0, 0),
		Especialidade:        "Psicologia Clínica",
		RegistroProfissional: "CRP12345",
	}

	mockRepo.On("BuscarPorEmail", dtoIn.Email).Return(nil, gorm.ErrRecordNotFound)

	result, err := servico.RegistrarProfissional(dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrSenhaFraca, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_RegistrarProfissional_MenorDeIdade(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dtoIn := &dtos.RegistrarProfissionalDTOIn{
		Nome:                 "Dr. João Silva",
		Email:                "joao@example.com",
		Senha:                "Senha123!",
		DataNascimento:       time.Now().AddDate(-17, 0, 0),
		Especialidade:        "Psicologia Clínica",
		RegistroProfissional: "CRP12345",
	}

	mockRepo.On("BuscarPorEmail", dtoIn.Email).Return(nil, gorm.ErrRecordNotFound)

	result, err := servico.RegistrarProfissional(dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrProfissionalMenorDeIdade, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ========== Testes para RegistrarPaciente ==========

func TestUsuarioServico_RegistrarPaciente_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dependente := false
	dtoIn := &dtos.RegistrarPacienteDTOIn{
		Nome:           "Maria Silva",
		Email:          "maria@example.com",
		Senha:          "Senha123!",
		DataNascimento: time.Now().AddDate(-25, 0, 0),
		Dependente:     &dependente,
	}

	mockRepo.On("BuscarPorEmail", dtoIn.Email).Return(nil, gorm.ErrRecordNotFound)
	mockRepo.On("CriarUsuario", mock.Anything, mock.Anything).Return(nil)
	mockRepo.On("CriarPaciente", mock.Anything, mock.Anything).Return(nil)

	result, err := servico.RegistrarPaciente(dtoIn)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dtoIn.Nome, result.Usuario.Nome)
	assert.Equal(t, dtoIn.Email, result.Usuario.Email)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_RegistrarPaciente_Dependente_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dependente := true
	dtoIn := &dtos.RegistrarPacienteDTOIn{
		Nome:               "Pedro Silva",
		Email:              "pedro@example.com",
		Senha:              "Senha123!",
		DataNascimento:     time.Now().AddDate(-10, 0, 0),
		Dependente:         &dependente,
		NomeResponsavel:    "José Silva",
		ContatoResponsavel: "11987654321",
	}

	mockRepo.On("BuscarPorEmail", dtoIn.Email).Return(nil, gorm.ErrRecordNotFound)
	mockRepo.On("CriarUsuario", mock.Anything, mock.Anything).Return(nil)
	mockRepo.On("CriarPaciente", mock.Anything, mock.Anything).Return(nil)

	result, err := servico.RegistrarPaciente(dtoIn)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dtoIn.Nome, result.Usuario.Nome)
	assert.Equal(t, dtoIn.Email, result.Usuario.Email)
	assert.NotNil(t, result.Dependente)
	assert.True(t, *result.Dependente)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_RegistrarPaciente_EmailJaCadastrado(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dtoIn := &dtos.RegistrarPacienteDTOIn{
		Nome:           "Maria Silva",
		Email:          "maria@example.com",
		Senha:          "Senha123!",
		DataNascimento: time.Now().AddDate(-25, 0, 0),
	}

	usuarioExistente := &dominio.Usuario{ID: 1, Email: dtoIn.Email}
	mockRepo.On("BuscarPorEmail", dtoIn.Email).Return(usuarioExistente, nil)

	result, err := servico.RegistrarPaciente(dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrEmailJaCadastrado, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_RegistrarPaciente_DependenteSemResponsavel(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dependente := true
	dtoIn := &dtos.RegistrarPacienteDTOIn{
		Nome:           "Pedro Silva",
		Email:          "pedro@example.com",
		Senha:          "Senha123!",
		DataNascimento: time.Now().AddDate(-10, 0, 0),
		Dependente:     &dependente,
	}

	mockRepo.On("BuscarPorEmail", dtoIn.Email).Return(nil, gorm.ErrRecordNotFound)

	result, err := servico.RegistrarPaciente(dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrResponsavelVazio, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ========== Testes para Login ==========

func TestUsuarioServico_Login_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	senha := "Senha123!"
	hashSenha, _ := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

	usuario := &dominio.Usuario{
		ID:          1,
		Email:       "joao@example.com",
		Senha:       string(hashSenha),
		TipoUsuario: 2,
	}

	mockRepo.On("BuscarPorEmail", usuario.Email).Return(usuario, nil)

	token, err := servico.Login(usuario.Email, senha)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_Login_UsuarioNaoEncontrado(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	mockRepo.On("BuscarPorEmail", "invalido@example.com").Return(nil, gorm.ErrRecordNotFound)

	token, err := servico.Login("invalido@example.com", "Senha123!")

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrUsuarioNaoEncontrado, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_Login_SenhaInvalida(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	senhaCorreta := "Senha123!"
	hashSenha, _ := bcrypt.GenerateFromPassword([]byte(senhaCorreta), bcrypt.DefaultCost)

	usuario := &dominio.Usuario{
		ID:          1,
		Email:       "joao@example.com",
		Senha:       string(hashSenha),
		TipoUsuario: 2,
	}

	mockRepo.On("BuscarPorEmail", usuario.Email).Return(usuario, nil)

	token, err := servico.Login(usuario.Email, "SenhaErrada!")

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrCrendenciaisInvalidas, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

// ========== Testes para BuscarUsuarioPorID ==========

func TestUsuarioServico_BuscarUsuarioPorID_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	usuario := &dominio.Usuario{
		ID:    1,
		Nome:  "João Silva",
		Email: "joao@example.com",
	}

	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)

	result, err := servico.BuscarUsuarioPorID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, usuario.Nome, result.Nome)
	assert.Equal(t, usuario.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_BuscarUsuarioPorID_NaoEncontrado(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	mockRepo.On("BuscarUsuarioPorID", uint(999)).Return(nil, gorm.ErrRecordNotFound)

	result, err := servico.BuscarUsuarioPorID(999)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrUsuarioNaoEncontrado, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ========== Testes para ProprioPerfilPaciente ==========

func TestUsuarioServico_ProprioPerfilPaciente_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	paciente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 1,
		Usuario: dominio.Usuario{
			ID:    1,
			Nome:  "Maria Silva",
			Email: "maria@example.com",
		},
		DataNascimento: time.Now().AddDate(-25, 0, 0),
	}

	mockRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(1)).Return(paciente, nil)

	result, err := servico.ProprioPerfilPaciente(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, paciente.Usuario.Nome, result.Usuario.Nome)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_ProprioPerfilPaciente_NaoEncontrado(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	mockRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(999)).Return(nil, gorm.ErrRecordNotFound)

	result, err := servico.ProprioPerfilPaciente(999)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrUsuarioNaoEncontrado, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ========== Testes para ProprioPerfilProfissional ==========

func TestUsuarioServico_ProprioPerfilProfissional_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: 1,
		Usuario: dominio.Usuario{
			ID:    1,
			Nome:  "Dr. João Silva",
			Email: "joao@example.com",
		},
		Especialidade:        "Psicologia Clínica",
		RegistroProfissional: "CRP12345",
	}

	mockRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(1)).Return(profissional, nil)

	result, err := servico.ProprioPerfilProfissional(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, profissional.Usuario.Nome, result.Usuario.Nome)
	assert.Equal(t, profissional.Especialidade, result.Especialidade)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_ProprioPerfilProfissional_NaoEncontrado(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	mockRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(999)).Return(nil, gorm.ErrRecordNotFound)

	result, err := servico.ProprioPerfilProfissional(999)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrUsuarioNaoEncontrado, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ========== Testes para AtualizarPerfil ==========

func TestUsuarioServico_AtualizarPerfil_UsuarioSimples_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	usuario := &dominio.Usuario{
		ID:          1,
		TipoUsuario: 1,
		Nome:        "João Silva",
		Email:       "joao@example.com",
	}

	dtoIn := &dtos.AtualizarPerfilDTOIn{
		Nome:    "João Silva Updated",
		Contato: "11987654321",
		Bio:     "Nova bio",
	}

	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)
	mockRepo.On("Atualizar", mock.Anything, mock.Anything).Return(nil)

	err := servico.AtualizarPerfil(1, dtoIn)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_AtualizarPerfil_Profissional_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	usuario := &dominio.Usuario{
		ID:          1,
		TipoUsuario: 2,
		Nome:        "Dr. João Silva",
		Email:       "joao@example.com",
	}

	profissional := &dominio.Profissional{
		ID:                   1,
		UsuarioID:            1,
		Especialidade:        "Psicologia Clínica",
		RegistroProfissional: "CRP12345",
		DataNascimento:       time.Now().AddDate(-30, 0, 0),
	}

	dtoIn := &dtos.AtualizarPerfilDTOIn{
		Nome:                 "Dr. João Silva Updated",
		Especialidade:        "Neuropsicologia",
		RegistroProfissional: "CRP54321",
	}

	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)
	mockRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(1)).Return(profissional, nil)
	mockRepo.On("Atualizar", mock.Anything, mock.Anything).Return(nil)
	mockRepo.On("AtualizarProfissional", mock.Anything, mock.Anything).Return(nil)

	err := servico.AtualizarPerfil(1, dtoIn)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_AtualizarPerfil_Paciente_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	usuario := &dominio.Usuario{
		ID:          1,
		TipoUsuario: 3,
		Nome:        "Maria Silva",
		Email:       "maria@example.com",
	}

	dataNascimento := time.Now().AddDate(-25, 0, 0)
	paciente := &dominio.Paciente{
		ID:             1,
		UsuarioID:      1,
		DataNascimento: dataNascimento,
		Dependente:     false,
	}

	novaDataNascimento := time.Now().AddDate(-26, 0, 0)
	dtoIn := &dtos.AtualizarPerfilDTOIn{
		Nome:           "Maria Silva Updated",
		DataNascimento: &novaDataNascimento,
	}

	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)
	mockRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(1)).Return(paciente, nil)
	mockRepo.On("Atualizar", mock.Anything, mock.Anything).Return(nil)
	mockRepo.On("AtualizarPaciente", mock.Anything, mock.Anything).Return(nil)

	err := servico.AtualizarPerfil(1, dtoIn)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_AtualizarPerfil_NomeVazio_Erro(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	usuario := &dominio.Usuario{
		ID:          1,
		TipoUsuario: 1,
		Nome:        "João Silva",
		Email:       "joao@example.com",
	}

	dtoIn := &dtos.AtualizarPerfilDTOIn{
		Nome: "",
	}

	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)

	err := servico.AtualizarPerfil(1, dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrNomeVazio, err)
	mockRepo.AssertExpectations(t)
}

// ========== Testes para AlterarSenha ==========

func TestUsuarioServico_AlterarSenha_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	senhaAtual := "Senha123!"
	hashSenha, _ := bcrypt.GenerateFromPassword([]byte(senhaAtual), bcrypt.DefaultCost)

	usuario := &dominio.Usuario{
		ID:    1,
		Email: "joao@example.com",
		Senha: string(hashSenha),
	}

	dtoIn := &dtos.AlterarSenhaDTOIn{
		SenhaAtual:  senhaAtual,
		NovaSenha:   "NovaSenha123!",
		NovaSenhaRe: "NovaSenha123!",
	}

	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)
	mockRepo.On("Atualizar", mock.Anything, mock.Anything).Return(nil)

	err := servico.AlterarSenha(1, dtoIn)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_AlterarSenha_SenhasNaoConferem(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	dtoIn := &dtos.AlterarSenhaDTOIn{
		SenhaAtual:  "Senha123!",
		NovaSenha:   "NovaSenha123!",
		NovaSenhaRe: "SenhaDiferente123!",
	}

	err := servico.AlterarSenha(1, dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrSenhaNaoConfere, err)
}

func TestUsuarioServico_AlterarSenha_SenhaAtualInvalida(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	senhaAtual := "Senha123!"
	hashSenha, _ := bcrypt.GenerateFromPassword([]byte(senhaAtual), bcrypt.DefaultCost)

	usuario := &dominio.Usuario{
		ID:    1,
		Email: "joao@example.com",
		Senha: string(hashSenha),
	}

	dtoIn := &dtos.AlterarSenhaDTOIn{
		SenhaAtual:  "SenhaErrada!",
		NovaSenha:   "NovaSenha123!",
		NovaSenhaRe: "NovaSenha123!",
	}

	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)

	err := servico.AlterarSenha(1, dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrCrendenciaisInvalidas, err)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_AlterarSenha_NovaSenhaFraca(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	senhaAtual := "Senha123!"
	hashSenha, _ := bcrypt.GenerateFromPassword([]byte(senhaAtual), bcrypt.DefaultCost)

	usuario := &dominio.Usuario{
		ID:    1,
		Email: "joao@example.com",
		Senha: string(hashSenha),
	}

	dtoIn := &dtos.AlterarSenhaDTOIn{
		SenhaAtual:  senhaAtual,
		NovaSenha:   "123",
		NovaSenhaRe: "123",
	}

	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)

	err := servico.AlterarSenha(1, dtoIn)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrSenhaFraca, err)
	mockRepo.AssertExpectations(t)
}

// ========== Testes para ListarPacientesDoProfissional ==========

func TestUsuarioServico_ListarPacientesDoProfissional_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: 1,
	}

	pacientes := []dominio.Paciente{
		{
			ID:        1,
			UsuarioID: 2,
			Usuario: dominio.Usuario{
				ID:    2,
				Nome:  "Paciente 1",
				Email: "paciente1@example.com",
			},
		},
		{
			ID:        2,
			UsuarioID: 3,
			Usuario: dominio.Usuario{
				ID:    3,
				Nome:  "Paciente 2",
				Email: "paciente2@example.com",
			},
		},
	}

	mockRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(1)).Return(profissional, nil)
	mockRepo.On("BuscarPacientesDoProfissional", mock.Anything, uint(1)).Return(pacientes, nil)

	result, err := servico.ListarPacientesDoProfissional(1)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Paciente 1", result[0].Usuario.Nome)
	assert.Equal(t, "Paciente 2", result[1].Usuario.Nome)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_ListarPacientesDoProfissional_ProfissionalNaoEncontrado(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	mockRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(999)).Return(nil, gorm.ErrRecordNotFound)

	result, err := servico.ListarPacientesDoProfissional(999)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrUsuarioNaoEncontrado, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
	mockRepo.AssertExpectations(t)
}

// ========== Testes para DeletarPerfil ==========

func TestUsuarioServico_DeletarPerfil_Sucesso(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	usuario := &dominio.Usuario{
		ID:    1,
		Nome:  "João Silva",
		Email: "joao@example.com",
	}

	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)
	mockRepo.On("DeletarUsuario", mock.Anything, uint(1)).Return(nil)

	err := servico.DeletarPerfil(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_DeletarPerfil_UsuarioNaoEncontrado(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	mockRepo.On("BuscarUsuarioPorID", uint(999)).Return(nil, gorm.ErrRecordNotFound)

	err := servico.DeletarPerfil(999)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrUsuarioNaoEncontrado, err)
	mockRepo.AssertExpectations(t)
}

func TestUsuarioServico_DeletarPerfil_ErroAoDeletar(t *testing.T) {
	mockRepo := new(MockUsuarioRepositorio)
	db := setupTestDB(t)
	servico := servicos.NovoUsuarioServico(db, mockRepo)

	usuario := &dominio.Usuario{
		ID:    1,
		Nome:  "João Silva",
		Email: "joao@example.com",
	}

	erroDB := errors.New("erro ao deletar")
	mockRepo.On("BuscarUsuarioPorID", uint(1)).Return(usuario, nil)
	mockRepo.On("DeletarUsuario", mock.Anything, uint(1)).Return(erroDB)

	err := servico.DeletarPerfil(1)

	assert.Error(t, err)
	assert.Equal(t, erroDB, err)
	mockRepo.AssertExpectations(t)
}
