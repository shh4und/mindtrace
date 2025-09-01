package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type PacienteControlador struct {
	usuarioServico servicos.UsuarioServico
}

func NovoPacienteControlador(us servicos.UsuarioServico) *PacienteControlador {
	return &PacienteControlador{usuarioServico: us}
}

type RegistrarPacienteRequest struct {
	Nome                 string     `json:"nome" binding:"required"`
	Email                string     `json:"email" binding:"required,email"`
	Senha                string     `json:"senha" binding:"required,min=8"`
	Dependente           *bool      `json:"dependente" binding:"required"`
	Idade                int8       `json:"idade" binding:"required"`
	DataInicioTratamento *time.Time `json:"data_inicio_tratamento"`
	CPF                  string     `json:"cpf" binding:"required"`
	NomeResponsavel      string     `json:"nome_responsavel"`
	ContatoResponsavel   string     `json:"contato_responsavel"`
	Contato              string     `json:"contato"`
}

type ProprioPacienteRequest struct {
	Nome                 string     `json:"nome"`
	Email                string     `json:"email"`
	Dependente           *bool      `json:"dependente"`
	Idade                int8       `json:"idade"`
	DataInicioTratamento *time.Time `json:"data_inicio_tratamento"`
	CPF                  string     `json:"cpf"`
	NomeResponsavel      string     `json:"nome_responsavel"`
	ContatoResponsavel   string     `json:"contato_responsavel"`
	Contato              string     `json:"contato"`
}

func (pc *PacienteControlador) Registrar(c *gin.Context) {
	var req RegistrarPacienteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if len(req.Senha) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Senha inválida. Use 8 ou mais caracteres com letras, números e os símbolos: !@#$%^&*"})
		return
	}

	passwordRegex := `^[a-zA-Z0-9!@#$%^&*].{8,}$`
	if match, _ := regexp.MatchString(passwordRegex, req.Senha); !match {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Senha inválida. Use 8 ou mais caracteres com letras, números e os símbolos: !@#$%^&*"})
		return
	}

	dto := servicos.RegistrarPacienteDTO{
		Nome:                 req.Nome,
		Email:                req.Email,
		Senha:                req.Senha,
		Dependente:           *req.Dependente,
		Idade:                req.Idade,
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

	c.JSON(http.StatusCreated, paciente)
}

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

	proprioPaciente := ProprioPacienteRequest{
		Nome:                 paciente.Usuario.Nome,
		Email:                paciente.Usuario.Email,
		Dependente:           &paciente.Dependente,
		Idade:                paciente.Idade,
		DataInicioTratamento: paciente.DataInicioTratamento,
		CPF:                  paciente.Usuario.CPF,
		NomeResponsavel:      paciente.NomeResponsavel,
		ContatoResponsavel:   paciente.ContatoResponsavel,
		Contato:              paciente.Usuario.Contato,
	}

	c.JSON(http.StatusOK, proprioPaciente)
}
