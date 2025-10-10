# MindTrace

[![Go Version](https://img.shields.io/badge/Go-1.25.1-blue.svg)](https://golang.org/)
[![Vue Version](https://img.shields.io/badge/Vue.js-3.5.18-green.svg)](https://vuejs.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://www.docker.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-blue.svg)](https://www.postgresql.org/)

MindTrace is a comprehensive full-stack web application designed for mental health tracking and management. It enables healthcare professionals to monitor patient mood records, generate reports, and manage notifications, while patients can log their daily moods and view personalized insights.

## 🚀 Key Features

- **User Management**: Registration and authentication for both patients and healthcare professionals
- **Mood Tracking**: Patients can record daily mood entries with timestamps and notes
- **Reports & Analytics**: Professionals can generate mood trend reports for their patients
- **Automated Alerts**: Intelligent notifications for professionals based on patient data patterns
- **Invitation System**: Professionals can send account linking invitations to patients
- **Dual Dashboards**: Separate, role-based interfaces for patients and professionals
- **Data Visualization**: Interactive charts and graphs for mood trend analysis

## 🛠 Technology Stack

### Backend
- **Language**: Go (Golang) 1.25.1
- **Web Framework**: Gin v1.10.1
- **ORM**: GORM v1.30.1
- **Database**: PostgreSQL 17 (production) / SQLite (development)
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **Testing**: Testify v1.10.0
- **Architecture**: Clean Architecture (Domain-Driven Design)

### Frontend
- **Language**: JavaScript (ES6+)
- **Framework**: Vue.js 3.5.18 (Composition API)
- **Build Tool**: Vite v7.0.6
- **State Management**: Pinia v3.0.3
- **HTTP Client**: Axios v1.11.0
- **Styling**: TailwindCSS v4.1.11
- **Charts**: ApexCharts v5.3.4 with vue3-apexcharts
- **Icons**: FontAwesome v7.0.0

### Infrastructure & DevOps
- **Containerization**: Docker
- **Orchestration**: Docker Compose
- **Database Admin**: PgAdmin 4
- **Reverse Proxy**: Nginx
- **Version Control**: Git

## 🏗 Project Architecture

MindTrace implements **Clean Architecture** principles with clear separation of concerns across multiple layers:

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
┌─────────────────────────────────────────────┼─────────────┘
│                 Domain Layer                │
│  ┌─────────────────┐                        │
│  │  Entities       │                        │
│  │  (Business      │                        │
│  │   Models)       │                        │
│  └─────────────────┘                        │
└─────────────────────────────────────────────────────────┘
```

### Architectural Principles
- **Dependency Inversion**: Inner layers don't depend on outer layers
- **Single Responsibility**: Each layer has a distinct purpose
- **Interface Segregation**: Repository interfaces define clear contracts
- **Testability**: Dependency injection enables comprehensive testing

## 🚀 Getting Started

### Prerequisites
- Docker and Docker Compose
- Node.js 22.17.1 (optional, for local frontend development)
- Go 1.25.1 (optional, for local backend development)

### Quick Start with Docker (Recommended)

1. **Clone the repository**:
   ```bash
   git clone https://github.com/shh4und/mindtrace.git
   cd mindtrace
   ```

2. **Create environment file**:
   Create a `.env` file in the root directory:
   ```env
   POSTGRES_USER=your_db_user
   POSTGRES_PASSWORD=your_db_password
   POSTGRES_DB=mindtrace
   PGADMIN_DEFAULT_EMAIL=admin@example.com
   PGADMIN_DEFAULT_PASSWORD=admin_password
   JWT_SECRET=your_jwt_secret_key
   ```

3. **Start the application**:
   ```bash
   # For development (with hot reload)
   docker-compose -f docker-compose.yml -f docker-compose.override.yml up --build

   # For production
   docker-compose -f docker-compose.yml -f docker-compose.prod.yml up --build
   ```

4. **Access the application**:
   - Frontend: http://localhost
   - Backend API: http://localhost/api/v1
   - PgAdmin: http://localhost:5050

### Local Development Setup

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

## 📁 Project Structure

```
mindtrace/
├── backend/                          # Go backend application
│   ├── cmd/api/                      # Application entry point
│   ├── interno/                      # Internal packages
│   │   ├── aplicacao/                # Application layer
│   │   │   ├── controladores/        # HTTP controllers
│   │   │   ├── middlewares/          # HTTP middlewares
│   │   │   └── servicos/             # Business services
│   │   ├── dominio/                  # Domain layer
│   │   │   ├── alerta.go             # Alert entity
│   │   │   ├── convite.go            # Invitation entity
│   │   │   ├── registro_humor.go     # Mood record entity
│   │   │   └── usuario.go            # User entities
│   │   └── persistencia/             # Persistence layer
│   │       ├── postgres/             # PostgreSQL implementations
│   │       ├── repositorios/         # Repository interfaces
│   │       └── sqlite/               # SQLite implementations
│   ├── Dockerfile                    # Production container
│   ├── Dockerfile.dev                # Development container
│   ├── go.mod                        # Go modules
│   └── go.sum                        # Dependencies
├── frontend/                         # Vue.js frontend application
│   ├── src/
│   │   ├── components/               # Reusable Vue components
│   │   ├── views/                    # Page components
│   │   ├── router/                   # Vue Router configuration
│   │   ├── store/                    # Pinia state management
│   │   └── services/                 # API services
│   ├── Dockerfile                    # Production container
│   ├── Dockerfile.dev                # Development container
│   ├── package.json                  # Node dependencies
│   └── vite.config.js                # Vite configuration
├── docker-compose.yml                # Base Docker Compose
├── docker-compose.override.yml       # Development overrides
├── docker-compose.prod.yml           # Production configuration
├── docker-compose.sqlite.yml         # SQLite configuration
├── schema.sql                        # Database schema
├── seed.sh                           # Database seeding script
└── README.md                         # This file
```

## 🔧 Development Workflow

### Branching Strategy
- `main`: Production-ready code
- `feature/*`: New features and enhancements
- `docs/*`: New or updated documentation
- `bugfix/*`: Bug fixes
- `hotfix/*`: Critical production fixes

### Development Process
1. **Create Feature Branch**: `git checkout -b feature/new-feature`
2. **Make Changes**: Implement features following Clean Architecture principles
3. **Run Tests**: Execute test suites for both backend and frontend
4. **Code Review**: Submit pull request for review
5. **Merge**: Squash merge to main after approval

### Database Management
- **Migrations**: Automatic via GORM AutoMigrate
- **Seeding**: Use `seed.sh` for initial data
- **Backup**: Regular PostgreSQL backups in production

## 📋 Coding Standards

### Go Backend Standards
- **Formatting**: `gofmt` and `goimports` for consistent formatting
- **Naming**: PascalCase for exported, camelCase for unexported
- **Error Handling**: Explicit error returns, no panics in production
- **Documentation**: Godoc comments for all exported functions
- **Testing**: Table-driven tests with testify assertions

### Vue.js Frontend Standards
- **Composition API**: Use Vue 3 Composition API over Options API
- **Component Naming**: PascalCase for component files
- **State Management**: Pinia stores for global state
- **Styling**: TailwindCSS utility classes
- **TypeScript**: Consider migration for better type safety

### General Standards
- **Commits**: Conventional commits (`feat:`, `fix:`, `docs:`)
- **Documentation**: Update README and docs for significant changes
- **Security**: Input validation, JWT for authentication
- **Performance**: Optimize database queries, lazy loading for components

## 🧪 Testing

### Backend Testing
- **Framework**: Testify for assertions and test organization
- **Coverage**: Unit tests for services, repositories, and controllers
- **Mocking**: Interface-based dependency injection enables easy mocking
- **Database**: SQLite for fast, isolated integration tests

### Frontend Testing
- **Framework**: Vue Test Utils (planned)
- **Coverage**: Component and service testing
- **E2E**: Playwright or Cypress for end-to-end testing (planned)

### Running Tests
```bash
# Backend tests
cd backend
go test ./...

# Frontend tests (when implemented)
cd frontend
npm run test
```

## 🤝 Contributing

We welcome contributions! Please follow these guidelines:

### Development Setup
1. Fork the repository
2. Create a feature branch
3. Make your changes following the coding standards
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

### Code Examples
- **Controller Pattern**: Thin controllers that delegate to services
- **Service Pattern**: Business logic separated from HTTP concerns
- **Repository Pattern**: Data access abstracted through interfaces
- **Component Composition**: Reusable Vue components with clear props/emits

### Pull Request Process
1. **Title**: Use conventional commit format
2. **Description**: Explain what and why, reference issues
3. **Testing**: Include test coverage for changes
4. **Documentation**: Update docs if needed
5. **Review**: Address reviewer feedback

## 📚 Additional Documentation

- [Project Architecture Blueprint](./docs/Project_Architecture_Blueprint.md) - Detailed architectural documentation
- [API Documentation](./frontend/swagger-output.json) - OpenAPI specification
<!-- - [Database Schema](./schema.sql) - PostgreSQL database schema -->

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with Go, Vue.js, and modern web technologies
- Inspired by mental health tracking best practices
- Community contributions and open-source ecosystem

---

**MindTrace** - Empowering mental health professionals and patients through technology.