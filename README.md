# Calculator App Backend

REST API для сохранения и вычисления математических выражений. Приложение написано на Go, использует Echo для HTTP-роутинга, GORM для работы с базой данных и PostgreSQL для хранения данных.

## Возможности

- Создание вычисления по математическому выражению
- Получение списка сохраненных вычислений
- Обновление существующего вычисления
- Удаление вычисления
- Автоматическая миграция таблицы при запуске приложения

## Стек

- Go 1.26.3
- Echo
- GORM
- PostgreSQL
- govaluate

## Структура проекта

```text
cmd/
  main.go                         # Точка входа в приложение
internal/
  calculationService/
    orm.go                        # Модели вычислений
    repository.go                 # Слой доступа к базе данных
    service.go                    # Бизнес-логика
  db/
    db.go                         # Инициализация базы данных
  handlers/
    calculationHandlers.go        # HTTP-обработчики
```

## Переменные окружения

Создай локальный файл `.env` на основе примера:

```powershell
Copy-Item .env.example .env
```

Затем обнови `.env`, указав свои данные для подключения к PostgreSQL:

```env
DATABASE_DSN=host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable
SERVER_ADDR=:8080
```

Переменные:

| Название | Обязательная | Описание |
| --- | --- | --- |
| `DATABASE_DSN` | Да | Строка подключения к PostgreSQL для GORM |
| `SERVER_ADDR` | Нет | Адрес запуска HTTP-сервера. По умолчанию `:8080` |

Не коммить файл `.env`. В репозиторий должен попадать только `.env.example`.

## Запуск

Установить зависимости:

```powershell
go mod download
```

Запустить приложение:

```powershell
go run ./cmd
```

API будет доступно по адресу:

```text
http://localhost:8080
```

## API

### Создать вычисление

```http
POST /calculations
Content-Type: application/json
```

Тело запроса:

```json
{
  "expression": "2 + 2 * 3"
}
```

Ответ:

```json
{
  "id": "generated-uuid",
  "expression": "2 + 2 * 3",
  "result": "8"
}
```

Пример:

```powershell
curl.exe -X POST http://localhost:8080/calculations `
  -H "Content-Type: application/json" `
  -d "{\"expression\":\"2 + 2 * 3\"}"
```

### Получить список вычислений

```http
GET /calculations
```

Пример:

```powershell
curl.exe http://localhost:8080/calculations
```

### Обновить вычисление

```http
PATCH /calculations/{id}
Content-Type: application/json
```

Тело запроса:

```json
{
  "expression": "10 / 2"
}
```

Пример:

```powershell
curl.exe -X PATCH http://localhost:8080/calculations/{id} `
  -H "Content-Type: application/json" `
  -d "{\"expression\":\"10 / 2\"}"
```

### Удалить вычисление

```http
DELETE /calculations/{id}
```

Пример:

```powershell
curl.exe -X DELETE http://localhost:8080/calculations/{id}
```

## Разработка

Отформатировать код:

```powershell
gofmt -w ./cmd ./internal
```

Запустить проверки:

```powershell
go test ./...
```
