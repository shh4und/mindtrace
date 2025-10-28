package tests

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/mappers"
	"mindtrace/backend/interno/dominio"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ========== Testes para Mappers de Saída ==========

func TestUsuarioParaDTOOut(t *testing.T) {
	now := time.Now()
	usuario := &dominio.Usuario{
		ID:        1,
		Email:     "teste@example.com",
		Nome:      "João Silva",
		Contato:   "11987654321",
		Bio:       "Psicólogo clínico",
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := mappers.UsuarioParaDTOOut(usuario)

	assert.NotNil(t, result)
	assert.Equal(t, usuario.ID, result.ID)
	assert.Equal(t, usuario.Email, result.Email)
	assert.Equal(t, usuario.Nome, result.Nome)
	assert.Equal(t, usuario.Contato, result.Contato)
	assert.Equal(t, usuario.Bio, result.Bio)
	assert.Equal(t, usuario.CreatedAt, result.CreatedAt)
	assert.Equal(t, usuario.UpdatedAt, result.UpdatedAt)
}

func TestProfissionalParaDTOOut_ComDadosCompletos(t *testing.T) {
	now := time.Now()
	dataNascimento := time.Now().AddDate(-30, 0, 0)

	profissional := &dominio.Profissional{
		ID: 1,
		Usuario: dominio.Usuario{
			ID:        1,
			Email:     "prof@example.com",
			Nome:      "Dr. João Silva",
			Contato:   "11987654321",
			Bio:       "Psicólogo clínico",
			CreatedAt: now,
			UpdatedAt: now,
		},
		DataNascimento:       dataNascimento,
		Especialidade:        "Psicologia Clínica",
		RegistroProfissional: "CRP12345",
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	result := mappers.ProfissionalParaDTOOut(profissional)

	assert.NotNil(t, result)
	assert.Equal(t, profissional.ID, result.ID)
	assert.Equal(t, profissional.Usuario.Nome, result.Usuario.Nome)
	assert.Equal(t, profissional.Usuario.Email, result.Usuario.Email)
	assert.Equal(t, profissional.DataNascimento, result.DataNascimento)
	assert.Equal(t, profissional.Especialidade, result.Especialidade)
	assert.Equal(t, profissional.RegistroProfissional, result.RegistroProfissional)
	assert.Equal(t, profissional.CreatedAt, result.CreatedAt)
	assert.Equal(t, profissional.UpdatedAt, result.UpdatedAt)
}

func TestProfissionalParaDTOOut_ComNil(t *testing.T) {
	result := mappers.ProfissionalParaDTOOut(nil)
	assert.Nil(t, result)
}

func TestPacienteParaDTOOut_ComDadosCompletos(t *testing.T) {
	now := time.Now()
	dataNascimento := time.Now().AddDate(-25, 0, 0)

	paciente := &dominio.Paciente{
		ID: 1,
		Usuario: dominio.Usuario{
			ID:        2,
			Email:     "paciente@example.com",
			Nome:      "Maria Silva",
			Contato:   "11987654321",
			CreatedAt: now,
			UpdatedAt: now,
		},
		DataNascimento: dataNascimento,
		Dependente:     false,
		Profissionais: []dominio.Profissional{
			{
				ID: 1,
				Usuario: dominio.Usuario{
					ID:    1,
					Nome:  "Dr. João",
					Email: "joao@example.com",
				},
				Especialidade:        "Psicologia",
				RegistroProfissional: "CRP123",
				DataNascimento:       time.Now().AddDate(-30, 0, 0),
			},
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := mappers.PacienteParaDTOOut(paciente)

	assert.NotNil(t, result)
	assert.Equal(t, paciente.ID, result.ID)
	assert.Equal(t, paciente.Usuario.Nome, result.Usuario.Nome)
	assert.Equal(t, paciente.Usuario.Email, result.Usuario.Email)
	assert.Equal(t, paciente.DataNascimento, result.DataNascimento)
	assert.NotNil(t, result.Dependente)
	assert.False(t, *result.Dependente)
	assert.Len(t, result.Profissionais, 1)
	assert.Equal(t, "Dr. João", result.Profissionais[0].Usuario.Nome)
}

func TestPacienteParaDTOOut_Dependente(t *testing.T) {
	now := time.Now()
	dataNascimento := time.Now().AddDate(-10, 0, 0)

	paciente := &dominio.Paciente{
		ID: 1,
		Usuario: dominio.Usuario{
			ID:        2,
			Email:     "crianca@example.com",
			Nome:      "Pedro Silva",
			CreatedAt: now,
			UpdatedAt: now,
		},
		DataNascimento:     dataNascimento,
		Dependente:         true,
		NomeResponsavel:    "José Silva",
		ContatoResponsavel: "11987654321",
		Profissionais:      []dominio.Profissional{},
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	result := mappers.PacienteParaDTOOut(paciente)

	assert.NotNil(t, result)
	assert.NotNil(t, result.Dependente)
	assert.True(t, *result.Dependente)
	assert.Empty(t, result.Profissionais)
}

func TestPacienteParaDTOOut_ComNil(t *testing.T) {
	result := mappers.PacienteParaDTOOut(nil)
	assert.Nil(t, result)
}

func TestRegistroHumorParaDTOOut(t *testing.T) {
	now := time.Now()
	registro := &dominio.RegistroHumor{
		ID:               1,
		PacienteID:       2,
		NivelHumor:       4,
		HorasSono:        8,
		NivelEnergia:     7,
		NivelStress:      3,
		AutoCuidado:      "Exercício físico",
		Observacoes:      "Dia produtivo",
		DataHoraRegistro: now,
		CreatedAt:        now,
	}

	result := mappers.RegistroHumorParaDTOOut(registro)

	assert.NotNil(t, result)
	assert.Equal(t, registro.ID, result.ID)
	assert.Equal(t, registro.PacienteID, result.PacienteID)
	assert.Equal(t, registro.NivelHumor, result.NivelHumor)
	assert.Equal(t, registro.HorasSono, result.HorasSono)
	assert.Equal(t, registro.NivelEnergia, result.NivelEnergia)
	assert.Equal(t, registro.NivelStress, result.NivelStress)
	assert.Equal(t, registro.AutoCuidado, result.AutoCuidado)
	assert.Equal(t, registro.Observacoes, result.Observacoes)
	assert.Equal(t, registro.DataHoraRegistro, result.DataHoraRegistro)
	assert.Equal(t, registro.CreatedAt, result.CreatedAt)
}

func TestResumoPacienteParaDTOOut(t *testing.T) {
	now := time.Now()
	registro := &dominio.RegistroHumor{
		ID:               1,
		PacienteID:       2,
		NivelHumor:       4,
		Observacoes:      "Sentindo-se bem",
		DataHoraRegistro: now,
	}

	result := mappers.ResumoPacienteParaDTOOut(registro)

	assert.NotNil(t, result)
	assert.Equal(t, registro.DataHoraRegistro, result.Data)
	assert.Equal(t, registro.NivelHumor, result.Humor)
	assert.Equal(t, registro.Observacoes, result.Anotacao)
}

func TestPacientesParaDTOOut(t *testing.T) {
	now := time.Now()
	pacientes := []dominio.Paciente{
		{
			ID: 1,
			Usuario: dominio.Usuario{
				ID:    1,
				Nome:  "Paciente 1",
				Email: "pac1@example.com",
			},
			DataNascimento: time.Now().AddDate(-25, 0, 0),
			Dependente:     false,
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			ID: 2,
			Usuario: dominio.Usuario{
				ID:    2,
				Nome:  "Paciente 2",
				Email: "pac2@example.com",
			},
			DataNascimento: time.Now().AddDate(-30, 0, 0),
			Dependente:     false,
			CreatedAt:      now,
			UpdatedAt:      now,
		},
	}

	result := mappers.PacientesParaDTOOut(pacientes)

	assert.Len(t, result, 2)
	assert.Equal(t, "Paciente 1", result[0].Usuario.Nome)
	assert.Equal(t, "Paciente 2", result[1].Usuario.Nome)
}

func TestPacientesParaDTOOut_ListaVazia(t *testing.T) {
	pacientes := []dominio.Paciente{}
	result := mappers.PacientesParaDTOOut(pacientes)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestProfissionaisParaDTOOut(t *testing.T) {
	now := time.Now()
	profissionais := []dominio.Profissional{
		{
			ID: 1,
			Usuario: dominio.Usuario{
				ID:    1,
				Nome:  "Dr. João",
				Email: "joao@example.com",
			},
			DataNascimento:       time.Now().AddDate(-35, 0, 0),
			Especialidade:        "Psicologia Clínica",
			RegistroProfissional: "CRP123",
			CreatedAt:            now,
			UpdatedAt:            now,
		},
		{
			ID: 2,
			Usuario: dominio.Usuario{
				ID:    2,
				Nome:  "Dra. Maria",
				Email: "maria@example.com",
			},
			DataNascimento:       time.Now().AddDate(-40, 0, 0),
			Especialidade:        "Neuropsicologia",
			RegistroProfissional: "CRP456",
			CreatedAt:            now,
			UpdatedAt:            now,
		},
	}

	result := mappers.ProfissionaisParaDTOOut(profissionais)

	assert.Len(t, result, 2)
	assert.Equal(t, "Dr. João", result[0].Usuario.Nome)
	assert.Equal(t, "Psicologia Clínica", result[0].Especialidade)
	assert.Equal(t, "Dra. Maria", result[1].Usuario.Nome)
	assert.Equal(t, "Neuropsicologia", result[1].Especialidade)
}

func TestProfissionaisParaDTOOut_ListaVazia(t *testing.T) {
	profissionais := []dominio.Profissional{}
	result := mappers.ProfissionaisParaDTOOut(profissionais)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

// ========== Testes para Mappers de Entrada ==========

func TestRegistrarUsuarioDTOInParaEntidade(t *testing.T) {
	dtoIn := &dtos.RegistrarUsuarioDTOIn{
		Email: "novo@example.com",
		Senha: "SenhaSegura123!",
		Nome:  "Novo Usuário",
	}

	result := mappers.RegistrarUsuarioDTOInParaEntidade(dtoIn)

	assert.NotNil(t, result)
	assert.Equal(t, dtoIn.Email, result.Email)
	assert.Equal(t, dtoIn.Senha, result.Senha)
	assert.Equal(t, dtoIn.Nome, result.Nome)
	assert.Equal(t, uint8(0), result.TipoUsuario) // Valor padrão
}

func TestRegistrarProfissionalDTOInParaEntidade(t *testing.T) {
	dataNascimento := time.Now().AddDate(-30, 0, 0)
	dtoIn := &dtos.RegistrarProfissionalDTOIn{
		Nome:                 "Dr. João Silva",
		Email:                "joao@example.com",
		Senha:                "SenhaSegura123!",
		DataNascimento:       dataNascimento,
		Especialidade:        "Psicologia Clínica",
		RegistroProfissional: "CRP12345",
		CPF:                  "12345678901",
		Contato:              "11987654321",
	}

	usuario, profissional := mappers.RegistrarProfissionalDTOInParaEntidade(dtoIn)

	// Validar Usuario
	assert.NotNil(t, usuario)
	assert.Equal(t, dtoIn.Nome, usuario.Nome)
	assert.Equal(t, dtoIn.Email, usuario.Email)
	assert.Equal(t, dtoIn.Senha, usuario.Senha)
	assert.Equal(t, uint8(2), usuario.TipoUsuario)
	assert.Equal(t, dtoIn.CPF, usuario.CPF)
	assert.Equal(t, dtoIn.Contato, usuario.Contato)

	// Validar Profissional
	assert.NotNil(t, profissional)
	assert.Equal(t, dtoIn.DataNascimento, profissional.DataNascimento)
	assert.Equal(t, dtoIn.Especialidade, profissional.Especialidade)
	assert.Equal(t, dtoIn.RegistroProfissional, profissional.RegistroProfissional)
}

func TestRegistrarPacienteDTOInParaEntidade(t *testing.T) {
	dataNascimento := time.Now().AddDate(-25, 0, 0)
	dataInicioTratamento := time.Now().AddDate(0, -1, 0)
	dependente := false

	dtoIn := &dtos.RegistrarPacienteDTOIn{
		Nome:                 "Maria Silva",
		Email:                "maria@example.com",
		Senha:                "SenhaSegura123!",
		Dependente:           &dependente,
		DataNascimento:       dataNascimento,
		DataInicioTratamento: &dataInicioTratamento,
		HistoricoSaude:       "Nenhum",
		CPF:                  "98765432100",
		Contato:              "11987654321",
	}

	usuario, paciente := mappers.RegistrarPacienteDTOInParaEntidade(dtoIn)

	// Validar Usuario
	assert.NotNil(t, usuario)
	assert.Equal(t, dtoIn.Nome, usuario.Nome)
	assert.Equal(t, dtoIn.Email, usuario.Email)
	assert.Equal(t, dtoIn.Senha, usuario.Senha)
	assert.Equal(t, uint8(3), usuario.TipoUsuario)
	assert.Equal(t, dtoIn.CPF, usuario.CPF)
	assert.Equal(t, dtoIn.Contato, usuario.Contato)

	// Validar Paciente
	assert.NotNil(t, paciente)
	assert.Equal(t, dtoIn.DataNascimento, paciente.DataNascimento)
	assert.Equal(t, *dtoIn.Dependente, paciente.Dependente)
	assert.Equal(t, dtoIn.DataInicioTratamento, paciente.DataInicioTratamento)
}

func TestRegistrarPacienteDTOInParaEntidade_Dependente(t *testing.T) {
	dataNascimento := time.Now().AddDate(-10, 0, 0)
	dependente := true

	dtoIn := &dtos.RegistrarPacienteDTOIn{
		Nome:               "Pedro Silva",
		Email:              "pedro@example.com",
		Senha:              "SenhaSegura123!",
		Dependente:         &dependente,
		DataNascimento:     dataNascimento,
		CPF:                "11122233344",
		NomeResponsavel:    "José Silva",
		ContatoResponsavel: "11987654321",
	}

	usuario, paciente := mappers.RegistrarPacienteDTOInParaEntidade(dtoIn)

	assert.NotNil(t, usuario)
	assert.NotNil(t, paciente)
	assert.True(t, paciente.Dependente)
	assert.Equal(t, dtoIn.NomeResponsavel, paciente.NomeResponsavel)
	assert.Equal(t, dtoIn.ContatoResponsavel, paciente.ContatoResponsavel)
}

func TestCriarRegistroHumorDTOInParaEntidade(t *testing.T) {
	now := time.Now()
	dtoIn := &dtos.CriarRegistroHumorDTOIn{
		NivelHumor:       4,
		HorasSono:        8,
		NivelStress:      3,
		NivelEnergia:     7,
		AutoCuidado:      "Exercício físico",
		Observacoes:      "Dia produtivo",
		DataHoraRegistro: now,
	}
	pacienteID := uint(5)

	result := mappers.CriarRegistroHumorDTOInParaEntidade(dtoIn, pacienteID)

	assert.NotNil(t, result)
	assert.Equal(t, pacienteID, result.PacienteID)
	assert.Equal(t, dtoIn.NivelHumor, result.NivelHumor)
	assert.Equal(t, dtoIn.HorasSono, result.HorasSono)
	assert.Equal(t, dtoIn.NivelStress, result.NivelStress)
	assert.Equal(t, dtoIn.NivelEnergia, result.NivelEnergia)
	assert.Equal(t, dtoIn.AutoCuidado, result.AutoCuidado)
	assert.Equal(t, dtoIn.Observacoes, result.Observacoes)
	assert.Equal(t, dtoIn.DataHoraRegistro, result.DataHoraRegistro)
}

func TestConviteParaDTOOut(t *testing.T) {
	now := time.Now()
	dataExpiracao := now.AddDate(0, 0, 7)

	convite := &dominio.Convite{
		Token:         "abc123def456",
		DataExpiracao: dataExpiracao,
		Usado:         false,
		CreatedAt:     now,
	}

	result := mappers.ConviteParaDTOOut(convite)

	assert.NotNil(t, result)
	assert.Equal(t, convite.Token, result.Token)
	assert.Equal(t, convite.DataExpiracao, result.DataExpiracao)
	assert.Equal(t, convite.Usado, result.Usado)
	assert.Equal(t, convite.CreatedAt, result.CreatedAt)
}

func TestConviteParaDTOOut_Usado(t *testing.T) {
	now := time.Now()
	dataExpiracao := now.AddDate(0, 0, -1) // Expirado

	convite := &dominio.Convite{
		Token:         "xyz789",
		DataExpiracao: dataExpiracao,
		Usado:         true,
		CreatedAt:     now.AddDate(0, 0, -10),
	}

	result := mappers.ConviteParaDTOOut(convite)

	assert.NotNil(t, result)
	assert.True(t, result.Usado)
	assert.Equal(t, convite.Token, result.Token)
}

func TestConviteParaDTOOut_ComNil(t *testing.T) {
	result := mappers.ConviteParaDTOOut(nil)
	assert.Nil(t, result)
}

// ========== Testes de Casos Extremos ==========

func TestPacienteParaDTOOut_SemProfissionais(t *testing.T) {
	now := time.Now()
	paciente := &dominio.Paciente{
		ID: 1,
		Usuario: dominio.Usuario{
			ID:    1,
			Nome:  "Paciente Solo",
			Email: "solo@example.com",
		},
		DataNascimento: time.Now().AddDate(-25, 0, 0),
		Dependente:     false,
		Profissionais:  []dominio.Profissional{},
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	result := mappers.PacienteParaDTOOut(paciente)

	assert.NotNil(t, result)
	assert.Empty(t, result.Profissionais)
}

func TestProfissionalParaDTOOut_SemContato(t *testing.T) {
	profissional := &dominio.Profissional{
		ID: 1,
		Usuario: dominio.Usuario{
			ID:      1,
			Nome:    "Dr. Sem Contato",
			Email:   "semcontato@example.com",
			Contato: "", // Vazio
		},
		DataNascimento:       time.Now().AddDate(-30, 0, 0),
		Especialidade:        "Psicologia",
		RegistroProfissional: "CRP999",
	}

	result := mappers.ProfissionalParaDTOOut(profissional)

	assert.NotNil(t, result)
	assert.Empty(t, result.Usuario.Contato)
}

func TestRegistroHumorParaDTOOut_SemObservacoes(t *testing.T) {
	registro := &dominio.RegistroHumor{
		ID:               1,
		PacienteID:       2,
		NivelHumor:       3,
		HorasSono:        7,
		NivelEnergia:     5,
		NivelStress:      6,
		AutoCuidado:      "Meditação",
		Observacoes:      "", // Vazio
		DataHoraRegistro: time.Now(),
	}

	result := mappers.RegistroHumorParaDTOOut(registro)

	assert.NotNil(t, result)
	assert.Empty(t, result.Observacoes)
}
