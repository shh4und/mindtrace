package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AutControlador gerencia requisicoes HTTP relacionadas a autenticacao
type AutControlador struct {
	usuarioServico servicos.UsuarioServico
}

// NovoAutControlador cria uma nova instancia de AutControlador com o UsuarioServico fornecido
func NovoAutControlador(us servicos.UsuarioServico) *AutControlador {
	return &AutControlador{usuarioServico: us}
}

// LoginRequest representa o payload da requisicao para login
type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

// Login lida com o login do usuario
// Valida a entrada e chama o servico para autenticar o usuario
func (ac *AutControlador) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	token, err := ac.usuarioServico.Login(req.Email, req.Senha)
	if err != nil {
		// Retorna 401 para credenciais invalidas
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
