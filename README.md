# Проект "Визуализатор реферального графа"

Этот проект представляет собой сервис для визуализации связанных сущностей (пользователей) в виде графа.

## Технологический стек

- **Backend:** Go (гексагональная архитектура)
- **Frontend:** React + Vite (с библиотекой React-Sigma для визуализации)
- **База данных:** Neo4j
- **Кэш:** Redis
- **Инфраструктура:** Docker, Docker Compose, Traefik

## Быстрый старт

1.  **Клонируйте репозиторий:**
    ```bash
    git clone <your-repo-url>
    cd my-referral-graph
    ```

2.  **Настройте окружение:**
    Скопируйте файл с переменными окружения. Для локальной разработки менять ничего не нужно.
    ```bash
    cp .env.example .env
    ```

3.  **Добавьте локальные домены в `hosts`:**
    Чтобы Traefik мог правильно маршрутизировать запросы, добавьте следующие строки в ваш файл `/etc/hosts` (на Windows `C:\Windows\System32\drivers\etc\hosts`):
    ```
    127.0.0.1 api.myapp.local
    127.0.0.1 app.myapp.local
    127.0.0.1 playground.myapp.local
    ```

4.  **Запустите все сервисы:**
    Эта команда соберет и запустит все контейнеры в фоновом режиме.
    ```bash
    docker-compose up -d
    ```

5.  **Откройте приложения в браузере:**
    - **Основное приложение:** [http://app.myapp.local](http://app.myapp.local)
    - **"Песочница":** [http://playground.myapp.local](http://playground.myapp.local)
    - **API Health Check:** [http://api.myapp.local/api/health](http://api.myapp.local/api/health)
    - **Neo4j Browser:** [http://localhost:7474](http://localhost:7474) (логин/пароль: `neo4j` / `supersecretpassword`)
    - **Traefik Dashboard:** [http://localhost:8081](http://localhost:8081)

## Разработка

### Backend
- Код находится в папке `/backend`.
- Благодаря `air`, любые изменения в `.go` файлах автоматически пересобирают и перезапускают бэкенд-сервис.

### Frontend
- Код основного приложения в `/frontend/app`, "песочницы" - в `/frontend/playground`.
- Vite обеспечивает Hot Module Replacement (HMR) для мгновенного обновления в браузере.

### Использование моков
Для разработки фронтенда без реальной базы данных можно включить мок-режим на бэкенде.

1.  Откройте файл `.env`.
2.  Измените `USE_MOCKS=false` на `USE_MOCKS=true`.
3.  Перезапустите бэкенд-контейнер:
    ```bash
    docker-compose restart backend
    ```
Теперь API `/api/v1/graph` будет возвращать захардкоженные данные.

## Остановка окружения

Чтобы остановить все сервисы:
```bash
docker-compose down
```
Чтобы остановить и удалить volumes (данные Neo4j будут потеряны):
```bash
docker-compose down -v
```
