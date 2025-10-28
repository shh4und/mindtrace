package tests

import (
	"mindtrace/backend/interno/dominio"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ========== Testes para Convite ==========

func TestConvite_ValidarToken(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		wantErr error
	}{
		{
			name:    "token valido com 10 caracteres",
			token:   "1234567890",
			wantErr: nil,
		},
		{
			name:    "token valido com 32 caracteres (hex)",
			token:   "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6",
			wantErr: nil,
		},
		{
			name:    "token valido longo",
			token:   "abc123def456ghi789jkl012mno345pqr678stu901vwx234yz",
			wantErr: nil,
		},
		{
			name:    "token vazio",
			token:   "",
			wantErr: dominio.ErrTokenConviteVazio,
		},
		{
			name:    "token com menos de 10 caracteres",
			token:   "abc123",
			wantErr: dominio.ErrTokenConviteInvalido,
		},
		{
			name:    "token com 9 caracteres",
			token:   "123456789",
			wantErr: dominio.ErrTokenConviteInvalido,
		},
		{
			name:    "token com 1 caractere",
			token:   "a",
			wantErr: dominio.ErrTokenConviteInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convite := &dominio.Convite{Token: tt.token}
			err := convite.ValidarToken()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestConvite_ValidarDataExpiracao(t *testing.T) {
	now := time.Now()
	futuro := now.Add(24 * time.Hour)
	passado := now.Add(-24 * time.Hour)

	tests := []struct {
		name          string
		dataExpiracao time.Time
		wantErr       error
	}{
		{
			name:          "data expiracao valida (amanha)",
			dataExpiracao: futuro,
			wantErr:       nil,
		},
		{
			name:          "data expiracao valida (daqui 1 hora)",
			dataExpiracao: now.Add(1 * time.Hour),
			wantErr:       nil,
		},
		{
			name:          "data expiracao valida (daqui 1 minuto)",
			dataExpiracao: now.Add(1 * time.Minute),
			wantErr:       nil,
		},
		{
			name:          "data expiracao vazia",
			dataExpiracao: time.Time{},
			wantErr:       dominio.ErrDataExpiracaoVazia,
		},
		{
			name:          "data expiracao no passado (ontem)",
			dataExpiracao: passado,
			wantErr:       dominio.ErrDataExpiracaoNoPassado,
		},
		{
			name:          "data expiracao no passado (1 hora atras)",
			dataExpiracao: now.Add(-1 * time.Hour),
			wantErr:       dominio.ErrDataExpiracaoNoPassado,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convite := &dominio.Convite{DataExpiracao: tt.dataExpiracao}
			err := convite.ValidarDataExpiracao()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestConvite_Validar(t *testing.T) {
	now := time.Now()
	futuro := now.Add(24 * time.Hour)

	tests := []struct {
		name    string
		convite dominio.Convite
		wantErr error
	}{
		{
			name: "convite valido completo",
			convite: dominio.Convite{
				Token:         "abc123def456",
				DataExpiracao: futuro,
				Usado:         false,
			},
			wantErr: nil,
		},
		{
			name: "convite valido com token longo",
			convite: dominio.Convite{
				Token:         "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6",
				DataExpiracao: futuro,
				Usado:         false,
			},
			wantErr: nil,
		},
		{
			name: "convite com token vazio",
			convite: dominio.Convite{
				Token:         "",
				DataExpiracao: futuro,
				Usado:         false,
			},
			wantErr: dominio.ErrTokenConviteVazio,
		},
		{
			name: "convite com token curto",
			convite: dominio.Convite{
				Token:         "abc123",
				DataExpiracao: futuro,
				Usado:         false,
			},
			wantErr: dominio.ErrTokenConviteInvalido,
		},
		{
			name: "convite com data expiracao vazia",
			convite: dominio.Convite{
				Token:         "abc123def456",
				DataExpiracao: time.Time{},
				Usado:         false,
			},
			wantErr: dominio.ErrDataExpiracaoVazia,
		},
		{
			name: "convite com data expiracao no passado",
			convite: dominio.Convite{
				Token:         "abc123def456",
				DataExpiracao: now.Add(-1 * time.Hour),
				Usado:         false,
			},
			wantErr: dominio.ErrDataExpiracaoNoPassado,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.convite.Validar()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestConvite_EstaValido(t *testing.T) {
	now := time.Now()
	futuro := now.Add(24 * time.Hour)
	passado := now.Add(-24 * time.Hour)

	tests := []struct {
		name     string
		convite  dominio.Convite
		expected bool
	}{
		{
			name: "convite valido (nao usado, nao expirado)",
			convite: dominio.Convite{
				Token:         "abc123def456",
				DataExpiracao: futuro,
				Usado:         false,
			},
			expected: true,
		},
		{
			name: "convite invalido (ja usado)",
			convite: dominio.Convite{
				Token:         "abc123def456",
				DataExpiracao: futuro,
				Usado:         true,
			},
			expected: false,
		},
		{
			name: "convite invalido (expirado)",
			convite: dominio.Convite{
				Token:         "abc123def456",
				DataExpiracao: passado,
				Usado:         false,
			},
			expected: false,
		},
		{
			name: "convite invalido (usado e expirado)",
			convite: dominio.Convite{
				Token:         "abc123def456",
				DataExpiracao: passado,
				Usado:         true,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.convite.EstaValido()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvite_EstaExpirado(t *testing.T) {
	now := time.Now()
	futuro := now.Add(24 * time.Hour)
	passado := now.Add(-24 * time.Hour)

	tests := []struct {
		name     string
		convite  dominio.Convite
		expected bool
	}{
		{
			name: "convite nao expirado (futuro)",
			convite: dominio.Convite{
				DataExpiracao: futuro,
			},
			expected: false,
		},
		{
			name: "convite nao expirado (daqui 1 hora)",
			convite: dominio.Convite{
				DataExpiracao: now.Add(1 * time.Hour),
			},
			expected: false,
		},
		{
			name: "convite expirado (passado)",
			convite: dominio.Convite{
				DataExpiracao: passado,
			},
			expected: true,
		},
		{
			name: "convite expirado (1 hora atras)",
			convite: dominio.Convite{
				DataExpiracao: now.Add(-1 * time.Hour),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.convite.EstaExpirado()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvite_JaFoiUtilizado(t *testing.T) {
	tests := []struct {
		name     string
		usado    bool
		expected bool
	}{
		{
			name:     "convite nao utilizado",
			usado:    false,
			expected: false,
		},
		{
			name:     "convite ja utilizado",
			usado:    true,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convite := dominio.Convite{Usado: tt.usado}
			result := convite.JaFoiUtilizado()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvite_UtilizarConvite(t *testing.T) {
	convite := &dominio.Convite{
		Token:         "abc123def456",
		DataExpiracao: time.Now().Add(24 * time.Hour),
		Usado:         false,
		PacienteID:    nil,
	}

	pacienteID := uint(123)

	// Verifica estado inicial
	assert.False(t, convite.Usado)
	assert.Nil(t, convite.PacienteID)

	// Utiliza o convite
	convite.UtilizarConvite(pacienteID)

	// Verifica que foi marcado como usado
	assert.True(t, convite.Usado)
	assert.NotNil(t, convite.PacienteID)
	assert.Equal(t, pacienteID, *convite.PacienteID)
}

func TestConvite_UtilizarConvite_ComDiferentesPacientes(t *testing.T) {
	tests := []struct {
		name       string
		pacienteID uint
	}{
		{
			name:       "paciente ID 1",
			pacienteID: 1,
		},
		{
			name:       "paciente ID 999",
			pacienteID: 999,
		},
		{
			name:       "paciente ID 42",
			pacienteID: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convite := &dominio.Convite{
				Token:         "abc123def456",
				DataExpiracao: time.Now().Add(24 * time.Hour),
				Usado:         false,
				PacienteID:    nil,
			}

			convite.UtilizarConvite(tt.pacienteID)

			assert.True(t, convite.Usado)
			assert.Equal(t, tt.pacienteID, *convite.PacienteID)
		})
	}
}

// ========== Testes de Casos Extremos ==========

func TestConvite_ValidarComTokenExatamente10Caracteres(t *testing.T) {
	convite := &dominio.Convite{
		Token:         "1234567890", // Exatamente 10
		DataExpiracao: time.Now().Add(24 * time.Hour),
	}

	err := convite.Validar()
	assert.NoError(t, err)
}

func TestConvite_EstaValidoLimiteExpiracao(t *testing.T) {
	// Testa convite que expira em 1 segundo
	convite := &dominio.Convite{
		Token:         "abc123def456",
		DataExpiracao: time.Now().Add(1 * time.Second),
		Usado:         false,
	}

	assert.True(t, convite.EstaValido())
}

func TestConvite_CombinacoesDeEstado(t *testing.T) {
	now := time.Now()
	futuro := now.Add(24 * time.Hour)
	passado := now.Add(-24 * time.Hour)

	tests := []struct {
		name           string
		usado          bool
		dataExpiracao  time.Time
		esperaValido   bool
		esperaExpirado bool
		esperaUsado    bool
	}{
		{
			name:           "novo convite valido",
			usado:          false,
			dataExpiracao:  futuro,
			esperaValido:   true,
			esperaExpirado: false,
			esperaUsado:    false,
		},
		{
			name:           "convite usado mas nao expirado",
			usado:          true,
			dataExpiracao:  futuro,
			esperaValido:   false,
			esperaExpirado: false,
			esperaUsado:    true,
		},
		{
			name:           "convite expirado mas nao usado",
			usado:          false,
			dataExpiracao:  passado,
			esperaValido:   false,
			esperaExpirado: true,
			esperaUsado:    false,
		},
		{
			name:           "convite usado e expirado",
			usado:          true,
			dataExpiracao:  passado,
			esperaValido:   false,
			esperaExpirado: true,
			esperaUsado:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convite := &dominio.Convite{
				Token:         "abc123def456",
				DataExpiracao: tt.dataExpiracao,
				Usado:         tt.usado,
			}

			assert.Equal(t, tt.esperaValido, convite.EstaValido())
			assert.Equal(t, tt.esperaExpirado, convite.EstaExpirado())
			assert.Equal(t, tt.esperaUsado, convite.JaFoiUtilizado())
		})
	}
}
