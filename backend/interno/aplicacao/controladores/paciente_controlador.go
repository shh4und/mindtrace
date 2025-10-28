package controladores

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PacienteControlador gerencia requisicoes HTTP relacionadas ao gerenciamento de pacientes
type PacienteControlador struct {
	usuarioServico servicos.UsuarioServico
}

// NovoPacienteControlador cria uma nova instancia de PacienteControlador com o UsuarioServico fornecido
func NovoPacienteControlador(us servicos.UsuarioServico) *PacienteControlador {
	return &PacienteControlador{usuarioServico: us}
}

// Registrar lida com o registro de um novo paciente
// Valida a entrada analisa datas verifica a forca da senha e chama o servico para registrar o paciente
func (pc *PacienteControlador) Registrar(c *gin.Context) {
	var req dtos.RegistrarPacienteDTOIn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	pacienteOut, err := pc.usuarioServico.RegistrarPaciente(&req)
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

	c.JSON(http.StatusCreated, pacienteOut)
}

// ProprioPerfilPaciente recupera o perfil do paciente autenticado
// Extrai o ID do usuario do contexto e busca os dados do paciente do servico
func (uc *PacienteControlador) ProprioPerfilPaciente(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	proprioPaciente, err := uc.usuarioServico.ProprioPerfilPaciente(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, proprioPaciente)
}
