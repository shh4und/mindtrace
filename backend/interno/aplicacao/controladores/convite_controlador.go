package controladores

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
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

	conviteOut, err := cc.conviteServico.GerarConvite(userID.(uint))
	if err != nil {
		if err == dominio.ErrUsuarioNaoEncontrado {
			c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, conviteOut)
}

// VincularPaciente vincula um paciente usando um token de convite
// Valida a entrada e chama o servico para realizar o vinculo
func (cc *ConviteControlador) VincularPaciente(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	var req dtos.VincularPacienteDTOIn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := cc.conviteServico.VincularPaciente(userID.(uint), req.Token)
	if err != nil {
		switch err {
		case dominio.ErrTokenConviteInvalido, dominio.ErrConviteExpirado, dominio.ErrConviteExpirado, dominio.ErrUsuarioNaoEncontrado:
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao vincular paciente"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Paciente vinculado com sucesso"})
}
