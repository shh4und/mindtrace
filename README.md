# MindTrace

[![Go Version](https://img.shields.io/badge/Go-1.25.1-blue.svg)](https://golang.org/)
[![Vue Version](https://img.shields.io/badge/Vue.js-3.5.18-green.svg)](https://vuejs.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://www.docker.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-blue.svg)](https://www.postgresql.org/)
[![Tests](https://img.shields.io/badge/Tests-281%20passing-brightgreen.svg)](#-testes)

MindTrace é uma aplicação web full-stack completa projetada para rastreamento e gerenciamento de saúde mental. Permite que profissionais de saúde monitorem registros de humor de pacientes, gerem relatórios e gerenciem notificações, enquanto pacientes podem registrar seus humores diários e visualizar insights personalizados.

## 🚀 Principais Funcionalidades

- **Gestão de Usuários**: Registro e autenticação para pacientes e profissionais de saúde
- **Rastreamento de Humor**: Pacientes podem registrar entradas diárias de humor com timestamps e notas
- **Questionários Padronizados**: Sistema completo de escalas psicométricas (PHQ-9, GAD-7, WHOQOL-BREF, WHO-5)
- **Monitoramento Ativo**: Análise automática de padrões de risco com alertas por e-mail para profissionais
- **Atribuição de Instrumentos**: Profissionais podem atribuir questionários a pacientes
- **Analytics Avançados**: Análise histórica com classificação de status (REGULAR, ATENÇÃO, PREOCUPANTE)
- **Relatórios e Análises**: Profissionais podem gerar relatórios de tendências de humor para seus pacientes
- **Sistema de Convites**: Profissionais podem enviar convites de vinculação de conta para pacientes
- **Dashboards Duplos**: Interfaces separadas baseadas em função para pacientes e profissionais
- **Visualização de Dados**: Gráficos e tabelas interativas para análise de tendências de humor

## 🛠 Stack Tecnológico

### Backend
- **Linguagem**: Go (Golang) 1.25.1
- **Framework Web**: Gin v1.10.1
- **ORM**: GORM v1.30.1
- **Tipos de Dados**: GORM Datatypes (JSONB)
- **Banco de Dados**: PostgreSQL 17 (produção) / SQLite (desenvolvimento)
- **Autenticação**: JWT (golang-jwt/jwt/v5)
- **Emails**: net/smtp (Standard Lib) com Templates HTML Embutidos
- **Testes**: Testify v1.10.0 - **281 testes unitários**
- **Arquitetura**: Clean Architecture (Domain-Driven Design)

### Frontend
- **Linguagem**: JavaScript (ES6+)
- **Framework**: Vue.js 3.5.18 (Composition API)
- **Build Tool**: Vite v7.0.6
- **Gerenciamento de Estado**: Pinia v3.0.3
- **Cliente HTTP**: Axios v1.11.0
- **Estilização**: TailwindCSS v4.1.11
- **Gráficos**: ApexCharts v5.3.4 com vue3-apexcharts v1.8.0
- **Ícones**: FontAwesome v7.0.0

### Infraestrutura & DevOps
- **Containerização**: Docker
- **Orquestração**: Docker Compose
- **CI/CD**: GitHub Actions
- **Registro de Containers**: Docker Hub
- **Provedor Cloud**: AWS EC2
- **Admin de Banco**: PgAdmin 4
- **Proxy Reverso**: Nginx
- **Controle de Versão**: Git

## 🏗 Arquitetura do Projeto

MindTrace implementa os princípios de **Clean Architecture** com clara separação de responsabilidades através de múltiplas camadas:

```
┌─────────────────────────────────────────────────────────────┐
│                    Camada de Apresentação                   │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────┐  │
│  │   Vue.js SPA    │  │   REST API      │  │   Nginx     │  │
│  │  (Frontend)     │  │   (Gin)         │  │ (Proxy      │  │
│  └─────────────────┘  └─────────────────┘  │  Reverso)   │  │
└─────────────────────────────────────────────┼─────────────┘
                                              │
┌─────────────────────────────────────────────┼─────────────┐
│                 Camada de Aplicação         │             │
│  ┌─────────────────┐  ┌─────────────────┐   │             │
│  │  Controladores  │  │   Serviços      │   │             │
│  │ (HTTP Handlers) │  │(Lógica Negócio) │   │             │
│  └─────────────────┘  └─────────────────┘   │             │
└─────────────────────────────────────────────┼─────────────┘
                                              │
┌─────────────────────────────────────────────┼─────────────┐
│                 Camada de Domínio           │             │
│  ┌─────────────────┐                        │             │
│  │  Entidades      │                        │             │
│  │  (Modelos de    │                        │             │
│  │   Negócio)      │                        │             │
│  └─────────────────┘                        │             │
└─────────────────────────────────────────────┼─────────────┘
                                              │
┌─────────────────────────────────────────────┼─────────────┐
│               Camada de Persistência        │             │
│  ┌─────────────────┐  ┌─────────────────┐   │             │
│  │ Interfaces de   │  │ Implementações  │   │             │
│  │ Repositório     │  │   Database      │   │             │
│  └─────────────────┘  └─────────────────┘   │             │
└─────────────────────────────────────────────────────────────┘
```

### Princípios Arquiteturais
- **Inversão de Dependência**: Camadas internas não dependem de camadas externas
- **Responsabilidade Única**: Cada camada tem um propósito distinto
- **Segregação de Interface**: Interfaces de repositório definem contratos claros
- **Testabilidade**: Injeção de dependência permite testes abrangentes (281 testes unitários)

## 🚀 Começando

### Pré-requisitos
- Docker e Docker Compose (para desenvolvimento local)
- Node.js 22.17.1 (opcional, para desenvolvimento frontend local)
- Go 1.25.1 (opcional, para desenvolvimento backend local)

### Configuração para Desenvolvimento Local

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/shh4und/mindtrace.git
   cd mindtrace
   ```

2. **Crie o arquivo de ambiente**:
   Crie um arquivo `.env` no diretório raiz:
   ```env
   POSTGRES_USER=seu_usuario_db
   POSTGRES_PASSWORD=sua_senha_db
   POSTGRES_DB=mindtrace
   PGADMIN_DEFAULT_EMAIL=admin@exemplo.com
   PGADMIN_DEFAULT_PASSWORD=senha_admin
   JWT_SECRET=sua_chave_secreta_jwt
   ```

3. **Inicie a aplicação**:
   ```bash
   # Para desenvolvimento (com hot reload)
   docker-compose -f docker-compose.yml -f docker-compose.override.yml up --build

   # Para desenvolvimento com SQLite (configuração mais leve)
   docker-compose -f docker-compose.sqlite.yml up --build
   ```

4. **Acesse a aplicação**:
   - Frontend: http://localhost
   - API Backend: http://localhost/api/v1
   - PgAdmin: http://localhost:5050

### Desenvolvimento Local (Sem Docker)

#### Backend (Go)
```bash
cd backend
go mod download
go run cmd/api/main.go
```

#### Frontend (Vue.js)
```bash
cd frontend
npm install
npm run dev
```

## 🚀 Deploy em Produção

MindTrace usa **CI/CD** com GitHub Actions para deploy automatizado na AWS:

### Processo de Deploy
1. **Push para branch main** dispara workflow do GitHub Actions
2. **Build** de imagens Docker para backend e frontend
3. **Push** de imagens para Docker Hub
4. **Deploy** na instância AWS EC2 via SSH
5. **Atualização** de containers com zero-downtime

### Configuração de Ambiente de Produção
- **Provedor Cloud**: AWS EC2
- **Registro de Containers**: Docker Hub
- **CI/CD**: GitHub Actions
- **Banco de Dados**: PostgreSQL (gerenciado)
- **Proxy Reverso**: Nginx

### Configuração de Deploy
O deploy em produção requer estes secrets no repositório GitHub:
- `DOCKER_HUB_USERNAME`: Nome de usuário Docker Hub
- `DOCKER_HUB_TOKEN`: Token de acesso Docker Hub
- `EC2_HOST`: IP/hostname da instância AWS EC2
- `EC2_USER`: Nome de usuário SSH do EC2
- `EC2_SSH_KEY`: Chave SSH privada para acesso EC2
- `FRONTEND_API_BASE_URL`: URL da API em produção (opcional, padrão localhost)

### Deploy Manual (se necessário)
```bash
# No servidor de produção
cd /home/ubuntu/mindtrace
git pull origin main
docker compose -f docker-compose.prod.yml --env-file .env.prod pull
docker compose -f docker-compose.prod.yml --env-file .env.prod up -d --remove-orphans
```

## 📁 Estrutura do Projeto

```
mindtrace/
├── backend/                          # Aplicação backend em Go
│   ├── cmd/api/                      # Entry point
│   │   └── main.go
│   ├── interno/                      # Pacotes internos
│   │   ├── dominio/                  # Camada de domínio (Entities)
│   │   │   ├── usuario.go            # Usuário, Profissional, Paciente
│   │   │   ├── registro_humor.go     # Registro de humor
│   │   │   ├── convite.go            # Sistema de convites
│   │   │   ├── relatorio.go          # DTOs de relatório
│   │   │   ├── notificacao.go        # Sistema de notificações
│   │   │   ├── instrumento.go        # ✨ Questionários padronizados
│   │   │   ├── atribuicao.go         # ✨ Atribuições de questionários
│   │   │   ├── resposta.go           # ✨ Respostas (JSONB)
│   │   │   └── tests/                # ✅ Testes de domínio (142 testes)
│   │   ├── aplicacao/                # Camada de aplicação
│   │   │   ├── controladores/        # HTTP Controllers
│   │   │   │   ├── aut_controlador.go
│   │   │   │   ├── usuario_controlador.go
│   │   │   │   ├── paciente_controlador.go
│   │   │   │   ├── profissional_controlador.go
│   │   │   │   ├── registro_humor_controlador.go
│   │   │   │   ├── convite_controlador.go
│   │   │   │   ├── relatorio_controlador.go
│   │   │   │   ├── resumo_controlador.go
│   │   │   │   └── instrumento_controlador.go  # ✨ API de questionários
│   │   │   ├── servicos/             # Business Logic
│   │   │   │   ├── usuario_servico.go
│   │   │   │   ├── registro_humor_servico.go
│   │   │   │   ├── convite_servico.go
│   │   │   │   ├── resumo_servico.go
│   │   │   │   ├── alerta_servico.go          # ✨ Sistema de alertas
│   │   │   │   ├── analise_servico.go         # ✨ Analytics avançados (Monitoramento)
│   │   │   │   ├── notificacao_servico.go
│   │   │   │   ├── email_servico.go           # ✨ Envio de emails (SMTP)
│   │   │   │   ├── instrumento_servico.go     # ✨ Lógica de questionários
│   │   │   │   ├── templates/                 # 📧 Templates HTML (embed)
│   │   │   │   │   ├── ativacao.html
│   │   │   │   │   ├── atribuicao.html
│   │   │   │   │   ├── convite_profissional.html
│   │   │   │   │   └── monitoramento.html
│   │   │   │   └── tests/            # ✅ Testes de serviços (74 testes)
│   │   │   ├── dtos/                 # Data Transfer Objects
│   │   │   │   └── tipos.go
│   │   │   ├── helpers/              # Helper functions
│   │   │   │   └── pdf.go            # Geração de PDF
│   │   │   ├── mappers/              # DTO ↔ Entity conversão
│   │   │   │   ├── utils.go
│   │   │   │   └── tests/            # ✅ Testes de mappers (23 testes)
│   │   │   └── middlewares/          # HTTP Middlewares
│   │   │       ├── aut_middleware.go
│   │   │       └── cors_middleware.go
│   │   └── persistencia/             # Camada de persistência
│   │       ├── repositorios/         # Repository interfaces
│   │       │   └── repositorios.go
│   │       ├── postgres/             # PostgreSQL implementation
│   │       │   ├── db.go
│   │       │   ├── usuario_repositorio.go
│   │       │   ├── registro_humor_repositorio.go
│   │       │   ├── convite_repositorio.go
│   │       │   ├── relatorio_repositorio.go
│   │       │   ├── notificacao_repositorio.go
│   │       │   └── instrumento_repositorio.go  # ✨ Repos questionários
│   │       ├── sqlite/               # SQLite implementation
│   │       │   ├── db.go
│   │       │   ├── usuario_repositorio.go
│   │       │   ├── registro_humor_repositorio.go
│   │       │   ├── convite_repositorio.go
│   │       │   ├── relatorio_repositorio.go
│   │       │   ├── notificacao_repositorio.go
│   │       │   └── instrumento_repositorio.go  # ✨ Repos questionários
│   │       └── seeds/                # Database seeding
│   │           └── instrumento_seed.go
│   ├── Dockerfile                    # Production container
│   ├── Dockerfile.dev                # Development container
│   ├── go.mod                        # Go modules
│   └── go.sum                        # Dependencies
├── frontend/                         # Aplicação frontend Vue.js
│   ├── src/
│   │   ├── views/                    # Views (Páginas)
│   │   │   ├── auth/                 # Autenticação
│   │   │   │   ├── Login.vue
│   │   │   │   ├── Cadastro.vue
│   │   │   │   └── ForgotPassword.vue
│   │   │   ├── dashboard-paciente/   # Dashboard Paciente
│   │   │   │   ├── PacienteDashboard.vue
│   │   │   │   ├── RegistroHumor.vue
│   │   │   │   ├── Resumo.vue
│   │   │   │   ├── VincularProfissional.vue
│   │   │   │   ├── QuestionariosAtribuidos.vue    # ✨ Lista de questionários
│   │   │   │   └── ResponderQuestionario.vue      # ✨ Interface de resposta
│   │   │   ├── dashboard-profissional/ # Dashboard Profissional
│   │   │   │   ├── ProfissionalDashboard.vue
│   │   │   │   ├── GerarConvite.vue
│   │   │   │   ├── ListaPacientes.vue
│   │   │   │   ├── AtribuirQuestionario.vue       # ✨ Atribuir questionários
│   │   │   │   └── QuestionariosAtribuidos.vue    # ✨ Gerenciar atribuições
│   │   │   ├── shared/               # Views compartilhadas
│   │   │   │   ├── EditarPerfil.vue
│   │   │   │   └── Relatorio.vue
│   │   │   └── Landpage.vue          # Página inicial
│   │   ├── components/               # Componentes reutilizáveis
│   │   │   ├── layout/               # Layout components
│   │   │   │   ├── NavbarPublic.vue
│   │   │   │   ├── Sidebar.vue
│   │   │   │   ├── SidebarPaciente.vue
│   │   │   │   ├── SidebarProfissional.vue
│   │   │   │   └── TopNavbar.vue
│   │   │   └── ui/                   # UI components
│   │   │       ├── BaseButton.vue
│   │   │       ├── BaseCard.vue
│   │   │       ├── BaseInput.vue
│   │   │       └── index.js
│   │   ├── composables/              # Composition API
│   │   │   ├── index.js
│   │   │   ├── useAuth.js            # Lógica de autenticação
│   │   │   └── useMoodForm.js        # Lógica de formulário
│   │   ├── services/                 # API Clients
│   │   │   └── api.js                # Cliente HTTP (Axios)
│   │   ├── router/                   # Vue Router
│   │   │   └── index.js
│   │   ├── store/                    # Pinia State Management
│   │   │   └── user.js               # User store
│   │   ├── utils/                    # Utilities
│   │   │   └── jwt.js                # JWT helper
│   │   ├── types/                    # Type definitions
│   │   ├── constants/                # Constants
│   │   ├── assets/                   # Static assets
│   │   ├── App.vue                   # Root component
│   │   └── main.js                   # Entry point
│   ├── Dockerfile                    # Container de produção
│   ├── Dockerfile.dev                # Container de desenvolvimento
│   ├── package.json                  # Dependências Node
│   └── vite.config.js                # Configuração Vite
├── docs/                             # Documentação
│   ├── ARQUITETURA_MINDTRACE.md      # 📘 Documento principal de arquitetura
│   └── TESTES_UNITARIOS_RELATORIO.md # Documentação de testes
├── seeders/                          # Scripts de seed do banco
│   ├── seed.sh                       # Script principal de seed
│   ├── seed.ps1                      # Script PowerShell
│   ├── monitoramento.sh              # Seed de monitoramento
│   ├── registro_humor.sh             # Seed de registros de humor
│   └── vincular.sh                   # Seed de vínculos
├── docker-compose.yml                # Docker Compose base
├── docker-compose.override.yml       # Overrides de desenvolvimento
├── docker-compose.prod.yml           # Configuração de produção
├── docker-compose.sqlite.yml         # Configuração SQLite
├── dump-mindtrace-public.sql         # Dump do banco de dados
└── README.md                         # Este arquivo
```

## 🔧 Fluxo de Desenvolvimento

### Estratégia de Branches
- `main`: Código pronto para produção (branch protegida com CI/CD)
- `feature/*`: Novas funcionalidades e melhorias
- `docs/*`: Documentação nova ou atualizada
- `bugfix/*`: Correções de bugs
- `hotfix/*`: Correções críticas de produção

### Processo de Desenvolvimento
1. **Criar Branch de Feature**: `git checkout -b feature/nova-funcionalidade`
2. **Fazer Alterações**: Implementar features seguindo princípios de Clean Architecture
3. **Executar Testes**: Executar suítes de teste para backend e frontend
4. **Code Review**: Submeter pull request para revisão
5. **Merge**: Squash merge para main após aprovação
6. **Auto-deploy**: GitHub Actions automaticamente faz build e deploy para produção

### Pipeline CI/CD
- **Trigger**: Push para branch `main` ou dispatch manual
- **Build**: Builds Docker multi-stage para imagens otimizadas
- **Test**: Testes automatizados (testes unitários backend)
- **Deploy**: Deploy zero-downtime para AWS EC2
- **Monitoramento**: Health checks de containers e agregação de logs

### Ativando CI/CD
O workflow de deploy está atualmente desabilitado (`.github/workflows/deploy.yml.disabled`). Para ativar:

1. Renomeie `.github/workflows/deploy.yml.disabled` para `.github/workflows/deploy.yml`
2. Configure os secrets necessários no seu repositório GitHub:
   - `DOCKER_HUB_USERNAME`
   - `DOCKER_HUB_TOKEN`
   - `EC2_HOST`
   - `EC2_USER`
   - `EC2_SSH_KEY`
   - `FRONTEND_API_BASE_URL` (opcional)

### Gerenciamento de Banco de Dados
- **Migrações**: Automáticas via GORM AutoMigrate
- **Seeding**: Use `seed.sh` para dados iniciais
- **Backup**: Backups regulares do PostgreSQL em produção

## 📊 Monitoramento & Observabilidade

### Monitoramento em Produção
- **Health Checks**: Endpoints de saúde dos containers
- **Logs**: Logging centralizado com Docker Compose
- **Métricas**: Monitoramento de performance da aplicação
- **Alertas**: Notificações automatizadas para problemas do sistema

### Rastreamento de Erros
- **Log de Erros**: Logging estruturado em produção
- **Tratamento de Exceções**: Respostas de erro elegantes
- **Informações de Debug**: Detalhes de erro específicos por ambiente

## 📋 Padrões de Código

### Padrões Backend Go
- **Formatação**: `gofmt` e `goimports` para formatação consistente
- **Nomenclatura**: PascalCase para exportados, camelCase para não exportados
- **Tratamento de Erros**: Retornos de erro explícitos, sem panics em produção
- **Documentação**: Comentários Godoc para todas as funções exportadas
- **Testes**: Testes table-driven com asserções testify

### Padrões Frontend Vue.js
- **Composition API**: Usar Vue 3 Composition API ao invés de Options API
- **Nomenclatura de Componentes**: PascalCase para arquivos de componentes
- **Gerenciamento de Estado**: Stores Pinia para estado global
- **Estilização**: Classes utilitárias TailwindCSS
- **TypeScript**: Considerar migração para melhor type safety

### Padrões Gerais
- **Commits**: Commits convencionais (`feat:`, `fix:`, `docs:`)
- **Documentação**: Atualizar README e docs para mudanças significativas
- **Segurança**: Validação de entrada, JWT para autenticação
- **Performance**: Otimizar queries de banco, lazy loading para componentes

## 🧪 Testes

### Infraestrutura de Testes Backend

O MindTrace possui uma robusta infraestrutura de testes com **281 testes unitários** distribuídos em três camadas:

#### Estatísticas de Testes

| Camada | Módulo | Testes | Status |
|--------|--------|--------|--------|
| **Domínio** | Usuario | 62 | ✅ 100% |
| **Domínio** | RegistroHumor | 45 | ✅ 100% |
| **Domínio** | Convite | 35 | ✅ 100% |
| **Serviços** | UsuarioServico | 28 | ✅ Completo |
| **Serviços** | RelatorioServico | 17 | ✅ Completo |
| **Serviços** | RegistroHumorServico | 13 | ✅ Completo |
| **Serviços** | ConviteServico | 13 | ✅ Completo |
| **Mappers** | Utils | 23 | ✅ Completo |
| **TOTAL** | **8 módulos** | **281** | ✅ **Todos passando** |

#### Padrões de Teste

**Table-Driven Tests:**
```go
tests := []struct {
    name    string
    input   string
    wantErr bool
}{
    {"válido", "email@exemplo.com", false},
    {"inválido", "invalido", true},
}
```

**Características:**
- **Framework**: Testify para asserções e organização de testes
- **Cobertura**: Testes unitários para serviços, repositórios e controladores
- **Mocking**: Injeção de dependência baseada em interfaces facilita mocking
- **Banco de Dados**: SQLite para testes de integração rápidos e isolados
- **Organização**: Testes em subdiretórios `/tests` dedicados
- **Tempo de Execução**: < 1 segundo para toda a suíte

### Testes Frontend
- **Framework**: Vue Test Utils (planejado)
- **Cobertura**: Testes de componentes e serviços
- **E2E**: Playwright ou Cypress para testes end-to-end (planejado)

### Executando Testes

```bash
# Testes backend
cd backend
go test ./interno/dominio/tests ./interno/aplicacao/servicos/tests ./interno/aplicacao/mappers/tests

# Testes com verbose
go test ./interno/dominio/tests ./interno/aplicacao/servicos/tests ./interno/aplicacao/mappers/tests -v

# Testes frontend (quando implementado)
cd frontend
npm run test
```

## 📚 Documentação Adicional

- **[Arquitetura MindTrace](./docs/ARQUITETURA_MINDTRACE.md)** - Documento principal de arquitetura (PT-BR)
- **[Relatório de Testes Unitários](./docs/TESTES_UNITARIOS_RELATORIO.md)** - Documentação detalhada dos testes
- **[Especificação API](./frontend/swagger-output.json)** - Especificação OpenAPI
- **[Schema do Banco de Dados](./schema_dump.sql)** - Schema PostgreSQL

**MindTrace** - Capacitando profissionais de saúde mental e pacientes através da tecnologia.
