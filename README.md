# Приложение по работе с grpc и kafka
Приложение представляет собой два сервиса:
1) Первый получает запрос по grpc, и отправляет его в kafka
2) Второй получает сообщение из kafka и записывает его в бд в транзакции


## Структура проекта

### `cmd`

- `producer` - сервис получателя запросов по grpc и отправителя в kafka
- `consumer` - сервис получателя запросов из kafka

### `api`

- `server_v1` - proto файлы для grpc

### `internal`

- `api` - реализация grpc-методов
- `app` - пакет для запуска приложения
- `server` - компонент grpc-сервера
- `models` - модели базы данных
- `repository` - репозиторий для взаимодействия с базой данных
- `service` - внутренние сервисы
    - `server` - бизнес-логика для сервиса producer
    - `server-consumer` - бизнес-логика для сервиса consumer
    - `kafka` - взаимодействие producer и consumer через очередь сообщений

### `pkg`

- `server_v1` - сгенерированные файлы для grpc
- `logger` - реализация логгера для всего проекта


## Инструкция по установке
Клонируйте репозиторий:
```
git clone https://github.com/StasikLeyshin/grpc-kafka-services.git
```
Переходите в клонированную директорию:
```
cd grpc-kafka-services
```
Разверните приложение в docker:
```
docker-compose up -d --build
```