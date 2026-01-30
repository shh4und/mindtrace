package main

import (
	"fmt"
	"log"
	"mindtrace/backend/interno/aplicacao/controladores"
	"mindtrace/backend/interno/aplicacao/middlewares"
	"mindtrace/backend/interno/aplicacao/servicos"
	"mindtrace/backend/interno/dominio"
	postgres_repo "mindtrace/backend/interno/persistencia/postgres"
	"mindtrace/backend/interno/persistencia/repositorios"
	"mindtrace/backend/interno/persistencia/seeds"
	"os"
	"strings"

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
	default:
		log.Fatalf("DB_DRIVER invalido: %s", dbDriver)
	}

	skipDBInit := os.Getenv("SKIP_DB_INIT") == "true"

	if !skipDBInit {
		// Executa migracoes automatizadas para alinhar esquema do banco
		err = db.AutoMigrate(
			&dominio.Usuario{},
			&dominio.RefreshToken{},
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
	}
	emailSvc := servicos.NovoEmailServico(db, usuarioRepo)
	// Inicializa servicos
	usuarioSvc := servicos.NovoUsuarioServico(db, usuarioRepo, emailSvc)
	analiseSvc := servicos.NovoAnaliseServico(db, registroHumorRepo, usuarioRepo, emailSvc)
	registroHumorSvc := servicos.NovoRegistroHumorServico(db, registroHumorRepo, usuarioRepo, analiseSvc)
	resumoSvc := servicos.NovoResumoServico(db, registroHumorRepo, usuarioRepo)
	conviteSvc := servicos.NovoConviteServico(db, conviteRepo, usuarioRepo, emailSvc)
	instrumentoSvc := servicos.NovoInstrumentoServico(db, instrumentoRepo, usuarioRepo, emailSvc)

	// Inicializa controladores
	profissionalCtrl := controladores.NovoProfissionalControlador(usuarioSvc)
	pacienteCtrl := controladores.NovoPacienteControlador(usuarioSvc)
	autCtrl := controladores.NovoAutControlador(usuarioSvc)
	usuarioCtrl := controladores.NovoUsuarioControlador(usuarioSvc, emailSvc)
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
	/*
	   TODO:
	     - INSERIR USUARIO COM EMAIL HASH NO DB
	     - USAR FLAG EstaAtivo PARA PERMITIR LOGIN
	     - ADICIONAR TIMEOUT += 48H PARA INVALIDAR TOKEN
	     - NOVO CONTROLADOR/ROTA PARA ATIVACAO DE EMAIL
	     - testar.
	*/
	api := roteador.Group("/api/v1")
	{
		// --- ROTAS PUBLICAS ---
		// Middleware de Rate Limit: 0.5 tokens/s (1 a cada 2s), burst de 5
		limiterPublico := middlewares.RateLimitMiddleware(0.5, 5)

		auth := api.Group("/entrar")
		auth.Use(limiterPublico)
		{
			auth.POST("/login", autCtrl.Login)
			auth.POST("/refresh", autCtrl.Refresh)
			auth.GET("/ativar", usuarioCtrl.AtivarConta)
			auth.POST("/ativar/reenviar", usuarioCtrl.ReenviarAtivacao)
		}

		profissionais := api.Group("/profissionais")
		profissionais.Use(limiterPublico)
		{
			// Registro de profissionais acessivel sem autenticacao
			profissionais.POST("/registrar", profissionalCtrl.Registrar)
		}

		pacientes := api.Group("/pacientes")
		pacientes.Use(limiterPublico)
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
				usuarios.GET("/paciente/profissionais", usuarioCtrl.ListarProfissionaisDoPaciente)
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
				convites.GET("/info", conviteCtrl.ObterInfo)
			}

			instrumentos := protegido.Group("/instrumentos")
			{
				instrumentos.GET("/listar-instrumentos", instrumentoCtrl.ListarInstrumentos)
				instrumentos.POST("/atribuir-instrumento", instrumentoCtrl.AtribuirInstrumento)
				instrumentos.GET("/listar-atribuicoes-paciente", instrumentoCtrl.ListarAtribuicoesPaciente)
				instrumentos.GET("/listar-atribuicoes-profissional", instrumentoCtrl.ListarAtribuicoesProfissional)
				instrumentos.GET("/atribuicao", instrumentoCtrl.ApresentarPerguntasAtribuicao)
				instrumentos.POST("/registrar-respostas", instrumentoCtrl.RegistrarRespostas)
				instrumentos.GET("/visualizar-respostas", instrumentoCtrl.VisualizarRespostas)

			}
		}
	}
	fmt.Printf("\n%s\n"+`Acesso com dados mockados:
- Profissional: 
	joao.silva@mindtrace.dev
- Pacientes:
	ana.costa@mindtrace.dev
	bruno.lima@mindtrace.dev
ano.costo@mindtrace.dev <- precisa ativar conta
- Senha:
	Password123!`+"\n%s\n", strings.Repeat("#", 50), strings.Repeat("#", 50))
	log.Println("servidor iniciado na porta 9090")
	roteador.Run(":9090")

}
