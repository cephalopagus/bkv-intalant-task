package main

import (
	"log/slog"
	"os"

	core_repository_postgres "github.com/cephalopagus/bkv-intalant-task/internal/core/repository/postgres"
)

func main() {
	cfg := core_repository_postgres.Load()
	db, err := core_repository_postgres.NewPostgresDB(cfg)
	if err != nil {
		slog.Error("failed to connect to db", "err", err)
		os.Exit(1)
	}

	slog.Info("connected to db", "host", cfg.DBHost, "db", cfg.DBName)
	slog.Info("connected to db",
		"host", cfg.DBHost,
		"port", cfg.DBPort,
		"user", cfg.DBUser,
		"db", cfg.DBName,
	)

	_ = db

}
