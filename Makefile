# Variáveis
APP_NAME = mindtrace
DOCKER_COMPOSE = docker compose
DATABASE = pgadmin-data 
SGBD = postgres-data

# Alvo padrão
.PHONY: help
help:
	@echo "Comandos disponíveis:"
	@echo "  build     - Constrói a aplicação"
	@echo "  clean     - Limpa arquivos temporários"
	@echo "  docker-up - Inicia serviços com Docker Compose"
	@echo "  docker-down - Para serviços com Docker Compose"

#    @echo "  run       - Executa a aplicação"


# Constrói a aplicação (exemplo genérico; ajuste para sua linguagem)
.PHONY: build
build:
	@echo "Buildando containers docker..."
	systemctl --user start docker-desktop
	@echo "Aguardando docker-desktop iniciar"
	until docker info >/dev/null 2>&1; do sleep 1; done
	$(DOCKER_COMPOSE) up --build -d

# Limpa arquivos
.PHONY: clean
clean-db-dirs:
	@echo "Resetando banco de dados..."
	rm -rf $(DATABASE) $(SGBD) 

# Inicia Docker Compose
.PHONY: docker-up
docker-up:
	@echo "Iniciando docker e app..."
	systemctl --user start docker-desktop
	@echo "Aguardando docker-desktop iniciar"
	until docker info >/dev/null 2>&1; do sleep 1; done
	$(DOCKER_COMPOSE) up -d

# Para Docker Compose
.PHONY: docker-down
docker-down:
	@echo "Encerrando docker e app..."
	$(DOCKER_COMPOSE) down -v

