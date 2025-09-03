# mindtrace-projeto/frontend/Dockerfile

# --- Estágio 1: Build ---
FROM node:22.17.1-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
ARG VITE_API_BASE_URL
ENV VITE_API_BASE_URL=${VITE_API_BASE_URL}
RUN npm run build

# --- Estágio 2: Produção ---
# Usa uma imagem de um servidor web leve (Nginx) para servir os arquivos estáticos
FROM nginx:1.29-alpine

# Copia os arquivos estáticos gerados no estágio de build para a pasta que o Nginx serve
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expõe a porta 80 (padrão do Nginx)
EXPOSE 80

# O Nginx inicia automaticamente, então não precisamos de um CMD