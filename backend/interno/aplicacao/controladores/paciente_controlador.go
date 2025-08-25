package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
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
	EhDependente         *bool      `json:"dependente" binding:"required"`
	Idade                int8       `json:"idade" binding:"required"`
	DataInicioTratamento *time.Time `json:"data_inicio_tratamento"`
	HistoricoSaude       string     `json:"historico_saude"`
	CPF                  string     `json:"cpf" binding:"required"`
}

func (pc *PacienteControlador) Registrar(c *gin.Context) {
	var req RegistrarPacienteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	dto := servicos.RegistrarPacienteDTO{
		Nome:                 req.Nome,
		Email:                req.Email,
		Senha:                req.Senha,
		EhDependente:         *req.EhDependente,
		Idade:                req.Idade,
		DataInicioTratamento: req.DataInicioTratamento,
		HistoricoSaude:       req.HistoricoSaude,
		CPF:                  req.CPF,
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
