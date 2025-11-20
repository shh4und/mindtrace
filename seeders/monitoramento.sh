#!/bin/bash

# Script para testar a rota de monitoramento
# Realiza monitoramento para os 4 primeiros pacientes do profissional joao.silva@mindtrace.com

API_BASE="http://localhost:8181/api/v1"

echo "🔗 Iniciando testes de monitoramento..."
echo

# Email e senha do profissional
PROF_EMAIL="joao.silva@mindtrace.com"
PROF_SENHA="Password123!"

# Pacientes associados ao profissional (primeiros 4)
PACIENTES=(
    "ana.paciente@mindtrace.com"
    "bruno.paciente@mindtrace.com"
    "carla.paciente@mindtrace.com"
    "diego.paciente@mindtrace.com"
)

# IDs dos pacientes (devem corresponder aos emails acima)
PACIENTE_IDS=(
    1
    2
    3
    4
)

# ============================================
# FUNÇÕES
# ============================================

# Função para fazer login
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

# Função para realizar monitoramento
realizar_monitoramento() {
    local token_aut="$1"
    local paciente_id="$2"
    local num_registros="$3"
    
    local response=$(curl -s -w "\n%{http_code}" -X GET "$API_BASE/monitoramento/paciente-lista?pacienteID=$paciente_id&numRegistros=$num_registros" \
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
# FLUXO PRINCIPAL
# ============================================

# 1. Fazer login do profissional
echo "🔑 Fazendo login do profissional: $PROF_EMAIL"
prof_token=$(login "$PROF_EMAIL" "$PROF_SENHA")

if [ -z "$prof_token" ]; then
    echo "❌ Erro ao fazer login do profissional"
    exit 1
fi

echo "✅ Login realizado com sucesso"
echo "   Token: ${prof_token:0:30}..."
echo

# 2. Realizar monitoramento para cada paciente
for i in {0..3}; do
    pac_email="${PACIENTES[$i]}"
    pac_id="${PACIENTE_IDS[$i]}"
    
    echo "📊 Monitoramento do Paciente: $pac_email (ID: $pac_id)"
    
    # Realizar monitoramento com 4 últimos registros
    monitoramento=$(realizar_monitoramento "$prof_token" "$pac_id" "4")
    
     if [ $? -eq 0 ]; then
        echo "✅ Monitoramento realizado com sucesso"
        echo
        echo "   Dados do Monitoramento:"
        echo "$monitoramento" | jq '.'
        echo
        
        # ✅ Verifica se há dados antes de tentar iterar
        total_registros=$(echo "$monitoramento" | jq '.dados_monitoramento | length // 0')
        
        if [ "$total_registros" -gt 0 ]; then
            echo "   Detalhes dos Registros:"
            echo "$monitoramento" | jq -r '.dados_monitoramento[] | 
                "     📅 \(.data) | 😊 Humor: \(.nivel_humor) | 😴 Sono: \(.horas_sono)h | ⚡ Energia: \(.nivel_energia) | 😰 Stress: \(.nivel_stress)"'
        else
            echo "   ⚠️  Nenhum registro de humor encontrado para este paciente"
        fi
        echo
    else
        echo "❌ Erro ao realizar monitoramento"
    fi
    
    echo "---"
    echo
done

echo "✨ Testes de monitoramento finalizados!"