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
	"gorm.io/gorm"
)

// ========== Mocks ==========

// MockRegistroHumorRepositorio simula o repositorio de registros de humor
type MockRegistroHumorRepositorio struct {
	mock.Mock
}

func (m *MockRegistroHumorRepositorio) CriarRegistroHumor(tx *gorm.DB, registro *dominio.RegistroHumor) error {
	args := m.Called(tx, registro)
	return args.Error(0)
}

func (m *MockRegistroHumorRepositorio) BuscarPorPacienteEPeriodo(pacienteID uint, inicio, fim time.Time) ([]*dominio.RegistroHumor, error) {
	args := m.Called(pacienteID, inicio, fim)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*dominio.RegistroHumor), args.Error(1)
}

func (m *MockRegistroHumorRepositorio) BuscarUltimoRegistroDePaciente(pacienteID uint) (*dominio.RegistroHumor, error) {
	args := m.Called(pacienteID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.RegistroHumor), args.Error(1)
}

func (m *MockRegistroHumorRepositorio) BuscarPorNUltimosRegistros(pacienteID uint, numLimite int) ([]*dominio.RegistroHumor, error) {
	args := m.Called(pacienteID, numLimite)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*dominio.RegistroHumor), args.Error(1)
}

// MockUsuarioRepositorioRH simula o repositorio de usuarios para testes de registro humor
type MockUsuarioRepositorioRH struct {
	mock.Mock
}

func (m *MockUsuarioRepositorioRH) CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error {
	return nil
}

func (m *MockUsuarioRepositorioRH) CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	return nil
}

func (m *MockUsuarioRepositorioRH) CriarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	return nil
}

func (m *MockUsuarioRepositorioRH) BuscarPorEmail(email string) (*dominio.Usuario, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRH) BuscarUsuarioPorID(id uint) (*dominio.Usuario, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRH) BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRH) BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRH) BuscarProfissionalPorUsuarioID(tx *gorm.DB, usuarioID uint) (*dominio.Profissional, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRH) BuscarPacientePorUsuarioID(tx *gorm.DB, usuarioID uint) (*dominio.Paciente, error) {
	args := m.Called(tx, usuarioID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Paciente), args.Error(1)
}

func (m *MockUsuarioRepositorioRH) BuscarPacientesDoProfissional(tx *gorm.DB, profissionalID uint) ([]dominio.Paciente, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRH) Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error {
	return nil
}

func (m *MockUsuarioRepositorioRH) AtualizarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	return nil
}

func (m *MockUsuarioRepositorioRH) AtualizarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	return nil
}

func (m *MockUsuarioRepositorioRH) DeletarUsuario(tx *gorm.DB, id uint) error {
	return nil
}

// MockAnaliseServico simula o servico de analise
type MockAnaliseServico struct {
	mock.Mock
}

func (m *MockAnaliseServico) GerarAnaliseHistorica(pacienteID uint, dias int) (*dtos.AnalisePacienteDTOOut, error) {
	args := m.Called(pacienteID, dias)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.AnalisePacienteDTOOut), args.Error(1)
}

func (m *MockAnaliseServico) ExecutarMonitoramento(pacienteID uint) error {
	args := m.Called(pacienteID)
	return args.Error(0)
}

// ========== Helper Functions ==========

func setupTestDBRegistroHumor(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Falha ao abrir banco de dados de teste: %v", err)
	}

	err = db.AutoMigrate(&dominio.Usuario{}, &dominio.Profissional{}, &dominio.Paciente{}, &dominio.RegistroHumor{})
	if err != nil {
		t.Fatalf("Falha ao migrar esquema: %v", err)
	}

	return db
}

// ========== Testes do Serviço ==========

func TestRegistroHumorServico_CriarRegistroHumor_Sucesso(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	mockAnaliseServico.On("ExecutarMonitoramento", mock.Anything).Return(nil).Maybe()

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Exercício físico",
		Observacoes:      "Dia produtivo",
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("CriarRegistroHumor", mock.Anything, mock.AnythingOfType("*dominio.RegistroHumor")).Return(nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, int16(4), resultado.NivelHumor)
	assert.Equal(t, int16(8), resultado.HorasSono)
	assert.Equal(t, int16(7), resultado.NivelEnergia)
	assert.Equal(t, int16(3), resultado.NivelStress)
	assert.Equal(t, "Exercício físico", resultado.AutoCuidado)
	assert.Equal(t, uint(1), resultado.PacienteID)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRegistroHumorServico_CriarRegistroHumor_PacienteNaoEncontrado(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Exercício físico",
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(999)).Return(nil, gorm.ErrRecordNotFound)

	resultado, err := servico.CriarRegistroHumor(dto, 999)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrUsuarioNaoEncontrado, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertNotCalled(t, "CriarRegistroHumor")
}

func TestRegistroHumorServico_CriarRegistroHumor_ErroAoBuscarPaciente(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Exercício físico",
		DataHoraRegistro: time.Now(),
	}

	erroGenerico := errors.New("erro de conexão com banco de dados")
	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(nil, erroGenerico)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.Error(t, err)
	assert.Equal(t, erroGenerico, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertNotCalled(t, "CriarRegistroHumor")
}

func TestRegistroHumorServico_CriarRegistroHumor_ValidacaoNivelHumorInvalido(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       0, // Inválido
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Exercício físico",
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrNivelHumorInvalido, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertNotCalled(t, "CriarRegistroHumor")
}

func TestRegistroHumorServico_CriarRegistroHumor_ValidacaoHorasSonoInvalido(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        15, // Inválido
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Exercício físico",
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrHorasSonoInvalido, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertNotCalled(t, "CriarRegistroHumor")
}

func TestRegistroHumorServico_CriarRegistroHumor_ValidacaoNivelEnergiaInvalido(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     0, // Inválido
		NivelStress:      3,
		AutoCuidado:      "Exercício físico",
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrNivelEnergiaInvalido, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertNotCalled(t, "CriarRegistroHumor")
}

func TestRegistroHumorServico_CriarRegistroHumor_ValidacaoNivelStressInvalido(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      11, // Inválido
		AutoCuidado:      "Exercício físico",
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrNivelStressInvalido, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertNotCalled(t, "CriarRegistroHumor")
}

func TestRegistroHumorServico_CriarRegistroHumor_ValidacaoAutoCuidadoVazio(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "", // Inválido
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrAutoCuidadoVazio, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertNotCalled(t, "CriarRegistroHumor")
}

func TestRegistroHumorServico_CriarRegistroHumor_ValidacaoDataHoraRegistroVazia(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Exercício físico",
		DataHoraRegistro: time.Time{}, // Inválido
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.Error(t, err)
	assert.Equal(t, dominio.ErrDataHoraRegistroVazia, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertNotCalled(t, "CriarRegistroHumor")
}

func TestRegistroHumorServico_CriarRegistroHumor_ErroAoCriarRegistro(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Exercício físico",
		DataHoraRegistro: time.Now(),
	}

	erroGenerico := errors.New("erro ao inserir no banco de dados")
	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("CriarRegistroHumor", mock.Anything, mock.AnythingOfType("*dominio.RegistroHumor")).Return(erroGenerico)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.Error(t, err)
	assert.Equal(t, erroGenerico, err)
	assert.Nil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRegistroHumorServico_CriarRegistroHumor_ComObservacoes(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	mockAnaliseServico.On("ExecutarMonitoramento", mock.Anything).Return(nil).Maybe()

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       5,
		HorasSono:        9,
		NivelEnergia:     8,
		NivelStress:      2,
		AutoCuidado:      "Meditação e yoga",
		Observacoes:      "Excelente dia, me senti muito bem após a sessão de terapia",
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("CriarRegistroHumor", mock.Anything, mock.AnythingOfType("*dominio.RegistroHumor")).Return(nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, "Excelente dia, me senti muito bem após a sessão de terapia", resultado.Observacoes)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRegistroHumorServico_CriarRegistroHumor_ValoresMinimos(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	mockAnaliseServico.On("ExecutarMonitoramento", mock.Anything).Return(nil).Maybe()

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       1,
		HorasSono:        0,
		NivelEnergia:     1,
		NivelStress:      1,
		AutoCuidado:      "Nenhum",
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("CriarRegistroHumor", mock.Anything, mock.AnythingOfType("*dominio.RegistroHumor")).Return(nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, int16(1), resultado.NivelHumor)
	assert.Equal(t, int16(0), resultado.HorasSono)
	assert.Equal(t, int16(1), resultado.NivelEnergia)
	assert.Equal(t, int16(1), resultado.NivelStress)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRegistroHumorServico_CriarRegistroHumor_ValoresMaximos(t *testing.T) {
	db := setupTestDBRegistroHumor(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRH)
	mockAnaliseServico := new(MockAnaliseServico)

	mockAnaliseServico.On("ExecutarMonitoramento", mock.Anything).Return(nil).Maybe()

	servico := servicos.NovoRegistroHumorServico(db, mockRegistroHumorRepo, mockUsuarioRepo, mockAnaliseServico)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	dto := dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       5,
		HorasSono:        12,
		NivelEnergia:     10,
		NivelStress:      10,
		AutoCuidado:      "Todas as atividades possíveis",
		DataHoraRegistro: time.Now(),
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", mock.Anything, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("CriarRegistroHumor", mock.Anything, mock.AnythingOfType("*dominio.RegistroHumor")).Return(nil)

	resultado, err := servico.CriarRegistroHumor(dto, 10)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, int16(5), resultado.NivelHumor)
	assert.Equal(t, int16(12), resultado.HorasSono)
	assert.Equal(t, int16(10), resultado.NivelEnergia)
	assert.Equal(t, int16(10), resultado.NivelStress)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertExpectations(t)
}
