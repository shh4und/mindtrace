package main

import (
	"log"
	"mindtrace/backend/interno/aplicacao/controladores"
	"mindtrace/backend/interno/aplicacao/middlewares"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	postgres_repo "mindtrace/backend/interno/persistencia/postgres"
	"mindtrace/backend/interno/persistencia/repositorios"
	"mindtrace/backend/interno/persistencia/seeds"
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

	skipDBInit := os.Getenv("SKIP_DB_INIT") == "true"

	if !skipDBInit {
		// Executa migracoes automatizadas para alinhar esquema do banco
		err = db.AutoMigrate(
			&dominio.Usuario{},
			&dominio.Profissional{},
			&dominio.Paciente{},
			&dominio.RegistroHumor{},
			&dominio.Notificacao{},
			&dominio.Convite{},
			&dominio.Instrumento{},
			&dominio.Pergunta{},
			&dominio.OpcaoEscala{},
			&dominio.Atribuicao{},
			&dominio.Resposta{},
		)
		if err != nil {
			log.Fatalf("falha ao migrar o banco de dados: %v", err)
		}

		// Instrumentos imutaveis seedados
		seeds.ExecutarSeeds(db)
		// Dados mock para ambiente de desenvolvimento
		seeds.ExecutarSeedsMock(db)
	}

	var usuarioRepo repositorios.UsuarioRepositorio
	var registroHumorRepo repositorios.RegistroHumorRepositorio
	var conviteRepo repositorios.ConviteRepositorio
	var instrumentoRepo repositorios.InstrumentoRepositorio

	// Seleciona implementacoes de repositorio conforme driver ativo
	switch dbDriver {
	case "postgres":
		usuarioRepo = postgres_repo.NovoGormUsuarioRepositorio(db)
		registroHumorRepo = postgres_repo.NovoGormRegistroHumorRepositorio(db)
		conviteRepo = postgres_repo.NovoGormConviteRepositorio(db)
		instrumentoRepo = postgres_repo.NovoGormInstrumentoRepositorio(db)
	case "sqlite":
		usuarioRepo = sqlite_repo.NovoGormUsuarioRepositorio(db)
		registroHumorRepo = sqlite_repo.NovoGormRegistroHumorRepositorio(db)
		conviteRepo = sqlite_repo.NovoGormConviteRepositorio(db)
	}

	// Inicializa servicos
	usuarioSvc := servicos.NovoUsuarioServico(db, usuarioRepo)
	analiseSvc := servicos.NovoAnaliseServico(db, registroHumorRepo, usuarioRepo)
	registroHumorSvc := servicos.NovoRegistroHumorServico(db, registroHumorRepo, usuarioRepo, analiseSvc)
	resumoSvc := servicos.NovoResumoServico(db, registroHumorRepo, usuarioRepo)
	conviteSvc := servicos.NovoConviteServico(db, conviteRepo, usuarioRepo)
	instrumentoSvc := servicos.NovoInstrumentoServico(db, instrumentoRepo, usuarioRepo)

	// Inicializa controladores
	profissionalCtrl := controladores.NovoProfissionalControlador(usuarioSvc)
	pacienteCtrl := controladores.NovoPacienteControlador(usuarioSvc)
	autCtrl := controladores.NovoAutControlador(usuarioSvc)
	usuarioCtrl := controladores.NovoUsuarioControlador(usuarioSvc)
	registroHumorCtrl := controladores.NovoRegistroHumorControlador(registroHumorSvc)
	relatorioCtrl := controladores.NovoRelatorioControlador(analiseSvc)
	resumoCtrl := controladores.NovoResumoControlador(resumoSvc)
	conviteCtrl := controladores.NovoConviteControlador(conviteSvc)
	instrumentoCtrl := controladores.NovoInstrumentoControlador(instrumentoSvc)

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
			auth.POST("/login", autCtrl.Login)
		}

		profissionais := api.Group("/profissionais")
		{
			// Registro de profissionais acessivel sem autenticacao
			profissionais.POST("/registrar", profissionalCtrl.Registrar)
		}

		pacientes := api.Group("/pacientes")
		{
			// Registro de pacientes disponivel sem token
			pacientes.POST("/registrar", pacienteCtrl.Registrar)
		}

		// --- ROTAS PROTEGIDAS ---
		// Todas as rotas deste grupo exigirao token jwt valido
		protegido := api.Group("/")
		protegido.Use(middlewares.AutMiddleware())
		{
			usuarios := protegido.Group("/usuarios")
			{
				usuarios.GET("/", usuarioCtrl.BuscarPerfil)
				usuarios.GET("/paciente", pacienteCtrl.ProprioPerfilPaciente)
				usuarios.GET("/profissional", profissionalCtrl.ProprioPerfilProfissional)
				usuarios.GET("/profissional/pacientes", usuarioCtrl.ListarPacientesDoProfissional)
				usuarios.PUT("/perfil", usuarioCtrl.AtualizarPerfil)
				usuarios.PUT("/perfil/alterar-senha", usuarioCtrl.AlterarSenha)
				usuarios.DELETE("/perfil/apagar-conta", usuarioCtrl.DeletarPerfil)
			}

			registroHumor := protegido.Group("/registro-humor")
			{
				registroHumor.POST("/", registroHumorCtrl.Criar)

			}

			relatorios := protegido.Group("/relatorios")
			{
				relatorios.GET("/", relatorioCtrl.GerarRelatorio)
				relatorios.GET("/paciente-lista", relatorioCtrl.GerarAnaliseHistorica)
			}

			resumo := protegido.Group("/resumo")
			{
				resumo.GET("/", resumoCtrl.GerarResumo)
			}

			convites := protegido.Group("/convites")
			{
				convites.POST("/gerar", conviteCtrl.GerarConvite)
				convites.POST("/vincular", conviteCtrl.VincularPaciente)
			}

			instrumentos := protegido.Group("/instrumentos")
			{
				instrumentos.GET("/listar-instrumentos", instrumentoCtrl.ListarInstrumentos)
				instrumentos.POST("/atribuir-instrumento", instrumentoCtrl.AtribuirInstrumento)
				instrumentos.GET("/listar-atribuicoes-paciente", instrumentoCtrl.ListarAtribuicoesPaciente)
				instrumentos.GET("/listar-atribuicoes-profissional", instrumentoCtrl.ListarAtribuicoesProfissional)
				instrumentos.GET("/atribuicao", instrumentoCtrl.ApresentarPerguntasAtribuicao)
				instrumentos.POST("/registrar-respostas", instrumentoCtrl.RegistrarRespostas)

			}
		}
	}

	log.Println("servidor iniciado na porta 9090")
	roteador.Run(":9090")
}
