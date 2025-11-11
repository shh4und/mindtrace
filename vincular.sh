#!/bin/bash

# Script para vincular pacientes aos profissionais
# 4 pacientes por profissional (20 pacientes / 5 profissionais)

API_BASE="http://localhost:8080/api/v1"

echo "🔗 Iniciando vinculação de pacientes aos profissionais..."
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
    
    # Tenta extrair o token de diferentes formatos de resposta
    local token=$(echo "$body" | jq -r '.dados.token // .token // empty' 2>/dev/null)
    
    echo "$token"
}

# Função para gerar convite com debug
gerar_convite() {
    local token="$1"
    local paciente_email="$2"
    
    local response=$(curl -s -w "\n%{http_code}" -X POST "$API_BASE/convites/gerar" \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $token" \
        -d "{
            \"email_paciente\": \"$paciente_email\"
        }")
    
    local http_code=$(echo "$response" | tail -n1)
    local body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" != "200" ]; then
        echo "[DEBUG] HTTP $http_code: $body" >&2
    fi
    
    echo "$body" | jq -r '.dados.token // .token // empty' 2>/dev/null
}

# Função para vincular paciente com debug
vincular_paciente() {
    local token_paciente="$1"
    local token_convite="$2"
    
    local response=$(curl -s -w "\n%{http_code}" -X POST "$API_BASE/convites/vincular" \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $token_paciente" \
        -d "{
            \"token\": \"$token_convite\"
        }")
    
    local http_code=$(echo "$response" | tail -n1)
    local body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" != "200" ]; then
        echo "[DEBUG] HTTP $http_code: $body" >&2
    fi
    
    echo "$body" | jq -r '.mensagem // .message // empty' 2>/dev/null
}

# Variável para controlar qual paciente vincular (4 por profissional)
paciente_index=0

# Loop pelos profissionais
for i in "${!PROFISSIONAIS[@]}"; do
    prof_email="${PROFISSIONAIS[$i]}"
    prof_nome=$(echo "$prof_email" | cut -d@ -f1 | sed 's/\./ /g')
    
    echo "👨‍⚕️  Profissional: $prof_nome ($prof_email)"
    
    # Fazer login do profissional
    echo "  🔑 Fazendo login do profissional..."
    prof_token=$(login "$prof_email" "Password123!")
    
    if [ -z "$prof_token" ]; then
        echo "  ❌ Erro ao fazer login do profissional"
        continue
    fi
    echo "  ✅ Login realizado com sucesso"
    echo "     Token: ${prof_token:0:30}..."
    
    # Vincular 4 pacientes a este profissional
    for j in {1..4}; do
        if [ $paciente_index -ge ${#PACIENTES[@]} ]; then
            break
        fi
        
        paciente_email="${PACIENTES[$paciente_index]}"
        paciente_nome=$(echo "$paciente_email" | cut -d@ -f1 | sed 's/\./ /g')
        
        echo "  👤 Paciente $j/4: $paciente_nome"
        
        # Gerar convite para o paciente
        echo "    📨 Gerando convite..."
        token_convite=$(gerar_convite "$prof_token" "$paciente_email")
        
        if [ -z "$token_convite" ]; then
            echo "    ❌ Erro ao gerar convite"
            paciente_index=$((paciente_index + 1))
            continue
        fi
        echo "    ✅ Convite gerado: ${token_convite:0:20}..."
        
        # Fazer login do paciente
        echo "    🔑 Fazendo login do paciente..."
        paciente_token=$(login "$paciente_email" "Password123!")
        
        if [ -z "$paciente_token" ]; then
            echo "    ❌ Erro ao fazer login do paciente"
            paciente_index=$((paciente_index + 1))
            continue
        fi
        echo "    ✅ Login do paciente realizado"
        
        # Vincular paciente ao profissional
        echo "    🔗 Vinculando paciente..."
        resultado=$(vincular_paciente "$paciente_token" "$token_convite")
        
        if [[ "$resultado" == *"sucesso"* ]] || [[ "$resultado" == *"vinculad"* ]]; then
            echo "    ✅ Paciente vinculado com sucesso!"
        else
            echo "    ⚠️  Resposta: $resultado"
        fi
        
        paciente_index=$((paciente_index + 1))
        echo
    done
done

echo
echo "✅ Processo de vinculação concluído!"
echo
echo "📊 Resumo:"
echo "   - 5 Profissionais"
echo "   - 20 Pacientes"
echo "   - 4 Pacientes por Profissional"
echo "   - Total: 20 Vinculações"