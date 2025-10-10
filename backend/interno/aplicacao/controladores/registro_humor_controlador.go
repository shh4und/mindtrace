package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"time"

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

// CriarRegistroHumorRequest representa o payload da requisicao para criar um registro de humor
type CriarRegistroHumorRequest struct {
	NivelHumor       int16     `json:"nivel_humor" binding:"required,gte=1"`
	HorasSono        int16     `json:"horas_sono" binding:"gte=0"`
	NivelStress      int16     `json:"nivel_stress" binding:"required,gte=1"`
	NivelEnergia     int16     `json:"nivel_energia" binding:"required,gte=1"`
	AutoCuidado      []string  `json:"auto_cuidado" binding:"required"`
	Observacoes      string    `json:"observacoes"`
	DataHoraRegistro time.Time `json:"data_hora_registro"`
}

// Criar cria um novo registro de humor para o usuario autenticado
// Valida a entrada e chama o servico para criar o registro
func (rhc *RegistroHumorControlador) Criar(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}
	var req CriarRegistroHumorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	dto := servicos.CriarRegistroHumorDTO{
		UsuarioID:        userID.(uint),
		NivelHumor:       req.NivelHumor,
		HorasSono:        req.HorasSono,
		NivelStress:      req.NivelStress,
		NivelEnergia:     req.NivelEnergia,
		AutoCuidado:      req.AutoCuidado,
		Observacoes:      req.Observacoes,
		DataHoraRegistro: req.DataHoraRegistro,
	}

	registro_humor, err := rhc.registroHumorServico.CriarRegistroHumor(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registro_humor)

}
