# mindtrace-projeto/frontend/Dockerfile

# --- Estágio 1: Build ---
# Usa uma imagem com Node.js para construir o projeto Vue
FROM node:22.17.1-alpine AS builder

WORKDIR /app

# Copia package.json e package-lock.json (ou yarn.lock, etc)
COPY package*.json ./

# Instala as dependências
RUN npm install

# Copia o restante do código do frontend
COPY . .

# Roda o script de build para gerar os arquivos estáticos
RUN npm run build

# --- Estágio 2: Produção ---
# Usa uma imagem de um servidor web leve (Nginx) para servir os arquivos estáticos
FROM nginx:1.29-alpine

# Copia os arquivos estáticos gerados no estágio de build para a pasta que o Nginx serve
COPY --from=builder /app/dist /usr/share/nginx/html

# Expõe a porta 80 (padrão do Nginx)
EXPOSE 80

# O Nginx inicia automaticamente, então não precisamos de um CMD