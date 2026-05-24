FROM golang:1.26.2-alpine AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

# Copia o monorepo completo para garantir o escopo do go.work e pacotes locais
COPY . .

WORKDIR /app/api-gateway

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o /app/api-gateway-bin ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/api-gateway-bin .

EXPOSE 8080

ENTRYPOINT ["./api-gateway-bin"]