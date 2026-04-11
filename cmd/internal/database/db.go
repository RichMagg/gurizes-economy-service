package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(connString string) *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
