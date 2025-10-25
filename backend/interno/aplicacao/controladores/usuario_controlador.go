package controladores

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UsuarioControlador gerencia requisicoes HTTP relacionadas ao gerenciamento de usuarios
type UsuarioControlador struct {
	usuarioServico servicos.UsuarioServico
}

// NovoUsuarioControlador cria uma nova instancia de UsuarioControlador com o UsuarioServico fornecido
func NovoUsuarioControlador(us servicos.UsuarioServico) *UsuarioControlador {
	return &UsuarioControlador{usuarioServico: us}
}

// BuscarPerfil busca o perfil do usuario autenticado
// Extrai o ID do usuario do contexto e chama o servico para buscar os dados
func (uc *UsuarioControlador) BuscarPerfil(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	usuarioOut, err := uc.usuarioServico.BuscarUsuarioPorID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuarioOut)
}

// AtualizarPerfil atualiza o perfil do usuario autenticado
// Valida a entrada e chama o servico para atualizar os dados
func (uc *UsuarioControlador) AtualizarPerfil(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	var req dtos.AtualizarPerfilDTOIn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := uc.usuarioServico.AtualizarPerfil(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Perfil atualizado com sucesso"})
}

// AlterarSenha altera a senha do usuario autenticado
// Valida a entrada e chama o servico para alterar a senha
func (uc *UsuarioControlador) AlterarSenha(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	var req dtos.AlterarSenhaDTOIn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := uc.usuarioServico.AlterarSenha(userID.(uint), &req)
	if err != nil {
		if err == servicos.ErrSenhaNaoConfere || err == servicos.ErrCrendenciaisInvalidas {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Senha alterada com sucesso"})
}

// ListarPacientesDoProfissional lista os pacientes associados ao profissional autenticado
// Extrai o ID do usuario e chama o servico para listar os pacientes
func (uc *UsuarioControlador) ListarPacientesDoProfissional(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	pacientesOut, err := uc.usuarioServico.ListarPacientesDoProfissional(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pacientesOut)
}

// DeletarPerfil deleta o perfil do usuario autenticado
// Extrai o ID do usuario e chama o servico para deletar a conta
func (uc *UsuarioControlador) DeletarPerfil(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	err := uc.usuarioServico.DeletarPerfil(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Conta deletada com sucesso"})
}
