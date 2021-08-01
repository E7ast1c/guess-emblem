package handlers

import (
	"guess-emblem/internal/config"
	"guess-emblem/internal/emblem"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func BuildHandlers(db sqlx.DB, cfg config.AppConfig) *gin.Engine {
	router := gin.Default()
 	elemService := emblem.Service(emblem.NewService(emblem.NewRepository(&db)))
	rg := router.Group("/")

	emblem.RegisterHandlers(rg, elemService)
	return router
}
