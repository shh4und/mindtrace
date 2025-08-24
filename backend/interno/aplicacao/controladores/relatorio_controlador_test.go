package controladores

import (
	"encoding/json"
	"errors"
	"fmt"
	"mindtrace/backend/interno/aplicacao/middlewares"
	"mindtrace/backend/interno/aplicacao/servicos"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// 1. Mock do Serviço de Relatório
// Implementa a interface RelatorioServico para que possamos controlar seu comportamento nos testes.
type mockRelatorioServico struct {
	dtoToReturn *servicos.RelatorioPacienteDTO
	errToReturn error
}

func (m *mockRelatorioServico) GerarRelatorioPaciente(pacienteID uint, filtroPeriodo int64) (*servicos.RelatorioPacienteDTO, error) {
	return m.dtoToReturn, m.errToReturn
}

// 2. Helper para gerar um Token JWT de teste
func generateTestToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	// IMPORTANTE: A chave secreta aqui deve ser a mesma usada no seu ambiente de teste.
	// Para este exemplo, vamos usar uma chave simples.
	jwtSecret := []byte("test_secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

// 3. Função de Setup do Roteador
// Cria um roteador Gin com o middleware e o controlador de relatório para os testes.
func setupRouter(mockService servicos.RelatorioServico, t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mock do JWT_SECRET para o middleware
	t.Setenv("JWT_SECRET", "test_secret")

	relatorioController := NovoRelatorioControlador(mockService)

	// Agrupando rotas protegidas
	protegido := router.Group("/api/v1")
	protegido.Use(middlewares.AutMiddleware())
	{

		protegido.GET("/relatorios", relatorioController.GerarRelatorio)

	}
	return router
}

func TestGerarRelatorio_HappyPath(t *testing.T) {
	// Arrange
	mockDto := &servicos.RelatorioPacienteDTO{MediaSono: 7.5}
	mockService := &mockRelatorioServico{dtoToReturn: mockDto}
	router := setupRouter(mockService, t)

	w := httptest.NewRecorder()
	token, _ := generateTestToken(1)
	req, _ := http.NewRequest("GET", "/api/v1/relatorios?periodo=30", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody servicos.RelatorioPacienteDTO
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, mockDto.MediaSono, responseBody.MediaSono)
}

func TestGerarRelatorio_Unauthorized(t *testing.T) {
	// Arrange
	mockService := &mockRelatorioServico{}
	router := setupRouter(mockService, t)

	w := httptest.NewRecorder()
	// Requisição sem token de autorização
	req, _ := http.NewRequest("GET", "/api/v1/relatorios?periodo=30", nil)

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGerarRelatorio_BadRequest(t *testing.T) {
	// Arrange
	mockService := &mockRelatorioServico{}
	router := setupRouter(mockService, t)

	w := httptest.NewRecorder()
	token, _ := generateTestToken(1)
	// Requisição com parâmetro de período inválido
	req, _ := http.NewRequest("GET", "/api/v1/relatorios?periodo=abc", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGerarRelatorio_ServiceError(t *testing.T) {
	// Arrange
	mockService := &mockRelatorioServico{
		errToReturn: errors.New("erro ao gerar relatório"),
	}
	router := setupRouter(mockService, t)

	w := httptest.NewRecorder()
	token, _ := generateTestToken(1)
	req, _ := http.NewRequest("GET", "/api/v1/relatorios?periodo=30", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
