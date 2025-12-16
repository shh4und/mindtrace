-- =============================================================================
-- SEEDER: Dados Mock para Desenvolvimento
-- Arquivo: seeds/dados_mock.sql
-- Data: 2025-11-26
-- Descrição: Dados mockados para popular o banco de dados em ambiente de dev
-- Senha padrão: Password123! (hash bcrypt incluído)
-- =============================================================================

BEGIN;

-- -----------------------------------------------------------------------------
-- 1. Inserção de Usuários (1 Profissional + 2 Pacientes)
-- Senha: Password123! -> $2a$10$JhWPWQ4Wy3at5lO3Rn7JmOHqC3jsoUEHsxy5X9HD.ZsQpHf2VHpCW
-- -----------------------------------------------------------------------------

-- Profissional (tipo_usuario = 2)
INSERT INTO usuarios (tipo_usuario, nome, email, senha, contato, bio, cpf, created_at, updated_at)
VALUES (
    2,
    'Dr. João Silva',
    'joao.silva@mindtrace.dev',
    '$2a$10$JhWPWQ4Wy3at5lO3Rn7JmOHqC3jsoUEHsxy5X9HD.ZsQpHf2VHpCW',
    '11999990001',
    'Psicólogo clínico com 10 anos de experiência em saúde mental.',
    '12345678901',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (email) DO NOTHING;

-- Paciente 1 (tipo_usuario = 3) - Adulto independente
INSERT INTO usuarios (tipo_usuario, nome, email, senha, contato, bio, cpf, created_at, updated_at)
VALUES (
    3,
    'Ana Costa',
    'ana.costa@mindtrace.dev',
    '$2a$10$JhWPWQ4Wy3at5lO3Rn7JmOHqC3jsoUEHsxy5X9HD.ZsQpHf2VHpCW',
    '11888880001',
    'Paciente em acompanhamento para ansiedade.',
    '98765432101',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (email) DO NOTHING;

-- Paciente 2 (tipo_usuario = 3) - Menor dependente
INSERT INTO usuarios (tipo_usuario, nome, email, senha, contato, bio, cpf, created_at, updated_at)
VALUES (
    3,
    'Bruno Lima',
    'bruno.lima@mindtrace.dev',
    '$2a$10$JhWPWQ4Wy3at5lO3Rn7JmOHqC3jsoUEHsxy5X9HD.ZsQpHf2VHpCW',
    '11777770001',
    'Paciente adolescente em acompanhamento.',
    '11122233344',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (email) DO NOTHING;


-- -----------------------------------------------------------------------------
-- 2. Inserção de Profissional
-- -----------------------------------------------------------------------------

INSERT INTO profissionais (usuario_id, data_nascimento, especialidade, registro_profissional, created_at, updated_at)
SELECT 
    id,
    '1985-04-15'::timestamp with time zone,
    'Psicologia Clínica',
    'CRP-06/12345',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM usuarios 
WHERE email = 'joao.silva@mindtrace.dev'
ON CONFLICT (usuario_id) DO NOTHING;


-- -----------------------------------------------------------------------------
-- 3. Inserção de Pacientes
-- -----------------------------------------------------------------------------

-- Paciente 1 - Ana (adulta, independente)
INSERT INTO pacientes (usuario_id, data_nascimento, dependente, nome_responsavel, contato_responsavel, data_inicio_tratamento, created_at, updated_at)
SELECT 
    id,
    '1995-03-20'::timestamp with time zone,
    FALSE,
    NULL,
    NULL,
    '2025-10-01'::timestamp with time zone,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM usuarios 
WHERE email = 'ana.costa@mindtrace.dev'
ON CONFLICT (usuario_id) DO NOTHING;

-- Paciente 2 - Bruno (menor, dependente)
INSERT INTO pacientes (usuario_id, data_nascimento, dependente, nome_responsavel, contato_responsavel, data_inicio_tratamento, created_at, updated_at)
SELECT 
    id,
    '2010-07-22'::timestamp with time zone,
    TRUE,
    'Maria Lima',
    '11666660001',
    '2025-10-15'::timestamp with time zone,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM usuarios 
WHERE email = 'bruno.lima@mindtrace.dev'
ON CONFLICT (usuario_id) DO NOTHING;


-- -----------------------------------------------------------------------------
-- 4. Vinculação Profissional-Paciente (Many-to-Many)
-- -----------------------------------------------------------------------------

INSERT INTO profissional_paciente (profissional_id, paciente_id)
SELECT p.id, pac.id
FROM profissionais p
CROSS JOIN pacientes pac
JOIN usuarios up ON p.usuario_id = up.id
JOIN usuarios upac ON pac.usuario_id = upac.id
WHERE up.email = 'joao.silva@mindtrace.dev'
  AND upac.email IN ('ana.costa@mindtrace.dev', 'bruno.lima@mindtrace.dev')
ON CONFLICT (profissional_id, paciente_id) DO NOTHING;


-- -----------------------------------------------------------------------------
-- 5. Inserção de Convites (1 usado, 1 ativo)
-- -----------------------------------------------------------------------------

-- Convite usado pela Ana
INSERT INTO convites (profissional_id, token, data_expiracao, usado, paciente_id, created_at, updated_at)
SELECT 
    p.id,
    'CONV-MOCK-USED-ANA-12345678',
    '2025-10-05'::timestamp with time zone,
    TRUE,
    pac.id,
    '2025-09-28'::timestamp with time zone,
    '2025-10-01'::timestamp with time zone
FROM profissionais p
JOIN usuarios up ON p.usuario_id = up.id
CROSS JOIN pacientes pac
JOIN usuarios upac ON pac.usuario_id = upac.id
WHERE up.email = 'joao.silva@mindtrace.dev'
  AND upac.email = 'ana.costa@mindtrace.dev'
ON CONFLICT (token) DO NOTHING;

-- Convite ativo (não usado, expira em 7 dias)
INSERT INTO convites (profissional_id, token, data_expiracao, usado, paciente_id, created_at, updated_at)
SELECT 
    p.id,
    'CONV-MOCK-ACTIVE-7891234567',
    (CURRENT_TIMESTAMP + INTERVAL '7 days'),
    FALSE,
    NULL,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM profissionais p
JOIN usuarios up ON p.usuario_id = up.id
WHERE up.email = 'joao.silva@mindtrace.dev'
ON CONFLICT (token) DO NOTHING;


-- -----------------------------------------------------------------------------
-- 6. Inserção de Registros de Humor (15 dias por paciente)
-- Valores variados para simular dados reais
-- -----------------------------------------------------------------------------

-- Registros de Humor - Ana Costa (15 dias)
INSERT INTO registros_humor (paciente_id, nivel_humor, horas_sono, nivel_energia, nivel_stress, auto_cuidado, observacoes, data_hora_registro, created_at)
SELECT 
    pac.id,
    vals.nivel_humor,
    vals.horas_sono,
    vals.nivel_energia,
    vals.nivel_stress,
    vals.auto_cuidado::jsonb,
    vals.observacoes,
    vals.data_hora_registro,
    vals.data_hora_registro
FROM pacientes pac
JOIN usuarios u ON pac.usuario_id = u.id
CROSS JOIN (
    VALUES
        (4, 7, 7, 3, '["Meditação"]', 'Dia produtivo no trabalho', (CURRENT_TIMESTAMP - INTERVAL '14 days')),
        (3, 6, 6, 5, '["Caminhada"]', 'Reunião estressante', (CURRENT_TIMESTAMP - INTERVAL '13 days')),
        (4, 8, 8, 2, '["Yoga", "Leitura"]', 'Final de semana relaxante', (CURRENT_TIMESTAMP - INTERVAL '12 days')),
        (5, 8, 7, 2, '["Academia", "Banho relaxante"]', 'Me senti muito bem hoje', (CURRENT_TIMESTAMP - INTERVAL '11 days')),
        (3, 5, 5, 6, '["Assistir série"]', 'Insônia leve', (CURRENT_TIMESTAMP - INTERVAL '10 days')),
        (2, 4, 4, 7, '[]', 'Dia difícil, muita ansiedade', (CURRENT_TIMESTAMP - INTERVAL '9 days')),
        (3, 6, 6, 5, '["Conversa com amigos"]', 'Melhorando aos poucos', (CURRENT_TIMESTAMP - INTERVAL '8 days')),
        (4, 7, 7, 4, '["Meditação guiada"]', 'Sessão de terapia hoje', (CURRENT_TIMESTAMP - INTERVAL '7 days')),
        (4, 7, 8, 3, '["Cozinhar"]', 'Dia tranquilo', (CURRENT_TIMESTAMP - INTERVAL '6 days')),
        (5, 8, 8, 2, '["Passeio no parque"]', 'Ótimo humor!', (CURRENT_TIMESTAMP - INTERVAL '5 days')),
        (4, 7, 7, 3, '["Leitura"]', 'Rotina estável', (CURRENT_TIMESTAMP - INTERVAL '4 days')),
        (3, 6, 5, 5, '["Música"]', 'Um pouco cansada', (CURRENT_TIMESTAMP - INTERVAL '3 days')),
        (4, 7, 7, 3, '["Alongamento"]', 'Bom dia de trabalho', (CURRENT_TIMESTAMP - INTERVAL '2 days')),
        (4, 8, 8, 2, '["Skincare"]', 'Me sentindo bem', (CURRENT_TIMESTAMP - INTERVAL '1 day')),
        (5, 8, 7, 2, '["Academia"]', 'Excelente dia!', CURRENT_TIMESTAMP)
) AS vals(nivel_humor, horas_sono, nivel_energia, nivel_stress, auto_cuidado, observacoes, data_hora_registro)
WHERE u.email = 'ana.costa@mindtrace.dev'
ON CONFLICT (paciente_id, nivel_humor, horas_sono, nivel_energia, nivel_stress, auto_cuidado, observacoes) DO NOTHING;

-- Registros de Humor - Bruno Lima (15 dias)
INSERT INTO registros_humor (paciente_id, nivel_humor, horas_sono, nivel_energia, nivel_stress, auto_cuidado, observacoes, data_hora_registro, created_at)
SELECT 
    pac.id,
    vals.nivel_humor,
    vals.horas_sono,
    vals.nivel_energia,
    vals.nivel_stress,
    vals.auto_cuidado::jsonb,
    vals.observacoes,
    vals.data_hora_registro,
    vals.data_hora_registro
FROM pacientes pac
JOIN usuarios u ON pac.usuario_id = u.id
CROSS JOIN (
    VALUES
        (3, 8, 6, 4, '["Videogame"]', 'Prova difícil na escola', (CURRENT_TIMESTAMP - INTERVAL '14 days')),
        (4, 9, 7, 3, '["Futebol"]', 'Dia legal na escola', (CURRENT_TIMESTAMP - INTERVAL '13 days')),
        (4, 10, 8, 2, '["Dormir"]', 'Sábado tranquilo', (CURRENT_TIMESTAMP - INTERVAL '12 days')),
        (5, 9, 8, 2, '["Passear com cachorro"]', 'Domingo em família', (CURRENT_TIMESTAMP - INTERVAL '11 days')),
        (3, 7, 5, 5, '["Assistir vídeos"]', 'Muita lição de casa', (CURRENT_TIMESTAMP - INTERVAL '10 days')),
        (2, 6, 4, 7, '[]', 'Briga com colega', (CURRENT_TIMESTAMP - INTERVAL '9 days')),
        (3, 7, 7, 5, '["Conversar com a mãe"]', 'Conversei sobre o problema', (CURRENT_TIMESTAMP - INTERVAL '8 days')),
        (4, 8, 8, 3, '["Desenhar"]', 'Resolvido o problema', (CURRENT_TIMESTAMP - INTERVAL '7 days')),
        (4, 8, 9, 3, '["Ler mangá"]', 'Dia normal', (CURRENT_TIMESTAMP - INTERVAL '6 days')),
        (5, 9, 9, 2, '["Jogar bola"]', 'Feriado divertido', (CURRENT_TIMESTAMP - INTERVAL '5 days')),
        (4, 8, 8, 3, '["Ouvir música"]', 'Voltando à rotina', (CURRENT_TIMESTAMP - INTERVAL '4 days')),
        (3, 7, 6, 5, '["Ver série"]', 'Cansado da escola', (CURRENT_TIMESTAMP - INTERVAL '3 days')),
        (4, 8, 8, 3, '["Jogar online"]', 'Dia bom', (CURRENT_TIMESTAMP - INTERVAL '2 days')),
        (4, 8, 9, 2, '["Descansar"]', 'Bem disposto', (CURRENT_TIMESTAMP - INTERVAL '1 day')),
        (5, 9, 8, 2, '["Esporte"]', 'Ótimo dia!', CURRENT_TIMESTAMP)
) AS vals(nivel_humor, horas_sono, nivel_energia, nivel_stress, auto_cuidado, observacoes, data_hora_registro)
WHERE u.email = 'bruno.lima@mindtrace.dev'
ON CONFLICT (paciente_id, nivel_humor, horas_sono, nivel_energia, nivel_stress, auto_cuidado, observacoes) DO NOTHING;


COMMIT;

-- =============================================================================
-- Resumo dos dados mockados:
-- - 1 Profissional: Dr. João Silva (joao.silva@mindtrace.dev)
-- - 2 Pacientes: Ana Costa e Bruno Lima
-- - Senha padrão para todos: Password123!
-- - Vínculo profissional-paciente estabelecido
-- - 2 Convites: 1 usado (Ana), 1 ativo
-- - 30 Registros de humor (15 por paciente, últimos 15 dias)
-- =============================================================================
