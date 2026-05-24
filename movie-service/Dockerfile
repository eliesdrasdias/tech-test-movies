FROM golang:1.26.2-alpine AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

# Copia o monorepo completo para garantir o escopo do go.work e pacotes locais
COPY . .

WORKDIR /app/movie-service

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o /app/movie-service-bin ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/movie-service-bin .
COPY --from=builder /app/data/ ./data

EXPOSE 50051

ENTRYPOINT ["./movie-service-bin"]