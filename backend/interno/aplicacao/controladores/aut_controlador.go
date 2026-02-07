package controladores

import (
	"mindtrace/backend/interno/aplicacao/dtos"
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

// Login lida com o login do usuario
// Valida a entrada e chama o servico para autenticar o usuario
func (ac *AutControlador) Login(c *gin.Context) {
	var req dtos.LoginDTOIn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	accessToken, refreshToken, err := ac.usuarioServico.Login(req.Email, req.Senha)
	if err != nil {
		// Retorna 401 para credenciais invalidas
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.TokenDTOOut{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// Refresh renova o access token usando um refresh token valido
func (ac *AutControlador) Refresh(c *gin.Context) {
	var req dtos.RefreshTokenDTOIn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	accessToken, refreshToken, err := ac.usuarioServico.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.TokenDTOOut{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
