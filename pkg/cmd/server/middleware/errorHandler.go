package middleware

import (
	"fmt"
	"net/http"

	"github.com/bastean/bingo/pkg/cmd/server/util"
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
		err = fmt.Errorf("internal error: please try again later")
	}

	c.JSON(code, util.JSONResponse(false, err.(error).Error(), struct{}{}))

	c.Abort()
}
