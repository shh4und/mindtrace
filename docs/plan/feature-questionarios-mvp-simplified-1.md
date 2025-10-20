---
title: "Questionários Personalizados com Templates - MVP v1.0"
version: "1.0"
date_created: "2025-10-16"
last_updated: "2025-10-16"
author: "Desenvolvedor"
status: "Planned"
priority: "Alta"
tags: ["feature", "questionnaires", "templates", "mvp"]
project: "MindTrace"
---

# 📋 Questionários Personalizados com Templates - MVP v1.0

![Status: Planned](https://img.shields.io/badge/status-Planned-blue) ![Priority: High](https://img.shields.io/badge/priority-High-red) ![Timeline: 20-30 days](https://img.shields.io/badge/timeline-20--30%20days-green)

## 📊 Visão Geral

Sistema MVP simplificado de questionários personalizados para MindTrace. Profissionais criam/editam questionários, atribuem a pacientes, coletam respostas e visualizam estatísticas.

**📊 Escopo**: 1 desenvolvedor | 20-30 dias | Projeto acadêmico

## 1️⃣ Requisitos e Restrições

### 1.1 Requisitos Críticos (MVP)

- **REQ-001**: Profissionais podem editar templates pré-definidos (PHQ-9, GAD-7, Bem-estar)
- **REQ-002**: Profissionais podem criar questionários customizados baseados em templates
- **REQ-003**: Profissionais podem salvar questionários customizados como templates próprios
- **REQ-004**: Profissionais podem atribuir questionários a pacientes específicos
- **REQ-005**: Pacientes recebem notificação de novo questionário
- **REQ-006**: Pacientes respondem questionários através de formulário dinâmico
- **REQ-007**: Respostas armazenadas em BD estruturado (JSON quando necessário)
- **REQ-008**: Profissionais visualizam estatísticas simples (médias, contadores, totalizadores)
- **REQ-009**: Dados de questionários aparecem nos relatórios existentes

### 1.2 Constraints & Simplificações (MVP)

- **CON-001**: MVP - apenas 3 tipos de pergunta: escala (1-10), múltipla escolha, texto livre
- **CON-002**: Sem lógica condicional (if/then/skip) na v1.0
- **CON-003**: Limite: 30 perguntas/questionário, 1000 respostas/período
- **CON-004**: Sem gráficos avançados - apenas tabelas e números
- **CON-005**: Sem mobile app - web responsivo apenas
- **CON-006**: Sem suporte a outras línguas na v1.0
- **CON-007**: Um desenvolvedor - máximo 30 dias

### 1.3 Padrões & Guidelines

- **PAT-001**: Repository Pattern para acesso a dados
- **PAT-002**: Service Layer para lógica de negócio
- **PAT-003**: Reusar componentes Vue/Tailwind existentes
- **GUD-001**: Código em português para domínio
- **GUD-002**: Nomenclatura clara: Questionario, Pergunta, Resposta

## 2️⃣ Fases de Implementação

### Fase 1: Banco & Domínio (3-4 dias) 🗄️

**GOAL-001**: Estrutura de dados básica

| Task | Description | Status |
|------|-------------|--------|
| TASK-001 | Criar struct `Questionario` (ID, Titulo, Descricao, ProfissionalID, IsTemplate, Categoria) em `backend/interno/dominio/questionario.go` | ⏳ |
| TASK-002 | Criar struct `Pergunta` (ID, QuestionarioID, Tipo enum, Texto, Ordem, OpcoesJSON) em `backend/interno/dominio/pergunta.go` | ⏳ |
| TASK-003 | Criar struct `QuestionarioAtribuido` (ID, QuestionarioID, PacienteID, DataAtribuicao, Status enum) | ⏳ |
| TASK-004 | Criar struct `Resposta` (ID, QuestionarioAtribuidoID, PerguntaID, PacienteID, ValorResposta, DataResposta) | ⏳ |
| TASK-005 | Criar migration SQL para todas as 4 tabelas (questionarios, perguntas, questionarios_atribuidos, respostas_questionario) | ⏳ |
| TASK-006 | Adicionar relacionamentos GORM: Questionario ↔ Profissional, Questionario ↔ Perguntas | ⏳ |
| TASK-007 | Criar seeds: 3 templates padrão pré-populados (PHQ-9, GAD-7, Bem-estar) | ⏳ |

### Fase 2: Backend - Repositórios (2-3 dias) 📦

**GOAL-002**: Camada de persistência

| Task | Description | Status |
|------|-------------|--------|
| TASK-008 | Implementar `RepositorioQuestionario`: CRUD + BuscarTemplates + BuscarPorProfissional | ⏳ |
| TASK-009 | Implementar `RepositorioPergunta`: CRUD + BuscarPorQuestionario + InsertBatch | ⏳ |
| TASK-010 | Implementar `RepositorioQuestionarioAtribuido`: CRUD + BuscarPorPaciente + BuscarPendentes | ⏳ |
| TASK-011 | Implementar `RepositorioResposta`: CRUD + CalcularEstatisticas (médias, contadores) | ⏳ |
| TASK-012 | Adicionar índices de BD para otimização: idx_questionarios_profissional, idx_respostas_paciente | ⏳ |

### Fase 3: Backend - Serviços (3-4 dias) ⚙️

**GOAL-003**: Lógica de negócio

| Task | Description | Status |
|------|-------------|--------|
| TASK-013 | Serviço `QuestionarioServico`: Criar, Editar, Deletar, ClonarTemplate, SalvarComoTemplate, ListarQuestionarios | ⏳ |
| TASK-014 | Validação: min 1 pergunta, max 30, tipos válidos, nome não vazio | ⏳ |
| TASK-015 | Serviço `AtribuicaoServico`: AtribuirQuestionario (validar vinculação prof-paciente) + permissões | ⏳ |
| TASK-016 | Serviço `RespostaServico`: SalvarResposta, ValidarResposta (por tipo), BuscarRespostas | ⏳ |
| TASK-017 | Serviço `EstatisticasServico`: CalcularMedia, CalcularTotalizadores, GerarDadosPorPeriodo | ⏳ |
| TASK-018 | Integração com `NotificacaoServico` existente: notificar paciente ao atribuir | ⏳ |

### Fase 4: Backend - APIs (2-3 dias) 🔌

**GOAL-004**: Endpoints RESTful

| Task | Description | Status |
|------|-------------|--------|
| TASK-019 | POST `/api/v1/questionarios` - criar questionário customizado | ⏳ |
| TASK-020 | GET `/api/v1/questionarios` - listar questionários do profissional com filtros | ⏳ |
| TASK-021 | GET `/api/v1/questionarios/:id` - detalhes do questionário | ⏳ |
| TASK-022 | PUT `/api/v1/questionarios/:id` - editar questionário | ⏳ |
| TASK-023 | DELETE `/api/v1/questionarios/:id` - deletar questionário | ⏳ |
| TASK-024 | GET `/api/v1/templates` - listar templates públicos | ⏳ |
| TASK-025 | POST `/api/v1/questionarios/:id/clonar` - clonar template em novo questionário | ⏳ |
| TASK-026 | POST `/api/v1/questionarios/:id/template` - salvar questionário como template privado | ⏳ |
| TASK-027 | POST `/api/v1/questionarios/:id/atribuir` - atribuir a paciente (JSON: {pacienteIDs: []}) | ⏳ |
| TASK-028 | GET `/api/v1/pacientes/:id/questionarios` - questionários pendentes e respondidos do paciente | ⏳ |
| TASK-029 | POST `/api/v1/questionarios-atribuidos/:id/responder` - salvar respostas de paciente | ⏳ |
| TASK-030 | GET `/api/v1/questionarios/:id/estatisticas` - dados agregados (médias, contadores) | ⏳ |
| TASK-031 | Adicionar autenticação JWT + validação de permissões em todos endpoints | ⏳ |

### Fase 5: Frontend - Services & State (2 dias) 🛠️

**GOAL-005**: Camada HTTP + State Management

| Task | Description | Status |
|------|-------------|--------|
| TASK-032 | Criar `questionarioService.js`: todos métodos HTTP para CRUD, templates, atribuir, responder | ⏳ |
| TASK-033 | Criar Pinia store: state, actions, getters para questionários e respostas | ⏳ |
| TASK-034 | Implementar tratamento de erro padronizado com toast notifications | ⏳ |

### Fase 6: Frontend - Prof Dashboard (4-5 dias) 📊

**GOAL-006**: Interface para profissional gerenciar questionários

| Task | Description | Status |
|------|-------------|--------|
| TASK-035 | Página `ProfissionalQuestionarios.vue` - listar questionários + filtros (categoria, meus, templates) | ⏳ |
| TASK-036 | Componente `FormQuestionario.vue` - criar/editar com lista dinâmica de perguntas drag-drop | ⏳ |
| TASK-037 | Componente `SeletorTemplate.vue` modal - selecionar e pré-popular com template | ⏳ |
| TASK-038 | Componente `ConstrutorPergunta.vue` - adicionar/editar/remover perguntas individuais | ⏳ |
| TASK-039 | Componente `ModalAtribuir.vue` - selecionar pacientes e atribuir questionário | ⏳ |
| TASK-040 | Página `EstatisticasQuestionarios.vue` - tabelas com médias e totalizadores por pergunta | ⏳ |
| TASK-041 | Adicionar rotas no Vue Router (`/dashboard-profissional/questionarios`) + Menu sidebar | ⏳ |

### Fase 7: Frontend - Patient Dashboard (3-4 dias) 👤

**GOAL-007**: Interface para paciente responder questionários

| Task | Description | Status |
|------|-------------|--------|
| TASK-042 | Página `QuestionariosPaciente.vue` - lista questionários pendentes e respondidos | ⏳ |
| TASK-043 | Componente `ResponderQuestionario.vue` - renderiza perguntas dinamicamente | ⏳ |
| TASK-044 | Renderização por tipo: radio/select (múltipla), range slider (escala), textarea (texto) | ⏳ |
| TASK-045 | Validação: perguntas obrigatórias, tipos de dados, ranges corretos | ⏳ |
| TASK-046 | Barra de progresso visual (X de Y perguntas respondidas) | ⏳ |
| TASK-047 | Tela de confirmação após envio + histórico de respostas respondidas | ⏳ |
| TASK-048 | Adicionar rotas no Vue Router (`/dashboard-paciente/questionarios`) + Menu sidebar | ⏳ |

### Fase 8: Integração & Relatórios (2 dias) 📈

**GOAL-008**: Integrar dados nos relatórios existentes

| Task | Description | Status |
|------|-------------|--------|
| TASK-049 | Atualizar `ServicosRelatorio`: adicionar widget de "Últimas Respostas a Questionários" | ⏳ |
| TASK-050 | Adicionar dados de questionários na seção de estatísticas/evolução do relatório | ⏳ |
| TASK-051 | Testar fluxo completo: prof cria → atribui → paciente responde → prof vê no relatório | ⏳ |

### Fase 9: Testes & QA (2-3 dias) ✅

**GOAL-009**: Qualidade & confiabilidade

| Task | Description | Status |
|------|-------------|--------|
| TASK-052 | Testes unitários backend: validações, cálculos, repositórios (min 80% cobertura) | ⏳ |
| TASK-053 | Testes de API: CRUD, permissões, validação, erro handling (httptest) | ⏳ |
| TASK-054 | Testes de UI: formulários, validação, responsivo (mobile, tablet, desktop) | ⏳ |
| TASK-055 | Teste E2E: Prof cria → atribui → Paciente responde → Prof ve stats no relatório | ⏳ |
| TASK-056 | Teste de segurança: SQL Injection, XSS em respostas texto, permissões prof-paciente | ⏳ |

### Fase 10: Documentação & Deploy (1-2 dias) 📚

**GOAL-010**: Documentar feature e preparar produção

| Task | Description | Status |
|------|-------------|--------|
| TASK-057 | Criar `/docs/QUESTIONARIOS_MVP.md` com guia de uso para profissionais | ⏳ |
| TASK-058 | Documentar endpoints em `/frontend/swagger-output.json` | ⏳ |
| TASK-059 | Atualizar README.md com features de questionários | ⏳ |
| TASK-060 | Criar migration script + rollback plan para produção | ⏳ |

## 3️⃣ Alternativas Descartadas

- **ALT-001**: Form builder de terceiros (Typeform, Google Forms) - ❌ custo + dependência + perda de controle dados saúde
- **ALT-002**: Questionários hardcoded - ❌ limita flexibilidade profissional para adaptar
- **ALT-003**: GraphQL API - ❌ aumenta complexidade, não necessário para MVP
- **ALT-004**: Gráficos avançados ApexCharts - ❌ tabelas simples são suficientes v1.0
- **ALT-005**: Lógica condicional (if/then/skip) - ❌ postpone para v2.0

## 4️⃣ Dependências Externas

- **DEP-001**: JWT auth middleware existente (`backend/interno/aplicacao/middlewares/`)
- **DEP-002**: Sistema notificações existente (`backend/interno/aplicacao/servicos/notificacao_servico.go`)
- **DEP-003**: Relacionamento Profissional-Paciente existente (tabela `profissional_paciente`)
- **DEP-004**: GORM ORM 1.30.1 para migrations (já usado)
- **DEP-005**: Vue 3 + Pinia 3.0.3 (já instalados)
- **DEP-006**: TailwindCSS 4.1.11 (já instalado)
- **DEP-007**: PostgreSQL 17 com JSON support
- **DEP-008**: Toast notifications sistema (já existe no projeto)

## 5️⃣ Estrutura de Arquivos

### Backend

**Novos arquivos**
- `backend/interno/dominio/questionario.go`
- `backend/interno/dominio/pergunta.go`
- `backend/interno/dominio/questionario_atribuido.go`
- `backend/interno/dominio/resposta_questionario.go`
- `backend/interno/persistencia/repositorios/repositorio_questionario.go`
- `backend/interno/persistencia/repositorio_pergunta.go`
- `backend/interno/persistencia/repositorio_atribuido.go`
- `backend/interno/persistencia/repositorio_resposta.go`
- `backend/interno/aplicacao/servicos/questionario_servico.go`
- `backend/interno/aplicacao/servicos/atribuicao_servico.go`
- `backend/interno/aplicacao/servicos/resposta_servico.go`
- `backend/interno/aplicacao/servicos/estatisticas_servico.go`
- `backend/interno/aplicacao/controladores/questionario_controlador.go`

**Arquivos modificados**
- `backend/cmd/api/main.go` - adicionar rotas
- `backend/interno/aplicacao/servicos/relatorio_servico.go` - integrar dados

### Frontend

**Novos arquivos**
- `frontend/src/services/questionarioService.js`
- `frontend/src/store/questionario.js`
- `frontend/src/store/resposta.js`
- `frontend/src/views/dashboard-profissional/ProfissionalQuestionarios.vue`
- `frontend/src/views/dashboard-profissional/EstatisticasQuestionarios.vue`
- `frontend/src/views/dashboard-paciente/QuestionariosPaciente.vue`
- `frontend/src/views/dashboard-paciente/ResponderQuestionario.vue`
- `frontend/src/components/FormQuestionario.vue`
- `frontend/src/components/ConstrutorPergunta.vue`
- `frontend/src/components/SeletorTemplate.vue`
- `frontend/src/components/ModalAtribuir.vue`

**Arquivos modificados**
- `frontend/src/router/index.js` - adicionar rotas
- `frontend/src/views/dashboard-profissional/ProfissionalDashboard.vue` - adicionar menu
- `frontend/src/views/dashboard-paciente/PacienteDashboard.vue` - adicionar menu

## 6️⃣ Testes Essenciais

| # | Teste | Esperado |
|---|-------|----------|
| TEST-001 | Criar questionário com 3 perguntas | ✓ valida e salva |
| TEST-002 | Editar questionário | ✓ atualiza campos |
| TEST-003 | Salvar como template | ✓ aparece em templates |
| TEST-004 | Atribuir a 2 pacientes | ✓ ambos recebem notificação |
| TEST-005 | Paciente responde | ✓ todas respostas são salvas |
| TEST-006 | Validação pergunta obrigatória vazia | ✓ rejeita |
| TEST-007 | Validação escala 1-10, valor 15 | ✓ rejeita |
| TEST-008 | Visualizar estatísticas | ✓ calcula médias corretas |
| TEST-009 | Deletar questionário | ✓ respostas não são deletadas |
| TEST-010 | Prof sem vínculo tenta atribuir | ✓ rejeita por segurança |
| TEST-011 | Teste responsivo mobile/tablet/desktop | ✓ funciona em todos |
| TEST-012 | Fluxo E2E completo end-to-end | ✓ sem erros |

## 7️⃣ Riscos & Mitigações

| Risk | Impacto | Mitigação |
|------|--------|-----------|
| Escopo cresce durante dev | Alto | Documentar MVP claramente, criar backlog v2.0 |
| Performance com muitas respostas | Médio | Paginação, índices de BD, caching |
| Validação clínica incorreta | Alto | Usar templates reconhecidos, validar com profissionais |
| Mudanças em schema produção | Alto | Testes em staging, backups, rollback plan |
| UI muito complexa | Médio | Começar simples (lista ordenável), iterar |
| Bugs em permissões | Alto | Testes de segurança, code review |

## 8️⃣ Suposições

- Profissionais têm conhecimento mínimo para criar questionários estruturados
- Pacientes têm acesso internet regular
- PostgreSQL comporta ~10k questionários, ~500k respostas (ano 1)
- Templates PHQ-9, GAD-7 são domínio público
- Profissionais preferem web (desktop-first, responsive)
- Notificações email são suficientes (sem push mobile)
- Português apenas na v1.0
- Sistema permissões atual (vinculação prof-paciente) é suficiente

## 9️⃣ Timeline Estimada

| Fase | Dias | Acumulado |
|------|------|-----------|
| 1: Banco & Domínio | 3-4 | 3-4 |
| 2: Repositórios | 2-3 | 5-7 |
| 3: Serviços | 3-4 | 8-11 |
| 4: APIs | 2-3 | 10-14 |
| 5: Services Frontend | 2 | 12-16 |
| 6: Prof Dashboard | 4-5 | 16-21 |
| 7: Patient Dashboard | 3-4 | 19-25 |
| 8: Integração/Relatórios | 2 | 21-27 |
| 9: Testes | 2-3 | 23-30 |
| 10: Documentação | 1-2 | 24-32 |
| **TOTAL** | **20-30 dias** | **20-30 dias** |

**Nota**: Fases podem ser paralelizadas. Com focus: ~25 dias

## 1️⃣0️⃣ Próximos Passos

1. ✅ Revisar e aprovar plano
2. ➡️ **Fase 1**: Criar structs de domínio
3. ➡️ **Fase 2**: Criar migrations SQL
4. ➡️ **Fase 3**: Seeds com templates padrão
5. ➡️ **Fase 4**: Implementar repositórios
6. ➡️ **Fase 5**: Testar CRUD básico
7. ➡️ **Fase 6**: Desenvolver serviços
8. ➡️ **Fase 7**: Criar APIs
9. ➡️ **Fase 8**: Build frontend
10. ➡️ **Fase 9**: Integração com relatórios
11. ➡️ **Fase 10**: Testes completos
12. ➡️ **Deploy**

---

## 📌 Metadados

| Campo | Valor |
|-------|-------|
| **Criado em** | 2025-10-16 |
| **Atualizado em** | 2025-10-20 |
| **Duração estimada** | 20-30 dias |
| **Prioridade** | Alta ⬆️ |
| **Contexto** | Disciplina Engenharia de Software - Projeto MindTrace |
| **Status** | Planned 🔵 |
