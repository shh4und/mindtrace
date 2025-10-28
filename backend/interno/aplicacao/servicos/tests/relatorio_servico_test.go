package tests

import (
	"errors"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ========== Mocks ==========

// MockRegistroHumorRepositorioRelatorio simula o repositorio de registros de humor
type MockRegistroHumorRepositorioRelatorio struct {
	mock.Mock
}

func (m *MockRegistroHumorRepositorioRelatorio) CriarRegistroHumor(tx *gorm.DB, registro *dominio.RegistroHumor) error {
	return nil
}

func (m *MockRegistroHumorRepositorioRelatorio) BuscarPorPacienteEPeriodo(pacienteID uint, inicio, fim time.Time) ([]*dominio.RegistroHumor, error) {
	args := m.Called(pacienteID, inicio, fim)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*dominio.RegistroHumor), args.Error(1)
}

func (m *MockRegistroHumorRepositorioRelatorio) BuscarUltimoRegistroDePaciente(pacienteID uint) (*dominio.RegistroHumor, error) {
	return nil, nil
}

// MockUsuarioRepositorioRelatorio simula o repositorio de usuarios
type MockUsuarioRepositorioRelatorio struct {
	mock.Mock
}

func (m *MockUsuarioRepositorioRelatorio) CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error {
	return nil
}

func (m *MockUsuarioRepositorioRelatorio) CriarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	return nil
}

func (m *MockUsuarioRepositorioRelatorio) CriarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	return nil
}

func (m *MockUsuarioRepositorioRelatorio) BuscarPorEmail(email string) (*dominio.Usuario, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRelatorio) BuscarUsuarioPorID(id uint) (*dominio.Usuario, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRelatorio) BuscarProfissionalPorID(tx *gorm.DB, id uint) (*dominio.Profissional, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRelatorio) BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRelatorio) BuscarProfissionalPorUsuarioID(tx *gorm.DB, usuarioID uint) (*dominio.Profissional, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRelatorio) BuscarPacientePorUsuarioID(tx *gorm.DB, usuarioID uint) (*dominio.Paciente, error) {
	args := m.Called(tx, usuarioID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dominio.Paciente), args.Error(1)
}

func (m *MockUsuarioRepositorioRelatorio) BuscarPacientesDoProfissional(tx *gorm.DB, profissionalID uint) ([]dominio.Paciente, error) {
	return nil, nil
}

func (m *MockUsuarioRepositorioRelatorio) Atualizar(tx *gorm.DB, usuario *dominio.Usuario) error {
	return nil
}

func (m *MockUsuarioRepositorioRelatorio) AtualizarProfissional(tx *gorm.DB, profissional *dominio.Profissional) error {
	return nil
}

func (m *MockUsuarioRepositorioRelatorio) AtualizarPaciente(tx *gorm.DB, paciente *dominio.Paciente) error {
	return nil
}

func (m *MockUsuarioRepositorioRelatorio) DeletarUsuario(tx *gorm.DB, id uint) error {
	return nil
}

// ========== Helper Functions ==========

func setupTestDBRelatorio(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Falha ao abrir banco de dados de teste: %v", err)
	}

	err = db.AutoMigrate(&dominio.Usuario{}, &dominio.Paciente{}, &dominio.RegistroHumor{})
	if err != nil {
		t.Fatalf("Falha ao migrar esquema: %v", err)
	}

	return db
}

// ========== Testes GerarRelatorioPaciente ==========

func TestRelatorioServico_GerarRelatorioPaciente_Sucesso(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			PacienteID:       1,
			NivelHumor:       4,
			HorasSono:        8,
			NivelEnergia:     7,
			NivelStress:      3,
			Observacoes:      "Dia bom",
			DataHoraRegistro: now.AddDate(0, 0, -2),
		},
		{
			PacienteID:       1,
			NivelHumor:       3,
			HorasSono:        6,
			NivelEnergia:     5,
			NivelStress:      5,
			Observacoes:      "Dia médio",
			DataHoraRegistro: now.AddDate(0, 0, -1),
		},
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", db, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("BuscarPorPacienteEPeriodo", uint(1), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(registros, nil)

	resultado, err := servico.GerarRelatorioPaciente(10, 7)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)

	// Verificar médias calculadas
	assert.Equal(t, 7.0, resultado.MediaSono)    // (8 + 6) / 2 = 7
	assert.Equal(t, 6.0, resultado.MediaEnergia) // (7 + 5) / 2 = 6
	assert.Equal(t, 4.0, resultado.MediaStress)  // (3 + 5) / 2 = 4

	// Verificar gráficos
	assert.Len(t, resultado.GraficoSono, 2)
	assert.Len(t, resultado.GraficoEnergia, 2)
	assert.Len(t, resultado.GraficoStress, 2)

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRelatorioServico_GerarRelatorioPaciente_PeriodoInvalido_Zero(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	resultado, err := servico.GerarRelatorioPaciente(10, 0)

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "periodo de filtro deve ser maior que 0", err.Error())
}

func TestRelatorioServico_GerarRelatorioPaciente_PeriodoInvalido_Negativo(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	resultado, err := servico.GerarRelatorioPaciente(10, -5)

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "periodo de filtro deve ser maior que 0", err.Error())
}

func TestRelatorioServico_GerarRelatorioPaciente_PeriodoExcedeLimite(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	resultado, err := servico.GerarRelatorioPaciente(10, 91)

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "periodo de filtro nao pode exceder 90 dias", err.Error())
}

func TestRelatorioServico_GerarRelatorioPaciente_PacienteNaoEncontrado(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", db, uint(999)).Return(nil, gorm.ErrRecordNotFound)

	resultado, err := servico.GerarRelatorioPaciente(999, 7)

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "paciente nao encontrado", err.Error())
	mockUsuarioRepo.AssertExpectations(t)
}

func TestRelatorioServico_GerarRelatorioPaciente_ErroAoBuscarPaciente(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	erroGenerico := errors.New("erro de conexão com banco de dados")
	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", db, uint(10)).Return(nil, erroGenerico)

	resultado, err := servico.GerarRelatorioPaciente(10, 7)

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, erroGenerico, err)
	mockUsuarioRepo.AssertExpectations(t)
}

func TestRelatorioServico_GerarRelatorioPaciente_SemRegistros(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	registrosVazios := []*dominio.RegistroHumor{}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", db, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("BuscarPorPacienteEPeriodo", uint(1), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(registrosVazios, gorm.ErrRecordNotFound)

	resultado, err := servico.GerarRelatorioPaciente(10, 7)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)

	// Médias devem ser zero quando não há registros
	assert.Equal(t, 0.0, resultado.MediaSono)
	assert.Equal(t, 0.0, resultado.MediaEnergia)
	assert.Equal(t, 0.0, resultado.MediaStress)

	// Gráficos devem ser vazios
	assert.Empty(t, resultado.GraficoSono)
	assert.Empty(t, resultado.GraficoEnergia)
	assert.Empty(t, resultado.GraficoStress)

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRelatorioServico_GerarRelatorioPaciente_UmRegistro(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	registros := []*dominio.RegistroHumor{
		{
			PacienteID:       1,
			NivelHumor:       5,
			HorasSono:        9,
			NivelEnergia:     8,
			NivelStress:      2,
			Observacoes:      "Excelente dia",
			DataHoraRegistro: time.Now(),
		},
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", db, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("BuscarPorPacienteEPeriodo", uint(1), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(registros, nil)

	resultado, err := servico.GerarRelatorioPaciente(10, 7)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)

	// Com 1 registro, média = valor do registro
	assert.Equal(t, 9.0, resultado.MediaSono)
	assert.Equal(t, 8.0, resultado.MediaEnergia)
	assert.Equal(t, 2.0, resultado.MediaStress)

	// Cada gráfico deve ter 1 ponto
	assert.Len(t, resultado.GraficoSono, 1)
	assert.Len(t, resultado.GraficoEnergia, 1)
	assert.Len(t, resultado.GraficoStress, 1)
}

func TestRelatorioServico_GerarRelatorioPaciente_CalculoMediasCorreto(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{HorasSono: 7, NivelEnergia: 5, NivelStress: 6, DataHoraRegistro: now},
		{HorasSono: 8, NivelEnergia: 7, NivelStress: 4, DataHoraRegistro: now},
		{HorasSono: 9, NivelEnergia: 8, NivelStress: 2, DataHoraRegistro: now},
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", db, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("BuscarPorPacienteEPeriodo", uint(1), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(registros, nil)

	resultado, err := servico.GerarRelatorioPaciente(10, 30)

	assert.NoError(t, err)

	// Médias: (7+8+9)/3=8, (5+7+8)/3=6.666..., (6+4+2)/3=4
	assert.InDelta(t, 8.0, resultado.MediaSono, 0.01)
	assert.InDelta(t, 6.666666, resultado.MediaEnergia, 0.01)
	assert.InDelta(t, 4.0, resultado.MediaStress, 0.01)
}

// ========== Testes GerarRelatorioPacienteDoProfissional ==========

func TestRelatorioServico_GerarRelatorioPacienteDoProfissional_Sucesso(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			PacienteID:       1,
			NivelHumor:       4,
			HorasSono:        8,
			NivelEnergia:     7,
			NivelStress:      3,
			DataHoraRegistro: now,
		},
	}

	mockRegistroHumorRepo.On("BuscarPorPacienteEPeriodo", uint(1), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(registros, nil)

	resultado, err := servico.GerarRelatorioPacienteDoProfissional(1, 7)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, 8.0, resultado.MediaSono)
	assert.Equal(t, 7.0, resultado.MediaEnergia)
	assert.Equal(t, 3.0, resultado.MediaStress)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRelatorioServico_GerarRelatorioPacienteDoProfissional_PeriodoInvalido_Zero(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	resultado, err := servico.GerarRelatorioPacienteDoProfissional(1, 0)

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "periodo de filtro deve ser maior que 0", err.Error())
}

func TestRelatorioServico_GerarRelatorioPacienteDoProfissional_PeriodoExcedeLimite(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	resultado, err := servico.GerarRelatorioPacienteDoProfissional(1, 100)

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "periodo de filtro nao pode exceder 90 dias", err.Error())
}

func TestRelatorioServico_GerarRelatorioPacienteDoProfissional_PacienteIDInvalido(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	resultado, err := servico.GerarRelatorioPacienteDoProfissional(0, 7)

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "id do paciente invalido", err.Error())
}

func TestRelatorioServico_GerarRelatorioPacienteDoProfissional_SemRegistros(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	registrosVazios := []*dominio.RegistroHumor{}

	mockRegistroHumorRepo.On("BuscarPorPacienteEPeriodo", uint(1), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(registrosVazios, gorm.ErrRecordNotFound)

	resultado, err := servico.GerarRelatorioPacienteDoProfissional(1, 7)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, 0.0, resultado.MediaSono)
	assert.Empty(t, resultado.GraficoSono)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRelatorioServico_GerarRelatorioPacienteDoProfissional_ErroAoBuscarRegistros(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	erroGenerico := errors.New("erro de conexão com banco de dados")
	mockRegistroHumorRepo.On("BuscarPorPacienteEPeriodo", uint(1), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(nil, erroGenerico)

	resultado, err := servico.GerarRelatorioPacienteDoProfissional(1, 7)

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, erroGenerico, err)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRelatorioServico_GerarRelatorioPacienteDoProfissional_MuitosRegistros(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	// Criar 30 registros
	now := time.Now()
	var registros []*dominio.RegistroHumor
	for i := 0; i < 30; i++ {
		registros = append(registros, &dominio.RegistroHumor{
			PacienteID:       1,
			NivelHumor:       int16(3 + (i % 3)),
			HorasSono:        int16(6 + (i % 5)),
			NivelEnergia:     int16(5 + (i % 6)),
			NivelStress:      int16(2 + (i % 7)),
			DataHoraRegistro: now.AddDate(0, 0, -i),
		})
	}

	mockRegistroHumorRepo.On("BuscarPorPacienteEPeriodo", uint(1), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(registros, nil)

	resultado, err := servico.GerarRelatorioPacienteDoProfissional(1, 30)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Len(t, resultado.GraficoSono, 30)
	assert.Len(t, resultado.GraficoEnergia, 30)
	assert.Len(t, resultado.GraficoStress, 30)
	assert.Greater(t, resultado.MediaSono, 0.0)
	mockRegistroHumorRepo.AssertExpectations(t)
}

func TestRelatorioServico_PeriodoLimiteMaximo(t *testing.T) {
	db := setupTestDBRelatorio(t)
	mockRegistroHumorRepo := new(MockRegistroHumorRepositorioRelatorio)
	mockUsuarioRepo := new(MockUsuarioRepositorioRelatorio)

	servico := servicos.NovoRelatorioServico(db, mockRegistroHumorRepo, mockUsuarioRepo)

	pacienteExistente := &dominio.Paciente{
		ID:        1,
		UsuarioID: 10,
	}

	registros := []*dominio.RegistroHumor{
		{
			HorasSono:        8,
			NivelEnergia:     7,
			NivelStress:      3,
			DataHoraRegistro: time.Now(),
		},
	}

	mockUsuarioRepo.On("BuscarPacientePorUsuarioID", db, uint(10)).Return(pacienteExistente, nil)
	mockRegistroHumorRepo.On("BuscarPorPacienteEPeriodo", uint(1), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(registros, nil)

	// Testar com exatamente 90 dias (limite máximo permitido)
	resultado, err := servico.GerarRelatorioPaciente(10, 90)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroHumorRepo.AssertExpectations(t)
}
