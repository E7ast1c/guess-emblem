package app

import (
	"context"
	"guess-emblem/internal/config"
	"guess-emblem/internal/emblem"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	Ctx context.Context

	AppConfig config.AppConfig
	Repo      emblem.Repository
	DbConn    *sqlx.DB
	ApiServer *gin.Engine
}

func NewContainer(ctx context.Context) *Container {
	return &Container{
		Ctx:       ctx,
		AppConfig: config.AppConfig{},
		Repo:      nil,
		DbConn:    nil,
		ApiServer: nil,
	}
}
