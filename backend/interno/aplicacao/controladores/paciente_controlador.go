package controladores

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"regexp"

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

	// dataNascimento, err := time.Parse("2006-01-02", req.DataNascimento)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"erro": "Formato de data inválido. Use YYYY-MM-DD"})
	// 	return
	// }

	if len(req.Senha) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Senha inválida. Use 8 ou mais caracteres com letras, números e os símbolos: !@#$%^&*"})
		return
	}
	/********************

		Verficicar a possibilidade das checagens feitas em controladores
		serem implementadas como regras de negocios nos dominios

	*********************/
	passwordRegex := `^[a-zA-Z0-9!@#$%^&*].{8,}$`
	if match, _ := regexp.MatchString(passwordRegex, req.Senha); !match {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Senha inválida. Use 8 ou mais caracteres com letras, números e os símbolos: !@#$%^&*"})
		return
	}

	pacienteOut, err := pc.usuarioServico.RegistrarPaciente(&req)
	if err != nil {
		if err == servicos.ErrEmailJaCadastrado {
			c.JSON(http.StatusConflict, gin.H{"erro": err.Error()})
		} else {
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
