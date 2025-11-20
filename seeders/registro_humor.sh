#!/bin/bash

# Script para criar registros de humor dos pacientes
# 4 pacientes por profissional (20 pacientes / 5 profissionais)

API_BASE="http://localhost:8181/api/v1"

echo "🔗 Iniciando registros de humor dos pacientes..."
echo

# Arrays com emails
PROFISSIONAIS=(
    "joao.silva@mindtrace.com"
    "maria.santos@mindtrace.com"
    "pedro.oliveira@mindtrace.com"
    "ana.costa@mindtrace.com"
    "carlos.lima@mindtrace.com"
)

PACIENTES=(
    "ana.paciente@mindtrace.com"
    "bruno.paciente@mindtrace.com"
    "carla.paciente@mindtrace.com"
    "diego.paciente@mindtrace.com"
    "elena.paciente@mindtrace.com"
    "fabio.paciente@mindtrace.com"
    "gabriela.paciente@mindtrace.com"
    "henrique.paciente@mindtrace.com"
    "isabela.paciente@mindtrace.com"
    "joao.paciente@mindtrace.com"
    "karla.paciente@mindtrace.com"
    "lucas.paciente@mindtrace.com"
    "mariana.paciente@mindtrace.com"
    "nicolas.paciente@mindtrace.com"
    "olivia.paciente@mindtrace.com"
    "paulo.paciente@mindtrace.com"
    "rafaela.paciente@mindtrace.com"
    "sofia.paciente@mindtrace.com"
    "thiago.paciente@mindtrace.com"
    "vanessa.paciente@mindtrace.com"
)

# ============================================
# FUNÇÕES
# ============================================

# Função para fazer login com debug
login() {
    local email="$1"
    local senha="$2"
    
    local response=$(curl -s -w "\n%{http_code}" -X POST "$API_BASE/entrar/login" \
        -H "Content-Type: application/json" \
        -d "{
            \"email\": \"$email\",
            \"senha\": \"$senha\"
        }")
    
    # Separa o body do status code
    local http_code=$(echo "$response" | tail -n1)
    local body=$(echo "$response" | sed '$d')
    
    # Debug
    if [ "$http_code" != "200" ]; then
        echo "[DEBUG] HTTP $http_code: $body" >&2
    fi
    
    # Extrai o token
    local token=$(echo "$body" | jq -r '.dados.token // .token // empty' 2>/dev/null)
    echo "$token"
}

# Função para criar registro de humor
criar_registro_humor() {
    local token_aut="$1"
    local nivel_humor="$2"
    local horas_sono="$3"
    local nivel_stress="$4"
    local nivel_energia="$5"
    local auto_cuidado="$6"
    local observacoes="$7"
    local data_hora_registro="$8"
    
    local response=$(curl -s -w "\n%{http_code}" -X POST "$API_BASE/registro-humor/" \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $token_aut" \
        -d "{
            \"nivel_humor\": $nivel_humor,
            \"horas_sono\": $horas_sono,
            \"nivel_stress\": $nivel_stress,
            \"nivel_energia\": $nivel_energia,
            \"auto_cuidado\": \"$auto_cuidado\",
            \"observacoes\": \"$observacoes\",
            \"data_hora_registro\": \"$data_hora_registro\"
        }")
    
    local http_code=$(echo "$response" | tail -n1)
    local body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" != "201" ] && [ "$http_code" != "200" ]; then
        echo "[DEBUG] HTTP $http_code: $body" >&2
        return 1
    fi
    
    echo "$body" | jq -r '.mensagem // .message // empty' 2>/dev/null
}

# Função para gerar relatório
gerar_relatorio() {
    local token_aut="$1"
    local periodo="$2"
    
    local response=$(curl -s -w "\n%{http_code}" -X GET "$API_BASE/relatorios/?periodo=$periodo" \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $token_aut")
    
    local http_code=$(echo "$response" | tail -n1)
    local body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" != "200" ]; then
        echo "[DEBUG] HTTP $http_code: $body" >&2
        return 1
    fi
    
    echo "$body"
}

# ============================================
# LOOP PRINCIPAL
# ============================================

for i in {0..3}; do
    pac_email="${PACIENTES[$i]}"
    pac_nome=$(echo "$pac_email" | cut -d@ -f1 | sed 's/\./ /g')
    
    echo "👨‍⚕️  Paciente: $pac_nome ($pac_email)"
    
    # Fazer login do Paciente
    echo "  🔑 Fazendo login do Paciente..."
    pac_token=$(login "$pac_email" "Password123!")
    
    if [ -z "$pac_token" ]; then
        echo "  ❌ Erro ao fazer login do Paciente"
        continue
    fi
    
    echo "  ✅ Login realizado com sucesso"
    echo "     Token: ${pac_token:0:30}..."
    echo
    
    # Definir valores de humor baseados no índice
    nivel_humor=$((1))
    nivel_sono=$((1))
    nivel_stress=$((4))
    nivel_energia=$((5))
    
    # Criar registro de humor
    echo "  📝 Criando registro de humor..."
    criar_registro_humor \
        "$pac_token" \
        "$nivel_humor" \
        "$nivel_sono" \
        "$nivel_stress" \
        "$nivel_energia" \
        "Nenhum" \
        "Registro de humor diário de $pac_nome" \
        "$(date '+%Y-%m-%dT%H:%M:%SZ')"
    
    if [ $? -eq 0 ]; then
        echo "  ✅ Registro de humor criado com sucesso"
    else
        echo "  ❌ Erro ao criar registro de humor"
        continue
    fi
    echo
    
    # Gerar relatório
    echo "  📊 Gerando relatório..."
    relatorio=$(gerar_relatorio "$pac_token" "7")
    
    if [ $? -eq 0 ]; then
        echo "  ✅ Relatório gerado com sucesso"
        echo "$relatorio" | jq '.'
    else
        echo "  ❌ Erro ao gerar relatório"
    fi
    echo "---"
    echo

done

echo "✨ Processo finalizado!"