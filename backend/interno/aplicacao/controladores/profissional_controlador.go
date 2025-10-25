package controladores

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"regexp"

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

	/********************

		Verficicar a possibilidade das checagens feitas em controladores
		serem implementadas como regras de negocios nos dominios

	*********************/

	if len(req.Senha) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Senha inválida. Use 8 ou mais caracteres com letras, números e os símbolos: !@#$%^&*"})
		return
	}

	passwordRegex := `^[a-zA-Z0-9!@#$%^&*].{8,}$`
	if match, _ := regexp.MatchString(passwordRegex, req.Senha); !match {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Senha inválida. Use 8 ou mais caracteres com letras, números e os símbolos: !@#$%^&*"})
		return
	}

	profissionalOut, err := pc.usuarioServico.RegistrarProfissional(&req)
	if err != nil {
		if err == servicos.ErrEmailJaCadastrado {
			c.JSON(http.StatusConflict, gin.H{"erro": err.Error()})
		} else {
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
