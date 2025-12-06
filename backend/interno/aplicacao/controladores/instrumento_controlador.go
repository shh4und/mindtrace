package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	"net/http"
	"strconv"

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
		return
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

func (ic *InstrumentoControlador) AtribuirInstrumento(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	pacienteIDStr := c.DefaultQuery("pacienteID", "0")
	if pacienteIDStr == "0" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID de paciente invalido"})
		return
	}
	pacienteID, err := strconv.Atoi(pacienteIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'pacienteID' invalido"})
		return
	}

	instrumentoIDStr := c.DefaultQuery("instrumentoID", "0")
	if instrumentoIDStr == "0" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID de instrumento invalido"})
		return
	}
	instrumentoID, err := strconv.Atoi(instrumentoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'instrumentoID' invalido"})
		return
	}

	instrumentoCodigoStr := c.DefaultQuery("instrumentoCodigo", "")

	err = ic.instrumentoServico.CriarAtribuicao(userID.(uint), uint(pacienteID), uint(instrumentoID), instrumentoCodigoStr)
	if err != nil {
		if err == dominio.ErrUsuarioNaoEncontrado {
			c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "atribuicao realizada com sucesso"})
}

func (ic *InstrumentoControlador) ListarAtribuicoesPaciente(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}
	atribuicoesOut, err := ic.instrumentoServico.ListarAtribuicoesPaciente(userID.(uint))
	if err != nil {
		if err == dominio.ErrUsuarioNaoEncontrado {
			c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, atribuicoesOut)
}

func (ic *InstrumentoControlador) ListarAtribuicoesProfissional(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}
	atribuicoesOut, err := ic.instrumentoServico.ListarAtribuicoesProfissional(userID.(uint))
	if err != nil {
		if err == dominio.ErrUsuarioNaoEncontrado {
			c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, atribuicoesOut)
}
