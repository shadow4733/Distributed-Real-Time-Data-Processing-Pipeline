# Указываем базовый образ для Go
FROM golang:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем весь исходный код в контейнер
COPY . ./

# Переходим в папку, где лежит main.go (папка cmd)
WORKDIR /app/src/cmd

# Собираем приложение
RUN go mod tidy && go build -o /app/myapp .

# Указываем команду для запуска приложения
CMD ["/app/myapp"]
