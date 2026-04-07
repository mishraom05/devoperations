package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mishraom05/golang-project/internal/env"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	//Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// DATABASE
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("Connected to DATABASE", "dsn", cfg.db.dsn)

	// APPLICATION
	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		//log.Printf("Server has failed to start, err: %s", err)
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
}
