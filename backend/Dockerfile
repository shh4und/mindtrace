# mindtrace-projeto/backend/Dockerfile

# --- Estágio 1: Build ---
# Usamos uma imagem que já tem o Go instalado para compilar nosso projeto
FROM golang:1.24.5-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de gerenciamento de dependências primeiro
# Isso aproveita o cache do Docker. Se go.mod e go.sum não mudarem,
# o Docker não vai baixar as dependências de novo.
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o resto do código do backend
COPY . .

# Compila a aplicação.
# CGO_ENABLED=0 cria um binário estático, necessário para rodar numa imagem base mínima sem libs C.
# -o /app/main cria o arquivo executável 'main' no diretório /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api/main.go

# --- Estágio 2: Produção ---
# Começamos uma nova imagem, muito menor, pois não precisamos mais do compilador do Go.
# DESCOMENTE PARA PRODUCAO / COMENTE PARA DESENVOLVIMENTO (CODE COMPLETIONS)
# FROM alpine:latest

WORKDIR /app

# Copia apenas o executável compilado do estágio 'builder'
# DESCOMENTE PARA PRODUCAO / COMENTE PARA DESENVOLVER (CODE COMPLETIONS)
# COPY --from=builder /app/main .

# Expõe a porta que a sua API Gin vai usar (ex: 8080)
EXPOSE 8080

# Comando para iniciar a aplicação quando o container rodar
CMD [ "./main" ]