Relatório de Auditoria Técnica: MindTrace

Data: 30 de Janeiro de 2026
Foco: Segurança, Confiabilidade e Conformidade (LGPD)
Status Geral: ⚠️ Requer Atenção Imediata em pontos críticos de segurança e dados sensíveis.

---

1. Segurança e Autenticação

🔴 Crítico: Gestão de Sessão Vulnerável

- Problema: O sistema utiliza JWTs (provavelmente de longa duração, 24h) sem mecanismo de Refresh Token ou revogação.
- Risco: Se um token for roubado, o atacante tem acesso total por 24h, mesmo que a senha seja trocada.
- Código Afetado: AutControlador e AutMiddleware.
- Ação Necessária: Implementar par Access Token (curta duração, ex: 15min) + Refresh Token (banco de dados).

🟠 Médio: Ausência de Rate Limiting

- Problema: Não há middleware de limitação de requisições configurado no main.go.
- Risco: Endpoints públicos como /api/v1/entrar/login e /api/v1/entrar/ativar/reenviar estão expostos a ataques de força bruta e abuso de
  envio de e-mails (custo financeiro/reputação de IP).

---

2. Confiabilidade e Estabilidade

🔴 Crítico: Risco de Panic e Crash Silencioso

- Problema: O envio de e-mail em analise_servico.go (linha 134) é feito em uma go func sem proteção de recover.
- Evidência:

```
  1     go func(emailProf, nomeProf, nomePac, st string) {
  2         s.emailServico.EnviarEmailAlertaMonitoramento(...)
  3     }(...)
```

- Risco: Se o serviço de e-mail falhar com um panic (ex: ponteiro nulo inesperado), toda a aplicação API cairá, derrubando o serviço para
  todos os usuários.

🟠 Médio: Bloqueio de Thread (Starvation)

- Problema: smtp.SendMail em email_servico.go é bloqueante e não utiliza context.WithTimeout.
- Risco: Se o servidor SMTP demorar para responder (network lag), a goroutine fica travada indefinidamente. Em carga alta, isso pode
  esgotar os recursos do servidor.

---

3. Conformidade LGPD e Proteção de Dados

🔴 Crítico: Vazamento de Dados Sensíveis em Logs

- Problema: O AnaliseServico loga dados de saúde do paciente no console (stdout).
- Evidência: backend/interno/aplicacao/servicos/analise_servico.go:

```
1     // L148
2     log.Printf("... Dados: mediaHumor: %.2f, mediaStress: %.2f ...", ...)
```

- Violação: Isso expõe dados sensíveis (Art. 5º, II da LGPD) em logs que podem ser persistidos em texto plano por ferramentas de
  infraestrutura (Docker logs, AWS CloudWatch), violando o princípio de minimização.

🟠 Médio: Insegurança em Templates de E-mail (XSS)

- Problema: O EmailServico utiliza o pacote text/template para gerar e-mails HTML.
- Risco: text/template não sanitiza entradas automaticamente. Se um usuário malicioso colocar um script no próprio nome `( <script>...)`, e
  esse nome for renderizado no e-mail de um profissional, o script pode ser executado no cliente de e-mail (XSS).
- Solução: Migrar urgentemente para html/template.

🟡 Baixo: Falta de Rastreabilidade de Consentimento

- Problema: A entidade Usuario não possui campos para versionamento de Termos de Uso ou data de aceite.
- Risco: Dificuldade jurídica em provar consentimento explícito em caso de auditoria.

---

📝 Plano de Implementação (Hardening)

Sugiro seguir esta ordem para a branch chore/test-refresh ou tech/security-hardening.

Fase 1: Estancar Sangramentos (Confiabilidade)

1. Refatorar `EmailServico`:
   - Substituir text/template por html/template.
   - Envolver smtp.SendMail em um canal com select e time.After (ou usar uma lib que suporte Context).
2. Proteger Goroutines:
   - Criar um wrapper GoSafe(fn func()) que executa a função e faz defer recover() logando o erro sem crashar o app. Aplicar em
     AnaliseServico.
3. Sanitizar Logs:
   - Remover dados de métricas (humor, stress) dos logs de AnaliseServico. Logar apenas PacienteID e o fato de que o monitoramento
     ocorreu.

Fase 2: Segurança de Acesso

1. Rate Limiter:
   - Adicionar gin-limiter (ou similar) no main.go, configurando limites estritos para rotas de /entrar e limites mais relaxados para
     rotas autenticadas.
2. Refatorar Auth (Maior Esforço):
   - Criar tabela refresh_tokens.
   - Alterar login para retornar access_token (15m) e refresh_token (7d).
   - Criar endpoint POST /refresh.

Fase 3: Conformidade LGPD

1. Auditoria de Banco:
   - Adicionar campos TermosAceitosEm e TermosVersao na tabela usuarios.
2. Endpoint de Exclusão:
   - Implementar lógica de "Anonimização" no UsuarioServico (sobrescrever nome/email com hash e deletar soft).

---

🛡️ Resumo das Alterações Realizadas

1. Confiabilidade e Estabilidade (Fase 1)

- Envio de E-mail Seguro: Implementei um timeout de 10 segundos no envio SMTP para evitar bloqueio de threads.
- Proteção contra Crashes: Adicionei defer recover() na goroutine de alertas de monitoramento. Se o serviço de e-mail falhar, a API não
  cai mais.
- Logs Limpos: Removi logs que expunham dados sensíveis de saúde e o corpo de e-mails com tokens no console.

2. Segurança de Acesso (Fase 2)

- Rate Limiting: Implementei um middleware de proteção contra força bruta.
  - Rotas críticas (/entrar, /registrar) agora têm limite de 1 requisição a cada 2 segundos (burst de 5).
- Autenticação Reforçada (JWT + Refresh Token):
  - Criei a tabela refresh_tokens.
  - O endpoint de login agora retorna um par de chaves: access_token (15 min) e refresh_token (7 dias).
  - Criei a rota POST /api/v1/entrar/refresh para renovação segura de sessão.

3. Conformidade LGPD (Fase 3)

- Consentimento Explícito:
  - Adicionei campos TermosAceitosEm e TermosVersao no banco.
  - O cadastro agora exige o envio de termos_aceitos: true para prosseguir.
- Direito ao Esquecimento (Anonimização):
  - Implementei a rota DELETE /api/v1/usuarios/perfil/anonimizar.
  - Diferente do "apagar conta" (que destrói tudo), a anonimização substitui Nome, Email e CPF por dados fictícios e desativa o acesso,
    mas mantém os dados de saúde (registros_humor) no banco para fins estatísticos, respeitando a privacidade.

⚠️ Próximos Passos

1.  Atualizar o Frontend:
    - O Login agora retorna { access_token, refresh_token }. Você precisará ajustar o armazenamento no localStorage e criar um interceptor
      no axios para chamar o /refresh quando der erro 401.
    - No cadastro, você precisa enviar o campo "termos_aceitos": true no JSON.
2.  Rodar Migrações:
    - Na próxima vez que subir o backend, o GORM criará automaticamente a tabela refresh_tokens e adicionará as novas colunas em usuarios.
