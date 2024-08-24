package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func InitializeDB(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	db = pool
	fmt.Println("Connected to Database")
	return nil
}

func GetDB() *pgxpool.Pool {
	return db
}
