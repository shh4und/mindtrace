#!/bin/bash
# ==============================================================================
# update-ec2-secrets.sh - Atualiza secrets do GitHub para AWS Academy
# ==============================================================================
# Uso:
#   Na EC2:     ./update-ec2-secrets.sh
#   Local:      ./update-ec2-secrets.sh <IP_DA_EC2>
#
# Requisitos:
#   - GitHub CLI (gh) instalado e autenticado
#   - Permissão 'repo' no token do GitHub
# ==============================================================================

set -e

REPO="shh4und/mindtrace"

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}╔════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║        MindTrace - AWS Academy Secret Updater              ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════════════════════════╝${NC}"
echo ""

# Detecta IP: parâmetro ou metadata da EC2
if [ -n "$1" ]; then
    PUBLIC_IP="$1"
    echo -e "${YELLOW}📍 Usando IP fornecido: ${PUBLIC_IP}${NC}"
else
    echo -e "${YELLOW}🔍 Buscando IP via EC2 metadata...${NC}"
    PUBLIC_IP=$(curl -s --connect-timeout 2 http://169.254.169.254/latest/meta-data/public-ipv4 2>/dev/null || echo "")
fi

if [ -z "$PUBLIC_IP" ]; then
    echo -e "${RED}❌ Não foi possível obter o IP público.${NC}"
    echo ""
    echo "Uso:"
    echo "  Na EC2:  ./update-ec2-secrets.sh"
    echo "  Local:   ./update-ec2-secrets.sh <IP_DA_EC2>"
    echo ""
    echo "Exemplo:"
    echo "  ./update-ec2-secrets.sh 54.123.45.67"
    exit 1
fi

echo -e "${GREEN}✓ IP Público: ${PUBLIC_IP}${NC}"
echo ""

# Verifica se gh está instalado
if ! command -v gh &> /dev/null; then
    echo -e "${RED}❌ GitHub CLI (gh) não está instalado.${NC}"
    echo ""
    echo "Instale com:"
    echo "  Ubuntu/Debian: sudo apt install gh"
    echo "  macOS:         brew install gh"
    echo "  Ou visite:     https://cli.github.com/"
    exit 1
fi

# Verifica se está autenticado
if ! gh auth status &> /dev/null; then
    echo -e "${RED}❌ GitHub CLI não está autenticado.${NC}"
    echo ""
    echo "Execute: gh auth login"
    exit 1
fi

echo -e "${BLUE}🔄 Atualizando secrets do repositório ${REPO}...${NC}"
echo ""

# Atualiza EC2_HOST
echo -n "   EC2_HOST: "
echo "$PUBLIC_IP" | gh secret set EC2_HOST --repo "$REPO" 2>/dev/null
echo -e "${GREEN}✓${NC}"

# Atualiza FRONTEND_API_BASE_URL (mantido para compatibilidade)
echo -n "   FRONTEND_API_BASE_URL: "
echo "http://${PUBLIC_IP}/api/v1" | gh secret set FRONTEND_API_BASE_URL --repo "$REPO" 2>/dev/null
echo -e "${GREEN}✓${NC}"

echo ""
echo -e "${GREEN}╔════════════════════════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║                    🎉 Sucesso!                             ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}Próximos passos:${NC}"
echo ""
echo "  1. Trigger o deploy:"
echo "     ${YELLOW}gh workflow run deploy.yml --repo $REPO${NC}"
echo "     Ou faça push para a branch main"
echo ""
echo "  2. Verifique o deploy:"
echo "     ${YELLOW}gh run watch --repo $REPO${NC}"
echo ""
echo "  3. Acesse a aplicação:"
echo "     ${GREEN}http://${PUBLIC_IP}${NC}"
echo ""
echo -e "${BLUE}Dica:${NC} Se precisar verificar os logs na EC2:"
echo "     ssh ubuntu@${PUBLIC_IP} 'docker logs mindtrace-backend'"
echo ""
