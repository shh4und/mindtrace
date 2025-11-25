package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InstrumentoControlador struct {
	instrumentoServico servicos.InstrumentoServico
}

// NovoInstrumentoControlador cria uma nova instancia de InstrumentoControlador com o InstrumentoServico fornecido
func NovoInstrumentoControlador(is servicos.InstrumentoServico) *InstrumentoControlador {
	return &InstrumentoControlador{instrumentoServico: is}
}

func (ic *InstrumentoControlador) ListarInstrumentos(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
	}

	instrumentosOut, err := ic.instrumentoServico.ListarInstrumentos(userID.(uint))
	if err != nil {
		if err == dominio.ErrUsuarioNaoEncontrado {
			c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, instrumentosOut)
}
