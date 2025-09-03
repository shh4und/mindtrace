#!/usr/bin/env bash
set -e
cd /home/ubuntu/mindtrace
git fetch origin main
git reset --hard origin/main
DOCKER_HUB_USERNAME=${DOCKER_HUB_USERNAME:-shh4und} \
docker compose -f docker-compose.prod.yml --env-file .env.prod pull
DOCKER_HUB_USERNAME=${DOCKER_HUB_USERNAME:-shh4und} \
docker compose -f docker-compose.prod.yml --env-file .env.prod up -d --remove-orphans
docker image prune -f
echo "Atualizado."