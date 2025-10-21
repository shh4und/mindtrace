---
title: "Refatoração Clean Architecture - DTOs e Mappers - MVP v1.0"
version: "1.0"
date_created: "2025-10-21"
last_updated: "2025-10-21"
author: "Desenvolvedor"
status: "Planned"
priority: "Alta"
tags: ["architecture", "refactor", "clean-code", "dtos", "mappers", "mvp"]
project: "MindTrace"
---

# 🏗️ Refatoração Clean Architecture - DTOs e Mappers - MVP v1.0

![Status: Planned](https://img.shields.io/badge/status-Planned-blue) 
![Priority: High](https://img.shields.io/badge/priority-High-red) 
![Timeline: 15-20 days](https://img.shields.io/badge/timeline-15--20%20days-green)

---

## 📊 Visão Geral

Este plano refatora o backend do MindTrace para implementar completamente o padrão **Clean Architecture** com **DTOs (Data Transfer Objects)** e **Mappers**, além de enriquecer a **Domain Layer** com lógica de negócio. O objetivo é melhorar segurança, testabilidade, manutenibilidade e desacoplamento entre camadas.

**Contexto**: O backend já possui uma boa estrutura de camadas, mas necessita de melhorias em:
1. Domain Layer: Apenas atributos, sem lógica de negócio
2. Falta de DTOs: Expõe entidades completas (incluso senhas)
3. Sem mappers: Conversão manual e propensa a erros
4. Responsabilidades misturadas: Serviços fazem lógica + coordenação

---

## 1️⃣ Requisitos e Restrições

### Requisitos Funcionais

- **REQ-001**: Implementar DTOs Input para todas as operações de criação/atualização
- **REQ-002**: Implementar DTOs Output para retorno de dados seguro
- **REQ-003**: Criar Mappers automáticos entre entidades e DTOs
- **REQ-004**: Enriquecer Domain Layer com validações de negócio
- **REQ-005**: Nunca expor campos sensíveis (senha, datas privadas) nas DTOs Output
- **REQ-006**: Suportar conversão bidirecional (Entrada ↔ Saída)

### Requisitos Técnicos

- **TEC-001**: DTOs devem estar em `aplicacao/dtos/`
- **TEC-002**: Mappers devem estar em `aplicacao/mappers/`
- **TEC-003**: Domain Layer em `dominio/` (enriquecer existentes)
- **TEC-004**: Validações usando tags `binding:"required,email,min=8"` no Gin
- **TEC-005**: Domain Layer com métodos de validação (`Validar()`, `ValidarEmail()`, etc)
- **TEC-006**: Utilizar package `errors` padrão do Go para domínio
- **TEC-007**: Manter GORM apenas em persistência, não em domínio

### Constraints

- **CON-001**: Não quebrar APIs existentes na fase 1 (backward compatibility)
- **CON-002**: Manter banco de dados sem alterações
- **CON-003**: Implementação não deve afetar repositórios PostgreSQL/SQLite
- **CON-004**: Testes devem rodar sem dependências externas (mocks)

### Padrões

- **PAT-001**: Nomenclatura: `NomeDTOIn` para entrada, `NomeDTOOut` para saída
- **PAT-002**: Mappers sempre retornam pointers `*dto`, exceto slices
- **PAT-003**: Validações de negócio como métodos na entidade (ex: `usuario.ValidarEmail()`)
- **PAT-004**: Erros do domínio como `var Err...` no topo do arquivo

### Guidelines

- **GUD-001**: DTOs Input com validações de binding (Gin)
- **GUD-002**: DTOs Output nunca retornam `json:"-"` (segurança)
- **GUD-003**: Métodos de mapper sempre em ordem: entrada → domínio → saída
- **GUD-004**: Domain Layer 100% independente de GORM

---

## 2️⃣ Fases de Implementação

### Fase 1️⃣: Preparação e Estrutura Base

**GOAL-001**: Criar estrutura de diretórios e templates base para DTOs e Mappers

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-001 | Criar diretório `aplicacao/dtos/` | ⏳ |
| TASK-002 | Criar diretório `aplicacao/mappers/` | ⏳ |
| TASK-003 | Criar arquivo base `dtos/tipos.go` com DTOs comuns (erros, paginação) | ⏳ |
| TASK-004 | Criar arquivo base `mappers/utils.go` com funções helper | ⏳ |
| TASK-005 | Documentar padrões em comentários Go | ⏳ |

---

### Fase 2️⃣: Domain Layer - Enriquecer com Lógica

**GOAL-002**: Adicionar validações e métodos de negócio no Domain Layer

#### 🗄️ Usuario (dominio/usuario.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-006 | Adicionar método `ValidarEmail()` com regex | ⏳ |
| TASK-007 | Adicionar método `ValidarSenha(senhaPlana string)` - mín 8 chars | ⏳ |
| TASK-008 | Adicionar método `ValidarNome()` - não vazio | ⏳ |
| TASK-009 | Adicionar método `Validar()` - orquestra validações | ⏳ |
| TASK-010 | Adicionar var errors no topo: `ErrEmailInvalido`, `ErrSenhaFraca`, etc | ⏳ |

#### 👤 Profissional (dominio/usuario.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-011 | Adicionar método `ValidarRegistroProfissional()` - formato/tamanho | ⏳ |
| TASK-012 | Adicionar método `ValidarEspecialidade()` - não vazio | ⏳ |
| TASK-013 | Adicionar método `Validar()` completo | ⏳ |
| TASK-014 | Adicionar métodos `Ativar()` e `Desativar()` | ⏳ |

#### 🧑‍⚕️ Paciente (dominio/usuario.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-015 | Adicionar método `CalcularIdade()` int | ⏳ |
| TASK-016 | Adicionar método `EhMaiorDeIdade()` bool | ⏳ |
| TASK-017 | Adicionar método `Validar()` com verificação de idade | ⏳ |

#### 🎭 RegistroHumor (dominio/registro_humor.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-018 | Adicionar método `ValidarHumor()` - 1-5 ou 1-10 | ⏳ |
| TASK-019 | Adicionar método `Validar()` | ⏳ |

#### 🎁 Convite (dominio/convite.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-020 | Adicionar método `ValidarToken()` - não vazio, tamanho | ⏳ |
| TASK-021 | Adicionar método `EstaExpirado()` bool | ⏳ |

---

### Fase 3️⃣: DTOs Input (Recebem dados do cliente)

**GOAL-003**: Implementar DTOs de entrada com validações Gin

#### 📨 Usuário DTOs (dtos/usuario_dtos.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-022 | Criar `CriarUsuarioDTOIn` com campos: email, senha, nome | ⏳ |
| TASK-023 | Criar `CriarProfissionalDTOIn` com campos de profissional | ⏳ |
| TASK-024 | Criar `CriarPacienteDTOIn` com campos de paciente | ⏳ |
| TASK-025 | Criar `AtualizarUsuarioDTOIn` | ⏳ |
| TASK-026 | Criar `AlterarSenhaDTOIn` | ⏳ |
| TASK-027 | Adicionar `binding` tags (required, email, min, max) | ⏳ |

#### 🎭 Humor DTOs (dtos/humor_dtos.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-028 | Criar `CriarRegistroHumorDTOIn` com humor, notas | ⏳ |
| TASK-029 | Criar `AtualizarRegistroHumorDTOIn` | ⏳ |

#### 🎁 Convite DTOs (dtos/convite_dtos.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-030 | Criar `CriarConviteDTOIn` com email/paciente | ⏳ |
| TASK-031 | Criar `AceitarConviteDTOIn` com token | ⏳ |

---

### Fase 4️⃣: DTOs Output (Retornam dados seguro)

**GOAL-004**: Implementar DTOs de saída sem dados sensíveis

#### 📨 Usuário DTOs Output (dtos/usuario_dtos.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-032 | Criar `UsuarioDTOOut` sem senha, datas privadas | ⏳ |
| TASK-033 | Criar `ProfissionalDTOOut` com usuario aninhado | ⏳ |
| TASK-034 | Criar `PacienteDTOOut` com idade calculada | ⏳ |
| TASK-035 | Criar `ListaProfissionalDTOOut` com paginação | ⏳ |
| TASK-036 | Criar `ListaPacienteDTOOut` com paginação | ⏳ |

#### 🎭 Humor DTOs Output (dtos/humor_dtos.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-037 | Criar `RegistroHumorDTOOut` seguro | ⏳ |
| TASK-038 | Criar `ListaRegistroHumorDTOOut` com paginação | ⏳ |

#### 🎁 Convite DTOs Output (dtos/convite_dtos.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-039 | Criar `ConviteDTOOut` seguro | ⏳ |

---

### Fase 5️⃣: Mappers - Converter Entidades ↔ DTOs

**GOAL-005**: Implementar mappers bidirecional

#### 🗺️ Usuario Mappers (mappers/usuario_mapper.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-040 | `UsuarioParaDTOOut(*Usuario) *UsuarioDTOOut` | ⏳ |
| TASK-041 | `ProfissionalParaDTOOut(*Profissional) *ProfissionalDTOOut` | ⏳ |
| TASK-042 | `PacienteParaDTOOut(*Paciente) *PacienteDTOOut` | ⏳ |
| TASK-043 | `CriarUsuarioDTOInParaEntidade(*CriarUsuarioDTOIn) *Usuario` | ⏳ |
| TASK-044 | `CriarProfissionalDTOInParaEntidade(*CriarProfissionalDTOIn) (*Usuario, *Profissional)` | ⏳ |
| TASK-045 | `CriarPacienteDTOInParaEntidade(*CriarPacienteDTOIn) (*Usuario, *Paciente)` | ⏳ |

#### 🗺️ Humor Mappers (mappers/humor_mapper.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-046 | `RegistroHumorParaDTOOut(*RegistroHumor) *RegistroHumorDTOOut` | ⏳ |
| TASK-047 | `CriarRegistroHumorDTOInParaEntidade(*CriarRegistroHumorDTOIn) *RegistroHumor` | ⏳ |

#### 🗺️ Convite Mappers (mappers/convite_mapper.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-048 | `ConviteParaDTOOut(*Convite) *ConviteDTOOut` | ⏳ |
| TASK-049 | `CriarConviteDTOInParaEntidade(*CriarConviteDTOIn) *Convite` | ⏳ |

---

### Fase 6️⃣: Atualizar Serviços

**GOAL-006**: Refatorar serviços para usar DTOs e Mappers

#### ⚙️ UsuarioServico (aplicacao/servicos/usuario_servico.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-050 | Atualizar `RegistrarProfissional()` para receber DTO Input | ⏳ |
| TASK-051 | Atualizar `RegistrarPaciente()` para receber DTO Input | ⏳ |
| TASK-052 | Chamar validações do domínio em cada serviço | ⏳ |
| TASK-053 | Retornar DTOs Output dos serviços (ou manter void + mappers nos controllers) | ⏳ |

#### ⚙️ HumorServico (aplicacao/servicos/registro_humor_servico.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-054 | Atualizar para receber/retornar DTOs | ⏳ |
| TASK-055 | Chamar validações do domínio | ⏳ |

#### ⚙️ ConviteServico (aplicacao/servicos/convite_servico.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-056 | Atualizar para receber/retornar DTOs | ⏳ |

---

### Fase 7️⃣: Atualizar Controladores

**GOAL-007**: Refatorar controladores para usar DTOs e Mappers

#### 🎮 UsuarioControlador (aplicacao/controladores/usuario_controlador.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-057 | Atualizar `CriarPaciente()` para receber/retornar DTOs | ⏳ |
| TASK-058 | Atualizar `CriarProfissional()` para receber/retornar DTOs | ⏳ |
| TASK-059 | Atualizar `BuscarPerfil()` para retornar DTOs | ⏳ |
| TASK-060 | Atualizar `AtualizarPerfil()` para receber/retornar DTOs | ⏳ |
| TASK-061 | Adicionar mapping usando mappers nos retornos | ⏳ |

#### 🎮 HumorControlador (aplicacao/controladores/registro_humor_controlador.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-062 | Atualizar `CriarRegistro()` para receber/retornar DTOs | ⏳ |
| TASK-063 | Atualizar `BuscarRegistros()` para retornar DTOs | ⏳ |

#### 🎮 ConviteControlador (aplicacao/controladores/convite_controlador.go)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-064 | Atualizar `CriarConvite()` para receber/retornar DTOs | ⏳ |

---

### Fase 8️⃣: Testes Unitários

**GOAL-008**: Criar testes para DTOs, Mappers e Domain validations

#### ✅ Domain Tests (interno/dominio/tests/)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-065 | Testes para `Usuario.ValidarEmail()` - casos válido/inválido | ⏳ |
| TASK-066 | Testes para `Usuario.ValidarSenha()` - tamanho mínimo | ⏳ |
| TASK-067 | Testes para `Profissional.ValidarRegistroProfissional()` | ⏳ |
| TASK-068 | Testes para `Paciente.CalcularIdade()` e `EhMaiorDeIdade()` | ⏳ |
| TASK-069 | Testes para `RegistroHumor.ValidarHumor()` - 1-5 | ⏳ |

#### ✅ Mapper Tests (interno/aplicacao/mappers/tests/)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-070 | Testes para `UsuarioParaDTOOut()` - sem exposição de senha | ⏳ |
| TASK-071 | Testes para `PacienteParaDTOOut()` - idade calculada corretamente | ⏳ |
| TASK-072 | Testes para conversão bidirecional Usuario In → Entidade → Out | ⏳ |

#### ✅ Service Tests (interno/aplicacao/servicos/tests/)

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-073 | Testes para `RegistrarPaciente()` com DTO | ⏳ |
| TASK-074 | Testes para validações sendo chamadas | ⏳ |

---

### Fase 9️⃣: Documentação e Cleanup

**GOAL-009**: Documentar padrões e limpar código

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-075 | Criar `docs/architecture/DTOs_AND_MAPPERS.md` | ⏳ |
| TASK-076 | Atualizar `README.md` com estrutura nova | ⏳ |
| TASK-077 | Adicionar exemplos de uso em comentários Go | ⏳ |
| TASK-078 | Code review e merge para main | ⏳ |

---

## 3️⃣ Alternativas Descartadas

- **ALT-001**: Usar biblioteca externa de DTO generation (protobuf/thrift) - Rejeitado por complexidade desnecessária em MVP
- **ALT-002**: Validações no controller com if-chains - Rejeitado por violação de Clean Architecture
- **ALT-003**: Sem separação Input/Output DTO - Rejeitado por questões de segurança (exposição de dados)
- **ALT-004**: Mappers automáticos via reflection - Rejeitado por performance e debugging

---

## 4️⃣ Dependências

- **DEP-001**: Go 1.21+ (já instalado)
- **DEP-002**: Gin v1.10+ (já presente em `go.mod`)
- **DEP-003**: GORM v1.30+ (já presente)
- **DEP-004**: Nenhuma dependência externa nova necessária
- **DEP-005**: Biblioteca padrão Go: `errors`, `regexp`, `time`

---

## 5️⃣ Estrutura de Arquivos

### 🆕 Novos Arquivos

```
backend/interno/aplicacao/
├── dtos/
│   ├── tipos.go                    # DTOs comuns
│   ├── usuario_dtos.go             # Usuario/Profissional/Paciente DTOs
│   ├── humor_dtos.go               # RegistroHumor DTOs
│   ├── convite_dtos.go             # Convite DTOs
│   └── README.md                   # Documentação DTOs
├── mappers/
│   ├── utils.go                    # Funções helper
│   ├── usuario_mapper.go           # Usuario Mappers
│   ├── humor_mapper.go             # Humor Mappers
│   ├── convite_mapper.go           # Convite Mappers
│   └── README.md                   # Documentação Mappers
```

### 📝 Arquivos Modificados

```
backend/interno/
├── dominio/
│   ├── usuario.go                  # Adicionar validações
│   ├── registro_humor.go           # Adicionar validações
│   ├── convite.go                  # Adicionar validações
├── aplicacao/
│   ├── servicos/
│   │   ├── usuario_servico.go      # Integrar DTOs
│   │   ├── registro_humor_servico.go
│   │   └── convite_servico.go
│   ├── controladores/
│   │   ├── usuario_controlador.go  # Integrar DTOs
│   │   ├── registro_humor_controlador.go
│   │   └── convite_controlador.go
```

---

## 6️⃣ Testes Essenciais

### Testes de Domain Layer

| # | Teste | Esperado |
|---|-------|----------|
| TEST-001 | `Usuario.ValidarEmail()` com email válido | ✓ nil error |
| TEST-002 | `Usuario.ValidarEmail()` com email inválido | ✓ ErrEmailInvalido |
| TEST-003 | `Usuario.ValidarSenha()` com senha < 8 chars | ✓ ErrSenhaFraca |
| TEST-004 | `Paciente.CalcularIdade()` com data conhecida | ✓ idade correta |
| TEST-005 | `Paciente.EhMaiorDeIdade()` com < 18 anos | ✓ false |
| TEST-006 | `RegistroHumor.ValidarHumor()` com humor > 5 | ✓ erro |

### Testes de Mappers

| # | Teste | Esperado |
|---|-------|----------|
| TEST-007 | `UsuarioParaDTOOut()` com senha no usuário | ✓ DTO sem senha |
| TEST-008 | `PacienteParaDTOOut()` | ✓ idade calculada automaticamente |
| TEST-009 | `CriarPacienteDTOInParaEntidade()` com data string | ✓ Paciente.DataNascimento preenchido |

### Testes de Integração

| # | Teste | Esperado |
|---|-------|----------|
| TEST-010 | POST `/pacientes` com DTO válido | ✓ 201, retorna PacienteDTOOut |
| TEST-011 | POST `/pacientes` com email inválido | ✓ 400, erro de validação |
| TEST-012 | GET `/pacientes/:id` | ✓ PacienteDTOOut sem dados sensíveis |

---

## 7️⃣ Riscos & Mitigações

| Risk | Impacto | Mitigação |
|------|---------|-----------|
| **RISK-001**: Quebra de compatibilidade com clientes existentes | Alto | Manter endpoints antigos em paralelo, deprecation notice |
| **RISK-002**: DTOs Output com dados confidenciais por erro | Alto | Code review obrigatório, testes checando ausência de `json:"-"` |
| **RISK-003**: Mappers com lógica complexa | Médio | Testes unitários 100%, manter simples |
| **RISK-004**: Performance com muitos mappers | Médio | Benchmarks em fase 9, otimizar se necessário |
| **RISK-005**: Domain Layer sem mudanças (apenas DTOs) | Médio | Checklist obrigatório antes de merge |

---

## 8️⃣ Suposições

- **ASS-001**: Equipe familiarizada com Go interfaces e pointers
- **ASS-002**: Testes podem usar mocks via interface (já existe `repositorios.go`)
- **ASS-003**: Banco de dados não será alterado nesta refatoração
- **ASS-004**: Controllers atuais podem ser modificados sem breaking changes significativas
- **ASS-005**: DTOs serão suficientes (sem necessidade de GraphQL/Protocol Buffers)
- **ASS-006**: Validações de negócio encapsuladas no domínio são suficientes

---

## 9️⃣ Timeline Estimada

| Fase | Dias | Acumulado | Responsável |
|------|------|-----------|-------------|
| 1: Preparação e Estrutura | 1-2 | 1-2 | Desenvolvedor |
| 2: Domain Layer | 2-3 | 3-5 | Desenvolvedor |
| 3: DTOs Input | 2-3 | 5-8 | Desenvolvedor |
| 4: DTOs Output | 2-3 | 7-11 | Desenvolvedor |
| 5: Mappers | 1-2 | 8-13 | Desenvolvedor |
| 6: Atualizar Serviços | 1-2 | 9-15 | Desenvolvedor |
| 7: Atualizar Controladores | 1-2 | 10-17 | Desenvolvedor |
| 8: Testes | 2-3 | 12-20 | Desenvolvedor |
| 9: Documentação e Cleanup | 1 | 13-21 | Desenvolvedor |
| **TOTAL** | **15-20 dias** | **15-20 dias** | - |

---

## 1️⃣0️⃣ Próximos Passos

1. ✅ **Revisar este plano** com a equipe - Validar escopo e timeline
2. 🔀 **Criar branch feature** `feature/refactor-clean-architecture-dtos-mappers`
3. 📂 **Iniciar Fase 1** - Criar estrutura de diretórios
4. 🧪 **Setup testes** - Criar estrutura de testes paralela
5. 📝 **Documentar** - Criar exemplos conforme avança
6. 🔍 **Code review contínuo** - Revisar cada fase
7. 🧩 **Integração gradual** - Testar com endpoints existentes
8. 📊 **Validação** - Verificar que segurança/performance não degradou
9. 🎯 **Release** - Merge para main e deploy

---

## 📌 Metadados

| Campo | Valor |
|-------|-------|
| **Criado em** | 2025-10-21 |
| **Atualizado em** | 2025-10-21 |
| **Duração estimada** | 15-20 dias |
| **Prioridade** | Alta ⬆️ |
| **Contexto** | Disciplina Engenharia de Software - Projeto MindTrace |
| **Status** | Planned 🔵 |
| **Branch** | `feature/refactor-clean-architecture-dtos-mappers` |
| **Próximas versões** | v1.1 (Add OpenAPI/Swagger), v1.2 (Add validation middleware) |

