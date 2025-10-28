package dominio

import (
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
			wantErr: ErrEmailInvalido,
		},
		{
			name:    "email invalido sem dominio",
			email:   "usuario@",
			wantErr: ErrEmailInvalido,
		},
		{
			name:    "email invalido sem usuario",
			email:   "@example.com",
			wantErr: ErrEmailInvalido,
		},
		{
			name:    "email vazio",
			email:   "",
			wantErr: ErrEmailInvalido,
		},
		{
			name:    "email com espacos",
			email:   "user name@example.com",
			wantErr: ErrEmailInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usuario{Email: tt.email}
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
			wantErr:    ErrSenhaFraca,
		},
		{
			name:       "senha vazia",
			senhaPlana: "",
			wantErr:    ErrSenhaFraca,
		},
		{
			name:       "senha com 7 caracteres",
			senhaPlana: "Senha12",
			wantErr:    ErrSenhaFraca,
		},
		{
			name:       "senha com 8 caracteres exatos",
			senhaPlana: "Senha123",
			wantErr:    ErrSenhaInvalida,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usuario{}
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
			wantErr: ErrNomeVazio,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usuario{Nome: tt.nome}
			err := u.ValidarNome()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUsuario_Validar(t *testing.T) {
	tests := []struct {
		name    string
		usuario Usuario
		wantErr error
	}{
		{
			name: "usuario valido",
			usuario: Usuario{
				Nome:  "João Silva",
				Email: "joao@example.com",
			},
			wantErr: nil,
		},
		{
			name: "usuario com email invalido",
			usuario: Usuario{
				Nome:  "João Silva",
				Email: "email-invalido",
			},
			wantErr: ErrEmailInvalido,
		},
		{
			name: "usuario com nome vazio",
			usuario: Usuario{
				Nome:  "",
				Email: "joao@example.com",
			},
			wantErr: ErrNomeVazio,
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
			wantErr:  ErrRegistroProfissionalVazio,
		},
		{
			name:     "registro com menos de 4 caracteres",
			registro: "123",
			wantErr:  ErrRegistroProfissionalInvalido,
		},
		{
			name:     "registro com mais de 12 caracteres",
			registro: "1234567890123",
			wantErr:  ErrRegistroProfissionalInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profissional{RegistroProfissional: tt.registro}
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
			wantErr:       ErrEspecialidadeVazia,
		},
		{
			name:          "especialidade com menos de 3 caracteres",
			especialidade: "Ps",
			wantErr:       ErrEspecialidadeInvalida,
		},
		{
			name:          "especialidade com mais de 255 caracteres",
			especialidade: string(make([]byte, 256)),
			wantErr:       ErrEspecialidadeInvalida,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profissional{Especialidade: tt.especialidade}
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
			wantErr:        ErrDataNascimentoVazia,
		},
		{
			name:           "profissional menor de idade (17 anos)",
			dataNascimento: time.Now().AddDate(-17, 0, 0),
			wantErr:        ErrProfissionalMenorDeIdade,
		},
		{
			name:           "profissional menor de idade (1 dia antes de 18 anos)",
			dataNascimento: time.Now().AddDate(-18, 0, 1),
			wantErr:        ErrProfissionalMenorDeIdade,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profissional{DataNascimento: tt.dataNascimento}
			err := p.ValidarDataNascimento()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestProfissional_Validar(t *testing.T) {
	dataNascimentoValida := time.Now().AddDate(-25, 0, 0)

	tests := []struct {
		name         string
		profissional Profissional
		wantErr      error
	}{
		{
			name: "profissional valido",
			profissional: Profissional{
				Usuario: Usuario{
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
			profissional: Profissional{
				Usuario: Usuario{
					Nome:  "Dr. João Silva",
					Email: "email-invalido",
				},
				DataNascimento:       dataNascimentoValida,
				Especialidade:        "Psicologia Clínica",
				RegistroProfissional: "CRP12345",
			},
			wantErr: ErrEmailInvalido,
		},
		{
			name: "profissional com registro vazio",
			profissional: Profissional{
				Usuario: Usuario{
					Nome:  "Dr. João Silva",
					Email: "joao@example.com",
				},
				DataNascimento:       dataNascimentoValida,
				Especialidade:        "Psicologia Clínica",
				RegistroProfissional: "",
			},
			wantErr: ErrRegistroProfissionalVazio,
		},
		{
			name: "profissional com especialidade vazia",
			profissional: Profissional{
				Usuario: Usuario{
					Nome:  "Dr. João Silva",
					Email: "joao@example.com",
				},
				DataNascimento:       dataNascimentoValida,
				Especialidade:        "",
				RegistroProfissional: "CRP12345",
			},
			wantErr: ErrEspecialidadeVazia,
		},
		{
			name: "profissional menor de idade",
			profissional: Profissional{
				Usuario: Usuario{
					Nome:  "Dr. João Silva",
					Email: "joao@example.com",
				},
				DataNascimento:       time.Now().AddDate(-17, 0, 0),
				Especialidade:        "Psicologia Clínica",
				RegistroProfissional: "CRP12345",
			},
			wantErr: ErrProfissionalMenorDeIdade,
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
	paciente1 := Paciente{ID: 1}
	paciente2 := Paciente{ID: 2}

	profissional := Profissional{
		Pacientes: []Paciente{paciente1, paciente2},
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
			wantErr:        ErrDataNascimentoPacienteVazia,
		},
		{
			name:           "data de nascimento no futuro",
			dataNascimento: time.Now().AddDate(1, 0, 0),
			wantErr:        ErrDataNascimentoPacienteNoFuturo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Paciente{DataNascimento: tt.dataNascimento}
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
			wantErr:            ErrResponsavelVazio,
		},
		{
			name:               "paciente dependente sem contato do responsavel",
			dependente:         true,
			nomeResponsavel:    "Maria Silva",
			contatoResponsavel: "",
			wantErr:            ErrContatoResponsavelVazio,
		},
		{
			name:               "paciente dependente com contato do responsavel invalido",
			dependente:         true,
			nomeResponsavel:    "Maria Silva",
			contatoResponsavel: "119876543",
			wantErr:            ErrContatoResponsavelInvalido,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Paciente{
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
			wantErr:              ErrDataInicioTratamentoNoFuturo,
		},
		{
			name:                 "data de inicio de tratamento anterior ao nascimento",
			dataNascimento:       dataNascimento,
			dataInicioTratamento: &dataAnteriorNascimento,
			wantErr:              ErrDataInicioTratamentoAnteriorNascimento,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Paciente{
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
		paciente Paciente
		wantErr  error
	}{
		{
			name: "paciente valido nao dependente",
			paciente: Paciente{
				Usuario: Usuario{
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
			paciente: Paciente{
				Usuario: Usuario{
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
			paciente: Paciente{
				Usuario: Usuario{
					Nome:  "João Silva",
					Email: "email-invalido",
				},
				DataNascimento: dataNascimentoValida,
			},
			wantErr: ErrEmailInvalido,
		},
		{
			name: "paciente com data de nascimento vazia",
			paciente: Paciente{
				Usuario: Usuario{
					Nome:  "João Silva",
					Email: "joao@example.com",
				},
				DataNascimento: time.Time{},
			},
			wantErr: ErrDataNascimentoPacienteVazia,
		},
		{
			name: "paciente dependente sem responsavel",
			paciente: Paciente{
				Usuario: Usuario{
					Nome:  "Maria Silva",
					Email: "maria@example.com",
				},
				DataNascimento: dataNascimentoValida,
				Dependente:     true,
			},
			wantErr: ErrResponsavelVazio,
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
	prof1 := Profissional{ID: 1}
	prof2 := Profissional{ID: 2}

	paciente := Paciente{
		Profissionais: []Profissional{prof1, prof2},
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
