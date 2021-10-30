package handlers

import (
	"guess-emblem/internal/config"
	"guess-emblem/internal/emblem"

	"github.com/gin-gonic/gin"
)

func BuildHandlers(repo emblem.Repository, cfg config.AppConfig) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

 	elemService := emblem.Service(emblem.NewService(repo))

	rg := router.Group("/")
	emblem.RegisterHandlers(rg, elemService)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		c.Next()
	}
}
