package controladores

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/mappers"
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

// PacienteControlador gerencia requisicoes HTTP relacionadas ao gerenciamento de pacientes
type PacienteControlador struct {
	usuarioServico servicos.UsuarioServico
}

// NovoPacienteControlador cria uma nova instancia de PacienteControlador com o UsuarioServico fornecido
func NovoPacienteControlador(us servicos.UsuarioServico) *PacienteControlador {
	return &PacienteControlador{usuarioServico: us}
}

// RegistrarPacienteRequest representa o payload da requisicao para registrar um novo paciente
type RegistrarPacienteRequest struct {
	Nome                 string     `json:"nome" binding:"required"`
	Email                string     `json:"email" binding:"required,email"`
	Senha                string     `json:"senha" binding:"required,min=8"`
	Dependente           *bool      `json:"dependente" binding:"required"`
	DataNascimento       string     `json:"data_nascimento" binding:"required"`
	DataInicioTratamento *time.Time `json:"data_inicio_tratamento"`
	CPF                  string     `json:"cpf" binding:"required"`
	NomeResponsavel      string     `json:"nome_responsavel"`
	ContatoResponsavel   string     `json:"contato_responsavel"`
	Contato              string     `json:"contato"`
}

// ProprioPacienteRequest representa o payload da resposta para o perfil proprio do paciente
type ProprioPacienteRequest struct {
	Nome                 string     `json:"nome"`
	Email                string     `json:"email"`
	Dependente           *bool      `json:"dependente"`
	DataNascimento       time.Time  `json:"data_nascimento"`
	DataInicioTratamento *time.Time `json:"data_inicio_tratamento"`
	CPF                  string     `json:"cpf"`
	NomeResponsavel      string     `json:"nome_responsavel"`
	ContatoResponsavel   string     `json:"contato_responsavel"`
	Contato              string     `json:"contato"`
}

// Registrar lida com o registro de um novo paciente
// Valida a entrada analisa datas verifica a forca da senha e chama o servico para registrar o paciente
func (pc *PacienteControlador) Registrar(c *gin.Context) {
	var req RegistrarPacienteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	dataNascimento, err := time.Parse("2006-01-02", req.DataNascimento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Formato de data inválido. Use YYYY-MM-DD"})
		return
	}

	if len(req.Senha) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Senha inválida. Use 8 ou mais caracteres com letras, números e os símbolos: !@#$%^&*"})
		return
	}
	/********************

		Verficicar a possibilidade das checagens feitas em controladores
		serem implementadas como regras de negocios nos dominios

	*********************/
	passwordRegex := `^[a-zA-Z0-9!@#$%^&*].{8,}$`
	if match, _ := regexp.MatchString(passwordRegex, req.Senha); !match {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Senha inválida. Use 8 ou mais caracteres com letras, números e os símbolos: !@#$%^&*"})
		return
	}

	dto := dtos.RegistrarPacienteDTOIn{
		Nome:                 req.Nome,
		Email:                req.Email,
		Senha:                req.Senha,
		DataNascimento:       dataNascimento,
		Dependente:           *req.Dependente,
		DataInicioTratamento: req.DataInicioTratamento,
		CPF:                  req.CPF,
		NomeResponsavel:      req.NomeResponsavel,
		ContatoResponsavel:   req.ContatoResponsavel,
		Contato:              req.Contato,
	}

	paciente, err := pc.usuarioServico.RegistrarPaciente(dto)
	if err != nil {
		if err == servicos.ErrEmailJaCadastrado {
			c.JSON(http.StatusConflict, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}
		return
	}

	pacienteOut := mappers.PacienteParaDTOOut(paciente)

	c.JSON(http.StatusCreated, pacienteOut)
}

// ProprioPerfilPaciente recupera o perfil do paciente autenticado
// Extrai o ID do usuario do contexto e busca os dados do paciente do servico
func (uc *PacienteControlador) ProprioPerfilPaciente(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	paciente, err := uc.usuarioServico.ProprioPerfilPaciente(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	proprioPaciente := mappers.PacienteParaDTOOut(paciente)
	c.JSON(http.StatusOK, proprioPaciente)
}
