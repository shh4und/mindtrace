---
goal: Implementação de Questionários e Escalas Padronizadas (MVP)
version: 1.0
date_created: 2025-11-20
status: Planned
tags: [feature, backend, frontend, mvp, clinical]
---

# Introduction

![Status: Planned](https://img.shields.io/badge/status-Planned-blue)

Este plano detalha a implementação do módulo de questionários e escalas psicométricas no MindTrace. O foco deste MVP é a disponibilização de instrumentos padronizados e validados cientificamente (como PHQ-9, GAD-7, WHOQOL-BREF) de forma imutável, garantindo a integridade clínica dos dados. O sistema permitirá a atribuição destes instrumentos a pacientes, a coleta de respostas e o cálculo automático de scores.

## 1. Requirements & Constraints

- **REQ-001**: O sistema deve disponibilizar instrumentos padronizados (PHQ-9, GAD-7, WHOQOL-BREF) pré-carregados.
- **CON-001**: Instrumentos padronizados devem ser **imutáveis**. Não será permitida a edição de textos, perguntas ou lógica de cálculo pelos usuários para manter a validade estatística.
- **REQ-002**: Profissionais devem poder atribuir instrumentos a seus pacientes vinculados.
- **REQ-003**: O sistema deve registrar as respostas dos pacientes e calcular o score final automaticamente baseando-se no algoritmo específico de cada instrumento.
- **REQ-004**: Armazenamento híbrido: Estrutura relacional para metadados e relatórios, e JSONB para armazenamento bruto das respostas (flexibilidade).
- **REQ-005**: O sistema deve permitir a visualização dos resultados e histórico de respostas.

## 2. Implementation Steps

### Implementation Phase 1: Persistência e Domínio

- GOAL-001: Estruturar o banco de dados e as entidades de domínio para suportar instrumentos, atribuições e respostas.

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-001 | Criar migrações SQL para tabelas `instrumentos`, `perguntas`, `opcoes` (com suporte a versionamento/imutabilidade). | | |
| TASK-002 | Criar migrações SQL para tabelas `atribuicoes` (vínculo paciente-instrumento) e `respostas` (com coluna JSONB). | | |
| TASK-003 | Implementar structs de domínio em `backend/interno/dominio/` (`Instrumento`, `Pergunta`, `Atribuicao`, `Resposta`). | | |
| TASK-004 | Criar seeders (SQL ou Go) para popular o banco com PHQ-9, GAD-7 e WHOQOL-BREF (dados hardcoded). | | |

### Implementation Phase 2: Lógica de Aplicação (Backend)

- GOAL-002: Implementar a lógica de negócios, incluindo repositórios, serviços e algoritmos de cálculo de score (Strategy Pattern).

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-005 | Implementar `InstrumentoRepositorio` (foco em leitura e listagem de padronizados). | | |
| TASK-006 | Implementar `AtribuicaoRepositorio` e `RespostaRepositorio` (operações transacionais). | | |
| TASK-007 | Implementar padrão Strategy para algoritmos de score (`ScoringStrategy`, `Phq9Strategy`, `Gad7Strategy`, etc.). | | |
| TASK-008 | Implementar `InstrumentoServico` (listagem), `AtribuicaoServico` (criação de vínculo) e `RespostaServico` (processamento e cálculo). | | |

### Implementation Phase 3: API e Integração Frontend

- GOAL-003: Expor os serviços via API REST e implementar a interface de usuário para listagem, resposta e visualização.

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-009 | Criar `InstrumentoControlador` com endpoints `GET /instrumentos` e `GET /instrumentos/{id}` (detalhes completos para renderização). | | |
| TASK-010 | Criar `AtribuicaoControlador` (`POST /atribuicoes`, `GET /pacientes/{id}/pendencias`) e `RespostaControlador` (`POST /respostas`). | | |
| TASK-011 | Implementar Frontend: Tela de listagem de instrumentos disponíveis para o profissional. | | |
| TASK-012 | Implementar Frontend: Componente de Formulário Dinâmico (renderiza perguntas/opções baseado no JSON do instrumento). | | |
| TASK-013 | Implementar Frontend: Visualização de resultados (Score calculado e interpretação). | | |

## 3. Alternatives

- **ALT-001**: **Permitir edição de templates.** Considerada e rejeitada. Permitir que profissionais alterem perguntas de escalas validadas (ex: mudar "tristeza" para "chateação" no PHQ-9) invalida o instrumento cientificamente e torna o score sem sentido clínico.
- **ALT-002**: **Criação de questionários personalizados.** Adiada para versões futuras. O foco atual é garantir a aplicação de métricas baseadas em evidência (Data-Critical).

## 4. Dependencies

- **DEP-001**: GORM (Go) para persistência de dados.
- **DEP-002**: PostgreSQL com suporte a JSONB.
- **DEP-003**: Frontend Framework (Vue.js) para renderização dinâmica.

## 5. Files

- `backend/interno/dominio/instrumento.go`
- `backend/interno/dominio/atribuicao.go`
- `backend/interno/dominio/resposta.go`
- `backend/interno/persistencia/instrumento_repositorio.go`
- `backend/interno/persistencia/atribuicao_repositorio.go`
- `backend/interno/persistencia/resposta_repositorio.go`
- `backend/interno/aplicacao/servicos/instrumento_servico.go`
- `backend/interno/aplicacao/servicos/resposta_servico.go`
- `backend/interno/aplicacao/algoritmos/scoring.go`
- `backend/cmd/api/controladores/instrumento_controlador.go`
- `seeders/instrumentos_padrao.sql`

## 6. Testing

- **TEST-001**: Testes unitários para cada `ScoringStrategy` (garantir que PHQ-9 com somatório X dê o resultado Y).
- **TEST-002**: Teste de integração para o fluxo: Atribuir -> Responder -> Verificar Persistência e Score.
- **TEST-003**: Validação de imutabilidade (tentar alterar um instrumento via repositório deve falhar ou não ser exposto).

## 7. Risks & Assumptions

- **RISK-001**: Profissionais podem solicitar escalas específicas não incluídas no seed inicial.
- **ASSUMPTION-001**: Os instrumentos PHQ-9, GAD-7 e WHOQOL-BREF cobrem a maior parte das necessidades iniciais de monitoramento.

## 8. Related Specifications / Further Reading

- [Documentação PHQ-9](https://www.phqscreeners.com/)
- [Documentação GAD-7](https://gad7.org/)
