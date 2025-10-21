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

// main inicializa servidor http organiza dependencias e configura rotas principais
func main() {
	var db *gorm.DB
	var err error

	dbDriver := os.Getenv("DB_DRIVER")

	switch dbDriver {
	case "postgres":
		db, err = postgres_repo.NewDB()
		if err != nil {
			log.Fatalf("falha ao conectar ao postgres: %v", err)
		}
	case "sqlite":
		db, err = sqlite_repo.NewDB()
		if err != nil {
			log.Fatalf("falha ao conectar ao sqlite: %v", err)
		}
	default:
		log.Fatalf("DB_DRIVER invalido: %s", dbDriver)
	}

	// Executa migracoes automatizadas para alinhar esquema do banco
	err = db.AutoMigrate(
		&dominio.Usuario{},
		&dominio.Profissional{},
		&dominio.Paciente{},
		&dominio.RegistroHumor{},
		&dominio.Notificacao{},
		&dominio.Convite{},
	)
	if err != nil {
		log.Fatalf("falha ao migrar o banco de dados: %v", err)
	}

	// Seleciona implementacoes de repositorio conforme driver ativo
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

	// Inicializa servicos e controladores da camada de aplicacao
	usuarioService := servicos.NovoUsuarioServico(db, usuarioRepo)
	profissionalController := controladores.NovoProfissionalControlador(usuarioService)
	pacienteController := controladores.NovoPacienteControlador(usuarioService)
	autController := controladores.NovoAutControlador(usuarioService)
	usuarioController := controladores.NovoUsuarioControlador(usuarioService)

	registroHumorService := servicos.NovoRegistroHumorServico(db, registroHumorRepo, usuarioRepo)
	registroHumorController := controladores.NovoRegistroHumorControlador(registroHumorService)

	relatorioService := servicos.NovoRelatorioServico(db, registroHumorRepo, usuarioRepo)
	relatorioController := controladores.NovoRelatorioControlador(relatorioService)

	resumoService := servicos.NovoResumoServico(db, registroHumorRepo, usuarioRepo)
	resumoController := controladores.NovoResumoControlador(resumoService)

	conviteService := servicos.NovoConviteServico(db, conviteRepo, usuarioRepo)
	conviteController := controladores.NovoConviteControlador(conviteService)

	// Configura roteador http com middlewares e grupos de rotas
	roteador := gin.Default()
	roteador.SetTrustedProxies([]string{"127.0.0.1"})
	// Inclui middleware cors padrao aceitando chamadas do frontend
	roteador.Use(middlewares.CORSMiddleware())

	api := roteador.Group("/api/v1")
	{
		// --- ROTAS PUBLICAS ---
		auth := api.Group("/entrar")
		{
			auth.POST("/login", autController.Login)
		}

		profissionais := api.Group("/profissionais")
		{
			// Registro de profissionais acessivel sem autenticacao
			profissionais.POST("/registrar", profissionalController.Registrar)
		}

		pacientes := api.Group("/pacientes")
		{
			// Registro de pacientes disponivel sem token
			pacientes.POST("/registrar", pacienteController.Registrar)
		}

		// --- ROTAS PROTEGIDAS ---
		// Todas as rotas deste grupo exigirao token jwt valido
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

			resumo := protegido.Group("/resumo")
			{
				resumo.GET("/", resumoController.GerarResumo)
			}

			convites := protegido.Group("/convites")
			{
				convites.POST("/gerar", conviteController.GerarConvite)
				convites.POST("/vincular", conviteController.VincularPaciente)
			}
		}
	}

	log.Println("servidor iniciado na porta 8080")
	roteador.Run(":8080")
}
