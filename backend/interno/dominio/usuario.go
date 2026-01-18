package dominio

import (
	"errors"
	"regexp"
	"time"

	"gorm.io/gorm"
)

// Constantes para tipos de usuario
const (
	TipoUsuarioProfissional uint8 = 2
	TipoUsuarioPaciente     uint8 = 3
)

var (
	ErrEmailJaCadastrado     = errors.New("e-mail existente")
	ErrCrendenciaisInvalidas = errors.New("credenciais invalidas")
	ErrUsuarioNaoEncontrado  = errors.New("usuario nao encontrado")
	ErrSenhaNaoConfere       = errors.New("a nova senha e a senha de confirmacao nao conferem")
	ErrEmailInvalido         = errors.New("email invalido")
	ErrSenhaFraca            = errors.New("senha deve ter no minimo 8 caracteres")
	ErrSenhaInvalida         = errors.New("senha com caracteres invalidos")
	ErrNomeVazio             = errors.New("nome nao pode estar vazio")
)

// Usuario e a base para todos os tipos de usuarios.
type Usuario struct {
	ID             uint    `gorm:"primaryKey"`
	TipoUsuario    uint8   `gorm:"type:smallint;not null;check:tipo_usuario >= 1"`
	Nome           string  `gorm:"type:varchar(255);not null"`
	Email          string  `gorm:"type:varchar(255);unique;not null"`
	Senha          string  `gorm:"type:text;not null"`
	Contato        string  `gorm:"type:varchar(11)"`
	Bio            string  `gorm:"type:text"`
	CPF            string  `gorm:"type:varchar(11);unique"`
	EstaAtivo      bool    `gorm:"type:boolean;default:false"`
	EmailVerifHash *string `gorm:"type:varchar(64)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

// TipoUsuarioParaString converte o tipo numerico para string
func TipoUsuarioParaString(tipo uint8) string {
	switch tipo {
	case TipoUsuarioProfissional:
		return "profissional"
	case TipoUsuarioPaciente:
		return "paciente"
	default:
		return "desconhecido"
	}
}

// StringParaTipoUsuario converte string para tipo numerico
func StringParaTipoUsuario(tipo string) uint8 {
	switch tipo {
	case "profissional":
		return TipoUsuarioProfissional
	case "paciente":
		return TipoUsuarioPaciente
	default:
		return 0
	}
}

// Metodos de validacao - LOGICA DE NEGOCIO
func (u *Usuario) ValidarEmail() error {
	regex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !regex.MatchString(u.Email) {
		return ErrEmailInvalido
	}
	return nil
}

func (u *Usuario) ValidarSenha(senhaPlana string) error {
	if len(senhaPlana) < 8 {
		return ErrSenhaFraca
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*].{8,}$`)
	if !regex.MatchString(senhaPlana) {
		return ErrSenhaInvalida
	}
	return nil
}

func (u *Usuario) ValidarNome() error {
	if u.Nome == "" {
		return ErrNomeVazio
	}
	return nil
}

// Validacao completa
func (u *Usuario) Validar() error {
	if err := u.ValidarEmail(); err != nil {
		return err
	}
	if err := u.ValidarNome(); err != nil {
		return err
	}
	return nil
}

// Profissional tem seus proprios dados e uma referencia ao Usuario.
type Profissional struct {
	ID                   uint    `gorm:"primaryKey"`
	UsuarioID            uint    `gorm:"unique;not null"`
	Usuario              Usuario `gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE"`
	DataNascimento       time.Time
	Especialidade        string     `gorm:"type:varchar(255);not null"`
	RegistroProfissional string     `gorm:"type:varchar(12);unique;not null"`
	Pacientes            []Paciente `gorm:"many2many:profissional_paciente;constraint:OnDelete:CASCADE;"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func (Profissional) TableName() string {
	return "profissionais"
}

// Erros de validacao - Profissional
var (
	ErrRegistroProfissionalVazio    = errors.New("registro profissional nao pode estar vazio")
	ErrRegistroProfissionalInvalido = errors.New("registro profissional deve ter entre 4 e 12 caracteres")
	ErrEspecialidadeVazia           = errors.New("especialidade nao pode estar vazia")
	ErrEspecialidadeInvalida        = errors.New("especialidade deve ter entre 3 e 255 caracteres")
	ErrDataNascimentoVazia          = errors.New("data de nascimento e obrigatoria")
	ErrProfissionalMenorDeIdade     = errors.New("profissional deve ter no minimo 18 anos")
)

// Metodos de validacao - LOGICA DE NEGOCIO (Profissional)
func (p *Profissional) ValidarRegistroProfissional() error {
	if p.RegistroProfissional == "" {
		return ErrRegistroProfissionalVazio
	}
	if len(p.RegistroProfissional) < 4 || len(p.RegistroProfissional) > 12 {
		return ErrRegistroProfissionalInvalido
	}
	return nil
}

func (p *Profissional) ValidarEspecialidade() error {
	if p.Especialidade == "" {
		return ErrEspecialidadeVazia
	}
	if len(p.Especialidade) < 3 || len(p.Especialidade) > 255 {
		return ErrEspecialidadeInvalida
	}
	return nil
}

func (p *Profissional) ValidarDataNascimento() error {
	if p.DataNascimento.IsZero() {
		return ErrDataNascimentoVazia
	}
	if p.DataNascimento.After(time.Now().AddDate(-18, 0, 0)) {
		return ErrProfissionalMenorDeIdade
	}
	return nil
}

// Validacao completa do Profissional
func (p *Profissional) Validar() error {
	if err := p.Usuario.Validar(); err != nil {
		return err
	}
	if err := p.ValidarRegistroProfissional(); err != nil {
		return err
	}
	if err := p.ValidarEspecialidade(); err != nil {
		return err
	}
	if err := p.ValidarDataNascimento(); err != nil {
		return err
	}
	return nil
}

// PossuiPaciente verifica se o profissional ja esta associado a um paciente
func (p *Profissional) PossuiPaciente(pacienteID uint) bool {
	for _, pac := range p.Pacientes {
		if pac.ID == pacienteID {
			return true
		}
	}
	return false
}

// Paciente tem seus proprios dados e uma referencia ao Usuario.
type Paciente struct {
	ID                   uint    `gorm:"primaryKey"`
	UsuarioID            uint    `gorm:"unique;not null"`
	Usuario              Usuario `gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE"`
	DataNascimento       time.Time
	Dependente           bool
	NomeResponsavel      string `gorm:"type:varchar(255)"`
	ContatoResponsavel   string `gorm:"type:varchar(11)"`
	DataInicioTratamento *time.Time
	Profissionais        []Profissional `gorm:"many2many:profissional_paciente;constraint:OnDelete:CASCADE;"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func (Paciente) TableName() string {
	return "pacientes"
}

// Erros de validacao - Paciente
var (
	ErrDataNascimentoPacienteVazia            = errors.New("data de nascimento e obrigatoria")
	ErrDataNascimentoPacienteNoFuturo         = errors.New("data de nascimento nao pode ser no futuro")
	ErrResponsavelVazio                       = errors.New("paciente dependente deve ter nome do responsavel")
	ErrContatoResponsavelVazio                = errors.New("paciente dependente deve ter contato do responsavel")
	ErrContatoResponsavelInvalido             = errors.New("contato do responsavel invalido")
	ErrDataInicioTratamentoNoFuturo           = errors.New("data de inicio do tratamento nao pode ser no futuro")
	ErrDataInicioTratamentoAnteriorNascimento = errors.New("data de inicio do tratamento nao pode ser anterior a data de nascimento")
)

// Metodos de validacao - LOGICA DE NEGOCIO (Paciente)
func (pc *Paciente) ValidarDataNascimento() error {
	if pc.DataNascimento.IsZero() {
		return ErrDataNascimentoPacienteVazia
	}
	if pc.DataNascimento.After(time.Now()) {
		return ErrDataNascimentoPacienteNoFuturo
	}
	return nil
}

func (pc *Paciente) ValidarResponsavel() error {
	if pc.Dependente && pc.NomeResponsavel == "" {
		return ErrResponsavelVazio
	}
	if pc.Dependente && pc.ContatoResponsavel == "" {
		return ErrContatoResponsavelVazio
	}
	if pc.Dependente && len(pc.ContatoResponsavel) < 10 {
		return ErrContatoResponsavelInvalido
	}
	return nil
}

func (pc *Paciente) ValidarDataInicioTratamento() error {
	if pc.DataInicioTratamento != nil {
		if pc.DataInicioTratamento.After(time.Now()) {
			return ErrDataInicioTratamentoNoFuturo
		}
		if pc.DataInicioTratamento.Before(pc.DataNascimento) {
			return ErrDataInicioTratamentoAnteriorNascimento
		}
	}
	return nil
}

// Validacao completa do Paciente
func (pc *Paciente) Validar() error {
	if err := pc.Usuario.Validar(); err != nil {
		return err
	}
	if err := pc.ValidarDataNascimento(); err != nil {
		return err
	}
	if err := pc.ValidarResponsavel(); err != nil {
		return err
	}
	if err := pc.ValidarDataInicioTratamento(); err != nil {
		return err
	}
	return nil
}

// PossuiProfissional verifica se o paciente ja esta associado a um profissional
func (pc *Paciente) PossuiProfissional(profissionalID uint) bool {
	for _, prof := range pc.Profissionais {
		if prof.ID == profissionalID {
			return true
		}
	}
	return false
}
