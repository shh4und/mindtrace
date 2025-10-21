package dtos

import "time"

// CriarRegistroHumorDTO representa os dados para criar um registro de humor
type CriarRegistroHumorDTOin struct {
	UsuarioID        uint
	NivelHumor       int16
	HorasSono        int16
	NivelStress      int16
	NivelEnergia     int16
	AutoCuidado      []string
	Observacoes      string
	DataHoraRegistro time.Time
}

// PontoDeDadosDTOout representa um ponto de dados para graficos
type PontoDeDadosDTOout struct {
	Data     time.Time `json:"data"`
	Valor    int16     `json:"valor"`
	Humor    int16     `json:"humor"`
	Anotacao string    `json:"anotacao,omitempty"`
}

// RelatorioPacienteDTO representa o relatorio de um paciente
type RelatorioPacienteDTOout struct {
	GraficoSono    []PontoDeDadosDTOout `json:"grafico_sono"`
	GraficoEnergia []PontoDeDadosDTOout `json:"grafico_energia"`
	GraficoStress  []PontoDeDadosDTOout `json:"grafico_stress"`
	MediaSono      float64              `json:"media_sono"`
	MediaEnergia   float64              `json:"media_energia"`
	MediaStress    float64              `json:"media_stress"`
}

// ResumoPacienteDTOout representa o resumo de um paciente <=> ultimo registro
type ResumoPacienteDTOout struct {
	Data     time.Time `json:"data"`
	Humor    int16     `json:"humor"`
	Anotacao string    `json:"anotacao,omitempty"`
}

// RegistrarProfissionalDTOin representa os dados para registrar um profissional
type RegistrarProfissionalDTOin struct {
	Nome                 string
	Email                string
	Senha                string
	DataNascimento       time.Time
	Especialidade        string
	RegistroProfissional string
	CPF                  string
	Contato              string
}

// RegistrarPacienteDTOin representa os dados para registrar um paciente
type RegistrarPacienteDTOin struct {
	Nome                 string
	Email                string
	Senha                string
	Dependente           bool
	DataNascimento       time.Time
	DataInicioTratamento *time.Time
	HistoricoSaude       string
	CPF                  string
	NomeResponsavel      string
	ContatoResponsavel   string
	Contato              string
}

// AtualizarPerfilDTOin representa os dados para atualizar o perfil do usuario
type AtualizarPerfilDTOin struct {
	Nome    string `json:"nome" binding:"required"`
	Contato string `json:"contato"`
	Bio     string `json:"bio"`
	// Campos para Profissional
	Especialidade        string `json:"especialidade,omitempty"`
	RegistroProfissional string `json:"registro_profissional,omitempty"`
	// IdadeProfissional    *int8  `json:"idade_profissional,omitempty"` // Se aplicavel
	// Campos para Paciente
	DataNascimento     *time.Time `json:"data_nascimento,omitempty"`
	Dependente         *bool      `json:"dependente,omitempty"`
	NomeResponsavel    string     `json:"nome_responsavel,omitempty"`
	ContatoResponsavel string     `json:"contato_responsavel,omitempty"`
}

// AlterarSenhaDTOin representa os dados para alterar a senha
type AlterarSenhaDTOin struct {
	SenhaAtual  string `json:"senha_atual" binding:"required"`
	NovaSenha   string `json:"nova_senha" binding:"required,min=8"`
	NovaSenhaRe string `json:"nova_senha_re" binding:"required,min=8"`
}
