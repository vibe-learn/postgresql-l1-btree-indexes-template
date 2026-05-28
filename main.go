// Package main is the postgresql lesson `l1_btree_indexes` homework scaffold for Vibe Learn.
//
// Задача: orders(1M строк): p50/p95/p99 трёх запросов + EXPLAIN Plan Type до/после CREATE INDEX.
// Реализуй функции ниже — сигнатуры и тестовая поверхность фиксированы;
// CI (.github/workflows/ci.yml) гоняет `go vet` и `go test ./...`.
// Подробности и критерии приёмки — в README.md.
//
// Драйвер: github.com/jackc/pgx/v5 (+ pgxpool). DATABASE_URL — DSN из env.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Latencies — собранные перцентили для бенчмарка запроса.
type Latencies struct{ P50, P95, P99 time.Duration }

// StandbyInfo — строка из pg_stat_replication для выбора кандидата на promote.
type StandbyInfo struct {
	ClientAddr string
	ReplayLSN  string
	State      string
}

// ----- config -----

// envOr returns the env var for `key` if set, else `fallback`.
func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// DatabaseURL — DSN PostgreSQL. Дефолт совпадает с docker-compose.yml.
func DatabaseURL() string {
	return envOr("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
}

// Connect открывает пул pgx из DATABASE_URL.
func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, DatabaseURL())
}

// ----- TODO #1: ExplainPlanType -----
//
// EXPLAIN (FORMAT JSON) запроса, вытащить верхний "Node Type" (Seq Scan / Index Scan)
func ExplainPlanType(ctx context.Context, pool *pgxpool.Pool, sql string, args ...any) (planType string, err error) {
	// TODO: implement
	panic("ExplainPlanType: not implemented")
}

// ----- TODO #2: BenchQuery -----
//
// прогнать запрос N раз, вернуть p50/p95/p99
func BenchQuery(ctx context.Context, pool *pgxpool.Pool, sql string, args ...any) (Latencies, error) {
	// TODO: implement
	panic("BenchQuery: not implemented")
}

// ----- TODO #3: ApplyIndexFix -----
//
// CREATE INDEX CONCURRENTLY ... (идемпотентно: IF NOT EXISTS)
func ApplyIndexFix(ctx context.Context, pool *pgxpool.Pool) error {
	// TODO: implement
	panic("ApplyIndexFix: not implemented")
}

// _refs keeps imports live while the TODO bodies are unimplemented stubs.
// Удали эту функцию, когда реализуешь TODO выше.
var _refs = []any{
	Latencies{},
	StandbyInfo{},
	time.Second,
}

// ----- main entry -----

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Printf("Vibe Learn — postgresql lesson %s scaffold up", "l1_btree_indexes")
	log.Printf("DATABASE_URL: %s", DatabaseURL())
	log.Printf("Реализуй TODO-функции, затем `go test ./...`. README.md содержит задачу.")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Graceful shutdown so `go run .` is interactive — Ctrl-C exits cleanly.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Printf("shutdown signal received")
		cancel()
	}()
	<-ctx.Done()
}
