# Estágio 1: Compilar o binário
FROM golang:latest AS builder

WORKDIR /app

# Copia o go.mod e go.sum para dentro do container
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante da aplicação
COPY . .

# Compila o binário
RUN go build -o movie-service-bin ./movie-service/cmd/main.go

# Estágio 2: Rodar a aplicação
FROM debian:bookworm-slim

WORKDIR /app

# Copia o binário
COPY --from=builder /app/movie-service-bin .

# Copia os arquivos de configuração
COPY --from=builder /app/data ./data

# Expo
EXPOSE 50051

# Roda a aplicação
CMD ["./movie-service-bin"]