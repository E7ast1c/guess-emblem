package main

import (
	"context"
	"fmt"
	"guess-emblem/internal/app"
	"guess-emblem/internal/config"
	"guess-emblem/internal/db/postgres"
	"guess-emblem/internal/emblem"
	"guess-emblem/internal/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	c := app.Container{}
	c.AppConfig = config.NewAppConfig().GetValues()
	c.DbConn = postgres.MustPGConnect(c.Ctx, c.AppConfig.DBConfig)
	c.Repo = emblem.NewRepository(c.DbConn)
	c.ApiServer = handlers.BuildHandlers(c.Repo, c.AppConfig)
	if err := c.ApiServer.Run(fmt.Sprintf("%s:%s",
		c.AppConfig.APIServer.Address,
		c.AppConfig.APIServer.Port)); err != nil {
		logrus.Fatalf("run api server failed %s", err)
	}

	_ = app.NewContainer(ctx)
}

