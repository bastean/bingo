package handler

import (
	"github.com/bastean/bingo/pkg/cmd/server/component/layout"
	"github.com/bastean/bingo/pkg/cmd/server/component/page/home"
	"github.com/bastean/bingo/pkg/cmd/server/component/script"
	"github.com/gin-gonic/gin"
)

func IndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		layout.Index(home.FormBinary(), home.Page(), script.Empty()).Render(c.Request.Context(), c.Writer)
	}
}
