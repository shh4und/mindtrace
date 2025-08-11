package main

import (
	"log"
	"mindtrace/backend/internal/application/controllers"
	"mindtrace/backend/internal/application/services"
	"mindtrace/backend/internal/domain"
	postgres_repo "mindtrace/backend/internal/persistence/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost usuario=postgres password=postgres dbname=mindtrace port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&domain.Usuario{}, &domain.Profissional{}, &domain.Paciente{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	usuarioRepo := postgres_repo.NewGormUsuarioRepository(db)
	usuarioService := services.NewUsuarioService(db, usuarioRepo)
	profissionalController := controllers.NewProfissionalController(usuarioService)

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		professionals := api.Group("/professionals")
		{
			professionals.POST("/register", profissionalController.Register)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend is running correctly!",
		})
	})

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}
