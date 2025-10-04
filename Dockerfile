FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем основной код
COPY . .

# Сборка бинарника
RUN go build -o url-shortener main.go

FROM alpine:3.18

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/url-shortener .
# Создаем непривилегированного пользователя
RUN adduser -D -u 1000 appuser
USER appuser

# Открываем порт
EXPOSE 8080

# Запуск приложения
CMD ["./url-shortener"]
