# Etapa de build
FROM golang:1.20-alpine AS builder

# Configura a pasta de trabalho
WORKDIR /app

# Copia o go.mod e go.sum e instala dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante dos arquivos
COPY . .

# Compila o binário do aplicativo
RUN go build -o bufalo-wallet-app ./cmd/main.go

# Etapa final
FROM alpine:latest

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/bufalo-wallet-app /bufalo-wallet-app

# Define a porta que o serviço irá expor
EXPOSE 8080

# Comando para rodar o serviço
ENTRYPOINT ["/bufalo-wallet-app"]
