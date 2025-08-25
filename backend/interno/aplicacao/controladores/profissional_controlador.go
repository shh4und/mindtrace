package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfissionalControlador struct {
	usuarioServico servicos.UsuarioServico
}

func NovoProfissionalControlador(us servicos.UsuarioServico) *ProfissionalControlador {
	return &ProfissionalControlador{usuarioServico: us}
}

type RegistrarProfissionalRequest struct {
	Nome                 string `json:"nome" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Senha                string `json:"senha" binding:"required,min=8"`
	Especialidade        string `json:"especialidade" binding:"required"`
	RegistroProfissional string `json:"registro_profissional" binding:"required"`
	CPF 				 string `json:"cpf" binding:"required"`
}

func (pc *ProfissionalControlador) Registrar(c *gin.Context) {
	var req RegistrarProfissionalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	dto := servicos.RegistrarProfissionalDTO{
		Nome:                 req.Nome,
		Email:                req.Email,
		Senha:                req.Senha,
		Especialidade:        req.Especialidade,
		RegistroProfissional: req.RegistroProfissional,
		CPF:				  req.CPF,
	}

	profissional, err := pc.usuarioServico.RegistrarProfissional(dto)
	if err != nil {
		if err == servicos.ErrEmailJaCadastrado {
			c.JSON(http.StatusConflict, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, profissional)
}
