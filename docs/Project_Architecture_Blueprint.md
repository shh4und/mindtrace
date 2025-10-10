# Project Architecture Blueprint

## 1. Architecture Detection and Analysis

### Technology Stack Analysis
MindTrace is a full-stack web application built with the following technology stacks:

**Backend:**
- **Language**: Go (Golang) 1.25.1
- **Web Framework**: Gin v1.10.1
- **ORM**: GORM v1.30.1 with PostgreSQL and SQLite drivers
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **Testing**: Testify v1.10.0
- **Additional Libraries**: CORS middleware, bcrypt for password hashing

**Frontend:**
- **Language**: JavaScript (ES6+)
- **Framework**: Vue.js 3.5.18 (Composition API)
- **Build Tool**: Vite v7.0.6
- **State Management**: Pinia v3.0.3
- **HTTP Client**: Axios v1.11.0
- **Styling**: TailwindCSS v4.1.11
- **Charts**: ApexCharts v5.3.4 with vue3-apexcharts
- **Icons**: FontAwesome v7.0.0
- **Utilities**: VueUse v13.7.0

**Database:**
- **Primary**: PostgreSQL (production)
- **Alternative**: SQLite (development/testing)
- **Schema Management**: GORM auto-migration

**Infrastructure:**
- **Containerization**: Docker
- **Orchestration**: Docker Compose
- **Database Admin**: PgAdmin
- **Reverse Proxy**: Nginx (frontend)

### Architectural Pattern Determination
The codebase implements **Clean Architecture** in the backend, with clear separation of concerns across multiple layers:

- **Domain Layer** (`interno/dominio/`): Contains business entities and core business rules
- **Application Layer** (`interno/aplicacao/`): Contains application services, controllers, and middlewares
- **Persistence Layer** (`interno/persistencia/`): Contains repository interfaces and implementations

The frontend follows a **Component-based Architecture** with Vue.js, using:
- **Views**: Page-level components
- **Components**: Reusable UI components
- **Store**: Centralized state management with Pinia
- **Services**: API communication layer
- **Router**: Client-side routing with Vue Router

## 2. Architectural Overview

MindTrace implements Clean Architecture principles to ensure separation of concerns, testability, and maintainability. The architecture emphasizes:

- **Dependency Inversion**: Higher-level modules don't depend on lower-level modules
- **Single Responsibility**: Each layer has a distinct purpose
- **Interface Segregation**: Repository interfaces define contracts without implementation details
- **Layer Isolation**: Dependencies flow inward toward the domain layer

### Guiding Principles
1. **Domain-Centric Design**: Business logic resides in the domain layer, independent of external concerns
2. **Dependency Injection**: All dependencies are injected, enabling testability and flexibility
3. **Repository Pattern**: Data access is abstracted through interfaces
4. **Separation of Concerns**: UI, business logic, and data access are strictly separated
5. **Configuration-Driven**: Database and environment settings are configurable

### Architectural Boundaries
- **Domain Layer**: Pure business logic, no external dependencies
- **Application Layer**: Orchestrates domain objects, handles HTTP requests/responses
- **Persistence Layer**: Implements data access interfaces, technology-specific implementations
- **Infrastructure**: External concerns (web framework, database drivers, containers)

## 3. Architecture Visualization

### High-Level Architectural Overview
```
┌─────────────────────────────────────────────────────────────┐
│                    Presentation Layer                       │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────┐  │
│  │   Vue.js SPA    │  │   REST API      │  │   Nginx     │  │
│  │  (Frontend)     │  │   (Gin)         │  │ (Reverse    │  │
│  └─────────────────┘  └─────────────────┘  │  Proxy)    │  │
└─────────────────────────────────────────────┼─────────────┘
                                              │
┌─────────────────────────────────────────────┼─────────────┐
│                 Application Layer           │             │
│  ┌─────────────────┐  ┌─────────────────┐   │             │
│  │  Controllers    │  │   Services      │   │             │
│  │  (HTTP Handlers)│  │ (Business Logic)│   │             │
│  └─────────────────┘  └─────────────────┘   │             │
└─────────────────────────────────────────────┼─────────────┘
                                              │
┌─────────────────────────────────────────────┼─────────────┐
│                 Domain Layer                │             │
│  ┌─────────────────┐                        │             │
│  │  Entities       │                        │             │
│  │  (Business      │                        │             │
│  │   Models)       │                        │             │
│  └─────────────────┘                        │             │
└─────────────────────────────────────────────┼─────────────┘
                                              │
┌─────────────────────────────────────────────┼─────────────┘
│                 Persistence Layer                       │
│  ┌─────────────────┐  ┌─────────────────┐               │
│  │ Repository      │  │   Database      │               │
│  │ Interfaces      │  │ Implementations │               │
│  └─────────────────┘  └─────────────────┘               │
└─────────────────────────────────────────────────────────┘
```

### Component Interaction Diagram
```
Frontend (Vue.js) ───HTTP───→ Controllers (Gin) ──→ Services ──→ Domain Entities
       │                           │                    │             │
       │                           │                    │             │
       └─Pinia Store               └─Middlewares        └─Repository  └─Database
                                                          Interfaces     (PostgreSQL/SQLite)
```

### Data Flow Diagram
```
User Input → Vue Components → Pinia Store → API Service (Axios) → REST API
                                                                          │
                                                                          ↓
HTTP Request → Gin Router → Middleware (Auth) → Controller → Service → Repository Interface
                                                                          │
                                                                          ↓
Repository Implementation → GORM → Database → Results → Domain Entities → JSON Response
                                                                          │
                                                                          ↓
JSON Response → Axios → Pinia Store → Vue Components → UI Update
```

## 4. Core Architectural Components

### Backend Components

#### Domain Layer (`interno/dominio/`)
**Purpose and Responsibility:**
- Contains the core business entities and rules
- Defines the data structures that represent business concepts
- Independent of any external frameworks or technologies

**Internal Structure:**
- `Usuario`: Base user entity with common fields
- `Profissional`: Professional user type extending Usuario
- `Paciente`: Patient user type extending Usuario
- `RegistroHumor`: Mood tracking records
- `Alerta`: Alert notifications
- `Convite`: Invitation system entities

**Interaction Patterns:**
- Entities are used by application services
- No direct dependencies on external libraries
- GORM tags define database mapping

#### Application Layer (`interno/aplicacao/`)
**Purpose and Responsibility:**
- Orchestrates business operations using domain entities
- Handles HTTP request/response transformation
- Implements cross-cutting concerns (auth, validation, logging)

**Internal Structure:**
- **Controllers**: HTTP request handlers (`controladores/`)
- **Services**: Business logic coordinators (`servicos/`)
- **Middlewares**: Cross-cutting concerns (`middlewares/`)

**Interaction Patterns:**
- Controllers depend on services via dependency injection
- Services depend on repository interfaces
- Middlewares intercept HTTP requests

#### Persistence Layer (`interno/persistencia/`)
**Purpose and Responsibility:**
- Abstracts data access operations
- Provides technology-agnostic data access interfaces
- Implements database-specific operations

**Internal Structure:**
- **Repository Interfaces** (`repositorios/`): Define data access contracts
- **PostgreSQL Implementation** (`postgres/`): Production database implementation
- **SQLite Implementation** (`sqlite/`): Development/testing database implementation

**Interaction Patterns:**
- Services depend on repository interfaces (dependency inversion)
- Multiple database implementations for different environments
- Uses GORM for ORM operations

### Frontend Components

#### Views (`src/views/`)
**Purpose and Responsibility:**
- Page-level components representing different application screens
- Compose smaller components to create complete pages
- Handle routing and navigation

**Internal Structure:**
- `Landpage.vue`: Public landing page
- `Login.vue`, `Cadastro.vue`: Authentication views
- `PacienteDashboard.vue`: Patient dashboard
- `ProfissionalDashboard.vue`: Professional dashboard

#### Store (`src/store/`)
**Purpose and Responsibility:**
- Centralized state management for the application
- Manages user authentication state and profile data
- Provides reactive state updates across components

**Internal Structure:**
- `user.js`: User authentication and profile management

#### Services (`src/services/`)
**Purpose and Responsibility:**
- Handles API communication with the backend
- Provides a clean interface for HTTP operations
- Manages authentication tokens and request interceptors

**Internal Structure:**
- `api.js`: Axios-based API client with authentication

## 5. Architectural Layers and Dependencies

### Layer Structure
MindTrace follows strict layering with inward-pointing dependencies:

```
┌─────────────────────────────────┐
│         Controllers             │ ← Application Layer
│         (HTTP Interface)        │
├─────────────────────────────────┤
│         Services                │ ← Application Layer
│      (Business Logic)           │
├─────────────────────────────────┤
│       Repository Interfaces     │ ← Persistence Abstraction
├─────────────────────────────────┤
│         Domain Entities         │ ← Domain Layer (Core)
├─────────────────────────────────┤
│   Repository Implementations    │ ← Infrastructure Layer
│     (PostgreSQL/SQLite)         │
└─────────────────────────────────┘
```

### Dependency Rules
1. **Controllers** depend on **Services** (not repositories directly)
2. **Services** depend on **Repository Interfaces** (not implementations)
3. **Repository Interfaces** depend on **Domain Entities**
4. **Repository Implementations** depend on **Domain Entities** and database drivers
5. **Domain Entities** have no external dependencies

### Abstraction Mechanisms
- **Repository Pattern**: Interfaces define data access contracts
- **Dependency Injection**: Constructor injection for all dependencies
- **Configuration**: Environment variables for database selection
- **Middleware Pattern**: Cross-cutting concerns applied via HTTP pipeline

## 6. Data Architecture

### Domain Model Structure
The domain model follows a hierarchical inheritance pattern:

```
Usuario (Base Entity)
├── Profissional
│   ├── Especialidade
│   ├── RegistroProfissional
│   └── Pacientes (Many-to-Many)
└── Paciente
    ├── DataNascimento
    ├── HistoricoSaude
    └── RegistrosHumor (One-to-Many)
```

### Entity Relationships
- **One-to-One**: Usuario → Profissional/Paciente
- **One-to-Many**: Paciente → RegistrosHumor, Profissional → Alertas
- **Many-to-Many**: Profissional ↔ Paciente (via junction table)

### Data Access Patterns
- **Repository Pattern**: All data access through repository interfaces
- **Unit of Work**: Transaction management via GORM DB instances
- **Query Objects**: Time-based queries for mood records
- **Eager Loading**: Related entities loaded with GORM Preload

### Data Transformation
- **DTO Pattern**: Request/Response DTOs in services
- **Entity Mapping**: GORM handles entity-to-table mapping
- **JSON Serialization**: Gin handles JSON request/response transformation

### Data Validation
- **Service Layer Validation**: Business rules enforced in services
- **Binding Validation**: Gin validates request structure
- **Database Constraints**: GORM enforces referential integrity

## 7. Cross-Cutting Concerns Implementation

### Authentication & Authorization
**Implementation:**
- JWT tokens stored in localStorage (frontend)
- Authorization header interceptor in Axios
- Gin middleware validates JWT tokens
- User ID extracted and injected into request context

**Security Boundary Patterns:**
- Public routes: `/entrar/login`, `/profissionais/registrar`, `/pacientes/registrar`
- Protected routes require valid JWT
- Role-based access (patient vs professional dashboards)

### Error Handling & Resilience
**Exception Handling Patterns:**
- Service methods return errors, controllers handle HTTP responses
- Database errors wrapped and returned as service errors
- Frontend displays user-friendly error messages

**Fallback Strategies:**
- Token expiration triggers logout and redirect
- Network errors show toast notifications
- Invalid requests return 400/401/403 status codes

### Logging & Monitoring
**Instrumentation Patterns:**
- Console logging in frontend for debugging
- Gin default logging for HTTP requests
- Error logging in services

**Diagnostic Information:**
- Request/response logging via Gin middleware
- Error details in JSON responses
- User actions tracked via API calls

### Validation
**Input Validation Strategies:**
- Gin binding validation for request structure
- Service-level business rule validation
- Database-level constraint validation

**Validation Responsibility:**
- Structural validation: Controllers (Gin binding)
- Business validation: Services
- Data integrity: Database constraints

### Configuration Management
**Configuration Sources:**
- Environment variables for database connection
- Build-time variables for API base URL
- Runtime configuration via Docker Compose

**Environment-Specific Configuration:**
- Production: PostgreSQL database
- Development: SQLite database
- Configurable via `DB_DRIVER` environment variable

## 8. Service Communication Patterns

### Service Boundary Definitions
- **REST API**: HTTP-based communication between frontend and backend
- **Resource-Based URLs**: `/pacientes/`, `/profissionais/`, `/registros-humor/`
- **HTTP Methods**: GET, POST, PUT, DELETE following REST conventions

### Communication Protocols
- **HTTP/HTTPS**: Standard web protocols
- **JSON**: Request/response format
- **JWT**: Authentication token format

### Synchronous Communication
- All API calls are synchronous HTTP requests
- Frontend waits for backend responses
- Error handling via HTTP status codes

### API Versioning Strategy
- URL prefix: `/api/v1/` (configurable)
- No explicit versioning in current implementation
- Backward compatibility maintained through additive changes

## 9. Technology-Specific Architectural Patterns

### Go Architectural Patterns
#### Framework Integration (Gin)
- Router grouping for API versioning
- Middleware pipeline for authentication
- Request binding and validation
- JSON response formatting

#### ORM Integration (GORM)
- Auto-migration for schema management
- Transaction management via DB instances
- Association loading with Preload
- Soft deletes with gorm.DeletedAt

#### Dependency Injection
- Constructor injection for services and repositories
- Interface-based dependencies enable testing
- Configuration-driven database selection

### Vue.js Architectural Patterns
#### Component Composition
- Composition API for reactive logic
- Component communication via props/emits
- Reusable components with slots

#### State Management (Pinia)
- Centralized user state
- Reactive state updates
- Actions for async operations
- Getters for computed state

#### Routing (Vue Router)
- Declarative route definitions
- Navigation guards for authentication
- Route-based code splitting

## 10. Implementation Patterns

### Interface Design Patterns
**Interface Segregation:**
```go
type UsuarioRepositorio interface {
    CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error
    BuscarPorEmail(email string) (*dominio.Usuario, error)
    // ... other methods
}
```

**Abstraction Levels:**
- Repository interfaces define data access contracts
- Service interfaces could be added for further abstraction
- Domain entities remain technology-agnostic

### Service Implementation Patterns
**Service Lifetime:**
- Services instantiated once per application
- Dependencies injected at creation time
- Stateless design for thread safety

**Operation Implementation:**
```go
func (us *UsuarioServico) BuscarUsuarioPorID(id uint) (*dominio.Usuario, error) {
    return us.usuarioRepo.BuscarUsuarioPorID(id)
}
```

### Repository Implementation Patterns
**Transaction Management:**
```go
func (ur *UsuarioRepositorioPostgres) CriarUsuario(tx *gorm.DB, usuario *dominio.Usuario) error {
    return tx.Create(usuario).Error
}
```

**Query Patterns:**
- Direct GORM queries for simple operations
- Custom queries for complex filtering (time-based mood records)

### Controller Implementation Patterns
**Request Handling:**
```go
func (uc *UsuarioControlador) BuscarPerfil(c *gin.Context) {
    userID, _ := c.Get("userID")
    usuario, err := uc.usuarioServico.BuscarUsuarioPorID(userID.(uint))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
        return
    }
    c.JSON(http.StatusOK, usuario)
}
```

**Parameter Validation:**
- Gin binding for automatic validation
- Custom validation in services for business rules

### Domain Model Implementation
**Entity Implementation:**
```go
type Usuario struct {
    ID          uint           `gorm:"primaryKey"`
    TipoUsuario string         `json:"tipo_usuario" gorm:"type:varchar(50);not null"`
    Nome        string         `json:"nome" gorm:"type:varchar(255);not null"`
    // ... other fields
}
```

**Relationship Mapping:**
- GORM tags define foreign keys and associations
- JSON tags control API serialization

## 11. Testing Architecture

### Testing Strategies
- **Unit Tests**: Test individual functions and methods
- **Integration Tests**: Test component interactions
- **API Tests**: Test HTTP endpoints (not currently implemented)

### Test Boundaries
- **Unit**: Test services with mock repositories
- **Integration**: Test controllers with real repositories
- **System**: Test complete API workflows

### Testing Tools
- **Testify**: Assertion library for Go tests
- **Mock Generation**: Manual mocks for repository interfaces
- **Test Databases**: SQLite for fast, isolated testing

### Test Data Strategies
- **Factory Pattern**: Create test data in test helpers
- **Cleanup**: Database cleanup between tests
- **Isolation**: Each test runs in isolation

## 12. Deployment Architecture

### Containerization Topology
```
┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │   Backend       │
│   (Nginx/Vue)   │    │   (Go/Gin)      │
│   Port 80       │◄──►│   Port 8080     │
└─────────────────┘    └─────────────────┘
         │                       │
         └─────────┬─────────────┘
                   │
         ┌─────────────────┐
         │   Database      │
         │ (PostgreSQL)    │
         │   Port 5432     │
         └─────────────────┘
```

### Environment Adaptations
- **Development**: Single docker-compose with SQLite
- **Production**: Multi-service docker-compose with PostgreSQL
- **Database Admin**: PgAdmin for database management

### Runtime Dependencies
- **Go Binary**: Compiled statically for Linux containers
- **Node Build**: Frontend built at image creation time
- **Database Drivers**: GORM loads appropriate drivers

### Configuration Management
- **Environment Variables**: Database connection, JWT secrets
- **Docker Secrets**: Sensitive data in production
- **Build Args**: API base URL configuration

## 13. Extension and Evolution Patterns

### Feature Addition Patterns
**New Business Entities:**
1. Define entity in `dominio/`
2. Create repository interface in `repositorios/`
3. Implement repository in `postgres/` and `sqlite/`
4. Create service in `servicos/`
5. Create controller in `controladores/`
6. Add routes in main.go

**New Frontend Features:**
1. Create view component in `views/`
2. Add route in `router/index.js`
3. Update store if needed
4. Add API methods in `services/api.js`

### Modification Patterns
**Entity Changes:**
- Update domain entity
- Run auto-migration (handled by GORM)
- Update services and controllers as needed
- Maintain backward compatibility in API

**Service Changes:**
- Modify service methods
- Update dependent controllers
- Add tests for new functionality

### Integration Patterns
**External APIs:**
- Add new service layer for external communication
- Use HTTP client (net/http or external library)
- Handle errors and timeouts appropriately

**Third-Party Services:**
- Create adapter interfaces
- Implement adapters for different providers
- Configure via environment variables

## 14. Architectural Pattern Examples

### Layer Separation Examples
**Interface Definition and Implementation Separation:**
```go
// Repository interface (persistence abstraction)
type UsuarioRepositorio interface {
    BuscarPorEmail(email string) (*dominio.Usuario, error)
}

// Implementation (infrastructure concern)
func (ur *UsuarioRepositorioPostgres) BuscarPorEmail(email string) (*dominio.Usuario, error) {
    var usuario dominio.Usuario
    err := ur.db.Where("email = ?", email).First(&usuario).Error
    return &usuario, err
}
```

**Cross-Layer Communication:**
```go
// Controller (application layer)
func (uc *UsuarioControlador) BuscarPerfil(c *gin.Context) {
    userID, _ := c.Get("userID")
    usuario, err := uc.usuarioServico.BuscarUsuarioPorID(userID.(uint))
    // Service handles domain logic, repository handles data access
}
```

### Component Communication Examples
**Service Invocation:**
```go
// Service composition
type UsuarioServico struct {
    usuarioRepo repositorios.UsuarioRepositorio
}

func NovoUsuarioServico(repo repositorios.UsuarioRepositorio) *UsuarioServico {
    return &UsuarioServico{usuarioRepo: repo}
}
```

**Event Publishing (Conceptual):**
While not implemented, alerts could be published as events:
```go
// Conceptual event-driven pattern
type AlertaServico interface {
    PublicarAlerta(alerta *dominio.Alerta) error
}
```

### Extension Point Examples
**Plugin Registration:**
```go
// Conceptual plugin interface
type NotificacaoPlugin interface {
    Enviar(mensagem string, destinatario string) error
}
```

**Configuration-Driven Extension:**
```go
// Environment-driven database selection
switch dbDriver {
case "postgres":
    db, _ = postgres_repo.NewDB()
case "sqlite":
    db, _ = sqlite_repo.NewDB()
}
```

## 15. Architectural Decision Records

### Architectural Style Decisions
**Decision: Clean Architecture Adoption**
- **Context**: Need for maintainable, testable codebase with clear separation of concerns
- **Options Considered**: Traditional layered architecture, MVC
- **Chosen**: Clean Architecture with domain at center
- **Consequences**: 
  - Positive: High testability, framework independence, clear boundaries
  - Negative: Initial complexity, more interfaces to maintain

**Decision: Repository Pattern Implementation**
- **Context**: Need to abstract data access for testability and flexibility
- **Options**: Active Record, Data Mapper
- **Chosen**: Repository pattern with interfaces
- **Consequences**: Enables easy testing with mocks, supports multiple databases

### Technology Selection Decisions
**Decision: Go for Backend**
- **Context**: Need performant, maintainable backend with good concurrency
- **Options**: Node.js, Python, Java
- **Chosen**: Go for performance and simplicity
- **Consequences**: Fast compilation, good performance, steeper learning curve

**Decision: Vue.js for Frontend**
- **Context**: Need reactive, component-based frontend
- **Options**: React, Angular, Svelte
- **Chosen**: Vue.js for simplicity and flexibility
- **Consequences**: Easy to learn, good ecosystem, Composition API for complex logic

### Implementation Approach Decisions
**Decision: JWT for Authentication**
- **Context**: Need stateless authentication for REST API
- **Options**: Session-based, API keys
- **Chosen**: JWT tokens
- **Consequences**: Stateless, scalable, but token revocation complexity

**Decision: GORM for ORM**
- **Context**: Need database abstraction with Go
- **Options**: Raw SQL, sqlx, ent
- **Chosen**: GORM for productivity
- **Consequences**: Rapid development, but some performance overhead

## 16. Architecture Governance

### Architectural Consistency Maintenance
- **Code Reviews**: Ensure new code follows architectural patterns
- **Interface Contracts**: Repository interfaces prevent direct database access
- **Layer Enforcement**: Dependency injection prevents layer violations
- **Naming Conventions**: Consistent naming across layers

### Automated Checks
- **Go Modules**: Dependency management
- **Docker Build**: Ensures containerization works
- **Auto-migration**: Database schema consistency

### Architectural Review Processes
- **Pull Request Reviews**: Architectural compliance checking
- **Design Discussions**: Major changes reviewed by team
- **Refactoring Sessions**: Regular architecture cleanup

## 17. Blueprint for New Development

### Development Workflow
1. **Domain Modeling**: Define entities and relationships
2. **Repository Design**: Create interfaces and implementations
3. **Service Implementation**: Add business logic
4. **Controller Creation**: Add HTTP endpoints
5. **Frontend Integration**: Update UI and API calls
6. **Testing**: Add unit and integration tests

### Implementation Templates
**New Entity:**
```go
type NovaEntidade struct {
    ID        uint      `gorm:"primaryKey"`
    Campo1    string    `json:"campo1" gorm:"type:varchar(255);not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

**New Repository Interface:**
```go
type NovaEntidadeRepositorio interface {
    Criar(tx *gorm.DB, entidade *dominio.NovaEntidade) error
    BuscarPorID(id uint) (*dominio.NovaEntidade, error)
}
```

**New Service:**
```go
type NovaEntidadeServico struct {
    repo repositorios.NovaEntidadeRepositorio
}

func NovoNovaEntidadeServico(repo repositorios.NovaEntidadeRepositorio) *NovaEntidadeServico {
    return &NovaEntidadeServico{repo: repo}
}
```

### Common Pitfalls
- **Layer Violations**: Don't access repositories directly from controllers
- **Business Logic in Controllers**: Keep controllers thin, move logic to services
- **Tight Coupling**: Use interfaces, not concrete implementations
- **Missing Validation**: Validate at all layers (binding, service, database)

---

*This architecture blueprint was generated on October 10, 2025. Regular updates recommended as the system evolves.*