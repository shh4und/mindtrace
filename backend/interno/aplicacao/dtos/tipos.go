package dtos

import "time"

type RegistrarUsuarioDTOIn struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required,min=8"`
	Nome  string `json:"nome" binding:"required"`
}

// CriarRegistroHumorDTOOut representa os dados para criar um registro de humor
type CriarRegistroHumorDTOIn struct {
	NivelHumor       int16     `json:"nivel_humor" binding:"required,min=1,max=5"`
	HorasSono        int16     `json:"horas_sono" binding:"required,min=0,max=12"`
	NivelStress      int16     `json:"nivel_stress" binding:"required,min=1,max=10"`
	NivelEnergia     int16     `json:"nivel_energia" binding:"required,min=1,max=10"`
	AutoCuidado      string    `json:"auto_cuidado" binding:"required"`
	Observacoes      string    `json:"observacoes"`
	DataHoraRegistro time.Time `json:"data_hora_registro"`
}

// RegistrarProfissionalDTOIn representa os dados para criar um profissional
type RegistrarProfissionalDTOIn struct {
	Nome                 string    `json:"nome" binding:"required"`
	Email                string    `json:"email" binding:"required,email"`
	Senha                string    `json:"senha" binding:"required,min=8"`
	DataNascimento       time.Time `json:"data_nascimento" binding:"required"`
	Especialidade        string    `json:"especialidade" binding:"required"`
	RegistroProfissional string    `json:"registro_profissional" binding:"required"`
	CPF                  string    `json:"cpf" binding:"required"`
	Contato              string    `json:"contato"`
}

// RegistrarPacienteDTOIn representa os dados para criar um paciente
type RegistrarPacienteDTOIn struct {
	Nome                 string     `json:"nome" binding:"required"`
	Email                string     `json:"email" binding:"required,email"`
	Senha                string     `json:"senha" binding:"required,min=8"`
	Dependente           bool       `json:"dependente"`
	DataNascimento       time.Time  `json:"data_nascimento" binding:"required"`
	DataInicioTratamento *time.Time `json:"data_inicio_tratamento,omitempty"`
	HistoricoSaude       string     `json:"historico_saude,omitempty"`
	CPF                  string     `json:"cpf" binding:"required"`
	NomeResponsavel      string     `json:"nome_responsavel,omitempty"`
	ContatoResponsavel   string     `json:"contato_responsavel,omitempty"`
	Contato              string     `json:"contato"`
}

// AtualizarPerfilDTOIn representa os dados para atualizar o perfil do usuario
type AtualizarPerfilDTOIn struct {
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

// AlterarSenhaDTOIn representa os dados para alterar a senha
type AlterarSenhaDTOIn struct {
	SenhaAtual  string `json:"senha_atual" binding:"required"`
	NovaSenha   string `json:"nova_senha" binding:"required,min=8"`
	NovaSenhaRe string `json:"nova_senha_re" binding:"required,min=8"`
}

// PontoDeDadosDTOOut representa um ponto de dados para graficos
type PontoDeDadosDTOOut struct {
	Data     time.Time `json:"data"`
	Valor    int16     `json:"valor"`
	Humor    int16     `json:"humor"`
	Anotacao string    `json:"anotacao,omitempty"`
}

// RelatorioPacienteDTOOut representa o relatorio de um paciente
type RelatorioPacienteDTOOut struct {
	GraficoSono    []PontoDeDadosDTOOut `json:"grafico_sono"`
	GraficoEnergia []PontoDeDadosDTOOut `json:"grafico_energia"`
	GraficoStress  []PontoDeDadosDTOOut `json:"grafico_stress"`
	MediaSono      float64              `json:"media_sono"`
	MediaEnergia   float64              `json:"media_energia"`
	MediaStress    float64              `json:"media_stress"`
}

// ResumoPacienteDTOOut representa o resumo de um paciente <=> ultimo registro
type ResumoPacienteDTOOut struct {
	Data     time.Time `json:"data"`
	Humor    int16     `json:"humor"`
	Anotacao string    `json:"anotacao,omitempty"`
}

// RegistroHumorDTOOut representa um registro de humor completo
type RegistroHumorDTOOut struct {
	ID               uint      `json:"id"`
	PacienteID       uint      `json:"paciente_id"`
	NivelHumor       int16     `json:"nivel_humor"`
	HorasSono        int16     `json:"horas_sono"`
	NivelEnergia     int16     `json:"nivel_energia"`
	NivelStress      int16     `json:"nivel_stress"`
	AutoCuidado      string    `json:"auto_cuidado"`
	Observacoes      string    `json:"observacoes,omitempty"`
	DataHoraRegistro time.Time `json:"data_hora_registro"`
	CreatedAt        time.Time `json:"created_at"`
}

type UsuarioDTOOut struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Nome      string    `json:"nome"`
	Contato   string    `json:"contato,omitempty"`
	Bio       string    `json:"bio,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProfissionalDTOOut struct {
	ID                   uint          `json:"id"`
	Usuario              UsuarioDTOOut `json:"usuario"`
	DataNascimento       time.Time     `json:"data_nascimento"`
	Especialidade        string        `json:"especialidade"`
	RegistroProfissional string        `json:"registro_profissional"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
}

type PacienteDTOOut struct {
	ID             uint                 `json:"id"`
	Usuario        UsuarioDTOOut        `json:"usuario"`
	DataNascimento time.Time            `json:"data_nascimento"`
	Dependente     bool                 `json:"dependente"`
	Profissionais  []ProfissionalDTOOut `json:"profissionais,omitempty"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
}
