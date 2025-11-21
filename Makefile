# Variáveis
APP_NAME = mindtrace
DOCKER_COMPOSE = docker compose
DATABASE = postgres-data

# Alvo padrão
.PHONY: help
help:
	@echo "Comandos disponíveis:"
	@echo "  build        - builda os containers docker"
	@echo "  clean-db     - apaga os arquivos do banco de dados (reseta o banco)"
	@echo "  up           - inicia os containers docker"
	@echo "  down         - para os containers docker"

.PHONY: build
build:
	@echo "Buildando containers docker..."
	systemctl --user start docker-desktop
	@echo "Aguardando docker-desktop iniciar"
	until docker info >/dev/null 2>&1; do sleep 1; done
	$(DOCKER_COMPOSE) up --build -d

.PHONY: clean
clean-db:
	@echo "Resetando banco de dados..."
	rm -rf $(DATABASE) 

.PHONY: docker-up
up:
	@echo "Iniciando docker e app..."
	systemctl --user start docker-desktop
	@echo "Aguardando docker-desktop iniciar"
	until docker info >/dev/null 2>&1; do sleep 1; done
	$(DOCKER_COMPOSE) up -d

.PHONY: docker-down
down:
	@echo "Encerrando docker e app..."
	$(DOCKER_COMPOSE) down -v

