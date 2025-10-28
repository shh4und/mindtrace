package tests

import (
	"mindtrace/backend/interno/dominio"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ========== Testes para RegistroHumor ==========

func TestRegistroHumor_ValidarNivelHumor(t *testing.T) {
	tests := []struct {
		name       string
		nivelHumor int16
		wantErr    error
	}{
		{
			name:       "nivel humor valido minimo (1)",
			nivelHumor: 1,
			wantErr:    nil,
		},
		{
			name:       "nivel humor valido medio (3)",
			nivelHumor: 3,
			wantErr:    nil,
		},
		{
			name:       "nivel humor valido maximo (5)",
			nivelHumor: 5,
			wantErr:    nil,
		},
		{
			name:       "nivel humor invalido menor que 1",
			nivelHumor: 0,
			wantErr:    dominio.ErrNivelHumorInvalido,
		},
		{
			name:       "nivel humor invalido negativo",
			nivelHumor: -1,
			wantErr:    dominio.ErrNivelHumorInvalido,
		},
		{
			name:       "nivel humor invalido maior que 5",
			nivelHumor: 6,
			wantErr:    dominio.ErrNivelHumorInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &dominio.RegistroHumor{NivelHumor: tt.nivelHumor}
			err := rh.ValidarNivelHumor()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRegistroHumor_ValidarHorasSono(t *testing.T) {
	tests := []struct {
		name      string
		horasSono int16
		wantErr   error
	}{
		{
			name:      "horas sono valido minimo (0)",
			horasSono: 0,
			wantErr:   nil,
		},
		{
			name:      "horas sono valido medio (8)",
			horasSono: 8,
			wantErr:   nil,
		},
		{
			name:      "horas sono valido maximo (12)",
			horasSono: 12,
			wantErr:   nil,
		},
		{
			name:      "horas sono invalido negativo",
			horasSono: -1,
			wantErr:   dominio.ErrHorasSonoInvalido,
		},
		{
			name:      "horas sono invalido maior que 12",
			horasSono: 13,
			wantErr:   dominio.ErrHorasSonoInvalido,
		},
		{
			name:      "horas sono invalido muito alto",
			horasSono: 24,
			wantErr:   dominio.ErrHorasSonoInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &dominio.RegistroHumor{HorasSono: tt.horasSono}
			err := rh.ValidarHorasSono()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRegistroHumor_ValidarNivelEnergia(t *testing.T) {
	tests := []struct {
		name         string
		nivelEnergia int16
		wantErr      error
	}{
		{
			name:         "nivel energia valido minimo (1)",
			nivelEnergia: 1,
			wantErr:      nil,
		},
		{
			name:         "nivel energia valido medio (5)",
			nivelEnergia: 5,
			wantErr:      nil,
		},
		{
			name:         "nivel energia valido maximo (10)",
			nivelEnergia: 10,
			wantErr:      nil,
		},
		{
			name:         "nivel energia invalido menor que 1",
			nivelEnergia: 0,
			wantErr:      dominio.ErrNivelEnergiaInvalido,
		},
		{
			name:         "nivel energia invalido negativo",
			nivelEnergia: -5,
			wantErr:      dominio.ErrNivelEnergiaInvalido,
		},
		{
			name:         "nivel energia invalido maior que 10",
			nivelEnergia: 11,
			wantErr:      dominio.ErrNivelEnergiaInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &dominio.RegistroHumor{NivelEnergia: tt.nivelEnergia}
			err := rh.ValidarNivelEnergia()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRegistroHumor_ValidarNivelStress(t *testing.T) {
	tests := []struct {
		name        string
		nivelStress int16
		wantErr     error
	}{
		{
			name:        "nivel stress valido minimo (1)",
			nivelStress: 1,
			wantErr:     nil,
		},
		{
			name:        "nivel stress valido medio (5)",
			nivelStress: 5,
			wantErr:     nil,
		},
		{
			name:        "nivel stress valido maximo (10)",
			nivelStress: 10,
			wantErr:     nil,
		},
		{
			name:        "nivel stress invalido menor que 1",
			nivelStress: 0,
			wantErr:     dominio.ErrNivelStressInvalido,
		},
		{
			name:        "nivel stress invalido negativo",
			nivelStress: -3,
			wantErr:     dominio.ErrNivelStressInvalido,
		},
		{
			name:        "nivel stress invalido maior que 10",
			nivelStress: 11,
			wantErr:     dominio.ErrNivelStressInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &dominio.RegistroHumor{NivelStress: tt.nivelStress}
			err := rh.ValidarNivelStress()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRegistroHumor_ValidarAutoCuidado(t *testing.T) {
	tests := []struct {
		name        string
		autoCuidado string
		wantErr     error
	}{
		{
			name:        "auto cuidado valido",
			autoCuidado: "Exercício físico",
			wantErr:     nil,
		},
		{
			name:        "auto cuidado valido com 3 caracteres",
			autoCuidado: "Ler",
			wantErr:     nil,
		},
		{
			name:        "auto cuidado valido longo",
			autoCuidado: "Meditação, caminhada e alongamento",
			wantErr:     nil,
		},
		{
			name:        "auto cuidado vazio",
			autoCuidado: "",
			wantErr:     dominio.ErrAutoCuidadoVazio,
		},
		{
			name:        "auto cuidado com menos de 3 caracteres",
			autoCuidado: "Ok",
			wantErr:     dominio.ErrAutoCuidadoInvalido,
		},
		{
			name:        "auto cuidado com 1 caractere",
			autoCuidado: "X",
			wantErr:     dominio.ErrAutoCuidadoInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &dominio.RegistroHumor{AutoCuidado: tt.autoCuidado}
			err := rh.ValidarAutoCuidado()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRegistroHumor_ValidarDataHoraRegistro(t *testing.T) {
	tests := []struct {
		name             string
		dataHoraRegistro time.Time
		wantErr          error
	}{
		{
			name:             "data hora registro valida (agora)",
			dataHoraRegistro: time.Now(),
			wantErr:          nil,
		},
		{
			name:             "data hora registro valida (ontem)",
			dataHoraRegistro: time.Now().AddDate(0, 0, -1),
			wantErr:          nil,
		},
		{
			name:             "data hora registro valida (1 hora atras)",
			dataHoraRegistro: time.Now().Add(-1 * time.Hour),
			wantErr:          nil,
		},
		{
			name:             "data hora registro vazia",
			dataHoraRegistro: time.Time{},
			wantErr:          dominio.ErrDataHoraRegistroVazia,
		},
		{
			name:             "data hora registro no futuro (1 hora)",
			dataHoraRegistro: time.Now().Add(1 * time.Hour),
			wantErr:          dominio.ErrDataHoraRegistroNoFuturo,
		},
		{
			name:             "data hora registro no futuro (1 dia)",
			dataHoraRegistro: time.Now().AddDate(0, 0, 1),
			wantErr:          dominio.ErrDataHoraRegistroNoFuturo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &dominio.RegistroHumor{DataHoraRegistro: tt.dataHoraRegistro}
			err := rh.ValidarDataHoraRegistro()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRegistroHumor_Validar(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name          string
		registroHumor dominio.RegistroHumor
		wantErr       error
	}{
		{
			name: "registro humor valido completo",
			registroHumor: dominio.RegistroHumor{
				NivelHumor:       4,
				HorasSono:        8,
				NivelEnergia:     7,
				NivelStress:      3,
				AutoCuidado:      "Exercício físico",
				Observacoes:      "Dia produtivo",
				DataHoraRegistro: now,
			},
			wantErr: nil,
		},
		{
			name: "registro humor valido minimo",
			registroHumor: dominio.RegistroHumor{
				NivelHumor:       1,
				HorasSono:        0,
				NivelEnergia:     1,
				NivelStress:      1,
				AutoCuidado:      "Ler",
				DataHoraRegistro: now,
			},
			wantErr: nil,
		},
		{
			name: "registro humor valido maximo",
			registroHumor: dominio.RegistroHumor{
				NivelHumor:       5,
				HorasSono:        12,
				NivelEnergia:     10,
				NivelStress:      10,
				AutoCuidado:      "Meditação prolongada",
				DataHoraRegistro: now,
			},
			wantErr: nil,
		},
		{
			name: "registro humor com nivel humor invalido",
			registroHumor: dominio.RegistroHumor{
				NivelHumor:       0,
				HorasSono:        8,
				NivelEnergia:     7,
				NivelStress:      3,
				AutoCuidado:      "Exercício físico",
				DataHoraRegistro: now,
			},
			wantErr: dominio.ErrNivelHumorInvalido,
		},
		{
			name: "registro humor com horas sono invalido",
			registroHumor: dominio.RegistroHumor{
				NivelHumor:       4,
				HorasSono:        15,
				NivelEnergia:     7,
				NivelStress:      3,
				AutoCuidado:      "Exercício físico",
				DataHoraRegistro: now,
			},
			wantErr: dominio.ErrHorasSonoInvalido,
		},
		{
			name: "registro humor com nivel energia invalido",
			registroHumor: dominio.RegistroHumor{
				NivelHumor:       4,
				HorasSono:        8,
				NivelEnergia:     0,
				NivelStress:      3,
				AutoCuidado:      "Exercício físico",
				DataHoraRegistro: now,
			},
			wantErr: dominio.ErrNivelEnergiaInvalido,
		},
		{
			name: "registro humor com nivel stress invalido",
			registroHumor: dominio.RegistroHumor{
				NivelHumor:       4,
				HorasSono:        8,
				NivelEnergia:     7,
				NivelStress:      11,
				AutoCuidado:      "Exercício físico",
				DataHoraRegistro: now,
			},
			wantErr: dominio.ErrNivelStressInvalido,
		},
		{
			name: "registro humor com auto cuidado vazio",
			registroHumor: dominio.RegistroHumor{
				NivelHumor:       4,
				HorasSono:        8,
				NivelEnergia:     7,
				NivelStress:      3,
				AutoCuidado:      "",
				DataHoraRegistro: now,
			},
			wantErr: dominio.ErrAutoCuidadoVazio,
		},
		{
			name: "registro humor com data hora registro vazia",
			registroHumor: dominio.RegistroHumor{
				NivelHumor:       4,
				HorasSono:        8,
				NivelEnergia:     7,
				NivelStress:      3,
				AutoCuidado:      "Exercício físico",
				DataHoraRegistro: time.Time{},
			},
			wantErr: dominio.ErrDataHoraRegistroVazia,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.registroHumor.Validar()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

// ========== Testes de Casos Extremos ==========

func TestRegistroHumor_TableName(t *testing.T) {
	rh := dominio.RegistroHumor{}
	assert.Equal(t, "registros_humor", rh.TableName())
}

func TestRegistroHumor_ValidarComObservacoesVazias(t *testing.T) {
	registro := dominio.RegistroHumor{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Meditação",
		Observacoes:      "", // Vazio é válido
		DataHoraRegistro: time.Now(),
	}

	err := registro.Validar()
	assert.NoError(t, err)
}

func TestRegistroHumor_ValidarComObservacoesLongas(t *testing.T) {
	longText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. " +
		"Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. " +
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris."

	registro := dominio.RegistroHumor{
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Terapia e exercícios",
		Observacoes:      longText,
		DataHoraRegistro: time.Now(),
	}

	err := registro.Validar()
	assert.NoError(t, err)
}

func TestRegistroHumor_ValidarTodosNiveisMaximos(t *testing.T) {
	registro := dominio.RegistroHumor{
		NivelHumor:       5,
		HorasSono:        12,
		NivelEnergia:     10,
		NivelStress:      10,
		AutoCuidado:      "Atividades completas",
		DataHoraRegistro: time.Now(),
	}

	err := registro.Validar()
	assert.NoError(t, err)
}

func TestRegistroHumor_ValidarTodosNiveisMinimos(t *testing.T) {
	registro := dominio.RegistroHumor{
		NivelHumor:       1,
		HorasSono:        0,
		NivelEnergia:     1,
		NivelStress:      1,
		AutoCuidado:      "Nenhum",
		DataHoraRegistro: time.Now(),
	}

	err := registro.Validar()
	assert.NoError(t, err)
}
