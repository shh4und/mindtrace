package servicos

import (
	"errors"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/postgres"
	"testing"
	"time"
)

// Mock para o repositório de registro de humor para testes de relatório
type mockRelatorioHumorRepo struct {
	postgres.RegistroHumorRepositorio // Embedding a nill interface to satisfy it
	registrosToReturn                 []*dominio.RegistroHumor
	errToReturn                       error
}

func (m *mockRelatorioHumorRepo) BuscarPorPacienteEPeriodo(pacienteID uint, inicio, fim time.Time) ([]*dominio.RegistroHumor, error) {
	if m.errToReturn != nil {
		return nil, m.errToReturn
	}
	return m.registrosToReturn, nil
}

func TestGerarRelatorioPaciente_HappyPath(t *testing.T) {
	mockData := []*dominio.RegistroHumor{
		{PacienteID: 1, HorasSono: 8, NivelEnergia: 7, NivelStress: 3, DataHoraRegistro: time.Now().AddDate(0, 0, -1)},
		{PacienteID: 1, HorasSono: 6, NivelEnergia: 5, NivelStress: 5, DataHoraRegistro: time.Now().AddDate(0, 0, -2)},
		{PacienteID: 1, HorasSono: 7, NivelEnergia: 6, NivelStress: 4, DataHoraRegistro: time.Now().AddDate(0, 0, -3)},
	}

	mockRepo := &mockRelatorioHumorRepo{
		registrosToReturn: mockData,
	}

	service := NovoRelatorioServico(mockRepo)

	relatorio, err := service.GerarRelatorioPaciente(1, 30)

	if err != nil {
		t.Fatalf("esperado nenhum erro, mas recebeu: %v", err)
	}
	if relatorio == nil {
		t.Fatal("esperado um relatorio, mas recebeu nulo")
	}

	esperadoMediaSono := (8.0 + 6.0 + 7.0) / 3.0
	if relatorio.MediaSono != esperadoMediaSono {
		t.Errorf("esperado MediaSono %.2f, mas recebeu %.2f", esperadoMediaSono, relatorio.MediaSono)
	}

	if len(relatorio.GraficoSono) != 3 {
		t.Errorf("esperado 3 pontos de dados para o sono, mas recebeu %d", len(relatorio.GraficoSono))
	}
}

func TestGerarRelatorioPaciente_SemRegistros(t *testing.T) {
	mockRepo := &mockRelatorioHumorRepo{
		registrosToReturn: []*dominio.RegistroHumor{}, // Retorna lista vazia
	}

	service := NovoRelatorioServico(mockRepo)

	relatorio, err := service.GerarRelatorioPaciente(1, 30)

	if err != nil {
		t.Fatalf("esperado nenhum erro, mas recebeu: %v", err)
	}
	if relatorio == nil {
		t.Fatal("esperado um relatorio não nulo, mas recebeu nulo")
	}

	if relatorio.MediaSono != 0 {
		t.Errorf("esperado MediaSono 0, mas recebeu %.2f", relatorio.MediaSono)
	}
	if len(relatorio.GraficoEnergia) != 0 {
		t.Errorf("esperado 0 pontos de dados para energia, mas recebeu %d", len(relatorio.GraficoEnergia))
	}
}

func TestGerarRelatorioPaciente_ErroNoRepositorio(t *testing.T) {
	erroEsperado := errors.New("erro simulado de banco de dados")
	mockRepo := &mockRelatorioHumorRepo{
		errToReturn: erroEsperado,
	}

	service := NovoRelatorioServico(mockRepo)

	_, err := service.GerarRelatorioPaciente(1, 30)

	if !errors.Is(err, erroEsperado) {
		t.Fatalf("esperado erro '%v', mas recebeu: %v", erroEsperado, err)
	}
}
