package tests

import (
	"errors"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// ========== Extensão dos Mocks para Monitoramento ==========

// Adicionar método BuscarPorNUltimosRegistros ao MockRegistroHumorRepositorio
func (m *MockRegistroHumorRepositorio) BuscarPorNUltimosRegistros(pacienteID uint, numLimite int) ([]*dominio.RegistroHumor, error) {
	args := m.Called(pacienteID, numLimite)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*dominio.RegistroHumor), args.Error(1)
}

// TestRealizarMonitoramentoPaciente_Sucesso testa o monitoramento com sucesso
func TestRealizarMonitoramentoPaciente_Sucesso(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 7

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			ID:               1,
			PacienteID:       pacID,
			NivelHumor:       3,
			HorasSono:        7,
			NivelEnergia:     6,
			NivelStress:      5,
			AutoCuidado:      "Meditação",
			Observacoes:      "Dia tranquilo",
			DataHoraRegistro: now.Add(-6 * 24 * time.Hour),
		},
		{
			ID:               2,
			PacienteID:       pacID,
			NivelHumor:       4,
			HorasSono:        8,
			NivelEnergia:     7,
			NivelStress:      4,
			AutoCuidado:      "Exercícios",
			Observacoes:      "Bom dia",
			DataHoraRegistro: now.Add(-5 * 24 * time.Hour),
		},
		{
			ID:               3,
			PacienteID:       pacID,
			NivelHumor:       2,
			HorasSono:        5,
			NivelEnergia:     4,
			NivelStress:      7,
			AutoCuidado:      "Leitura",
			Observacoes:      "Estressado",
			DataHoraRegistro: now.Add(-4 * 24 * time.Hour),
		},
		{
			ID:               4,
			PacienteID:       pacID,
			NivelHumor:       3,
			HorasSono:        6,
			NivelEnergia:     5,
			NivelStress:      6,
			AutoCuidado:      "Caminhada",
			DataHoraRegistro: now.Add(-3 * 24 * time.Hour),
		},
		{
			ID:               5,
			PacienteID:       pacID,
			NivelHumor:       4,
			HorasSono:        7,
			NivelEnergia:     7,
			NivelStress:      4,
			AutoCuidado:      "Yoga",
			DataHoraRegistro: now.Add(-2 * 24 * time.Hour),
		},
		{
			ID:               6,
			PacienteID:       pacID,
			NivelHumor:       3,
			HorasSono:        6,
			NivelEnergia:     6,
			NivelStress:      5,
			AutoCuidado:      "Música",
			DataHoraRegistro: now.Add(-1 * 24 * time.Hour),
		},
		{
			ID:               7,
			PacienteID:       pacID,
			NivelHumor:       5,
			HorasSono:        8,
			NivelEnergia:     8,
			NivelStress:      3,
			AutoCuidado:      "Arte",
			DataHoraRegistro: now,
		},
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(registros, nil)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, pacID, resultado.PacienteID)
	assert.Equal(t, profissional.ID, resultado.ProfissionalID)
	assert.Len(t, resultado.DadosMonitoramento, numLimiteRegistros+numLimiteRegistros) // inicializado + adicionado

	// Verificar médias
	// Soma: sono=47, humor=24, energia=43, stress=34
	// Média: sono=6.71, humor=3.43, energia=6.14, stress=4.86
	assert.InDelta(t, 6.71, resultado.MediaSono, 0.01)
	assert.InDelta(t, 3.43, resultado.MediaHumor, 0.01)
	assert.InDelta(t, 6.14, resultado.MediaEnergia, 0.01)
	assert.InDelta(t, 4.86, resultado.MediaStress, 0.01)
	assert.NotEmpty(t, resultado.TipoAlerta)

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_PeriodoInvalidoMenorQue1 testa validação de período menor que 1
func TestRealizarMonitoramentoPaciente_PeriodoInvalidoMenorQue1(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(1, 10, 0)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.EqualError(t, err, "periodo de filtro deve ser maior que 1")
}

// TestRealizarMonitoramentoPaciente_PeriodoInvalidoNegativo testa validação de período negativo
func TestRealizarMonitoramentoPaciente_PeriodoInvalidoNegativo(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(1, 10, -5)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.EqualError(t, err, "periodo de filtro deve ser maior que 1")
}

// TestRealizarMonitoramentoPaciente_PeriodoExcedeLimite testa validação de período acima de 14 dias
func TestRealizarMonitoramentoPaciente_PeriodoExcedeLimite(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(1, 10, 15)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.EqualError(t, err, "periodo de filtro nao pode exceder 14 dias")
}

// TestRealizarMonitoramentoPaciente_PacienteIDInvalido testa validação de paciente ID zero
func TestRealizarMonitoramentoPaciente_PacienteIDInvalido(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(1, 0, 7)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.EqualError(t, err, "id do paciente invalido")
}

// TestRealizarMonitoramentoPaciente_ProfissionalNaoEncontrado testa busca de profissional não encontrado
func TestRealizarMonitoramentoPaciente_ProfissionalNaoEncontrado(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(nil, gorm.ErrRecordNotFound)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, 7)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.EqualError(t, err, "paciente nao encontrado")

	mockUsuarioRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_ErroBuscarProfissional testa erro ao buscar profissional
func TestRealizarMonitoramentoPaciente_ErroBuscarProfissional(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(nil, errors.New("erro no banco"))

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, 7)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.EqualError(t, err, "erro no banco")

	mockUsuarioRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_SemRegistros testa monitoramento sem registros de humor
func TestRealizarMonitoramentoPaciente_SemRegistros(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 7

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return([]*dominio.RegistroHumor{}, gorm.ErrRecordNotFound)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Len(t, resultado.DadosMonitoramento, numLimiteRegistros) // apenas inicializado

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_ErroBuscarRegistros testa erro ao buscar registros
func TestRealizarMonitoramentoPaciente_ErroBuscarRegistros(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 7

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(nil, errors.New("erro no banco"))

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.EqualError(t, err, "erro no banco")

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_UmRegistro testa monitoramento com apenas 1 registro
func TestRealizarMonitoramentoPaciente_UmRegistro(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 2

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			ID:               1,
			PacienteID:       pacID,
			NivelHumor:       4,
			HorasSono:        8,
			NivelEnergia:     7,
			NivelStress:      5,
			AutoCuidado:      "Exercícios",
			Observacoes:      "Bom dia",
			DataHoraRegistro: now,
		},
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(registros, nil)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, 4.0, resultado.MediaSono)
	assert.Equal(t, 2.0, resultado.MediaHumor)
	assert.Equal(t, 3.5, resultado.MediaEnergia)
	assert.Equal(t, 2.5, resultado.MediaStress)

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_ValoresMinimos testa monitoramento com valores mínimos
func TestRealizarMonitoramentoPaciente_ValoresMinimos(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 3

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			ID:               1,
			PacienteID:       pacID,
			NivelHumor:       1,
			HorasSono:        0,
			NivelEnergia:     1,
			NivelStress:      1,
			AutoCuidado:      "Nada",
			DataHoraRegistro: now,
		},
		{
			ID:               2,
			PacienteID:       pacID,
			NivelHumor:       1,
			HorasSono:        0,
			NivelEnergia:     1,
			NivelStress:      1,
			AutoCuidado:      "Nada",
			DataHoraRegistro: now.Add(-1 * 24 * time.Hour),
		},
		{
			ID:               3,
			PacienteID:       pacID,
			NivelHumor:       1,
			HorasSono:        0,
			NivelEnergia:     1,
			NivelStress:      1,
			AutoCuidado:      "Nada",
			DataHoraRegistro: now.Add(-2 * 24 * time.Hour),
		},
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(registros, nil)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, 0.0, resultado.MediaSono)
	assert.Equal(t, 1.0, resultado.MediaHumor)
	assert.Equal(t, 1.0, resultado.MediaEnergia)
	assert.Equal(t, 1.0, resultado.MediaStress)
	assert.Equal(t, "PREOCUPANTE", resultado.TipoAlerta) // Valores muito baixos

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_ValoresMaximos testa monitoramento com valores máximos
func TestRealizarMonitoramentoPaciente_ValoresMaximos(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 3

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			ID:               1,
			PacienteID:       pacID,
			NivelHumor:       5,
			HorasSono:        12,
			NivelEnergia:     10,
			NivelStress:      10,
			AutoCuidado:      "Tudo",
			DataHoraRegistro: now,
		},
		{
			ID:               2,
			PacienteID:       pacID,
			NivelHumor:       5,
			HorasSono:        12,
			NivelEnergia:     10,
			NivelStress:      10,
			AutoCuidado:      "Tudo",
			DataHoraRegistro: now.Add(-1 * 24 * time.Hour),
		},
		{
			ID:               3,
			PacienteID:       pacID,
			NivelHumor:       5,
			HorasSono:        12,
			NivelEnergia:     10,
			NivelStress:      10,
			AutoCuidado:      "Tudo",
			DataHoraRegistro: now.Add(-2 * 24 * time.Hour),
		},
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(registros, nil)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, 12.0, resultado.MediaSono)
	assert.Equal(t, 5.0, resultado.MediaHumor)
	assert.Equal(t, 10.0, resultado.MediaEnergia)
	assert.Equal(t, 10.0, resultado.MediaStress)
	assert.Equal(t, "PREOCUPANTE", resultado.TipoAlerta) // Stress máximo (10) é preocupante

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_PadraoPreocupante testa alerta PREOCUPANTE
func TestRealizarMonitoramentoPaciente_PadraoPreocupante(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 3

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			ID:               1,
			PacienteID:       pacID,
			NivelHumor:       1, // Muito baixo
			HorasSono:        3, // Muito baixo
			NivelEnergia:     2, // Muito baixo
			NivelStress:      9, // Muito alto
			AutoCuidado:      "Pouco",
			DataHoraRegistro: now,
		},
		{
			ID:               2,
			PacienteID:       pacID,
			NivelHumor:       2,
			HorasSono:        4,
			NivelEnergia:     3,
			NivelStress:      8,
			AutoCuidado:      "Pouco",
			DataHoraRegistro: now.Add(-1 * 24 * time.Hour),
		},
		{
			ID:               3,
			PacienteID:       pacID,
			NivelHumor:       1,
			HorasSono:        3,
			NivelEnergia:     2,
			NivelStress:      10,
			AutoCuidado:      "Pouco",
			DataHoraRegistro: now.Add(-2 * 24 * time.Hour),
		},
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(registros, nil)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, "PREOCUPANTE", resultado.TipoAlerta)

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_PadraoAtencao testa alerta ATENCAO
func TestRealizarMonitoramentoPaciente_PadraoAtencao(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 3

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			ID:               1,
			PacienteID:       pacID,
			NivelHumor:       2, // Abaixo da referência mas acima do preocupante
			HorasSono:        5, // Abaixo da referência
			NivelEnergia:     5, // No limite
			NivelStress:      5, // No limite
			AutoCuidado:      "Algum",
			DataHoraRegistro: now,
		},
		{
			ID:               2,
			PacienteID:       pacID,
			NivelHumor:       3,
			HorasSono:        5,
			NivelEnergia:     5,
			NivelStress:      6,
			AutoCuidado:      "Algum",
			DataHoraRegistro: now.Add(-1 * 24 * time.Hour),
		},
		{
			ID:               3,
			PacienteID:       pacID,
			NivelHumor:       2,
			HorasSono:        6,
			NivelEnergia:     5,
			NivelStress:      5,
			AutoCuidado:      "Algum",
			DataHoraRegistro: now.Add(-2 * 24 * time.Hour),
		},
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(registros, nil)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Contains(t, []string{"ATENCAO", "PREOCUPANTE"}, resultado.TipoAlerta)

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_PadraoRegular testa alerta REGULAR
func TestRealizarMonitoramentoPaciente_PadraoRegular(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 3

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			ID:               1,
			PacienteID:       pacID,
			NivelHumor:       4,
			HorasSono:        7,
			NivelEnergia:     7,
			NivelStress:      4,
			AutoCuidado:      "Bom",
			DataHoraRegistro: now,
		},
		{
			ID:               2,
			PacienteID:       pacID,
			NivelHumor:       4,
			HorasSono:        8,
			NivelEnergia:     8,
			NivelStress:      3,
			AutoCuidado:      "Bom",
			DataHoraRegistro: now.Add(-1 * 24 * time.Hour),
		},
		{
			ID:               3,
			PacienteID:       pacID,
			NivelHumor:       5,
			HorasSono:        8,
			NivelEnergia:     8,
			NivelStress:      4,
			AutoCuidado:      "Ótimo",
			DataHoraRegistro: now.Add(-2 * 24 * time.Hour),
		},
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(registros, nil)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Equal(t, "REGULAR", resultado.TipoAlerta)

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_LimiteMaximo14Dias testa limite máximo de 14 dias
func TestRealizarMonitoramentoPaciente_LimiteMaximo14Dias(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 14 // Limite máximo permitido

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	now := time.Now()
	registros := make([]*dominio.RegistroHumor, 14)
	for i := 0; i < 14; i++ {
		registros[i] = &dominio.RegistroHumor{
			ID:               uint(i + 1),
			PacienteID:       pacID,
			NivelHumor:       3,
			HorasSono:        7,
			NivelEnergia:     6,
			NivelStress:      5,
			AutoCuidado:      "Rotina",
			DataHoraRegistro: now.Add(time.Duration(-i) * 24 * time.Hour),
		}
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(registros, nil)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	assert.Len(t, resultado.DadosMonitoramento, numLimiteRegistros+numLimiteRegistros) // inicializado + adicionado

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}

// TestRealizarMonitoramentoPaciente_RegistrosSemObservacoes testa registros sem observações
func TestRealizarMonitoramentoPaciente_RegistrosSemObservacoes(t *testing.T) {
	// Arrange
	mockDB := &gorm.DB{}
	mockRegistroRepo := new(MockRegistroHumorRepositorio)
	mockUsuarioRepo := new(MockUsuarioRepositorio)

	servico := servicos.NovoMonitoramentoServico(mockDB, mockRegistroRepo, mockUsuarioRepo)

	userID := uint(1)
	pacID := uint(10)
	numLimiteRegistros := 2

	profissional := &dominio.Profissional{
		ID:        1,
		UsuarioID: userID,
	}

	now := time.Now()
	registros := []*dominio.RegistroHumor{
		{
			ID:               1,
			PacienteID:       pacID,
			NivelHumor:       3,
			HorasSono:        7,
			NivelEnergia:     6,
			NivelStress:      5,
			AutoCuidado:      "Meditação",
			Observacoes:      "", // Sem observações
			DataHoraRegistro: now,
		},
		{
			ID:               2,
			PacienteID:       pacID,
			NivelHumor:       4,
			HorasSono:        8,
			NivelEnergia:     7,
			NivelStress:      4,
			AutoCuidado:      "Exercícios",
			Observacoes:      "", // Sem observações
			DataHoraRegistro: now.Add(-1 * 24 * time.Hour),
		},
	}

	mockUsuarioRepo.On("BuscarProfissionalPorUsuarioID", mockDB, userID).Return(profissional, nil)
	mockRegistroRepo.On("BuscarPorNUltimosRegistros", pacID, numLimiteRegistros).Return(registros, nil)

	// Act
	resultado, err := servico.RealizarMonitoramentoPaciente(userID, pacID, numLimiteRegistros)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	for i := numLimiteRegistros; i < len(resultado.DadosMonitoramento); i++ {
		assert.Empty(t, resultado.DadosMonitoramento[i].Observacoes)
	}

	mockUsuarioRepo.AssertExpectations(t)
	mockRegistroRepo.AssertExpectations(t)
}
