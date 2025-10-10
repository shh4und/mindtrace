package middlewares

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware cria um middleware para configurar CORS
// Permite requisicoes de origens especificadas no ambiente FRONTEND_ORIGINS
func CORSMiddleware() gin.HandlerFunc {
	originsEnv := strings.TrimSpace(os.Getenv("FRONTEND_ORIGINS"))
	var allowOrigins []string
	if originsEnv != "" {
		for _, o := range strings.Split(originsEnv, ",") {
			o = strings.TrimSpace(o)
			if o != "" {
				allowOrigins = append(allowOrigins, o)
			}
		}
	}
	if len(allowOrigins) == 0 {
		allowOrigins = []string{
			"http://localhost:5173",
		}
	}
	log.Printf("[CORS] AllowOrigins=%v", allowOrigins)

	cfg := cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return cors.New(cfg)
}
