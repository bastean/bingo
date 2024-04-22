package middleware

import (
	"net/http"

	"github.com/bastean/bingo/pkg/cmd/server/component/partial/alert"
	"github.com/bastean/bingo/pkg/context/shared/domain/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	var code int

	switch err.(type) {
	case errors.InvalidValue:
		code = http.StatusUnprocessableEntity
	default:
		code = http.StatusInternalServerError
	}

	c.Status(code)

	alert.Message("error", err.(error).Error()).Render(c.Request.Context(), c.Writer)

	c.Abort()
}
