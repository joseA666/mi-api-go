package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PostgresPool *pgxpool.Pool

func ConectarPostgres(url string) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}
	PostgresPool = pool
	log.Println("Postgres conectado")
}
