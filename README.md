        # postgresql — B-tree индексы: что, зачем, когда не помогают

        Homework-шаблон для урока **l1_btree_indexes** (B-tree индексы: что, зачем, когда не помогают) на платформе Vibe Learn.

        ## Что делать

        Дано: testcontainers PG + миграция, создающая orders(id, user_id, shop_id, status,
total_cents, created_at) с 1М синтетических строк. Реализуй на Go бенчмарк, который:
1) измеряет p50/p95/p99 трёх запросов (по user_id, по shop_id+created_at, по status);
2) парсит EXPLAIN ANALYZE для каждого и фиксирует Plan Type (Seq Scan / Index Scan);
3) применяет миграцию-фикс (CREATE INDEX CONCURRENTLY ...) и заново измеряет;
4) генерирует Markdown-отчёт «до vs после».
Тесты проверят корректность парсинга плана, latency-измерений и идемпотентность миграций.

## Контекст (из transfer-задачи урока)

У тебя в проде таблица `orders` на 50 млн строк. Жалоба: дашборд аналитики «продажи
по дням за месяц по конкретному магазину» отдаётся 12 секунд. Запрос:

```sql
SELECT date_trunc('day', created_at) AS day, count(*), sum(total_cents)
FROM orders
WHERE shop_id = 17
  AND created_at >= now() - INTERVAL '30 days'
  AND status = 'paid'
GROUP BY 1
ORDER BY 1;
```

## Recap из урока

- **B-tree — самый важный инструмент перфоманса.** Превращает Seq Scan в Index Scan, типичное ускорение в 100-1000×.
- **Композитный индекс работает по leftmost prefix.** Индекс (a,b,c) хорош для WHERE a, WHERE a AND b, WHERE a AND b AND c. WHERE b — нет.
- **CREATE INDEX CONCURRENTLY** — обязательно в проде. Без него таблица под блокировкой на минуты.
- **Индекс не бесплатен:** место + замедление записи + поддержка. Не лепи на каждую колонку — лепи под реальные запросы.
- **EXPLAIN ANALYZE — твой партнёр.** Перед оптимизацией смотри план; после — проверяй, что изменилось. Без замеров — гадание.

        ## Как работать

        1. Платформа Vibe Learn создаёт копию этого репо в твоём GitHub-аккаунте по клику «Начать домашку» на странице урока (через GitHub `/generate`, codecrafters-pattern).
        2. Склонируй копию локально, реализуй TODO в `main.go`, прогони тесты, запушь.
        3. CI (`.github/workflows/ci.yml`) запускает `go vet` + `go test ./...` на каждый push. Платформа слушает результат через webhook от GitHub Actions и обновляет статус домашки на странице урока.

        ## Локальное окружение

        - Go 1.22+
        - Docker + docker-compose — `docker compose up -d` поднимает single-node PostgreSQL 16 на `localhost:5432` с healthcheck. DSN: `postgres://postgres:postgres@localhost:5432/postgres`. Переопределяется через env `DATABASE_URL`.

        ## Запуск

        ```bash
        # Поднять локальный PostgreSQL
        docker compose up -d

        # Прогнать тесты (интеграционный включается через PG_INTEGRATION=1)
        go test ./...
        PG_INTEGRATION=1 go test ./...

        # Запустить main (печатает marker; замени stub на реализацию)
        go run .
        ```

        ## Заметка автора

        Это baseline-шаблон, сгенерированный платформой. Бизнес-сущность задачи (что конкретно реализовать в `main.go`, какие тесты сделать строгими) расширяется по ходу итераций — параллельно с углублением теории урока.
