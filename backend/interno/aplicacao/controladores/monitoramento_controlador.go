package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MonitoramentoControlador gerencia requisicoes HTTP relacionadas a monitoramentos
type MonitoramentoControlador struct {
	monitoramentoServico servicos.MonitoramentoServico
}

// NovoMonitoramentoControlador cria uma nova instancia de MonitoramentoControlador com o MonitoramentoServico fornecido
func NovoMonitoramentoControlador(ms servicos.MonitoramentoServico) *MonitoramentoControlador {
	return &MonitoramentoControlador{monitoramentoServico: ms}
}

// GerarMonitoramento gera um monitoramento para o paciente autenticado
// Extrai o periodo da query e chama o servico para gerar o monitoramento
func (mc *MonitoramentoControlador) RealizarMonitoramento(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuario nao encontrado no token"})
	}

	pacienteIDStr := c.DefaultQuery("pacienteID", "0")
	if pacienteIDStr == "0" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID de paciente invalido"})
	}
	pacienteID, err := strconv.Atoi(pacienteIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'pacienteID' invalido"})
	}

	numRegistrosStr := c.DefaultQuery("numRegistros", "4")
	numRegistros, err := strconv.Atoi(numRegistrosStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'numRegistros' invalido"})
	}
	monitoramentoOut, err := mc.monitoramentoServico.RealizarMonitoramentoPaciente(userID.(uint), uint(pacienteID), int(numRegistros))

	if monitoramentoOut == nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"erro": "sem dados de humor registrados para este numRegistros"})
	}

	c.JSON(http.StatusOK, monitoramentoOut)
}
