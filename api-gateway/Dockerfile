# Estágio 1: Compila o binário
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila o binário do API Gateway
RUN go build -o api-gateway-bin ./api-gateway/cmd/main.go

# Estágio 2: Imagem final (Pequena e leve)
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/api-gateway-bin .

# Expo
EXPOSE 8080

CMD ["./api-gateway-bin"]