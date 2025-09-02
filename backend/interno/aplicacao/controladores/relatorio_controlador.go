package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RelatorioControlador struct {
	relatorioServico servicos.RelatorioServico
}

func NovoRelatorioControlador(rs servicos.RelatorioServico) *RelatorioControlador {
	return &RelatorioControlador{relatorioServico: rs}
}

func (rc *RelatorioControlador) GerarRelatorio(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuario nao encontrado no token"})
	}

	periodoStr := c.DefaultQuery("periodo", "7")
	periodo, err := strconv.Atoi(periodoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'periodo' invalido"})
	}
	relatorio, err := rc.relatorioServico.GerarRelatorioPaciente(userID.(uint), int64(periodo))

	if relatorio == nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"erro": "sem dados de humor registrados para este periodo"})

	}

	c.JSON(http.StatusOK, relatorio)
}

func (rc *RelatorioControlador) GerarRelatorioPacienteDoProfissional(c *gin.Context) {
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID de usuario nao encontrado no token"})
	}

	pacienteIDStr := c.DefaultQuery("pacienteID", "0")
	if pacienteIDStr == "0" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID de paciente invalido"})
	}
	pacienteID, err := strconv.Atoi(pacienteIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'pacienteID' invalido"})
	}

	periodoStr := c.DefaultQuery("periodo", "7")
	periodo, err := strconv.Atoi(periodoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'periodo' invalido"})
	}
	relatorio, err := rc.relatorioServico.GerarRelatorioPacienteDoProfissional(uint(pacienteID), int64(periodo))

	if relatorio == nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"erro": "sem dados de humor registrados para este periodo"})

	}

	c.JSON(http.StatusOK, relatorio)
}
