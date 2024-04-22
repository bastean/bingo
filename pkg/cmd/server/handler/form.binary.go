package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bastean/bingo/pkg/cmd/server/service/builder"
	"github.com/bastean/bingo/pkg/context/binary/application/create"
	"github.com/gin-gonic/gin"
)

func FormBinary() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)

		fmt.Println(string(body))

		query := create.NewQuery("", "")

		response := builder.BinaryCreateHandler.Handle(query)

		filepath := response.FilePath

		download := fmt.Sprintf("/download?filepath=%s", filepath)

		c.Header("HX-Redirect", download)

		c.Status(http.StatusCreated)
	}
}
