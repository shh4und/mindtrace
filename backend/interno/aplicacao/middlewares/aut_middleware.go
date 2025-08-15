package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AutMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Cabeçalho de autorização não encontrado"})
			return
		}

		// O header deve ser no formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Formato do cabeçalho de autorização inválido"})
			return
		}

		tokenString := parts[1]
		jwtSecret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Token inválido"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Extrai o ID do usuário do token e o coloca no contexto do Gin
			// Isso é útil para os controladores saberem qual usuário está fazendo a requisição
			userIDFloat := claims["sub"].(float64)
			c.Set("userID", uint(userIDFloat))
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Token inválido"})
			return
		}

		c.Next() // Passa para o próximo handler
	}
}
