package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConviteControlador gerencia requisicoes HTTP relacionadas a convites
type ConviteControlador struct {
	conviteServico servicos.ConviteServico
}

// NovoConviteControlador cria uma nova instancia de ConviteControlador com o ConviteServico fornecido
func NovoConviteControlador(cs servicos.ConviteServico) *ConviteControlador {
	return &ConviteControlador{conviteServico: cs}
}

// GerarConvite gera um novo convite para o usuario autenticado
func (cc *ConviteControlador) GerarConvite(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	convite, err := cc.conviteServico.GerarConvite(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":          convite.Token,
		"data_expiracao": convite.DataExpiracao,
	})
}

// VincularRequest representa o payload da requisicao para vincular paciente
type VincularRequest struct {
	Token string `json:"token" binding:"required"`
}

// VincularPaciente vincula um paciente usando um token de convite
// Valida a entrada e chama o servico para realizar o vinculo
func (cc *ConviteControlador) VincularPaciente(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	var req VincularRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := cc.conviteServico.VincularPaciente(userID.(uint), req.Token)
	if err != nil {
		switch err {
		case servicos.ErrConviteNaoEncontrado, servicos.ErrConviteExpirado, servicos.ErrPerfilNaoEncontrado:
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao vincular paciente"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Paciente vinculado com sucesso!"})
}
