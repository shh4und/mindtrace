package main

import (
	"log"
	"mindtrace/backend/interno/aplicacao/controladores"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	postgres_repo "mindtrace/backend/interno/persistencia/postgres"

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

	err = db.AutoMigrate(&dominio.Usuario{}, &dominio.Profissional{}, &dominio.Paciente{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	usuarioRepo := postgres_repo.NewGormUsuarioRepositorio(db)
	usuarioService := servicos.NewUsuarioServico(db, usuarioRepo)
	profissionalController := controladores.NewProfissionalController(usuarioService)

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		professionals := api.Group("/professionals")
		{
			professionals.POST("/register", profissionalController.Registrar)
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
