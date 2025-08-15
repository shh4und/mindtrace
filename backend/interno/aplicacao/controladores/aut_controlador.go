package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AutControlador struct {
	usuarioServico servicos.UsuarioServico
}

func NovoAutControlador(us servicos.UsuarioServico) *AutControlador {
	return &AutControlador{usuarioServico: us}
}

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

func (ac *AutControlador) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	token, err := ac.usuarioServico.Login(req.Email, req.Senha)
	if err != nil {
		// Retorna 401 para credenciais inválidas
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
