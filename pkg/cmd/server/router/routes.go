package router

import (
	"github.com/bastean/bingo/pkg/cmd/server/handler"
)

func InitRoutes() {
	router.NoRoute(handler.NotRoute())

	router.GET("/", handler.IndexPage())
	router.POST("/", handler.FormBinary())
	router.GET("/download", handler.Download())
}
