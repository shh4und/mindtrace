package main

import (
	"log"
	"mindtrace/backend/interno/aplicacao/controladores"
	"mindtrace/backend/interno/aplicacao/middlewares"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	postgres_repo "mindtrace/backend/interno/persistencia/postgres"
	"mindtrace/backend/interno/persistencia/repositorios"
	sqlite_repo "mindtrace/backend/interno/persistencia/sqlite"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var err error

	dbDriver := os.Getenv("DB_DRIVER")

	switch dbDriver {
	case "postgres":
		db, err = postgres_repo.NewDB()
		if err != nil {
			log.Fatalf("failed to connect to postgres: %v", err)
		}
	case "sqlite":
		db, err = sqlite_repo.NewDB()
		if err != nil {
			log.Fatalf("failed to connect to sqlite: %v", err)
		}
	default:
		log.Fatalf("invalid DB_DRIVER: %s", dbDriver)
	}

	err = db.AutoMigrate(
		&dominio.Usuario{},
		&dominio.Profissional{},
		&dominio.Paciente{},
		&dominio.RegistroHumor{},
		&dominio.Alerta{},
		&dominio.Notificacao{},
		&dominio.Convite{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	var usuarioRepo repositorios.UsuarioRepositorio
	var registroHumorRepo repositorios.RegistroHumorRepositorio
	var conviteRepo repositorios.ConviteRepositorio

	switch dbDriver {
	case "postgres":
		usuarioRepo = postgres_repo.NovoGormUsuarioRepositorio(db)
		registroHumorRepo = postgres_repo.NovoGormRegistroHumorRepositorio(db)
		conviteRepo = postgres_repo.NovoGormConviteRepositorio(db)
	case "sqlite":
		usuarioRepo = sqlite_repo.NovoGormUsuarioRepositorio(db)
		registroHumorRepo = sqlite_repo.NovoGormRegistroHumorRepositorio(db)
		conviteRepo = sqlite_repo.NovoGormConviteRepositorio(db)
	}

	usuarioService := servicos.NovoUsuarioServico(db, usuarioRepo)
	profissionalController := controladores.NovoProfissionalControlador(usuarioService)
	pacienteController := controladores.NovoPacienteControlador(usuarioService)
	autController := controladores.NovoAutControlador(usuarioService)
	usuarioController := controladores.NovoUsuarioControlador(usuarioService)

	registroHumorService := servicos.NovoRegistroHumorServico(db, registroHumorRepo, usuarioRepo)
	registroHumorController := controladores.NovoRegistroHumorControlador(registroHumorService)

	relatorioService := servicos.NovoRelatorioServico(db, registroHumorRepo, usuarioRepo)
	relatorioController := controladores.NovoRelatorioControlador(relatorioService)

	conviteService := servicos.NovoConviteServico(db, conviteRepo, usuarioRepo)
	conviteController := controladores.NovoConviteControlador(conviteService)

	roteador := gin.Default()
	roteador.SetTrustedProxies([]string{"127.0.0.1"})
	// Adiciona o middleware de CORS
	roteador.Use(middlewares.CORSMiddleware())

	api := roteador.Group("/api/v1")
	{
		// --- ROTAS PÚBLICAS ---
		auth := api.Group("/entrar")
		{
			auth.POST("/login", autController.Login)
		}

		profissionais := api.Group("/profissionais")
		{
			// O registro de profissionais pode ser público
			profissionais.POST("/registrar", profissionalController.Registrar)
		}

		pacientes := api.Group("/pacientes")
		{
			// O registro de pacientes também
			pacientes.POST("/registrar", pacienteController.Registrar)
		}

		// --- ROTAS PROTEGIDAS ---
		// Todas as rotas dentro deste grupo exigirão um token JWT válido
		protegido := api.Group("/")
		protegido.Use(middlewares.AutMiddleware())
		{
			usuarios := protegido.Group("/usuarios")
			{
				usuarios.GET("/", usuarioController.BuscarPerfil)
				usuarios.GET("/paciente", pacienteController.ProprioPerfilPaciente)
				usuarios.GET("/profissional", profissionalController.ProprioPerfilProfissional)
				usuarios.GET("/profissional/pacientes", usuarioController.ListarPacientesDoProfissional)
				usuarios.PUT("/perfil", usuarioController.AtualizarPerfil)
				usuarios.PUT("/perfil/alterar-senha", usuarioController.AlterarSenha) // CONSERTAR
				// PARA FAZER - URGENTE
				usuarios.DELETE("/perfil/apagar-conta", usuarioController.DeletarPerfil)
			}

			registroHumor := protegido.Group("/registro-humor")
			{
				registroHumor.POST("/", registroHumorController.Criar)

			}

			relatorios := protegido.Group("/relatorios")
			{
				relatorios.GET("/", relatorioController.GerarRelatorio)
				relatorios.GET("/paciente-lista", relatorioController.GerarRelatorioPacienteDoProfissional)
			}

			convites := protegido.Group("/convites")
			{
				convites.POST("/gerar", conviteController.GerarConvite)
				convites.POST("/vincular", conviteController.VincularPaciente)
			}
		}
	}

	log.Println("Server is running on port 8080")
	roteador.Run(":8080")
}
