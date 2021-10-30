package postgres

import (
	"context"
	"guess-emblem/internal/config"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func MustPGConnect(ctx context.Context, dbConfig config.DBConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", dbConfig.URI)
	if err != nil {
		logrus.Fatalf("Unable to connect to database: %s\n", err)
	}
	return db
}

