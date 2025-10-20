---
goal: "Sistema de Alertas e Notificações por Email"
version: "1.0"
date_created: "2025-10-20"
last_updated: "2025-10-20"
owner: "Desenvolvedor"
status: 'Planned'
tags: ['feature', 'alerts', 'notifications', 'email', 'mvp']
---

# Sistema de Alertas e Notificações

![Status: Planned](https://img.shields.io/badge/status-Planned-blue)

Sistema MVP de alertas automáticos e notificações por email para MindTrace. Profissionais e pacientes recebem notificações sobre: padrões preocupantes nos dados, lembretes para completar formulários (Registro de Humor, Questionários) e eventos importantes do acompanhamento.

**Otimizado para**: 1 desenvolvedor | 15-20 dias | Projeto acadêmico | Email MVP (SMS/WhatsApp v2.0)

## 1. Requirements & Constraints

### Requisitos Funcionais (MVP)

- **REQ-001**: Sistema detecta automaticamente quando estatísticas caem abaixo de um limiar preocupante (ex: humor muito baixo, stress muito alto)
- **REQ-002**: Profissional recebe notificação email quando paciente tem padrão preocupante
- **REQ-003**: Paciente recebe notificação email lembrando de responder Registro de Humor diário
- **REQ-004**: Paciente recebe notificação email quando profissional atribui novo questionário
- **REQ-005**: Profissional recebe notificação email quando paciente completa questionário
- **REQ-006**: Sistema envia emails de forma assíncrona (não bloqueia requisição HTTP)
- **REQ-007**: Usuários podem desabilitar notificações por tipo (settings/preferências)
- **REQ-008**: Log de notificações enviadas para auditoria
- **REQ-009**: Notificações podem ser configuradas (dias/horas de envio, frequência)

### Requisitos Técnicos

- **TEC-001**: Backend segue arquitetura existente (Domain → Application → Persistence)
- **TEC-002**: Envio assíncrono com workers/goroutines (não framework pesado)
- **TEC-003**: SMTP para envio de email (pode ser Gmail, Sendgrid, Mailgun, servidor SMTP local)
- **TEC-004**: Fila de notificações em BD (tabela simples) ou Redis (optional)
- **TEC-005**: Scheduling de notificações periódicas (cron-like)
- **TEC-006**: Sem alteração no esquema de autenticação JWT

### Constraints & Simplificações

- **CON-001**: MVP - apenas EMAIL (sem SMS/WhatsApp)
- **CON-002**: Sem dashboard visual de notificações (apenas email + log em BD)
- **CON-003**: Sem retry inteligente (reenvio simples se falhar)
- **CON-004**: Sem templating avançado (email básico HTML simples)
- **CON-005**: Sem personalizações avançadas (nomes, dados do paciente nos templates)
- **CON-006**: Notificações diárias máximo (não horária ou minuto a minuto)
- **CON-007**: Um desenvolvedor - máximo 20 dias

### Padrões & Guidelines

- **PAT-001**: Repository Pattern para dados de notificações
- **PAT-002**: Service Layer para lógica de detecção e envio
- **PAT-003**: Factory Pattern para criar diferentes tipos de notificações
- **PAT-004**: Event-Driven (quando algo acontece, dispara notificação)
- **GUD-001**: Português para nomenclatura de domínio
- **GUD-002**: Nomes descritivos: Alerta, Notificacao, TipoNotificacao
- **GUD-003**: Templates de email em arquivos separados (reutilizável)

## 2. Implementation Steps

### Phase 1: Banco & Domínio (2-3 dias)

**GOAL-001**: Estrutura de dados para notificações e alertas

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-001 | Criar struct `TipoNotificacao` enum em `backend/interno/dominio/` (LEMBRETES_HUMOR, NOVO_QUESTIONARIO, RESULTADO_QUESTIONARIO, ALERTA_PREOCUPANTE) | | |
| TASK-002 | Criar struct `Notificacao` (ID, UsuarioID, TipoNotificacao, Titulo, Mensagem, DataCriacao, DataEnvio, Lido, Enviado, TentativasEnvio) | | |
| TASK-003 | Criar struct `Alerta` (ID, PacienteID, Tipo, Descricao, Severidade enum, DataDeteccao, Resolvido, ProfissionalNotificado) | | |
| TASK-004 | Criar struct `PreferenciaNotificacao` (ID, UsuarioID, TipoNotificacao, Habilitada, DiasSemana, HoraPreferida) | | |
| TASK-005 | Criar struct `ConfiguracaoAlerta` (ID, Tipo, Limiar, Metrica, CompetenciaProf) com valores padrão (ex: Humor < 2 = alerta) | | |
| TASK-006 | Migration SQL: tabelas notificacoes, alertas, preferencias_notificacao, configuracoes_alerta | | |
| TASK-007 | Relacionamentos GORM: Notificacao ↔ Usuario, Alerta ↔ Paciente, PreferenciaNotificacao ↔ Usuario | | |
| TASK-008 | Seeds: preferências padrão (todas habilitadas), configurações padrão de alertas | | |

### Phase 2: Serviço de Email (2-3 dias)

**GOAL-002**: Configuração de envio de email

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-009 | Criar struct `ConfigEmail` em `backend/interno/aplicacao/` com: Host, Port, Usuario, Senha, De, NomeRemetente (ler de .env) | | |
| TASK-010 | Criar serviço `EmailServico` em `backend/interno/aplicacao/servicos/` com método `EnviarEmail(para, assunto, html)` | | |
| TASK-011 | Implementar conexão SMTP usando `net/smtp` (Go stdlib, sem dependências externas) | | |
| TASK-012 | Implementar envio assíncrono com goroutines (não bloqueia requisição) | | |
| TASK-013 | Implementar fila simples em BD (tabela `fila_emails` com status: pendente, enviado, erro) | | |
| TASK-014 | Criar templates de email em arquivos `backend/templates/emails/`: lembretes_humor.html, novo_questionario.html, resultado_questionario.html, alerta_preocupante.html | | |
| TASK-015 | Implementar retry automático para emails que falharem (máx 3 tentativas) | | |
| TASK-016 | Adicionar logging de emails enviados/falhados | | |

### Phase 3: Detecção de Alertas (2-3 dias)

**GOAL-003**: Lógica para detectar padrões preocupantes

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-017 | Serviço `DeteccaoAlertasServico`: método `VerificarPadroesPaciente(pacienteID)` | | |
| TASK-018 | Implementar regras de detecção: Humor consecutivo baixo (< 2 por 3+ dias), Stress alto (8+ por 5+ dias), Sono reduzido (< 4h por 7+ dias) | | |
| TASK-019 | Implementar regra: Energia muito baixa (< 3 por 3+ dias) | | |
| TASK-020 | Implementar regra: Não respondeu Registro de Humor há 3+ dias (alerta mais brando) | | |
| TASK-021 | Implementar regra: Questionário pendente há 7+ dias (alerta ao paciente) | | |
| TASK-022 | Armazenar alertas gerados em BD com timestamp e severidade | | |
| TASK-023 | Marcar alerta como resolvido quando paciente dados melhoram (2+ dias bom) | | |

### Phase 4: Serviço de Notificações (2-3 dias)

**GOAL-004**: Lógica de criação e envio de notificações

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-024 | Serviço `NotificacaoServico`: CRUD de notificações + preferências | | |
| TASK-025 | Método `NotificarPacienteLembreteHumor(pacienteID)` - email + notificacao em BD | | |
| TASK-026 | Método `NotificarPacienteNovoQuestionario(pacienteID, questionarioID)` - disparado ao atribuir | | |
| TASK-027 | Método `NotificarProfissionalResultadoQuestionario(profissionalID, pacienteID, questionarioID)` - disparado ao responder | | |
| TASK-028 | Método `NotificarProfissionalAlertaPreocupante(profissionalID, pacienteID, alertaID)` - disparado ao detectar alerta | | |
| TASK-029 | Método `NotificarPacienteAlertaDetectado(pacienteID, alertaID)` - opcional, notifica paciente também | | |
| TASK-030 | Implementar verificação de preferências (se notificação está habilitada antes de enviar) | | |
| TASK-031 | Implementar verificação de dias/horas preferenciais (se houver, só envia naquele horário) | | |

### Phase 5: Gatilhos de Eventos (2-3 dias)

**GOAL-005**: Disparar notificações quando eventos ocorrem

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-032 | Ao salvar `RegistroHumor`: chamar `DeteccaoAlertasServico.VerificarPadroesPaciente()` e gerar alertas se necessário | | |
| TASK-033 | Ao salvar `Resposta` de questionário: chamar `NotificacaoServico.NotificarProfissionalResultado()` | | |
| TASK-034 | Ao atribuir `QuestionarioAtribuido`: chamar `NotificacaoServico.NotificarPacienteNovoQuestionario()` | | |
| TASK-035 | Ao detectar alerta: chamar `NotificacaoServico.NotificarProfissionalAlertaPreocupante()` | | |
| TASK-036 | Implementar goroutine assíncrona para envio de email (não bloqueia requisição HTTP) | | |
| TASK-037 | Implementar job periódico (diário) para: verificar alertas pendentes, enviar lembretes de humor | | |
| TASK-038 | Implementar job para limpar notificações antigas (>30 dias) | | |

### Phase 6: APIs para Gerenciar Preferências (1-2 dias)

**GOAL-006**: Endpoints para usuários configurarem notificações

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-039 | GET `/api/v1/usuarios/notificacoes/preferencias` - listar preferências do usuário autenticado | | |
| TASK-040 | PUT `/api/v1/usuarios/notificacoes/preferencias` - atualizar preferências (habilitar/desabilitar por tipo) | | |
| TASK-041 | GET `/api/v1/usuarios/notificacoes` - listar notificações do usuário (historico) | | |
| TASK-042 | PUT `/api/v1/usuarios/notificacoes/:id/lido` - marcar notificação como lida | | |
| TASK-043 | GET `/api/v1/alertas` (profissional) - listar alertas de seus pacientes | | |
| TASK-044 | GET `/api/v1/alertas/:id` - detalhes do alerta | | |
| TASK-045 | PUT `/api/v1/alertas/:id/resolvido` - marcar alerta como resolvido | | |
| TASK-046 | Adicionar autenticação JWT + validação de permissões | | |

### Phase 7: Frontend - Settings (2-3 dias)

**GOAL-007**: Interface para gerenciar preferências de notificação

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-047 | Criar página `ConfiguracaoNotificacoes.vue` em `frontend/src/views/` | | |
| TASK-048 | Componente `PreferenciaNotificacao.vue` - toggle habilitada/desabilitada por tipo | | |
| TASK-049 | Seletor de dia(s) da semana preferencial (segunda-sexta, fins de semana, etc) | | |
| TASK-050 | Time picker para hora preferencial (ex: 9h da manhã) | | |
| TASK-051 | Tabela com histórico de notificações enviadas (últimas 30) | | |
| TASK-052 | Service `notificacaoService.js` para chamar APIs | | |
| TASK-053 | Pinia store `useNotificacaoStore` para state management | | |
| TASK-054 | Adicionar rota no Vue Router + menu no dashboard | | |

### Phase 8: Frontend - Dashboard Prof (1-2 dias)

**GOAL-008**: Visualizar alertas de pacientes

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-055 | Adicionar widget "Alertas Recentes" no dashboard profissional com lista de alertas preocupantes | | |
| TASK-056 | Componente `CardAlerta.vue` mostrando: paciente, tipo de alerta, severidade, data | | |
| TASK-057 | Link para visualizar detalhes do paciente/alerta | | |
| TASK-058 | Badge de "alerta novo" se houver alertas não lidos | | |

### Phase 9: Testes & QA (2-3 dias)

**GOAL-009**: Qualidade & confiabilidade

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-059 | Testes unitários: detectar alerta quando humor < 2 por 3 dias consecutivos | | |
| TASK-060 | Testes unitários: não detectar alerta se humor melhora | | |
| TASK-061 | Testes unitários: envio de email com template correto | | |
| TASK-062 | Testes de integração: salvar RegistroHumor → verifica alertas → envia email | | |
| TASK-063 | Testes de API: GET preferências, PUT preferências, verificar autenticação | | |
| TASK-064 | Teste E2E: Paciente salva humor baixo → Profissional recebe email de alerta | | |
| TASK-065 | Teste segurança: usuário não consegue acessar notificações de outro usuário | | |
| TASK-066 | Teste email: verificar que email não é enviado se usuário desabilitar notificação | | |

### Phase 10: Documentação & Deployment (1-2 dias)

**GOAL-010**: Documentar feature e preparar produção

| Task | Description | Completed | Date |
|------|-------------|-----------|------|
| TASK-067 | Criar arquivo `.env.example` com variáveis de email (SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASS, EMAIL_FROM) | | |
| TASK-068 | Documentar em `/docs/ALERTAS_NOTIFICACOES.md`: como funciona, tipos de alertas, como configurar email | | |
| TASK-069 | Documentar endpoints de preferências em `/swagger-output.json` | | |
| TASK-070 | Criar migration script + rollback plan | | |
| TASK-071 | Testar em staging com email real (usar conta teste Gmail ou similar) | | |
| TASK-072 | Atualizar README.md com feature de alertas | | |

## 3. Alternativas Descartadas

- **ALT-001**: Usar serviço externo (SendGrid, Mailgun) - ❌ adiciona custo/dependência; SMTP direto é simples
- **ALT-002**: Implementar push notifications mobile - ❌ complexidade, postpone para v2.0
- **ALT-003**: Usar framework pesado (Bull, RabbitMQ) - ❌ BD simples + goroutines é suficiente MVP
- **ALT-004**: Dashboard visual em tempo real - ❌ email+BD é suficiente, UI pode ser adicionada depois
- **ALT-005**: SMS/WhatsApp na v1.0 - ❌ focar em email primeiro, mais simples

## 4. Dependencies

- **DEP-001**: Go stdlib (net/smtp, mime/quotedprintable) - nenhuma dependência externa adicional
- **DEP-002**: Banco de dados PostgreSQL existente
- **DEP-003**: JWT auth middleware existente
- **DEP-004**: Serviço de RegistroHumor existente (para disparar eventos)
- **DEP-005**: Serviço de Questionarios (já implementado ou em desenvolvimento)
- **DEP-006**: Vue 3 + Pinia (já instalados)
- **DEP-007**: TailwindCSS (já instalado)
- **DEP-008**: Servidor SMTP (Gmail, servidor local, ou provider)

## 5. Arquivos - Backend

**Novos**
- `backend/interno/dominio/notificacao.go`
- `backend/interno/dominio/alerta.go`
- `backend/interno/dominio/preferencia_notificacao.go`
- `backend/interno/dominio/tipo_notificacao.go` (enum)
- `backend/interno/persistencia/repositorios/repositorio_notificacao.go`
- `backend/interno/persistencia/repositorio_alerta.go`
- `backend/interno/persistencia/repositorio_preferencia.go`
- `backend/interno/aplicacao/servicos/email_servico.go`
- `backend/interno/aplicacao/servicos/notificacao_servico.go`
- `backend/interno/aplicacao/servicos/deteccao_alertas_servico.go`
- `backend/interno/aplicacao/controladores/notificacao_controlador.go`
- `backend/interno/aplicacao/controladores/alerta_controlador.go`
- `backend/interno/aplicacao/jobs/job_alertas_periodico.go`
- `backend/templates/emails/lembretes_humor.html`
- `backend/templates/emails/novo_questionario.html`
- `backend/templates/emails/resultado_questionario.html`
- `backend/templates/emails/alerta_preocupante.html`

**Modificados**
- `backend/cmd/api/main.go` - adicionar rotas, inicializar jobs
- `backend/interno/aplicacao/servicos/registro_humor_servico.go` - disparar detectação de alertas
- `backend/interno/aplicacao/servicos/resposta_servico.go` - notificar profissional
- `backend/interno/aplicacao/servicos/atribuicao_servico.go` - notificar paciente
- `.env.example` - adicionar variáveis SMTP

## 6. Arquivos - Frontend

**Novos**
- `frontend/src/views/ConfiguracaoNotificacoes.vue`
- `frontend/src/components/PreferenciaNotificacao.vue`
- `frontend/src/components/CardAlerta.vue`
- `frontend/src/services/notificacaoService.js`
- `frontend/src/store/notificacao.js`

**Modificados**
- `frontend/src/router/index.js` - adicionar rota settings
- `frontend/src/views/dashboard-profissional/ProfissionalDashboard.vue` - adicionar widget alertas
- Sidebar de ambos dashboards - adicionar link para preferências

## 7. Testes Essenciais

| # | Test | Esperado |
|---|------|----------|
| TEST-001 | Detectar alerta humor < 2 por 3 dias | ✓ gera alerta com severidade alta |
| TEST-002 | Detectar alerta stress > 8 por 5 dias | ✓ gera alerta |
| TEST-003 | Resolver alerta quando dados melhoram | ✓ marca como resolvido |
| TEST-004 | Não enviar email se notificação desabilitada | ✓ email não enviado |
| TEST-005 | Enviar email com template correto | ✓ email recebido com conteúdo correto |
| TEST-006 | Retry automático se email falha | ✓ tenta 3 vezes |
| TEST-007 | Atribuir questionário → paciente recebe email | ✓ email dentro de 1 minuto |
| TEST-008 | Paciente responde questionário → prof recebe email | ✓ email dentro de 1 minuto |
| TEST-009 | Profissional não consegue ver alertas de outro prof | ✓ segurança ok |
| TEST-010 | Usuário consegue mudar preferências | ✓ salva e respeita |
| TEST-011 | Email não enviado fora do horário preferencial | ✓ espera até horário correto |
| TEST-012 | Job periódico detecta alertas | ✓ cria alertas à 0h diariamente |

## 8. Riscos & Mitigações

| Risk | Impacto | Mitigação |
|------|--------|-----------|
| Email spam (muitos alertas) | Médio | Limitar frequência, preferências, deduplicação |
| SMTP falha/indisponível | Alto | Retry automático, fila de emails em BD, log detalhado |
| Algoritmo de detecção errado | Alto | Consultar literatura, validar com profissionais, testes abrangentes |
| Performance de detecção (muito lento) | Médio | Rodas em goroutine async, não bloqueia |
| Alertas falsos (muitos positivos) | Alto | Fine-tuning dos limiares baseado em feedback |

## 9. Assumptions

- Limiares de alertas definidos são razoáveis (podem ser ajustados depois)
- Profissionais têm email válido no sistema
- Pacientes têm email válido no sistema
- SMTP está disponível (Gmail, servidor local, ou provider)
- Banco de dados comporta logs de notificações
- Usuários recebem emails em inbox (não spam)
- Job periódico pode rodar a cada minuto (não trava o app)

## 10. Estimativa de Timeline

| Phase | Dias | Cumulative |
|-------|------|-----------|
| 1: Banco & Domínio | 2-3 | 2-3 |
| 2: Email | 2-3 | 4-6 |
| 3: Detecção Alertas | 2-3 | 6-9 |
| 4: Serviço Notificações | 2-3 | 8-12 |
| 5: Gatilhos Eventos | 2-3 | 10-15 |
| 6: APIs Preferências | 1-2 | 11-17 |
| 7: Frontend Settings | 2-3 | 13-20 |
| 8: Frontend Prof Dashboard | 1-2 | 14-22 |
| 9: Testes | 2-3 | 16-25 |
| 10: Documentação | 1-2 | 17-27 |
| **TOTAL** | **15-20** | **15-20** |

*Nota: Fases 2-5 podem ser paralelizadas. Com focus: ~18 dias*

## 11. Configuração de Email (Exemplo Gmail)

### Usando Gmail SMTP

```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=seu-email@gmail.com
SMTP_PASS=sua-senha-app (gerar em Google Account > Security)
EMAIL_FROM=noreply@mindtrace.com
EMAIL_FROM_NAME=MindTrace Alertas
```

### Usando Servidor Local (MailHog para dev)

```env
SMTP_HOST=localhost
SMTP_PORT=1025
SMTP_USER=
SMTP_PASS=
EMAIL_FROM=dev@mindtrace.local
```

## 12. Tipos de Alerta (Configuráveis)

| Tipo | Limiar | Severidade | Notifica |
|------|--------|-----------|----------|
| HUMOR_MUITO_BAIXO | Humor < 2 por 3+ dias | Alta | Prof + Paciente |
| STRESS_MUITO_ALTO | Stress > 8 por 5+ dias | Alta | Prof + Paciente |
| SONO_REDUZIDO | Sono < 4h por 7+ dias | Média | Prof |
| ENERGIA_BAIXA | Energia < 3 por 3+ dias | Média | Prof |
| HUMOR_NAO_RESPONDIDO | Sem registro há 3+ dias | Baixa | Paciente |
| QUESTIONARIO_PENDENTE | Sem resposta há 7+ dias | Média | Paciente |

## 13. Próximos Passos Imediatos

1. ✅ Revisar e aprovar plano
2. ➡️ Começar Phase 1: criar structs de domínio
3. ➡️ Criar migrations SQL
4. ➡️ Implementar EmailServico
5. ➡️ Criar templates de email
6. ➡️ Implementar DeteccaoAlertasServico
7. ➡️ Implementar NotificacaoServico
8. ➡️ Adicionar gatilhos nos serviços existentes
9. ➡️ Criar APIs para preferências
10. ➡️ Build frontend (settings + alertas)
11. ➡️ Testes completos
12. ➡️ Deploy

---

**Criado em**: 2025-10-20  
**Duração estimada**: 15-20 dias  
**Prioridade**: Alta  
**Contexto**: Disciplina Engenharia de Software - Projeto MindTrace  
**Próximas versões**: SMS (v2.0), WhatsApp (v3.0), Push Notifications (v4.0)
