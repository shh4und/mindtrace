package main

import (
	"fmt"
	"log"
	"mindtrace/backend/interno/aplicacao/controladores"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	postgres_repo "mindtrace/backend/interno/persistencia/postgres"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db_USER := os.Getenv("DB_USER")
	db_PASS := os.Getenv("DB_PASSWORD")
	db_NAME := os.Getenv("DB_DB")

	DSN := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		db_USER,
		db_PASS,
		db_NAME,
	)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(
		&dominio.Usuario{},
		&dominio.Profissional{},
		&dominio.Paciente{},
		&dominio.ResponsavelLegal{},
		&dominio.ProfissionalPaciente{}, // Tabela de junção
		&dominio.RegistroHumor{},
		&dominio.AnotacaoDiaria{},
		&dominio.Alerta{},
		&dominio.Notificacao{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	usuarioRepo := postgres_repo.NovoGormUsuarioRepositorio(db)
	usuarioService := servicos.NovoUsuarioServico(db, usuarioRepo)
	profissionalController := controladores.NovoProfissionalControlador(usuarioService)
	pacienteController := controladores.NovoPacienteControlador(usuarioService)

	roteador := gin.Default()

	api := roteador.Group("/api/v1")
	{
		profissionais := api.Group("/profissionais")
		{
			profissionais.POST("/registrar", profissionalController.Registrar)
		}

		pacientes := api.Group("/pacientes")
		{
			pacientes.POST("/registrar", pacienteController.Registrar)
		}
	}

	roteador.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend is running correctly!",
		})
	})

	log.Println("Server is running on port 8080")
	roteador.Run(":8080")
}
