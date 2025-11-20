package tests

import (
	"errors"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// ========== Mocks ==========

// MockConviteRepositorio simula o repositorio de convites
type MockConviteRepositorio struct {
	mock.Mock
}

func (m *MockConviteRepositorio) CriarConvite(tx *gorm.DB, convite *dominio.Convite) error {
	args := m.Called(tx, convite)
	return args.Error(0)
}

func (m *MockConviteRepositorio) BuscarConvitePorToken(tx *gorm.DB, token string) (*dominio.Convite, error) {
	args := m.Called(tx, token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Convite), args.Error(1)
}

func (m *MockConviteRepositorio) MarcarConviteComoUsado(tx *gorm.DB, convite *dominio.Convite) error {
	args := m.Called(tx, convite)
	return args.Error(0)
}

func (m *MockConviteRepositorio) BuscarConvitesAtivos(profissionalID uint) ([]*dominio.Convite, error) {
	args := m.Called(profissionalID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*dominio.Convite), args.Error(1)
}

// MockUsuarioRepositorioConvite simula o repositorio de usuarios para testes de convite
type MockUsuarioRepositorioConvite struct {
	mock.Mock
}

func (m *MockUsuarioRepositorioConvite) CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error {
	return nil
}

func (m *MockUsuarioRepositorioConvite) CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	return nil
}

func (m *MockUsuarioRepositorioConvite) CriarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	return nil
}

func (m *MockUsuarioRepositorioConvite) BuscarPorEmail(email string) (*dominio.Usuario, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioConvite) BuscarUsuarioPorID(id uint) (*dominio.Usuario, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioConvite) BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioConvite) BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioConvite) BuscarProfissionalPorUsuarioID(tx *gorm.DB, usuarioID uint) (*dominio.Profissional, error) {
	args := m.Called(tx, usuarioID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Profissional), args.Error(1)
}

func (m *MockUsuarioRepositorioConvite) BuscarPacientePorUsuarioID(tx *gorm.DB, usuarioID uint) (*dominio.Paciente, error) {
	args := m.Called(tx, usuarioID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Paciente), args.Error(1)
}

func (m *MockUsuarioRepositorioConvite) BuscarPacientesDoProfissional(tx *gorm.DB, profissionalID uint) ([]dominio.Paciente, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioConvite) Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error {
	return nil
}

func (m *MockUsuarioRepositorioConvite) AtualizarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	return nil
}

func (m *MockUsuarioRepositorioConvite) AtualizarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	return nil
}

func (m *MockUsuarioRepositorioConvite) DeletarUsuario(tx *gorm.DB, id uint) error {
	return nil
}

// ========== Helper Functions ==========

func setupTestDBConvite(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Falha ao abrir banco de dados de teste: %v", err)
	}

	err = db.AutoMigrate(&dominio.Usuario{}, &dominio.Profissional{}, &dominio.Paciente{}, &dominio.Convite{})
	if err != nil {
		t.Fatalf("Falha ao migrar esquema: %v", err)
	}

	return db
}

// ========== Testes do Serviço ==========

func TestConviteServico_GerarConvite_Sucesso(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	profissionalExistente := &dominio.Profissional{
		ID:        1,
		UsuarioID: 10,
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(10)).Return(profissionalExistente, nil)
	mockConviteRepo.On("CriarConvite", mock.Anything, mock.AnythingOfType("*dominio.Convite")).Return(nil)

	resultado, err := servico.GerarConvite(10)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.NotEmpty(t, resultado.Token)
	assert.False(t, resultado.Usado)
	assert.NotZero(t, resultado.DataExpiracao)
	mockUsuarioRepo.AssertExpectations(t)
	mockConviteRepo.AssertExpectations(t)
}

func TestConviteServico_GerarConvite_ProfissionalNaoEncontrado(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(999)).Return(nil, gorm.ErrRecordNotFound)

	resultado, err := servico.GerarConvite(999)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrUsuarioNaoEncontrado, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockConviteRepo.AssertNotCalled(t, "CriarConvite")
}

func TestConviteServico_GerarConvite_ErroAoBuscarProfissional(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	erroGenerico := errors.New("erro de conexão com banco de dados")
	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(10)).Return(nil, erroGenerico)

	resultado, err := servico.GerarConvite(10)

	assert.Error(t, err)
	assert.Equal(t, erroGenerico, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockConviteRepo.AssertNotCalled(t, "CriarConvite")
}

func TestConviteServico_GerarConvite_ErroAoCriarConvite(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	profissionalExistente := &dominio.Profissional{
		ID:        1,
		UsuarioID: 10,
	}

	erroGenerico := errors.New("erro ao inserir no banco de dados")
	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(10)).Return(profissionalExistente, nil)
	mockConviteRepo.On("CriarConvite", mock.Anything, mock.AnythingOfType("*dominio.Convite")).Return(erroGenerico)

	resultado, err := servico.GerarConvite(10)

	assert.Error(t, err)
	assert.Equal(t, erroGenerico, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockConviteRepo.AssertExpectations(t)
}

func TestConviteServico_GerarConvite_TokenAleatorio(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	profissionalExistente := &dominio.Profissional{
		ID:        1,
		UsuarioID: 10,
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mock.Anything, uint(10)).Return(profissionalExistente, nil)
	mockConviteRepo.On("CriarConvite", mock.Anything, mock.AnythingOfType("*dominio.Convite")).Return(nil).Times(3)

	// Gerar 3 convites e verificar que os tokens são diferentes
	tokens := make(map[string]bool)

	for i := 0; i < 3; i++ {
		resultado, err := servico.GerarConvite(10)
		assert.NoError(t, err)
		assert.NotNil(t, resultado)
		assert.NotEmpty(t, resultado.Token)

		// Verifica que o token tem pelo menos 10 caracteres (16 bytes em hex = 32 caracteres)
		assert.GreaterOrEqual(t, len(resultado.Token), 10)

		tokens[resultado.Token] = true
	}

	// Verifica que os 3 tokens são diferentes
	assert.Equal(t, 3, len(tokens))
}

func TestConviteServico_VincularPaciente_Sucesso(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	// Criar tabela de associação many-to-many
	db.Exec(`CREATE TABLE IF NOT EXISTS profissionais_pacientes (
		profissional_id INTEGER,
		paciente_id INTEGER,
		PRIMARY KEY (profissional_id, paciente_id)
	)`)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	conviteValido := &dominio.Convite{
		ID:             1,
		ProfissionalID: 1,
		Token:          "abc123def456",
		DataExpiracao:  time.Now().Add(24 * time.Hour),
		Usado:          false,
	}

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 20,
	}

	mockConviteRepo.On("BuscarConvitePorToken", mock.Anything, "abc123def456").Return(conviteValido, nil)
	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(20)).Return(pacienteExistente, nil)
	mockConviteRepo.On("MarcarConviteComoUsado", mock.Anything, conviteValido).Return(nil)

	err := servico.VincularPaciente(20, "abc123def456")

	assert.NoError(t, err)
	mockConviteRepo.AssertExpectations(t)
	mockUsuarioRepo.AssertExpectations(t)
}

func TestConviteServico_VincularPaciente_TokenNaoEncontrado(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	mockConviteRepo.On("BuscarConvitePorToken", mock.Anything, "token-invalido").Return(nil, gorm.ErrRecordNotFound)

	err := servico.VincularPaciente(20, "token-invalido")

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrTokenConviteInvalido, err)
	mockConviteRepo.AssertExpectations(t)
	mockUsuarioRepo.AssertNotCalled(t, "BuscarPacientePorUsuarioID")
}

func TestConviteServico_VincularPaciente_ConviteExpirado(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	conviteExpirado := &dominio.Convite{
		ID:             1,
		ProfissionalID: 1,
		Token:          "abc123def456",
		DataExpiracao:  time.Now().Add(-24 * time.Hour), // Expirado
		Usado:          false,
	}

	mockConviteRepo.On("BuscarConvitePorToken", mock.Anything, "abc123def456").Return(conviteExpirado, nil)

	err := servico.VincularPaciente(20, "abc123def456")

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrConviteExpirado, err)
	mockConviteRepo.AssertExpectations(t)
	mockUsuarioRepo.AssertNotCalled(t, "BuscarPacientePorUsuarioID")
}

func TestConviteServico_VincularPaciente_ConviteJaUtilizado(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	pacienteIDExistente := uint(99)
	conviteUsado := &dominio.Convite{
		ID:             1,
		ProfissionalID: 1,
		Token:          "abc123def456",
		DataExpiracao:  time.Now().Add(24 * time.Hour),
		Usado:          true,
		PacienteID:     &pacienteIDExistente,
	}

	mockConviteRepo.On("BuscarConvitePorToken", mock.Anything, "abc123def456").Return(conviteUsado, nil)

	err := servico.VincularPaciente(20, "abc123def456")

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrConviteJaUtilizado, err)
	mockConviteRepo.AssertExpectations(t)
	mockUsuarioRepo.AssertNotCalled(t, "BuscarPacientePorUsuarioID")
}

func TestConviteServico_VincularPaciente_PacienteNaoEncontrado(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	conviteValido := &dominio.Convite{
		ID:             1,
		ProfissionalID: 1,
		Token:          "abc123def456",
		DataExpiracao:  time.Now().Add(24 * time.Hour),
		Usado:          false,
	}

	mockConviteRepo.On("BuscarConvitePorToken", mock.Anything, "abc123def456").Return(conviteValido, nil)
	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(999)).Return(nil, gorm.ErrRecordNotFound)

	err := servico.VincularPaciente(999, "abc123def456")

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrUsuarioNaoEncontrado, err)
	mockConviteRepo.AssertExpectations(t)
	mockUsuarioRepo.AssertExpectations(t)
}

func TestConviteServico_VincularPaciente_ErroAoBuscarPaciente(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	conviteValido := &dominio.Convite{
		ID:             1,
		ProfissionalID: 1,
		Token:          "abc123def456",
		DataExpiracao:  time.Now().Add(24 * time.Hour),
		Usado:          false,
	}

	erroGenerico := errors.New("erro de conexão com banco de dados")
	mockConviteRepo.On("BuscarConvitePorToken", mock.Anything, "abc123def456").Return(conviteValido, nil)
	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(20)).Return(nil, erroGenerico)

	err := servico.VincularPaciente(20, "abc123def456")

	assert.Error(t, err)
	assert.Equal(t, erroGenerico, err)
	mockConviteRepo.AssertExpectations(t)
	mockUsuarioRepo.AssertExpectations(t)
}

func TestConviteServico_VincularPaciente_ErroAoMarcarComoUsado(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	// Criar tabela de associação many-to-many
	db.Exec(`CREATE TABLE IF NOT EXISTS profissionais_pacientes (
		profissional_id INTEGER,
		paciente_id INTEGER,
		PRIMARY KEY (profissional_id, paciente_id)
	)`)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	conviteValido := &dominio.Convite{
		ID:             1,
		ProfissionalID: 1,
		Token:          "abc123def456",
		DataExpiracao:  time.Now().Add(24 * time.Hour),
		Usado:          false,
	}

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 20,
	}

	erroGenerico := errors.New("erro ao atualizar convite")
	mockConviteRepo.On("BuscarConvitePorToken", mock.Anything, "abc123def456").Return(conviteValido, nil)
	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(20)).Return(pacienteExistente, nil)
	mockConviteRepo.On("MarcarConviteComoUsado", mock.Anything, conviteValido).Return(erroGenerico)

	err := servico.VincularPaciente(20, "abc123def456")

	assert.Error(t, err)
	assert.Equal(t, erroGenerico, err)
	mockConviteRepo.AssertExpectations(t)
	mockUsuarioRepo.AssertExpectations(t)
}

func TestConviteServico_VincularPaciente_ConviteExpiraEmSegundos(t *testing.T) {
	db := setupTestDBConvite(t)
	mockConviteRepo := new(MockConviteRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioConvite)

	servico := servicos.NovoConviteServico(db, mockConviteRepo, mockUsuarioRepo)

	// Convite que expira em poucos segundos (ainda válido)
	conviteQuaseExpirando := &dominio.Convite{
		ID:             1,
		ProfissionalID: 1,
		Token:          "abc123def456",
		DataExpiracao:  time.Now().Add(2 * time.Second),
		Usado:          false,
	}

	mockConviteRepo.On("BuscarConvitePorToken", mock.Anything, "abc123def456").Return(conviteQuaseExpirando, nil)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 20,
	}

	// Criar tabela de associação
	db.Exec(`CREATE TABLE IF NOT EXISTS profissionais_pacientes (
		profissional_id INTEGER,
		paciente_id INTEGER,
		PRIMARY KEY (profissional_id, paciente_id)
	)`)

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(20)).Return(pacienteExistente, nil)
	mockConviteRepo.On("MarcarConviteComoUsado", mock.Anything, conviteQuaseExpirando).Return(nil)

	err := servico.VincularPaciente(20, "abc123def456")

	// Ainda deve funcionar pois não expirou
	assert.NoError(t, err)
	mockConviteRepo.AssertExpectations(t)
	mockUsuarioRepo.AssertExpectations(t)
}
