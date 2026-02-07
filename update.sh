#!/usr/bin/env bash
set -e

# Uso: scp docker-compose.prod.yml usuario@ec2:~/mindtrace/ && ssh usuario@ec2 'bash ~/mindtrace/update.sh'
# Ou se o compose já está no servidor, apenas: ssh usuario@ec2 'bash ~/mindtrace/update.sh'

cd ~/mindtrace

if [ ! -f .env.prod ]; then
  echo "❌ .env.prod não encontrado. Crie antes de rodar."
  exit 1
fi

mkdir -p postgres-data

DOCKER_HUB_USERNAME=${DOCKER_HUB_USERNAME:-shh4und} \
docker compose -f docker-compose.prod.yml --env-file .env.prod pull
DOCKER_HUB_USERNAME=${DOCKER_HUB_USERNAME:-shh4und} \
docker compose -f docker-compose.prod.yml --env-file .env.prod up -d --remove-orphans
docker image prune -f
echo "✅ Atualizado."