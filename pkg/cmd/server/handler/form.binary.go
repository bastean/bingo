package handler

import (
	"fmt"
	"net/http"

	"github.com/bastean/bingo/pkg/cmd/server/service/builder"
	"github.com/bastean/bingo/pkg/cmd/server/util"
	"github.com/bastean/bingo/pkg/context/binary/application/create"
	"github.com/gin-gonic/gin"
)

func FormBinary() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := new(create.Query)

		c.BindJSON(query)

		response := builder.BinaryCreateHandler.Handle(query)

		filepath := response.FilePath

		download := fmt.Sprintf("/download?filepath=%s", filepath)

		c.JSON(http.StatusCreated, util.JSONResponse(true, "Download started!", struct{ Url string }{download}))
	}
}
