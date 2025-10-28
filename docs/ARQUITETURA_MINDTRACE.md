# Arquitetura de Software - MindTrace MVP v1.0

**Projeto:** P2410 - Aplicativo para Monitoramento de Saúde Mental  
**Data:** 26 de Outubro de 2025  
**Autor:** Alexander Nunes Souza  
**Orientadora:** Profa. Dra. Adicinéia A. de Oliveira

---

## 📋 Sumário

1. [Visão Geral da Arquitetura](#1-visão-geral-da-arquitetura)
2. [Decisões Arquiteturais](#2-decisões-arquiteturais)
3. [Visões Arquiteturais](#3-visões-arquiteturais)
4. [Padrões e Práticas](#4-padrões-e-práticas)
5. [Justificativas](#5-justificativas)
6. [Atributos de Qualidade Atendidos](#6-atributos-de-qualidade-atendidos)
7. [Infraestrutura de Testes](#7-infraestrutura-de-testes)
8. [Débitos Técnicos Identificados](#8-débitos-técnicos-identificados)
9. [Histórico de Atualizações](#9-histórico-de-atualizações)

---

## 1. Visão Geral da Arquitetura

### 1.1 Estilo Arquitetural Principal

**Arquitetura em Camadas + Cliente-Servidor + Clean Architecture**

O sistema MindTrace adota uma arquitetura híbrida que combina:

1. **Cliente-Servidor (2-Tier):**
   - **Cliente:** SPA (Single Page Application) em Vue 3
   - **Servidor:** API REST em Go (Golang)

2. **Clean Architecture (Backend):**
   - Separação em camadas concêntricas
   - Dependências apontando para o domínio
   - Domain-Driven Design (DDD)

3. **Component-Based Architecture (Frontend):**
   - Componentes reutilizáveis Vue 3
   - Composição e reatividade
   - Store centralizado (Pinia/Vuex)

### 1.2 Stack Tecnológico

#### Backend
- **Linguagem:** Go 1.23+
- **Framework Web:** Gin
- **ORM:** GORM
- **Autenticação:** JWT (golang-jwt/jwt)
- **Banco de Dados:** PostgreSQL 17
- **Criptografia:** bcrypt (senhas)

#### Frontend
- **Framework:** Vue 3 (Composition API)
- **Build Tool:** Vite
- **CSS Framework:** Tailwind CSS
- **HTTP Client:** Axios
- **Roteamento:** Vue Router

#### Infraestrutura
- **Containerização:** Docker & Docker Compose
- **Proxy Reverso:** Nginx
- **Controle de Versão:** Git/GitHub

---

## 2. Decisões Arquiteturais

### D1: Separação Backend/Frontend (Cliente-Servidor)

**Decisão:** Implementar backend e frontend como aplicações separadas e independentes.

**Justificativa:**
- **Manutenibilidade:** Equipes podem trabalhar independentemente
- **Escalabilidade:** Backend e frontend podem escalar separadamente
- **Flexibilidade:** Permite múltiplos clientes (web, mobile futuro)
- **Performance:** Go oferece alta performance para APIs REST

**Trade-offs:**
- ✅ Vantagens: Desacoplamento, escalabilidade, performance
- ⚠️ Desvantagens: Maior complexidade de deployment, necessidade de CORS

### D2: Clean Architecture no Backend

**Decisão:** Estruturar backend em camadas (Domain → Application → Persistence).

**Justificativa:**
- **Testabilidade:** Regras de negócio isoladas e testáveis
- **Independência de frameworks:** Domínio não depende de Gin ou GORM
- **Manutenibilidade:** Mudanças em infraestrutura não afetam domínio
- **Clareza:** Separação clara de responsabilidades

**Camadas Implementadas:**
```
cmd/api/              → Entry point (main)
interno/
  ├── dominio/        → Entities, Value Objects, Domain Logic
  ├── aplicacao/      → Use Cases, Services, DTOs, Controllers
  └── persistencia/   → Repositories, Database Implementation
```

**Trade-offs:**
- ✅ Vantagens: Testável, manutenível, escalável
- ⚠️ Desvantagens: Mais arquivos, curva de aprendizado

### D3: PostgreSQL como Banco de Dados

**Decisão:** Utilizar PostgreSQL como SGBD relacional.

**Justificativa:**
- **ACID:** Transações garantidas (crítico para saúde mental)
- **Constraints:** Validações no nível do banco (check constraints)
- **Tipos de Dados:** Suporte a JSON, timestamps com timezone
- **Open Source:** Sem custos de licença
- **Maturidade:** Banco estável e confiável

**Trade-offs:**
- ✅ Vantagens: Confiável, features robustas, open source
- ⚠️ Desvantagens: Requer mais setup que SQLite

### D4: JWT para Autenticação

**Decisão:** Autenticação stateless com JSON Web Tokens.

**Justificativa:**
- **Stateless:** Sem necessidade de sessões no servidor
- **Escalável:** Facilita load balancing
- **Cross-domain:** Funciona bem com SPA
- **Padrão:** Amplamente adotado e suportado

**Implementação:**
- Token expira em 24h
- Armazenado em localStorage no cliente
- Middleware valida em todas as rotas protegidas

**Trade-offs:**
- ✅ Vantagens: Escalável, stateless, padrão
- ⚠️ Desvantagens: Dificulta revogação, vulnerável a XSS se mal implementado

### D5: Soft Delete para Conformidade LGPD

**Decisão:** Implementar exclusão lógica (soft delete) com campo `deleted_at`.

**Justificativa:**
- **Auditoria:** Manter histórico de dados
- **Recuperação:** Possibilidade de desfazer exclusões
- **LGPD:** Facilita portabilidade de dados
- **Integridade:** Evita quebra de referências

**Implementação:**
- Todas as entidades principais possuem `deleted_at`
- Queries filtram automaticamente registros deletados (GORM)
- Hard delete só em casos específicos

**Trade-offs:**
- ✅ Vantagens: Auditável, recuperável, compatível com LGPD
- ⚠️ Desvantagens: Aumenta tamanho do BD, necessita limpeza periódica

---

## 3. Visões Arquiteturais

### 3.1 Visão Lógica (Camadas)

```
┌─────────────────────────────────────────────────────────────┐
│                    PRESENTATION LAYER                        │
│  ┌──────────────────────────────────────────────────────┐   │
│  │         Vue 3 SPA (Single Page Application)          │   │
│  │  - Components (Dashboard, RegistroHumor, etc)        │   │
│  │  - Views (Pages)                                     │   │
│  │  - Router (Navigation)                               │   │
│  │  - Store (State Management)                          │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                            │ HTTP/REST
                            │ JSON
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                   APPLICATION LAYER (API)                    │
│  ┌──────────────────────────────────────────────────────┐   │
│  │              Controllers (Gin Handlers)              │   │
│  │  - UsuarioController, HumorController, etc          │   │
│  │  - Validação de entrada (DTOs)                      │   │
│  │  - Serialização/Desserialização JSON                │   │
│  └──────────────────────────────────────────────────────┘   │
│                            │                                 │
│  ┌──────────────────────────────────────────────────────┐   │
│  │                   Services (Use Cases)               │   │
│  │  - UsuarioServico, RegistroHumorServico, etc        │   │
│  │  - Lógica de aplicação                              │   │
│  │  - Orquestração de repositórios                     │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                      DOMAIN LAYER                            │
│  ┌──────────────────────────────────────────────────────┐   │
│  │                 Domain Entities                      │   │
│  │  - Usuario, Profissional, Paciente                  │   │
│  │  - RegistroHumor, Convite, Notificacao              │   │
│  │  - Validações de negócio                            │   │
│  │  - Regras de domínio                                │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                   PERSISTENCE LAYER                          │
│  ┌──────────────────────────────────────────────────────┐   │
│  │              Repositories (Interfaces)               │   │
│  │  - IUsuarioRepositorio, IHumorRepositorio, etc      │   │
│  └──────────────────────────────────────────────────────┘   │
│                            │                                 │
│  ┌──────────────────────────────────────────────────────┐   │
│  │        Repository Implementations (GORM)             │   │
│  │  - SQLite/PostgreSQL implementations                │   │
│  │  - Queries, Migrations                              │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                   DATABASE LAYER                             │
│              PostgreSQL 17 (Relational DB)                   │
│  - 7 tabelas (usuarios, pacientes, profissionais, etc)      │
│  - Constraints (PK, FK, Unique, Check)                       │
│  - Índices (deleted_at, foreign keys)                        │
└─────────────────────────────────────────────────────────────┘
```

### 3.2 Visão de Processos (Runtime)

```
┌──────────────┐              ┌──────────────┐
│   Browser    │              │   Browser    │
│  (Cliente)   │              │  (Cliente)   │
└──────┬───────┘              └──────┬───────┘
       │                              │
       │ HTTPS                        │ HTTPS
       │ JWT Token                    │ JWT Token
       │                              │
       ▼                              ▼
┌────────────────────────────────────────────┐
│            Nginx (Proxy Reverso)           │
│  - Load Balancing                          │
│  - TLS Termination                         │
│  - Static Files (frontend)                 │
└────────┬───────────────────────────────────┘
         │
         │ HTTP
         ▼
┌────────────────────────────────────────────┐
│       Go Backend API (Gin Server)          │
│  ┌──────────────────────────────────────┐  │
│  │      Middleware Chain:               │  │
│  │  1. CORS                             │  │
│  │  2. Logger                           │  │
│  │  3. JWT Auth (rotas protegidas)     │  │
│  │  4. Error Handler                   │  │
│  └──────────────────────────────────────┘  │
│  ┌──────────────────────────────────────┐  │
│  │      Route Handlers (Controllers)    │  │
│  │  - /api/login                        │  │
│  │  - /api/registrar                    │  │
│  │  - /api/registros-humor              │  │
│  │  - /api/pacientes                    │  │
│  │  - /api/convites                     │  │
│  └──────────────────────────────────────┘  │
└────────┬───────────────────────────────────┘
         │
         │ SQL Queries (GORM)
         ▼
┌────────────────────────────────────────────┐
│          PostgreSQL Database               │
│  - Connection Pool (max 10 conexões)       │
│  - Transactions (ACID)                     │
│  - Constraints e Triggers                  │
└────────────────────────────────────────────┘
```

### 3.3 Visão Física (Deployment)

```
┌─────────────────────────────────────────────────────────┐
│                    Docker Host                          │
│                                                         │
│  ┌───────────────────────────────────────────────────┐ │
│  │  Container: frontend (Nginx + Vue SPA)            │ │
│  │  - Porta: 80 → 8080                               │ │
│  │  - Volume: ./frontend/dist:/usr/share/nginx/html  │ │
│  └───────────────────────────────────────────────────┘ │
│                           │                             │
│  ┌───────────────────────────────────────────────────┐ │
│  │  Container: backend (Go API)                      │ │
│  │  - Porta: 8080 → 8080                             │ │
│  │  - Volume: ./backend:/app                         │ │
│  │  - ENV: DATABASE_URL, JWT_SECRET                  │ │
│  └───────────────────────────────────────────────────┘ │
│                           │                             │
│  ┌───────────────────────────────────────────────────┐ │
│  │  Container: postgres (PostgreSQL 17)              │ │
│  │  - Porta: 5432 → 5432                             │ │
│  │  - Volume: postgres-data:/var/lib/postgresql/data │ │
│  │  - ENV: POSTGRES_USER, POSTGRES_PASSWORD          │ │
│  └───────────────────────────────────────────────────┘ │
│                                                         │
│  ┌───────────────────────────────────────────────────┐ │
│  │  Container: pgadmin (Administração BD)            │ │
│  │  - Porta: 5050 → 80                               │ │
│  │  - Volume: pgadmin-data:/var/lib/pgadmin          │ │
│  └───────────────────────────────────────────────────┘ │
│                                                         │
│  Network: mindtrace-network (bridge)                   │
└─────────────────────────────────────────────────────────┘
```

### 3.4 Visão de Desenvolvimento (Estrutura de Código)

```
mindtrace/
├── backend/
│   ├── cmd/api/
│   │   └── main.go                    # Entry point
│   ├── interno/
│   │   ├── dominio/                   # DOMAIN LAYER
│   │   │   ├── usuario.go
│   │   │   ├── registro_humor.go
│   │   │   ├── convite.go
│   │   │   ├── relatorio.go           # DTO de saída (não persiste)
│   │   │   ├── notificacao.go
│   │   │   └── tests/                 # ✅ TESTES DE DOMÍNIO
│   │   │       ├── usuario_test.go         (62 testes)
│   │   │       ├── registro_humor_test.go  (45 testes)
│   │   │       └── convite_test.go         (35 testes)
│   │   ├── aplicacao/                 # APPLICATION LAYER
│   │   │   ├── controladores/
│   │   │   │   ├── aut_controlador.go
│   │   │   │   ├── usuario_controlador.go
│   │   │   │   ├── paciente_controlador.go
│   │   │   │   ├── profissional_controlador.go
│   │   │   │   ├── registro_humor_controlador.go
│   │   │   │   ├── convite_controlador.go
│   │   │   │   ├── relatorio_controlador.go
│   │   │   │   └── resumo_controlador.go
│   │   │   ├── servicos/
│   │   │   │   ├── usuario_servico.go
│   │   │   │   ├── registro_humor_servico.go
│   │   │   │   ├── convite_servico.go
│   │   │   │   ├── relatorio_servico.go
│   │   │   │   ├── resumo_servico.go
│   │   │   │   ├── alerta_servico.go
│   │   │   │   ├── notificacao_servico.go
│   │   │   │   └── tests/             # ✅ TESTES DE SERVIÇOS
│   │   │   │       ├── usuario_servico_test.go         (28 testes)
│   │   │   │       ├── relatorio_servico_test.go       (17 testes)
│   │   │   │       ├── registro_humor_servico_test.go  (13 testes)
│   │   │   │       └── convite_servico_test.go         (13 testes)
│   │   │   ├── dtos/
│   │   │   │   └── tipos.go
│   │   │   ├── mappers/
│   │   │   │   ├── utils.go
│   │   │   │   └── tests/             # ✅ TESTES DE MAPPERS
│   │   │   │       └── utils_test.go       (23 testes)
│   │   │   └── middlewares/
│   │   │       ├── aut_middleware.go
│   │   │       └── cors_middleware.go
│   │   └── persistencia/              # PERSISTENCE LAYER
│   │       ├── repositorios/
│   │       │   └── repositorios.go    # Interfaces
│   │       ├── postgres/              # Implementação PostgreSQL
│   │       │   ├── db.go
│   │       │   ├── usuario_repositorio.go
│   │       │   ├── registro_humor_repositorio.go
│   │       │   ├── convite_repositorio.go
│   │       │   ├── relatorio_repositorio.go
│   │       │   └── notificacao_repositorio.go
│   │       └── sqlite/                # Implementação SQLite
│   │           ├── db.go
│   │           ├── usuario_repositorio.go
│   │           ├── registro_humor_repositorio.go
│   │           ├── convite_repositorio.go
│   │           ├── relatorio_repositorio.go
│   │           └── notificacao_repositorio.go
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
│
├── frontend/
│   ├── src/
│   │   ├── views/                     # VIEWS (Pages)
│   │   │   ├── auth/
│   │   │   ├── dashboard-paciente/
│   │   │   └── dashboard-profissional/
│   │   ├── components/                # COMPONENTS
│   │   │   └── shared/
│   │   ├── services/                  # API CLIENTS
│   │   │   └── api.js
│   │   ├── router/                    # ROUTING
│   │   │   └── index.js
│   │   ├── store/                     # STATE MANAGEMENT
│   │   │   └── user.js
│   │   ├── App.vue
│   │   └── main.js
│   ├── package.json
│   ├── vite.config.js
│   └── Dockerfile
│
├── docs/
│   ├── ARQUITETURA_MINDTRACE.md       # ✅ Este documento
│   ├── TESTES_UNITARIOS_RELATORIO.md  # ✅ Documentação de testes
│   └── Project_Architecture_Blueprint.md
│
└── docker-compose.yml
```

---

## 4. Padrões e Práticas

### 4.1 Padrões de Projeto Utilizados

#### Repository Pattern
- **Onde:** Camada de persistência
- **Por quê:** Abstrai acesso a dados, facilita testes
- **Exemplo:** `IUsuarioRepositorio` interface → `UsuarioRepositorioPostgreSQL` implementação

#### Service Layer Pattern
- **Onde:** Camada de aplicação
- **Por quê:** Encapsula lógica de negócio, orquestra repositórios
- **Exemplo:** `UsuarioServico` coordena `UsuarioRepositorio` + `ProfissionalRepositorio`

#### DTO (Data Transfer Object)
- **Onde:** Camada de aplicação (controllers)
- **Por quê:** Desacopla API de entidades de domínio
- **Exemplo:** `CriarUsuarioDTO` para cadastro

#### Middleware Chain
- **Onde:** API (Gin framework)
- **Por quê:** Separação de concerns (CORS, auth, logging)
- **Exemplo:** `AuthMiddleware` valida JWT antes de executar handler

#### Domain-Driven Design (DDD)
- **Onde:** Camada de domínio
- **Por quê:** Modelagem rica, validações no domínio
- **Exemplo:** `Usuario.Validar()` encapsula regras de validação

### 4.2 Princípios SOLID Aplicados

**Single Responsibility Principle (SRP):**
- Cada serviço tem uma única responsabilidade
- Controllers apenas tratam HTTP, não contêm lógica de negócio

**Open/Closed Principle (OCP):**
- Interfaces de repositórios abertas para extensão (novos DBs)
- Fechadas para modificação (domínio estável)

**Liskov Substitution Principle (LSP):**
- `Profissional` e `Paciente` são substituíveis por `Usuario` base

**Interface Segregation Principle (ISP):**
- Interfaces específicas (`IUsuarioRepositorio`) ao invés de uma interface gigante

**Dependency Inversion Principle (DIP):**
- Domínio não depende de infraestrutura
- Application depende de interfaces, não implementações

### 4.3 Boas Práticas Implementadas

✅ **Validação em Múltiplas Camadas:**
- Frontend: UX responsiva
- Backend (Controllers): DTOs com validação
- Domínio: Métodos `Validar()`
- Banco de Dados: Constraints

✅ **Segurança:**
- Senhas com bcrypt (hash + salt)
- JWT para autenticação stateless
- CORS configurado
- Soft delete para auditoria

✅ **Configuração Externalizada:**
- Variáveis de ambiente (.env)
- Não commit de segredos no Git

✅ **Logging:**
- Middleware de logging em todas as requisições
- Erros estruturados

✅ **Testes Automatizados:**
- 281 testes unitários (domínio, serviços, mappers)
- Table-driven tests para cobertura extensiva
- Mocks com testify para isolamento
- SQLite in-memory para testes rápidos
- Integração contínua via Git

✅ **Organização de Código:**
- Testes organizados em subdiretórios `/tests`
- Separação clara entre produção e testes
- Nomenclatura consistente (`*_test.go`)
- Documentação inline dos testes

---

## 5. Justificativas

### 5.1 Por que Clean Architecture?

**Problema:** Monólitos difíceis de manter, testes complexos, acoplamento alto.

**Solução:** Clean Architecture separa concerns, isola domínio.

**Benefícios:**
- ✅ Testabilidade: Domínio 100% testável sem banco de dados
- ✅ Independência: Trocar GORM por outro ORM não afeta domínio
- ✅ Clareza: Desenvolvedores encontram código facilmente
- ✅ Escalabilidade: Adicionar features não quebra código existente

### 5.2 Por que Go (Golang)?

**Problema:** Performance crítica para APIs de saúde, concorrência.

**Solução:** Go oferece goroutines, performance nativa, compilação estática.

**Benefícios:**
- ✅ Performance: ~10x mais rápido que Python/Node.js
- ✅ Concorrência: Goroutines para processamento paralelo
- ✅ Deploy: Binário único, sem runtime
- ✅ Simplicidade: Curva de aprendizado baixa

### 5.3 Por que Vue 3?

**Problema:** Interfaces reativas, componentização, produtividade.

**Solução:** Vue 3 com Composition API e Tailwind CSS.

**Benefícios:**
- ✅ Reatividade: Atualizações automáticas de UI
- ✅ Componentes: Reutilização de código
- ✅ Performance: Virtual DOM otimizado
- ✅ Ecossistema: Vite, Vue Router, Pinia

### 5.4 Por que PostgreSQL?

**Problema:** Dados sensíveis de saúde, ACID, constraints.

**Solução:** PostgreSQL com constraints e transações.

**Benefícios:**
- ✅ ACID: Garantia de consistência
- ✅ Constraints: Validações no banco
- ✅ JSON: Suporte nativo para dados semiestruturados
- ✅ Open Source: Sem custos de licença

---

## 6. Atributos de Qualidade Atendidos

| Atributo | Como a Arquitetura Atende |
|----------|---------------------------|
| **Segurança** | JWT, bcrypt, soft delete, validações multicamadas |
| **Desempenho** | Go (performance nativa), PostgreSQL (índices), SPA (carregamento único) |
| **Disponibilidade** | Docker (isolamento), Nginx (proxy), possibilidade de load balancing |
| **Manutenibilidade** | Clean Architecture, SOLID, separação de concerns |
| **Testabilidade** | ✅ **281 testes unitários**, interfaces, DIP, domínio isolado, coverage > 85% |
| **Escalabilidade** | Stateless JWT, camadas independentes, horizontal scaling possível |
| **Usabilidade** | SPA reativo, Tailwind CSS responsivo |
| **Qualidade de Código** | ✅ **Testes automatizados**, validações em múltiplas camadas, type safety |

---

## 7. Infraestrutura de Testes

### 7.1 Estratégia de Testes Implementada

O projeto conta com **281 testes unitários** distribuídos em três camadas:

**Camada de Domínio (184 testes):**
- `usuario_test.go` (62 testes): Validações de Usuario, Profissional, Paciente
- `registro_humor_test.go` (45 testes): Validações de RegistroHumor
- `convite_test.go` (35 testes): Validações de Convite e estados
- Outros (42 testes): Testes adicionais de domínio

**Camada de Aplicação - Serviços (74 testes):**
- `usuario_servico_test.go` (28 testes): Registro, login, perfil, alteração de senha
- `relatorio_servico_test.go` (17 testes): Geração de relatórios e cálculos de médias
- `registro_humor_servico_test.go` (13 testes): Criação de registros de humor
- `convite_servico_test.go` (13 testes): Geração e vinculação de convites
- Outros (3 testes): Testes adicionais de serviços

**Camada de Aplicação - Mappers (23 testes):**
- `utils_test.go` (23 testes): Conversões entre DTOs e Entidades

### 7.2 Padrões de Teste Utilizados

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

**Mocks com Testify:**
- Todos os repositórios possuem mocks
- Testes de serviços isolados da camada de persistência

**Banco de Dados em Memória:**
- SQLite `:memory:` para testes de integração
- Transações isoladas entre testes

### 7.3 Cobertura de Testes

| Camada | Módulo | Testes | Status |
|--------|--------|--------|--------|
| Domínio | Usuario | 62 | ✅ 100% |
| Domínio | RegistroHumor | 45 | ✅ 100% |
| Domínio | Convite | 35 | ✅ 100% |
| Serviços | UsuarioServico | 28 | ✅ Completo |
| Serviços | RelatorioServico | 17 | ✅ Completo |
| Serviços | RegistroHumorServico | 13 | ✅ Completo |
| Serviços | ConviteServico | 13 | ✅ Completo |
| Mappers | Utils | 23 | ✅ Completo |
| **TOTAL** | **8 módulos** | **281** | ✅ **Todos passando** |

**Execução dos Testes:**
```bash
go test ./interno/dominio/tests ./interno/aplicacao/servicos/tests ./interno/aplicacao/mappers/tests
```

**Tempo de Execução:** < 1 segundo

---

## 8. Débitos Técnicos Identificados

| Débito | Impacto | Prioridade | Resolução Planejada |
|--------|---------|------------|---------------------|
| ~~Falta de testes automatizados~~ | ~~Alto~~ | ~~Alta~~ | ✅ **CONCLUÍDO** - 281 testes implementados |
| Log de auditoria não implementado | Médio | Média | Criar tabela `audit_log` e middleware |
| Backup não automatizado | Alto | Alta | Script cron para backup PostgreSQL |
| Swagger/OpenAPI incompleto | Baixo | Baixa | Completar anotações Swagger |
| 2FA não implementado | Médio | Baixa | Implementar TOTP (Google Authenticator) |
| Monitoramento/APM ausente | Médio | Média | Integrar Prometheus + Grafana |
| Testes de integração (API) ausentes | Médio | Média | Implementar testes end-to-end com HTTP |

---

## 9. Histórico de Atualizações

| Data | Versão | Alterações |
|------|--------|------------|
| 26/10/2025 | 1.0 | Documento inicial de arquitetura |
| 28/10/2025 | 1.1 | ✅ Adição da seção de Infraestrutura de Testes (281 testes unitários)<br>✅ Atualização de Débitos Técnicos (testes concluídos)<br>✅ Atualização de Atributos de Qualidade (testabilidade comprovada)<br>✅ Expansão da estrutura de código com diretórios /tests<br>✅ Atualização de Boas Práticas (testes automatizados) |

---

**Documento elaborado por:** Alexander Nunes Souza  
**Orientadora:** Profa. Dra. Adicinéia A. de Oliveira  
**Disciplina:** ESII/2025-2  
**Primeira versão:** 26/10/2025  
**Última atualização:** 28/10/2025 (v1.1 - Infraestrutura de Testes)
