package tests

import (
	"mindtrace/backend/interno/dominio"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ========== Testes para Usuario ==========

func TestUsuario_ValidarEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr error
	}{
		{
			name:    "email valido simples",
			email:   "usuario@example.com",
			wantErr: nil,
		},
		{
			name:    "email valido com subdominio",
			email:   "usuario@mail.example.com",
			wantErr: nil,
		},
		{
			name:    "email valido com numeros e hifen",
			email:   "user-123@example-domain.com.br",
			wantErr: nil,
		},
		{
			name:    "email invalido sem @",
			email:   "usuarioexample.com",
			wantErr: dominio.ErrEmailInvalido,
		},
		{
			name:    "email invalido sem dominio",
			email:   "usuario@",
			wantErr: dominio.ErrEmailInvalido,
		},
		{
			name:    "email invalido sem usuario",
			email:   "@example.com",
			wantErr: dominio.ErrEmailInvalido,
		},
		{
			name:    "email vazio",
			email:   "",
			wantErr: dominio.ErrEmailInvalido,
		},
		{
			name:    "email com espacos",
			email:   "user name@example.com",
			wantErr: dominio.ErrEmailInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &dominio.Usuario{Email: tt.email}
			err := u.ValidarEmail()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUsuario_ValidarSenha(t *testing.T) {
	tests := []struct {
		name       string
		senhaPlana string
		wantErr    error
	}{
		{
			name:       "senha valida com 9 caracteres",
			senhaPlana: "Senha1234",
			wantErr:    nil,
		},
		{
			name:       "senha valida com caracteres especiais",
			senhaPlana: "S3nh@Forte!",
			wantErr:    nil,
		},
		{
			name:       "senha valida longa",
			senhaPlana: "SenhaForte123!@#",
			wantErr:    nil,
		},
		{
			name:       "senha com menos de 8 caracteres",
			senhaPlana: "Senh@1",
			wantErr:    dominio.ErrSenhaFraca,
		},
		{
			name:       "senha vazia",
			senhaPlana: "",
			wantErr:    dominio.ErrSenhaFraca,
		},
		{
			name:       "senha com 7 caracteres",
			senhaPlana: "Senha12",
			wantErr:    dominio.ErrSenhaFraca,
		},
		{
			name:       "senha com 8 caracteres exatos",
			senhaPlana: "Senha123",
			wantErr:    dominio.ErrSenhaInvalida,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &dominio.Usuario{}
			err := u.ValidarSenha(tt.senhaPlana)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUsuario_ValidarNome(t *testing.T) {
	tests := []struct {
		name    string
		nome    string
		wantErr error
	}{
		{
			name:    "nome valido",
			nome:    "João Silva",
			wantErr: nil,
		},
		{
			name:    "nome valido com acentuacao",
			nome:    "José María González",
			wantErr: nil,
		},
		{
			name:    "nome vazio",
			nome:    "",
			wantErr: dominio.ErrNomeVazio,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &dominio.Usuario{Nome: tt.nome}
			err := u.ValidarNome()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUsuario_Validar(t *testing.T) {
	tests := []struct {
		name    string
		usuario dominio.Usuario
		wantErr error
	}{
		{
			name: "usuario valido",
			usuario: dominio.Usuario{
				Nome:  "João Silva",
				Email: "joao@example.com",
			},
			wantErr: nil,
		},
		{
			name: "usuario com email invalido",
			usuario: dominio.Usuario{
				Nome:  "João Silva",
				Email: "email-invalido",
			},
			wantErr: dominio.ErrEmailInvalido,
		},
		{
			name: "usuario com nome vazio",
			usuario: dominio.Usuario{
				Nome:  "",
				Email: "joao@example.com",
			},
			wantErr: dominio.ErrNomeVazio,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.usuario.Validar()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

// ========== Testes para Profissional ==========

func TestProfissional_ValidarRegistroProfissional(t *testing.T) {
	tests := []struct {
		name     string
		registro string
		wantErr  error
	}{
		{
			name:     "registro valido com 4 caracteres",
			registro: "1234",
			wantErr:  nil,
		},
		{
			name:     "registro valido com 12 caracteres",
			registro: "123456789012",
			wantErr:  nil,
		},
		{
			name:     "registro valido com 8 caracteres",
			registro: "CRP12345",
			wantErr:  nil,
		},
		{
			name:     "registro vazio",
			registro: "",
			wantErr:  dominio.ErrRegistroProfissionalVazio,
		},
		{
			name:     "registro com menos de 4 caracteres",
			registro: "123",
			wantErr:  dominio.ErrRegistroProfissionalInvalido,
		},
		{
			name:     "registro com mais de 12 caracteres",
			registro: "1234567890123",
			wantErr:  dominio.ErrRegistroProfissionalInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dominio.Profissional{RegistroProfissional: tt.registro}
			err := p.ValidarRegistroProfissional()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestProfissional_ValidarEspecialidade(t *testing.T) {
	tests := []struct {
		name          string
		especialidade string
		wantErr       error
	}{
		{
			name:          "especialidade valida",
			especialidade: "Psicologia Clínica",
			wantErr:       nil,
		},
		{
			name:          "especialidade valida com 3 caracteres",
			especialidade: "Psi",
			wantErr:       nil,
		},
		{
			name:          "especialidade vazia",
			especialidade: "",
			wantErr:       dominio.ErrEspecialidadeVazia,
		},
		{
			name:          "especialidade com menos de 3 caracteres",
			especialidade: "Ps",
			wantErr:       dominio.ErrEspecialidadeInvalida,
		},
		{
			name:          "especialidade com mais de 255 caracteres",
			especialidade: string(make([]byte, 256)),
			wantErr:       dominio.ErrEspecialidadeInvalida,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dominio.Profissional{Especialidade: tt.especialidade}
			err := p.ValidarEspecialidade()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestProfissional_ValidarDataNascimento(t *testing.T) {
	tests := []struct {
		name           string
		dataNascimento time.Time
		wantErr        error
	}{
		{
			name:           "data de nascimento valida (18 anos exatos)",
			dataNascimento: time.Now().AddDate(-18, 0, 0),
			wantErr:        nil,
		},
		{
			name:           "data de nascimento valida (30 anos)",
			dataNascimento: time.Now().AddDate(-30, 0, 0),
			wantErr:        nil,
		},
		{
			name:           "data de nascimento vazia",
			dataNascimento: time.Time{},
			wantErr:        dominio.ErrDataNascimentoVazia,
		},
		{
			name:           "profissional menor de idade (17 anos)",
			dataNascimento: time.Now().AddDate(-17, 0, 0),
			wantErr:        dominio.ErrProfissionalMenorDeIdade,
		},
		{
			name:           "profissional menor de idade (1 dia antes de 18 anos)",
			dataNascimento: time.Now().AddDate(-18, 0, 1),
			wantErr:        dominio.ErrProfissionalMenorDeIdade,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dominio.Profissional{DataNascimento: tt.dataNascimento}
			err := p.ValidarDataNascimento()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestProfissional_Validar(t *testing.T) {
	dataNascimentoValida := time.Now().AddDate(-25, 0, 0)

	tests := []struct {
		name         string
		profissional dominio.Profissional
		wantErr      error
	}{
		{
			name: "profissional valido",
			profissional: dominio.Profissional{
				Usuario: dominio.Usuario{
					Nome:  "Dr. João Silva",
					Email: "joao@example.com",
				},
				DataNascimento:       dataNascimentoValida,
				Especialidade:        "Psicologia Clínica",
				RegistroProfissional: "CRP12345",
			},
			wantErr: nil,
		},
		{
			name: "profissional com email invalido",
			profissional: dominio.Profissional{
				Usuario: dominio.Usuario{
					Nome:  "Dr. João Silva",
					Email: "email-invalido",
				},
				DataNascimento:       dataNascimentoValida,
				Especialidade:        "Psicologia Clínica",
				RegistroProfissional: "CRP12345",
			},
			wantErr: dominio.ErrEmailInvalido,
		},
		{
			name: "profissional com registro vazio",
			profissional: dominio.Profissional{
				Usuario: dominio.Usuario{
					Nome:  "Dr. João Silva",
					Email: "joao@example.com",
				},
				DataNascimento:       dataNascimentoValida,
				Especialidade:        "Psicologia Clínica",
				RegistroProfissional: "",
			},
			wantErr: dominio.ErrRegistroProfissionalVazio,
		},
		{
			name: "profissional com especialidade vazia",
			profissional: dominio.Profissional{
				Usuario: dominio.Usuario{
					Nome:  "Dr. João Silva",
					Email: "joao@example.com",
				},
				DataNascimento:       dataNascimentoValida,
				Especialidade:        "",
				RegistroProfissional: "CRP12345",
			},
			wantErr: dominio.ErrEspecialidadeVazia,
		},
		{
			name: "profissional menor de idade",
			profissional: dominio.Profissional{
				Usuario: dominio.Usuario{
					Nome:  "Dr. João Silva",
					Email: "joao@example.com",
				},
				DataNascimento:       time.Now().AddDate(-17, 0, 0),
				Especialidade:        "Psicologia Clínica",
				RegistroProfissional: "CRP12345",
			},
			wantErr: dominio.ErrProfissionalMenorDeIdade,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.profissional.Validar()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestProfissional_PossuiPaciente(t *testing.T) {
	paciente1 := dominio.Paciente{ID: 1}
	paciente2 := dominio.Paciente{ID: 2}

	profissional := dominio.Profissional{
		Pacientes: []dominio.Paciente{paciente1, paciente2},
	}

	tests := []struct {
		name       string
		pacienteID uint
		want       bool
	}{
		{
			name:       "possui paciente com ID 1",
			pacienteID: 1,
			want:       true,
		},
		{
			name:       "possui paciente com ID 2",
			pacienteID: 2,
			want:       true,
		},
		{
			name:       "nao possui paciente com ID 3",
			pacienteID: 3,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := profissional.PossuiPaciente(tt.pacienteID)
			assert.Equal(t, tt.want, result)
		})
	}
}

// ========== Testes para Paciente ==========

func TestPaciente_ValidarDataNascimento(t *testing.T) {
	tests := []struct {
		name           string
		dataNascimento time.Time
		wantErr        error
	}{
		{
			name:           "data de nascimento valida (hoje)",
			dataNascimento: time.Now(),
			wantErr:        nil,
		},
		{
			name:           "data de nascimento valida (10 anos atras)",
			dataNascimento: time.Now().AddDate(-10, 0, 0),
			wantErr:        nil,
		},
		{
			name:           "data de nascimento vazia",
			dataNascimento: time.Time{},
			wantErr:        dominio.ErrDataNascimentoPacienteVazia,
		},
		{
			name:           "data de nascimento no futuro",
			dataNascimento: time.Now().AddDate(1, 0, 0),
			wantErr:        dominio.ErrDataNascimentoPacienteNoFuturo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dominio.Paciente{DataNascimento: tt.dataNascimento}
			err := p.ValidarDataNascimento()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestPaciente_ValidarResponsavel(t *testing.T) {
	tests := []struct {
		name               string
		dependente         bool
		nomeResponsavel    string
		contatoResponsavel string
		wantErr            error
	}{
		{
			name:               "paciente nao dependente sem responsavel",
			dependente:         false,
			nomeResponsavel:    "",
			contatoResponsavel: "",
			wantErr:            nil,
		},
		{
			name:               "paciente dependente com responsavel valido",
			dependente:         true,
			nomeResponsavel:    "Maria Silva",
			contatoResponsavel: "11987654321",
			wantErr:            nil,
		},
		{
			name:               "paciente dependente sem nome do responsavel",
			dependente:         true,
			nomeResponsavel:    "",
			contatoResponsavel: "11987654321",
			wantErr:            dominio.ErrResponsavelVazio,
		},
		{
			name:               "paciente dependente sem contato do responsavel",
			dependente:         true,
			nomeResponsavel:    "Maria Silva",
			contatoResponsavel: "",
			wantErr:            dominio.ErrContatoResponsavelVazio,
		},
		{
			name:               "paciente dependente com contato do responsavel invalido",
			dependente:         true,
			nomeResponsavel:    "Maria Silva",
			contatoResponsavel: "119876543",
			wantErr:            dominio.ErrContatoResponsavelInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dominio.Paciente{
				Dependente:         tt.dependente,
				NomeResponsavel:    tt.nomeResponsavel,
				ContatoResponsavel: tt.contatoResponsavel,
			}
			err := p.ValidarResponsavel()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestPaciente_ValidarDataInicioTratamento(t *testing.T) {
	dataNascimento := time.Now().AddDate(-20, 0, 0)
	dataValida := time.Now().AddDate(-1, 0, 0)
	dataFutura := time.Now().AddDate(1, 0, 0)
	dataAnteriorNascimento := dataNascimento.AddDate(-1, 0, 0)

	tests := []struct {
		name                 string
		dataNascimento       time.Time
		dataInicioTratamento *time.Time
		wantErr              error
	}{
		{
			name:                 "sem data de inicio de tratamento",
			dataNascimento:       dataNascimento,
			dataInicioTratamento: nil,
			wantErr:              nil,
		},
		{
			name:                 "data de inicio de tratamento valida",
			dataNascimento:       dataNascimento,
			dataInicioTratamento: &dataValida,
			wantErr:              nil,
		},
		{
			name:                 "data de inicio de tratamento no futuro",
			dataNascimento:       dataNascimento,
			dataInicioTratamento: &dataFutura,
			wantErr:              dominio.ErrDataInicioTratamentoNoFuturo,
		},
		{
			name:                 "data de inicio de tratamento anterior ao nascimento",
			dataNascimento:       dataNascimento,
			dataInicioTratamento: &dataAnteriorNascimento,
			wantErr:              dominio.ErrDataInicioTratamentoAnteriorNascimento,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dominio.Paciente{
				DataNascimento:       tt.dataNascimento,
				DataInicioTratamento: tt.dataInicioTratamento,
			}
			err := p.ValidarDataInicioTratamento()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestPaciente_Validar(t *testing.T) {
	dataNascimentoValida := time.Now().AddDate(-20, 0, 0)
	dataInicioValida := time.Now().AddDate(-1, 0, 0)

	tests := []struct {
		name     string
		paciente dominio.Paciente
		wantErr  error
	}{
		{
			name: "paciente valido nao dependente",
			paciente: dominio.Paciente{
				Usuario: dominio.Usuario{
					Nome:  "João Silva",
					Email: "joao@example.com",
				},
				DataNascimento:       dataNascimentoValida,
				Dependente:           false,
				DataInicioTratamento: &dataInicioValida,
			},
			wantErr: nil,
		},
		{
			name: "paciente valido dependente",
			paciente: dominio.Paciente{
				Usuario: dominio.Usuario{
					Nome:  "Maria Silva",
					Email: "maria@example.com",
				},
				DataNascimento:       dataNascimentoValida,
				Dependente:           true,
				NomeResponsavel:      "José Silva",
				ContatoResponsavel:   "11987654321",
				DataInicioTratamento: &dataInicioValida,
			},
			wantErr: nil,
		},
		{
			name: "paciente com email invalido",
			paciente: dominio.Paciente{
				Usuario: dominio.Usuario{
					Nome:  "João Silva",
					Email: "email-invalido",
				},
				DataNascimento: dataNascimentoValida,
			},
			wantErr: dominio.ErrEmailInvalido,
		},
		{
			name: "paciente com data de nascimento vazia",
			paciente: dominio.Paciente{
				Usuario: dominio.Usuario{
					Nome:  "João Silva",
					Email: "joao@example.com",
				},
				DataNascimento: time.Time{},
			},
			wantErr: dominio.ErrDataNascimentoPacienteVazia,
		},
		{
			name: "paciente dependente sem responsavel",
			paciente: dominio.Paciente{
				Usuario: dominio.Usuario{
					Nome:  "Maria Silva",
					Email: "maria@example.com",
				},
				DataNascimento: dataNascimentoValida,
				Dependente:     true,
			},
			wantErr: dominio.ErrResponsavelVazio,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.paciente.Validar()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestPaciente_PossuiProfissional(t *testing.T) {
	prof1 := dominio.Profissional{ID: 1}
	prof2 := dominio.Profissional{ID: 2}

	paciente := dominio.Paciente{
		Profissionais: []dominio.Profissional{prof1, prof2},
	}

	tests := []struct {
		name           string
		profissionalID uint
		want           bool
	}{
		{
			name:           "possui profissional com ID 1",
			profissionalID: 1,
			want:           true,
		},
		{
			name:           "possui profissional com ID 2",
			profissionalID: 2,
			want:           true,
		},
		{
			name:           "nao possui profissional com ID 3",
			profissionalID: 3,
			want:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := paciente.PossuiProfissional(tt.profissionalID)
			assert.Equal(t, tt.want, result)
		})
	}
}
