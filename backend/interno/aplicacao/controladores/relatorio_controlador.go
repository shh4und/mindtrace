package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RelatorioControlador gerencia requisicoes HTTP relacionadas a relatorios
type RelatorioControlador struct {
	analiseServico servicos.AnaliseServico
}

// NovoRelatorioControlador cria uma nova instancia de RelatorioControlador com o AnaliseServico fornecido
func NovoRelatorioControlador(rs servicos.AnaliseServico) *RelatorioControlador {
	return &RelatorioControlador{analiseServico: rs}
}

// GerarRelatorio gera um relatorio para o paciente autenticado
// Extrai o periodo da query e chama o servico para gerar o relatorio
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
	relatorio, err := rc.analiseServico.GerarAnaliseHistorica(userID.(uint), int(periodo))

	if relatorio == nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"erro": "sem dados de humor registrados para este periodo"})

	}

	c.JSON(http.StatusOK, relatorio)
}

// GerarAnaliseHistorica gera um relatorio para um paciente especifico pelo profissional
// Extrai o ID do paciente e periodo da query e chama o servico para gerar o relatorio
func (rc *RelatorioControlador) GerarAnaliseHistorica(c *gin.Context) {
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
	relatorio, err := rc.analiseServico.GerarAnaliseHistorica(uint(pacienteID), int(periodo))

	if relatorio == nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"erro": "sem dados de humor registrados para este periodo"})

	}

	c.JSON(http.StatusOK, relatorio)
}
