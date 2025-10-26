package controladores

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProfissionalControlador gerencia requisicoes HTTP relacionadas ao gerenciamento de profissionais
type ProfissionalControlador struct {
	usuarioServico servicos.UsuarioServico
}

// NovoProfissionalControlador cria uma nova instancia de ProfissionalControlador com o UsuarioServico fornecido
func NovoProfissionalControlador(us servicos.UsuarioServico) *ProfissionalControlador {
	return &ProfissionalControlador{usuarioServico: us}
}

// Registrar lida com o registro de um novo profissional
// Valida a entrada verifica a forca da senha e chama o servico para registrar o profissional
func (pc *ProfissionalControlador) Registrar(c *gin.Context) {
	var req dtos.RegistrarProfissionalDTOIn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	profissionalOut, err := pc.usuarioServico.RegistrarProfissional(&req)
	if err != nil {
		switch err {
		case dominio.ErrEmailJaCadastrado:
			c.JSON(http.StatusConflict, gin.H{"erro": err.Error()})
		case dominio.ErrSenhaFraca:
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		case dominio.ErrEmailInvalido:
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		case dominio.ErrNomeVazio:
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, profissionalOut)
}

// ProprioPerfilProfissional recupera o perfil do profissional autenticado
// Extrai o ID do usuario do contexto e busca os dados do profissional do servico
func (uc *ProfissionalControlador) ProprioPerfilProfissional(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	proprioProfissional, err := uc.usuarioServico.ProprioPerfilProfissional(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, proprioProfissional)
}
