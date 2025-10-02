# REST API

Простой и эффективный REST API сервер на Go с поддержкой базовых CRUD операций.

## Описание

Этот проект представляет собой RESTful API сервер, разработанный на языке Go. Сервер предоставляет API endpoints для работы с данными, поддерживает аутентификацию и использует PostgreSQL в качестве базы данных.

## Функциональность

- ✅ RESTful API endpoints
- ✅ Поддержка PostgreSQL
- ✅ JWT аутентификация
- ✅ Логирование с настраиваемым уровнем
- ✅ Конфигурация через TOML файл
- ✅ Сессии пользователей

## Требования

- Go 1.19+
- PostgreSQL 12+

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/TemniyKozhevnik/REST_API.git
cd REST_API
```

### 2. Установка зависимостей

```bash
go mod tidy
```

### 3. Настройка конфигурации

Создайте или отредактируйте файл конфигурации configs/apiserver.toml:
```toml
bind_addr = ":8080"
log_level = "debug"
database_url = "user=test host=localhost dbname=restapi_dev sslmode=disable password=test"
session_key = "test"
```

### 4. Настройка базы данных

Создайте или отредактируйте файл конфигурации configs/apiserver.toml:
```sql
CREATE DATABASE restapi_dev;
CREATE USER test WITH PASSWORD 'test';
GRANT ALL PRIVILEGES ON DATABASE restapi_dev TO test;
```

## Запуск проекта

```bash
go build -o rest-api cmd/apiserver/main.go
./rest-api
```


