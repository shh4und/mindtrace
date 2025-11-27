package controladores

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegistroHumorControlador gerencia requisicoes HTTP relacionadas a registros de humor
type RegistroHumorControlador struct {
	registroHumorServico servicos.RegistroHumorServico
}

// NovoRegistroHumorControlador cria uma nova instancia de RegistroHumorControlador com o RegistroHumorServico fornecido
func NovoRegistroHumorControlador(us servicos.RegistroHumorServico) *RegistroHumorControlador {
	return &RegistroHumorControlador{registroHumorServico: us}
}

// Criar cria um novo registro de humor para o usuario autenticado
// Valida a entrada e chama o servico para criar o registro
func (rhc *RegistroHumorControlador) Criar(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}
	var req dtos.CriarRegistroHumorDTOIn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	registro_humor, err := rhc.registroHumorServico.CriarRegistroHumor(&req, userID.(uint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registro_humor)

}
