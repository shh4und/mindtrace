package servicos

import (
	"errors"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/postgres"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Mock para o repositório de usuário
type mockUsuarioRepo struct {
	postgres.UsuarioRepositorio // Embedding a nill interface to satisfy it
	pacienteToReturn            *dominio.Paciente
	errToReturn                 error
}

func (m *mockUsuarioRepo) BuscarPacientePorID(tx *gorm.DB, id uint) (*dominio.Paciente, error) {
	if m.errToReturn != nil {
		return nil, m.errToReturn
	}
	if m.pacienteToReturn != nil && m.pacienteToReturn.UsuarioID == id {
		return m.pacienteToReturn, nil
	}
	return nil, gorm.ErrRecordNotFound
}

// Mock para o repositório de registro de humor
type mockRegistroHumorRepo struct {
	postgres.RegistroHumorRepositorio // Embedding a nill interface to satisfy it
	errToReturn                       error
}

func (m *mockRegistroHumorRepo) CriarRegistroHumor(tx *gorm.DB, registro *dominio.RegistroHumor) error {
	if m.errToReturn != nil {
		return m.errToReturn
	}
	registro.ID = 1 // Simulate DB assigning an ID
	return nil
}

func TestCriarRegistroHumor_HappyPath(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	mockUserRepo := &mockUsuarioRepo{
		pacienteToReturn: &dominio.Paciente{ID: 10, UsuarioID: 5},
	}
	mockHumorRepo := &mockRegistroHumorRepo{}

	service := NovoRegistroHumorServico(db, mockHumorRepo, mockUserRepo)

	dto := CriarRegistroHumorDTO{
		UsuarioID:  5,
		NivelHumor: 4,
		HorasSono:  8,
	}

	registro, err := service.CriarRegistroHumor(dto)

	if err != nil {
		t.Fatalf("esperado nenhum erro, mas recebeu: %v", err)
	}
	if registro == nil {
		t.Fatal("esperado um registro, mas recebeu nulo")
	}
	if registro.PacienteID != 10 {
		t.Errorf("esperado PacienteID 10, mas recebeu %d", registro.PacienteID)
	}
	if registro.HorasSono != 8 {
		t.Errorf("esperado HorasSono 8, mas recebeu %d", registro.HorasSono)
	}
}

func TestCriarRegistroHumor_PacienteNaoEncontrado(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	// Configura o mock para retornar "não encontrado"
	mockUserRepo := &mockUsuarioRepo{
		errToReturn: gorm.ErrRecordNotFound,
	}
	mockHumorRepo := &mockRegistroHumorRepo{}

	service := NovoRegistroHumorServico(db, mockHumorRepo, mockUserRepo)

	dto := CriarRegistroHumorDTO{
		UsuarioID: 999, // ID que não existe
	}

	_, err := service.CriarRegistroHumor(dto)

	if !errors.Is(err, ErrPacienteNaoEncontrado) {
		t.Fatalf("esperado erro ErrPacienteNaoEncontrado, mas recebeu: %v", err)
	}
}
