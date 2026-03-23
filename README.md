# Go + Svelte Web Application

Веб-приложение с Go бэкендом (Gin) и Svelte фронтендом.

## Структура проекта

```
go/
├── Makefile           # Команды для запуска
├── backend/           # Go сервер
│   ├── main.go
│   ├── data/          # База данных SQLite
│   ├── db/            # Работа с БД
│   ├── handlers/      # API обработчики
│   ├── models/        # Модели данных
│   └── router/        # Настройка роутера
└── frontend/          # Svelte приложение
    ├── src/
    │   ├── pages/     # Страницы
    │   ├── App.svelte
    │   └── main.js
    └── package.json
```

## Требования

- Go 1.21+
- Node.js 18+
- npm

## Быстрый старт

```bash
# Установка зависимостей
make install

# Запуск в режиме разработки (backend + frontend)
make run
```

После запуска:
- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080

## Команды Makefile

| Команда | Описание |
|---------|----------|
| `make run` | Запуск backend и frontend |
| `make backend` | Запуск только Go сервера |
| `make frontend` | Запуск только Vite dev сервера |
| `make build` | Сборка прод версий |
| `make install` | Установка npm зависимостей |
| `make clean` | Очистка артефактов сборки |

## API Endpoints

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/products` | Список товаров (пагинация, поиск, сортировка) |
| POST | `/api/products` | Создание товара |
| PUT | `/api/products/:id` | Обновление товара |
| DELETE | `/api/products/:id` | Удаление товара |

### Параметры API products

| Параметр | Описание | По умолчанию |
|----------|----------|--------------|
| `page` | Номер страницы | 1 |
| `pageSize` | Кол-во на странице | 10 |
| `search` | Поиск по названию | - |
| `sort` | Поле для сортировки (id, name, price) | id |
| `sortDir` | Направление (asc, desc) | asc |

## Технологии

- **Backend**: Go, Gin, SQLite
- **Frontend**: Svelte 4, Vite
- **База данных**: SQLite (10,000 тестовых товаров)

## Разработка

```bash
# Терминал 1 - Backend
cd backend
go run .

# Терминал 2 - Frontend
cd frontend
npm run dev
```

Frontend автоматически проксирует `/api` запросы на `http://localhost:8080`.
