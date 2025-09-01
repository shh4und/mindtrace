# MindTrace

MindTrace é uma aplicação web full-stack projetada para rastreamento e gerenciamento de saúde mental. Permite que profissionais monitorem registros de humor dos pacientes, gerem relatórios e gerenciem notificações, enquanto os pacientes podem registrar seus humores diários e visualizar insights personalizados.

## Recursos

- **Gerenciamento de Usuários**: Registro e autenticação para pacientes e profissionais.
- **Rastreamento de Humor**: Pacientes podem registrar seu humor diário com carimbos de data/hora e notas.
- **Relatórios e Análises**: Profissionais podem gerar relatórios sobre tendências de humor dos pacientes.
- **Notificações e Alertas**: Alertas automatizados para profissionais com base nos dados dos pacientes.
- **Convites**: Profissionais podem enviar convites para pacientes para vinculação de contas.
- **Dashboard**: Dashboards separados para pacientes e profissionais com visualização de dados relevante.

## Pilha Tecnológica

### Backend
- **Linguagem**: Go (Golang)
- **Framework**: Gin
- **Banco de Dados**: PostgreSQL com GORM ORM
- **Autenticação**: JWT
- **Arquitetura**: Arquitetura Limpa (camadas de Domínio, Aplicação, Persistência)

### Frontend
- **Linguagem**: JavaScript
- **Framework**: Vue.js 3
- **Ferramenta de Build**: Vite
- **Estilização**: TailwindCSS
- **Gerenciamento de Estado**: Pinia
- **Gráficos**: ApexCharts

### Infraestrutura
- **Containerização**: Docker
- **Orquestração**: Docker Compose
- **Gerenciamento de Banco de Dados**: PgAdmin

## Estrutura do Projeto

```
.
├── backend/                 # Aplicação backend em Go
│   ├── cmd/api/             # Ponto de entrada principal
│   ├── interno/             # Pacotes internos
│   │   ├── aplicacao/       # Camada de aplicação (controladores, serviços, middlewares)
│   │   ├── dominio/         # Camada de domínio (entidades)
│   │   └── persistencia/    # Camada de persistência (repositórios)
│   ├── Dockerfile           # Dockerfile de produção
│   ├── Dockerfile.dev       # Dockerfile de desenvolvimento
│   ├── go.mod               # Módulos Go
│   └── go.sum               # Dependências Go
├── frontend/                # Aplicação frontend em Vue.js
│   ├── src/                 # Código fonte
│   │   ├── components/      # Componentes Vue
│   │   ├── views/           # Visualizações de página
│   │   ├── router/          # Configuração de roteamento
│   │   ├── store/           # Gerenciamento de estado
│   │   └── services/        # Serviços de API
│   ├── Dockerfile           # Dockerfile de produção
│   ├── Dockerfile.dev       # Dockerfile de desenvolvimento
│   ├── package.json         # Dependências Node
│   └── vite.config.js       # Configuração Vite
├── docker-compose.yml       # Docker Compose para produção
├── docker-compose.override.yml  # Substituição Docker Compose para desenvolvimento
└── README.md                # Este arquivo
```

## Pré-requisitos

- Docker e Docker Compose
- Node.js (para desenvolvimento local do frontend, opcional)
- Go (para desenvolvimento local do backend, opcional)

## Instalação e Configuração

### Usando Docker (Recomendado)

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/shh4und/mindtrace.git
   cd mindtrace
   ```

2. **Crie o arquivo de ambiente**:
   Crie um arquivo `.env` no diretório raiz com as seguintes variáveis:
   ```
   POSTGRES_USER=seu_usuario_db
   POSTGRES_PASSWORD=sua_senha_db
   POSTGRES_DB=mindtrace
   PGADMIN_DEFAULT_EMAIL=admin@exemplo.com
   PGADMIN_DEFAULT_PASSWORD=senha_admin
   JWT_SECRET=seu_segredo_jwt
   ```

3. **Execute a aplicação**:
   Para desenvolvimento:
   ```bash
   docker-compose -f docker-compose.yml -f docker-compose.override.yml up --build
   ```

   Para produção:
   ```bash
   docker-compose up --build
   ```

4. **Acesse a aplicação**:
   - Frontend: http://localhost:5173
   - API Backend: http://localhost:8080
   - PgAdmin: http://localhost:8001

### Desenvolvimento Local

#### Backend
1. Navegue para o diretório backend:
   ```bash
   cd backend
   ```

2. Instale as dependências:
   ```bash
   go mod download
   ```

3. Execute a aplicação:
   ```bash
   go run ./cmd/api/main.go
   ```

#### Frontend
1. Navegue para o diretório frontend:
   ```bash
   cd frontend
   ```

2. Instale as dependências:
   ```bash
   npm install
   ```

3. Execute o servidor de desenvolvimento:
   ```bash
   npm run dev
   ```

## Uso

1. **Registrar/Login**: Crie uma conta como paciente ou profissional.
2. **Para Pacientes**:
   - Registre humores diários no dashboard.
   - Visualize histórico pessoal de humor e tendências.
3. **Para Profissionais**:
   - Gerencie pacientes e envie convites.
   - Visualize relatórios de pacientes e análises.
   - Gere alertas com base nos dados de humor.

## Documentação da API

O backend fornece APIs RESTful. Endpoints principais incluem:

- `POST /auth/login` - Login do usuário
- `POST /auth/register` - Registro do usuário
- `GET /mood` - Obter registros de humor
- `POST /mood` - Criar registro de humor
- `GET /reports` - Gerar relatórios

Para documentação detalhada da API, consulte o código do backend ou use ferramentas como Swagger se implementado.

## Contribuição

1. Faça um fork do repositório.
2. Crie uma branch de feature: `git checkout -b feature/sua-feature`
3. Faça commit das suas mudanças: `git commit -m 'Adicione alguma feature'`
4. Faça push para a branch: `git push origin feature/sua-feature`
5. Abra um pull request.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para detalhes.

## Contato

Para dúvidas ou suporte, entre em contato com o mantenedor do projeto.