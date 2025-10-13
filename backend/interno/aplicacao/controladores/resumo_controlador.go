package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResumoControlador gerencia requisicoes HTTP relacionadas a resumos
type ResumoControlador struct {
	resumoServico servicos.ResumoServico
}

// NovoResumoControlador cria uma nova instancia de ResumoControlador com o ResumoServico fornecido
func NovoResumoControlador(rs servicos.ResumoServico) *ResumoControlador {
	return &ResumoControlador{resumoServico: rs}
}

// GerarResumo gera um resumo para o paciente autenticado
// Extrai o periodo da query e chama o servico para gerar o resumo
func (rc *ResumoControlador) GerarResumo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuario nao encontrado no token"})
	}
	resumo, err := rc.resumoServico.GerarResumoPaciente(userID.(uint))

	if resumo == nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"erro": "Sem ocorrencias de registros de humor para este usuario"})

	}

	c.JSON(http.StatusOK, resumo)
}

// -- TODO --

// // GerarResumoPacienteDoProfissional gera um resumo para um paciente especifico pelo profissional
// // Extrai o ID do paciente e periodo da query e chama o servico para gerar o resumo
// func (rc *ResumoControlador) GerarResumoPacienteDoProfissional(c *gin.Context) {
// 	_, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID de usuario nao encontrado no token"})
// 	}

// 	pacienteIDStr := c.DefaultQuery("pacienteID", "0")
// 	if pacienteIDStr == "0" {
// 		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID de paciente invalido"})
// 	}
// 	pacienteID, err := strconv.Atoi(pacienteIDStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'pacienteID' invalido"})
// 	}

// 	periodoStr := c.DefaultQuery("periodo", "7")
// 	periodo, err := strconv.Atoi(periodoStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'periodo' invalido"})
// 	}
// 	resumo, err := rc.resumoServico.GerarResumoPacienteDoProfissional(uint(pacienteID), int64(periodo))

// 	if resumo == nil {
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"erro": "sem dados de humor registrados para este periodo"})

// 	}

// 	c.JSON(http.StatusOK, resumo)
// }
