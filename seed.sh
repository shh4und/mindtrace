#!/bin/bash

# Script para popular o banco de dados MindTrace via API
# Executa requests HTTP para registrar profissionais e pacientes

API_BASE="http://localhost:8080/api/v1"

echo "🚀 Iniciando seeding via API do MindTrace..."
echo "API Base: $API_BASE"
echo

# Função para registrar profissional
register_professional() {
    local name="$1"
    local email="$2"
    local specialty="$3"
    local registro="$4"
    local cpf="$5"
    local contato="$6"

    echo "📝 Registrando profissional: $name"

    curl -s -X POST "$API_BASE/profissionais/registrar" \
        -H "Content-Type: application/json" \
        -d "{
            \"nome\": \"$name\",
            \"email\": \"$email\",
            \"senha\": \"Password123!\",
            \"especialidade\": \"$specialty\",
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

# Registrar 3 profissionais
register_professional "Dr. João Silva" "joao.silva@mindtrace.com" "Psicólogo" "CRP-01/12345" "12345678901" "(11) 99999-0001"
register_professional "Dra. Maria Santos" "maria.santos@mindtrace.com" "Psiquiatra" "CRM-02/67890" "12345678902" "(11) 99999-0002"
register_professional "Dr. Pedro Oliveira" "pedro.oliveira@mindtrace.com" "Terapeuta" "CRP-03/54321" "12345678903" "(11) 99999-0003"

echo
echo "👥 Registrando pacientes..."

# Registrar 15 pacientes
register_patient "Ana Costa" "paciente1@mindtrace.com" "12345678904" "false" "1995-03-15" "(11) 77777-0001" "" ""
register_patient "Bruno Lima" "paciente2@mindtrace.com" "12345678905" "true" "2000-07-22" "(11) 77777-0002" "Responsável de Bruno Lima" "(11) 88888-0001"
register_patient "Carla Rocha" "paciente3@mindtrace.com" "12345678906" "true" "1998-11-08" "(11) 77777-0003" "Responsável de Carla Rocha" "(11) 88888-0002"
register_patient "Diego Fernandes" "paciente4@mindtrace.com" "12345678907" "false" "1992-05-30" "(11) 77777-0004" "" ""
register_patient "Elena Gomes" "paciente5@mindtrace.com" "12345678908" "false" "1997-09-12" "(11) 77777-0005" "" ""
register_patient "Fabio Alves" "paciente6@mindtrace.com" "12345678909" "true" "2001-01-25" "(11) 77777-0006" "Responsável de Fabio Alves" "(11) 88888-0003"
register_patient "Gabriela Pereira" "paciente7@mindtrace.com" "12345678910" "false" "1994-12-03" "(11) 77777-0007" "" ""
register_patient "Henrique Souza" "paciente8@mindtrace.com" "12345678911" "true" "1999-06-18" "(11) 77777-0008" "Responsável de Henrique Souza" "(11) 88888-0004"
register_patient "Isabela Martins" "paciente9@mindtrace.com" "12345678912" "false" "1996-08-27" "(11) 77777-0009" "" ""
register_patient "João Carvalho" "paciente10@mindtrace.com" "12345678913" "true" "2002-04-10" "(11) 77777-0010" "Responsável de João Carvalho" "(11) 88888-0005"
register_patient "Karla Ribeiro" "paciente11@mindtrace.com" "12345678914" "false" "1993-10-05" "(11) 77777-0011" "" ""
register_patient "Lucas Teixeira" "paciente12@mindtrace.com" "12345678915" "true" "2000-02-14" "(11) 77777-0012" "Responsável de Lucas Teixeira" "(11) 88888-0006"
register_patient "Mariana Santos" "paciente13@mindtrace.com" "12345678916" "false" "1998-07-29" "(11) 77777-0013" "" ""
register_patient "Nicolas Oliveira" "paciente14@mindtrace.com" "12345678917" "true" "1999-11-21" "(11) 77777-0014" "Responsável de Nicolas Oliveira" "(11) 88888-0007"
register_patient "Olivia Silva" "paciente15@mindtrace.com" "12345678918" "false" "1995-12-08" "(11) 77777-0015" "" ""

echo
echo "✅ Seeding concluído!"
echo
echo "📊 Resumo dos dados criados:"
echo "👨‍⚕️ 3 Profissionais registrados:"
echo "   - joao.silva@mindtrace.com (Psicólogo)"
echo "   - maria.santos@mindtrace.com (Psiquiatra)"
echo "   - pedro.oliveira@mindtrace.com (Terapeuta)"
echo
echo "👥 15 Pacientes registrados:"
echo "   - paciente1@mindtrace.com até paciente15@mindtrace.com"
echo
echo "🔑 Senha padrão para todas as contas: Password123!"
echo
echo "💡 Para testar, inicie o servidor e faça login com qualquer conta acima."