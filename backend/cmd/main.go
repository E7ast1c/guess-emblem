package main

import (
	"context"
	"guess-emblem/internal/config"
	"guess-emblem/internal/db/postgres"
	"guess-emblem/internal/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	appConfig := config.New()
	dbConn := postgres.MustPGConnect(ctx, appConfig.All().DBConfig)
	webEngine := handlers.BuildHandlers(*dbConn, appConfig)

	err := webEngine.Run("127.0.0.1:3000")
	if err != nil {
		logrus.Error(err)
	}
}

