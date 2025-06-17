# Этап сборки Go-приложения
FROM golang:1.22rc2 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o app ./cmd/library-app

# Финальный образ
FROM debian:bookworm-slim

# Установка зависимостей для скачивания и запуска
RUN apt-get update && \
    apt-get install -y curl ca-certificates tar && \
    apt-get clean

# Создаём рабочую директорию
WORKDIR /app

# Копируем Go-приложение
COPY --from=builder /app/app .

# Копируем миграции
COPY ./pkg/migrations ./migrations

# Скачиваем migrate CLI
RUN rm -f /app/migrate \
 && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz -o migrate.linux-amd64.tar.gz \
 && tar -xzf migrate.linux-amd64.tar.gz \
 && chmod +x /app/migrate \
 && rm migrate.linux-amd64.tar.gz

# Открываем порт
EXPOSE 8000

# Запуск: сначала миграции, потом само приложение
CMD /app/migrate -path ./migrations -database ${DB_DSN} up && ./app
