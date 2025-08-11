package controllers

import (
	"mindtrace/backend/internal/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfissionalController struct {
	usuarioService services.UsuarioService
}

func NewProfissionalController(us services.UsuarioService) *ProfissionalController {
	return &ProfissionalController{usuarioService: us}
}

type RegisterProfissionalRequest struct {
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required,min=8"`
	Specialty            string `json:"specialty" binding:"required"`
	ProfessionalRegistry string `json:"professional_registry" binding:"required"`
}

func (pc *ProfissionalController) Register(c *gin.Context) {
	var request RegisterProfissionalRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := services.RegisterProfissionalDTO{
		Name:                 request.Name,
		Email:                request.Email,
		Password:             request.Password,
		Specialty:            request.Specialty,
		ProfessionalRegistry: request.ProfessionalRegistry,
	}

	profissional, err := pc.usuarioService.RegisterProfissional(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating Profissional"})
		return
	}

	c.JSON(http.StatusCreated, profissional)
}
