package controladores

import (
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsuarioControlador struct {
	usuarioServico servicos.UsuarioServico
}

func NovoUsuarioControlador(us servicos.UsuarioServico) *UsuarioControlador {
	return &UsuarioControlador{usuarioServico: us}
}

func (uc *UsuarioControlador) BuscarPerfil(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	usuario, err := uc.usuarioServico.BuscarUsuarioPorID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (uc *UsuarioControlador) AtualizarPerfil(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	var req servicos.AtualizarPerfilDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	usuario, err := uc.usuarioServico.AtualizarPerfil(userID.(uint), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (uc *UsuarioControlador) AlterarSenha(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "ID do usuário não encontrado no token"})
		return
	}

	var req servicos.AlterarSenhaDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := uc.usuarioServico.AlterarSenha(userID.(uint), req)
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