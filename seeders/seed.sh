#!/bin/bash

# Script para popular o banco de dados MindTrace via API
# Executa requests HTTP para registrar profissionais e pacientes

API_BASE="http://localhost:8181/api/v1"

echo "🚀 Iniciando seeding via API do MindTrace..."
echo "API Base: $API_BASE"
echo

# Função para registrar profissional
register_professional() {
    local name="$1"
    local email="$2"
    local especialidade="$3"
    local registro="$4"
    local cpf="$5"
    local contato="$6"
    local data_nascimento="$7"

    echo "📝 Registrando profissional: $name"

    curl -s -X POST "$API_BASE/profissionais/registrar" \
        -H "Content-Type: application/json" \
        -d "{
            \"nome\": \"$name\",
            \"email\": \"$email\",
            \"senha\": \"Password123!\",
            \"data_nascimento\": \"$data_nascimento\",
            \"especialidade\": \"$especialidade\",
            \"registro_profissional\": \"$registro\",
            \"cpf\": \"$cpf\",
            \"contato\": \"$contato\"
        }" | jq -r '.mensagem // .erro' 2>/dev/null || echo "Request enviado"
}

# Função para registrar paciente
register_patient() {
    local name="$1"
    local email="$2"
    local cpf="$3"
    local dependente="$4"
    local data_nascimento="$5"
    local contato="$6"
    local nome_responsavel="$7"
    local contato_responsavel="$8"

    echo "📝 Registrando paciente: $name"

    local data="{
        \"nome\": \"$name\",
        \"email\": \"$email\",
        \"senha\": \"Password123!\",
        \"cpf\": \"$cpf\",
        \"dependente\": $dependente,
        \"data_nascimento\": \"$data_nascimento\",
        \"contato\": \"$contato\""

    if [ "$dependente" = "true" ] && [ -n "$nome_responsavel" ]; then
        data="$data,
        \"nome_responsavel\": \"$nome_responsavel\",
        \"contato_responsavel\": \"$contato_responsavel\""
    fi

    data="$data
    }"

    curl -s -X POST "$API_BASE/pacientes/registrar" \
        -H "Content-Type: application/json" \
        -d "$data" | jq -r '.mensagem // .erro' 2>/dev/null || echo "Request enviado"
}

echo "👨‍⚕️ Registrando profissionais..."
echo

# Registrar 5 profissionais
register_professional "Dr. João Silva" "joao.silva@mindtrace.com" "Psicologia Clínica" "CRP-01/12345" "12345678901" "11999990001" "1985-04-15T00:00:00Z"
register_professional "Dra. Maria Santos" "maria.santos@mindtrace.com" "Psiquiatria" "CRM-02/67890" "12345678902" "11999990002" "1980-08-22T00:00:00Z"
register_professional "Dr. Pedro Oliveira" "pedro.oliveira@mindtrace.com" "Terapia Cognitivo-Comportamental" "CRP-03/54321" "12345678903" "11999990003" "1988-12-10T00:00:00Z"
register_professional "Dra. Ana Paula Costa" "ana.costa@mindtrace.com" "Psicologia Infantil" "CRP-04/98765" "12345678920" "11999990004" "1990-03-28T00:00:00Z"
register_professional "Dr. Carlos Eduardo Lima" "carlos.lima@mindtrace.com" "Neuropsicologia" "CRP-05/11111" "12345678921" "11999990005" "1982-07-05T00:00:00Z"

echo
echo "👥 Registrando pacientes..."
echo

# Registrar 20 pacientes (mix de dependentes e independentes)
register_patient "Ana Costa" "ana.paciente@mindtrace.com" "12345678904" "false" "1995-03-15T00:00:00Z" "11777770001" "" ""
register_patient "Bruno Lima" "bruno.paciente@mindtrace.com" "12345678905" "true" "2010-07-22T00:00:00Z" "11777770002" "Maria Lima" "11888880001"
register_patient "Carla Rocha" "carla.paciente@mindtrace.com" "12345678906" "false" "1990-11-08T00:00:00Z" "11777770003" "" ""
register_patient "Diego Fernandes" "diego.paciente@mindtrace.com" "12345678907" "false" "1992-05-30T00:00:00Z" "11777770004" "" ""
register_patient "Elena Gomes" "elena.paciente@mindtrace.com" "12345678908" "false" "1997-09-12T00:00:00Z" "11777770005" "" ""
register_patient "Fabio Alves" "fabio.paciente@mindtrace.com" "12345678909" "true" "2012-01-25T00:00:00Z" "11777770006" "Roberto Alves" "11888880002"
register_patient "Gabriela Pereira" "gabriela.paciente@mindtrace.com" "12345678910" "false" "1994-12-03T00:00:00Z" "11777770007" "" ""
register_patient "Henrique Souza" "henrique.paciente@mindtrace.com" "12345678911" "true" "2011-06-18T00:00:00Z" "11777770008" "Patricia Souza" "11888880003"
register_patient "Isabela Martins" "isabela.paciente@mindtrace.com" "12345678912" "false" "1996-08-27T00:00:00Z" "11777770009" "" ""
register_patient "João Carvalho" "joao.paciente@mindtrace.com" "12345678913" "true" "2013-04-10T00:00:00Z" "11777770010" "Fernanda Carvalho" "11888880004"
register_patient "Karla Ribeiro" "karla.paciente@mindtrace.com" "12345678914" "false" "1993-10-05T00:00:00Z" "11777770011" "" ""
register_patient "Lucas Teixeira" "lucas.paciente@mindtrace.com" "12345678915" "false" "1998-02-14T00:00:00Z" "11777770012" "" ""
register_patient "Mariana Santos" "mariana.paciente@mindtrace.com" "12345678916" "false" "1991-07-29T00:00:00Z" "11777770013" "" ""
register_patient "Nicolas Oliveira" "nicolas.paciente@mindtrace.com" "12345678917" "true" "2014-11-21T00:00:00Z" "11777770014" "Sandra Oliveira" "11888880005"
register_patient "Olivia Silva" "olivia.paciente@mindtrace.com" "12345678918" "false" "1995-12-08T00:00:00Z" "11777770015" "" ""
register_patient "Paulo Mendes" "paulo.paciente@mindtrace.com" "12345678922" "false" "1989-03-17T00:00:00Z" "11777770016" "" ""
register_patient "Rafaela Nunes" "rafaela.paciente@mindtrace.com" "12345678923" "false" "1996-05-23T00:00:00Z" "11777770017" "" ""
register_patient "Sofia Castro" "sofia.paciente@mindtrace.com" "12345678924" "true" "2015-09-14T00:00:00Z" "11777770018" "Amanda Castro" "11888880006"
register_patient "Thiago Barbosa" "thiago.paciente@mindtrace.com" "12345678925" "false" "1993-11-30T00:00:00Z" "11777770019" "" ""
register_patient "Vanessa Dias" "vanessa.paciente@mindtrace.com" "12345678926" "false" "1994-06-19T00:00:00Z" "11777770020" "" ""

echo
echo "✅ Seeding concluído!"
echo
echo "📊 Resumo dos dados criados:"
echo
echo "👨‍⚕️ 5 Profissionais registrados:"
echo "   - joao.silva@mindtrace.com (Psicologia Clínica)"
echo "   - maria.santos@mindtrace.com (Psiquiatria)"
echo "   - pedro.oliveira@mindtrace.com (Terapia Cognitivo-Comportamental)"
echo "   - ana.costa@mindtrace.com (Psicologia Infantil)"
echo "   - carlos.lima@mindtrace.com (Neuropsicologia)"
echo
echo "👥 20 Pacientes registrados:"
echo "   - 14 pacientes independentes"
echo "   - 6 pacientes dependentes (menores de idade com responsáveis)"
echo
echo "📧 Exemplos de contas de pacientes:"
echo "   - ana.paciente@mindtrace.com (independente, 29 anos)"
echo "   - bruno.paciente@mindtrace.com (dependente, 15 anos)"
echo "   - carla.paciente@mindtrace.com (independente, 35 anos)"
echo "   - diego.paciente@mindtrace.com (independente, 33 anos)"
echo "   - elena.paciente@mindtrace.com (independente, 28 anos)"
echo
echo "🔑 Senha padrão para todas as contas: Password123!"
echo
echo "💡 Próximos passos:"
echo "   1. Faça login como profissional (ex: joao.silva@mindtrace.com)"
echo "   2. Gere convites para vincular pacientes"
echo "   3. Use os convites para associar pacientes aos profissionais"
echo "   4. Registre dados de humor dos pacientes para testar monitoramento"
echo
echo "🔗 Para vincular pacientes, use a API:"
echo "   POST /api/v1/convites/gerar (como profissional)"
echo "   POST /api/v1/convites/vincular (como paciente, com o token)"