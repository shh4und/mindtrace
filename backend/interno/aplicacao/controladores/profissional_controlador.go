package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"regexp"

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
	CPF                  string `json:"cpf" binding:"required"`
}

type ProprioProfissionalRequest struct {
	Nome                 string `json:"nome"`
	Email                string `json:"email" `
	Especialidade        string `json:"especialidade" `
	RegistroProfissional string `json:"registro_profissional" `
	CPF                  string `json:"cpf"`
}

func (pc *ProfissionalControlador) Registrar(c *gin.Context) {
	var req RegistrarProfissionalRequest
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

	dto := servicos.RegistrarProfissionalDTO{
		Nome:                 req.Nome,
		Email:                req.Email,
		Senha:                req.Senha,
		Especialidade:        req.Especialidade,
		RegistroProfissional: req.RegistroProfissional,
		CPF:                  req.CPF,
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

func (uc *ProfissionalControlador) ProprioPerfilProfissional(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	profissional, err := uc.usuarioServico.ProprioPerfilProfissional(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	proprioProfissional := ProprioProfissionalRequest{
		Nome:                 profissional.Usuario.Nome,
		Email:                profissional.Usuario.Email,
		CPF:                  profissional.Usuario.CPF,
		Especialidade:        profissional.Especialidade,
		RegistroProfissional: profissional.RegistroProfissional,
	}

	c.JSON(http.StatusOK, proprioProfissional)
}
