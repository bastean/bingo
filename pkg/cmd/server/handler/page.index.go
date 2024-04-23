package handler

import (
	"github.com/bastean/bingo/pkg/cmd/server/component/layout"
	"github.com/bastean/bingo/pkg/cmd/server/component/page"
	"github.com/gin-gonic/gin"
)

func IndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/")
		layout.Index(page.Home()).Render(c.Request.Context(), c.Writer)
	}
}