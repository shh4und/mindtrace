-- =============================================================================
-- SEEDER: Instrumentos Padronizados de Saúde Mental
-- Arquivo: seeders/instrumentos_padrao.sql
-- Data: 2025-11-21
-- Autor: Gerado via Gemini para MindTrace
-- =============================================================================

BEGIN;

-- -----------------------------------------------------------------------------
-- 1. Inserção dos Instrumentos (Metadados)
-- -----------------------------------------------------------------------------

INSERT INTO instrumentos (codigo, nome, descricao, algoritmo_pontuacao, versao, esta_ativo)
VALUES
(
    'whoqol_bref', 
    'WHOQOL-BREF', 
    'Instrumento abreviado de avaliação da qualidade de vida da Organização Mundial da Saúde. 26 questões divididas em 4 domínios: Físico, Psicológico, Relações Sociais e Meio Ambiente.', 
    'whoqol_bref', 
    1, 
    TRUE
),
(
    'gad_7', 
    'GAD-7', 
    'Escala de Transtorno de Ansiedade Generalizada. Ferramenta de rastreio e avaliação de gravidade de sintomas ansiosos.', 
    'gad_7', 
    1, 
    TRUE
),
(
    'phq_9', 
    'PHQ-9', 
    'Questionário sobre a Saúde do Paciente. Instrumento padrão para rastreio, diagnóstico e monitorização da gravidade da depressão.', 
    'phq_9', 
    1, 
    TRUE
),
(
    'who_5', 
    'WHO-5', 
    'Índice de Bem-Estar (5 itens). Escala curta para avaliação do bem-estar subjetivo positivo.', 
    'who_5', 
    1, 
    TRUE
)
ON CONFLICT (codigo) DO NOTHING;


-- -----------------------------------------------------------------------------
-- 2. Inserção das Opções de Resposta (Escalas Likert)
-- -----------------------------------------------------------------------------

-- WHOQOL-BREF (Escala 1-5)
INSERT INTO opcoes_escala (instrumento_id, valor, rotulo)
SELECT id, 1, 'Nada / Muito ruim / Muito insatisfeito' FROM instrumentos WHERE codigo = 'whoqol_bref'
UNION ALL
SELECT id, 2, 'Muito pouco / Ruim / Insatisfeito' FROM instrumentos WHERE codigo = 'whoqol_bref'
UNION ALL
SELECT id, 3, 'Médio / Nem ruim nem bom / Nem satisfeito nem insatisfeito' FROM instrumentos WHERE codigo = 'whoqol_bref'
UNION ALL
SELECT id, 4, 'Muito / Bom / Satisfeito' FROM instrumentos WHERE codigo = 'whoqol_bref'
UNION ALL
SELECT id, 5, 'Completamente / Muito bom / Muito satisfeito' FROM instrumentos WHERE codigo = 'whoqol_bref'
ON CONFLICT (instrumento_id, valor, rotulo) DO NOTHING;

-- GAD-7 (Escala 0-3)
INSERT INTO opcoes_escala (instrumento_id, valor, rotulo)
SELECT id, 0, 'Nenhuma vez' FROM instrumentos WHERE codigo = 'gad_7'
UNION ALL
SELECT id, 1, 'Vários dias' FROM instrumentos WHERE codigo = 'gad_7'
UNION ALL
SELECT id, 2, 'Mais da metade dos dias' FROM instrumentos WHERE codigo = 'gad_7'
UNION ALL
SELECT id, 3, 'Quase todos os dias' FROM instrumentos WHERE codigo = 'gad_7'
ON CONFLICT (instrumento_id, valor, rotulo) DO NOTHING;

-- PHQ-9 (Escala 0-3)
INSERT INTO opcoes_escala (instrumento_id, valor, rotulo)
SELECT id, 0, 'Nenhuma vez' FROM instrumentos WHERE codigo = 'phq_9'
UNION ALL
SELECT id, 1, 'Vários dias' FROM instrumentos WHERE codigo = 'phq_9'
UNION ALL
SELECT id, 2, 'Mais da metade dos dias' FROM instrumentos WHERE codigo = 'phq_9'
UNION ALL
SELECT id, 3, 'Quase todos os dias' FROM instrumentos WHERE codigo = 'phq_9'
ON CONFLICT (instrumento_id, valor, rotulo) DO NOTHING;

-- WHO-5 (Escala 0-5, onde 5 é o melhor)
INSERT INTO opcoes_escala (instrumento_id, valor, rotulo)
SELECT id, 5, 'Todo o tempo' FROM instrumentos WHERE codigo = 'who_5'
UNION ALL
SELECT id, 4, 'A maior parte do tempo' FROM instrumentos WHERE codigo = 'who_5'
UNION ALL
SELECT id, 3, 'Mais de metade do tempo' FROM instrumentos WHERE codigo = 'who_5'
UNION ALL
SELECT id, 2, 'Menos de metade do tempo' FROM instrumentos WHERE codigo = 'who_5'
UNION ALL
SELECT id, 1, 'Alguma parte do tempo' FROM instrumentos WHERE codigo = 'who_5'
UNION ALL
SELECT id, 0, 'Nunca / Nenhuma vez' FROM instrumentos WHERE codigo = 'who_5'
ON CONFLICT (instrumento_id, valor, rotulo) DO NOTHING;


-- -----------------------------------------------------------------------------
-- 3. Inserção das Perguntas (Itens)
-- -----------------------------------------------------------------------------

-- ================= WHOQOL-BREF =================

INSERT INTO perguntas (instrumento_id, ordem_item, dominio, conteudo, eh_pontuacao_invertida)
SELECT id, 1, 'Geral', 'Como você avaliaria sua qualidade de vida?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 2, 'Geral', 'Quão satisfeito(a) você está com a sua saúde?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 3, 'Físico', 'Em que medida você acha que sua dor (física) impede você de fazer o que você precisa?', TRUE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 4, 'Físico', 'O quanto você precisa de algum tratamento médico para levar sua vida diária?', TRUE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 5, 'Psicológico', 'O quanto você aproveita a vida?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 6, 'Psicológico', 'Em que medida você acha que a sua vida tem sentido?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 7, 'Psicológico', 'O quanto você consegue se concentrar?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 8, 'Meio Ambiente', 'Quão seguro(a) você se sente em sua vida diária?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 9, 'Meio Ambiente', 'Quão saudável é o seu ambiente fisico (clima, barulho, poluição, atrativos)?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 10, 'Físico', 'Você tem energia suficiente para seu dia-a-dia?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 11, 'Psicológico', 'Você é capaz de aceitar sua aparência fisica?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 12, 'Meio Ambiente', 'Você tem dinheiro suficiente para satisfazer suas necessidades?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 13, 'Meio Ambiente', 'Quão disponíveis para você estão as informações que precisa no seu dia-a-dia?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 14, 'Meio Ambiente', 'Em que medida você tem oportunidades de atividade de lazer?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 15, 'Físico', 'Quão bem você é capaz de se locomover?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 16, 'Físico', 'Quão satisfeito(a) você está com o seu sono?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 17, 'Físico', 'Quão satisfeito(a) você está com sua capacidade de desempenhar as atividades do seu dia-a-dia?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 18, 'Físico', 'Quão satisfeito(a) você está com sua capacidade para o trabalho?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 19, 'Psicológico', 'Quão satisfeito(a) você está consigo mesmo?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 20, 'Relações Sociais', 'Quão satisfeito(a) você está com suas relações pessoais (amigos, parentes, conhecidos, colegas)?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 21, 'Relações Sociais', 'Quão satisfeito(a) você está com sua vida sexual?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 22, 'Relações Sociais', 'Quão satisfeito(a) você está com o apoio que você recebe de seus amigos?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 23, 'Meio Ambiente', 'Quão satisfeito(a) você está com as condições do local onde mora?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 24, 'Meio Ambiente', 'Quão satisfeito(a) você está com o seu acesso aos serviços de saúde?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 25, 'Meio Ambiente', 'Quão satisfeito(a) você está com o seu meio de transporte?', FALSE FROM instrumentos WHERE codigo = 'whoqol_bref' UNION ALL
SELECT id, 26, 'Psicológico', 'Com que frequência você tem sentimentos negativos tais como mau humor, desespero, ansiedade, depressão?', TRUE FROM instrumentos WHERE codigo = 'whoqol_bref'
ON CONFLICT (instrumento_id, ordem_item, dominio, conteudo) DO NOTHING;

-- ================= GAD-7 =================

INSERT INTO perguntas (instrumento_id, ordem_item, conteudo, eh_pontuacao_invertida)
SELECT id, 1, 'Sentir-se nervoso, ansioso ou no limite', FALSE FROM instrumentos WHERE codigo = 'gad_7' UNION ALL
SELECT id, 2, 'Não ser capaz de parar ou controlar as preocupações', FALSE FROM instrumentos WHERE codigo = 'gad_7' UNION ALL
SELECT id, 3, 'Preocupar-se muito com diversas coisas', FALSE FROM instrumentos WHERE codigo = 'gad_7' UNION ALL
SELECT id, 4, 'Dificuldade para relaxar', FALSE FROM instrumentos WHERE codigo = 'gad_7' UNION ALL
SELECT id, 5, 'Ser tão inquieto que se torna difícil permanecer sentado', FALSE FROM instrumentos WHERE codigo = 'gad_7' UNION ALL
SELECT id, 6, 'Ficar facilmente irritado ou irritável', FALSE FROM instrumentos WHERE codigo = 'gad_7' UNION ALL
SELECT id, 7, 'Sentir medo como se algo horrível fosse acontecer', FALSE FROM instrumentos WHERE codigo = 'gad_7'
ON CONFLICT (instrumento_id, ordem_item, dominio, conteudo) DO NOTHING;

-- ================= PHQ-9 =================

INSERT INTO perguntas (instrumento_id, ordem_item, conteudo, eh_pontuacao_invertida)
SELECT id, 1, 'Pouco interesse ou pouco prazer em fazer as coisas', FALSE FROM instrumentos WHERE codigo = 'phq_9' UNION ALL
SELECT id, 2, 'Se sentir "para baixo", deprimido/a ou sem perspectiva', FALSE FROM instrumentos WHERE codigo = 'phq_9' UNION ALL
SELECT id, 3, 'Dificuldade para pegar no sono ou permanecer dormindo, ou dormir mais do que de costume', FALSE FROM instrumentos WHERE codigo = 'phq_9' UNION ALL
SELECT id, 4, 'Se sentir cansado/a ou com pouca energia', FALSE FROM instrumentos WHERE codigo = 'phq_9' UNION ALL
SELECT id, 5, 'Falta de apetite ou comendo demais', FALSE FROM instrumentos WHERE codigo = 'phq_9' UNION ALL
SELECT id, 6, 'Se sentir mal consigo mesmo/a ou achar que você é um fracasso ou que decepcionou sua família ou você mesmo/a', FALSE FROM instrumentos WHERE codigo = 'phq_9' UNION ALL
SELECT id, 7, 'Dificuldade para se concentrar nas coisas, como ler o jornal ou ver televisão', FALSE FROM instrumentos WHERE codigo = 'phq_9' UNION ALL
SELECT id, 8, 'Lentidão para se movimentar ou falar, a ponto das outras pessoas perceberem? Ou o oposto — estar tão agitado/a ou irrequieto/a', FALSE FROM instrumentos WHERE codigo = 'phq_9' UNION ALL
SELECT id, 9, 'Pensar em se ferir de alguma maneira ou que seria melhor estar morto/a', FALSE FROM instrumentos WHERE codigo = 'phq_9'
ON CONFLICT (instrumento_id, ordem_item, dominio, conteudo) DO NOTHING;

-- ================= WHO-5 =================

INSERT INTO perguntas (instrumento_id, ordem_item, conteudo, eh_pontuacao_invertida)
SELECT id, 1, 'Senti-me alegre e bem-disposto', FALSE FROM instrumentos WHERE codigo = 'who_5' UNION ALL
SELECT id, 2, 'Senti-me calmo e relaxado', FALSE FROM instrumentos WHERE codigo = 'who_5' UNION ALL
SELECT id, 3, 'Senti-me ativo e vigoroso', FALSE FROM instrumentos WHERE codigo = 'who_5' UNION ALL
SELECT id, 4, 'Acordei a sentir-me fresco e descansado', FALSE FROM instrumentos WHERE codigo = 'who_5' UNION ALL
SELECT id, 5, 'A minha vida diária tem sido preenchida por coisas que me interessam', FALSE FROM instrumentos WHERE codigo = 'who_5'
ON CONFLICT (instrumento_id, ordem_item, dominio, conteudo) DO NOTHING;

COMMIT;