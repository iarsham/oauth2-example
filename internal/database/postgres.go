package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/iarsham/oauth2-example/configs"
	_ "github.com/lib/pq"
	"time"
)

func OpenDB(cfg *configs.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", makeDsn(cfg))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(cfg.Postgres.ConnMaxIdleTime) * time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func makeDsn(cfg *configs.Config) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DB,
	)
}
