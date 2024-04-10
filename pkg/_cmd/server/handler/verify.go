package handler

import (
	"net/http"

	"github.com/bastean/bingo/pkg/cmd/server/service/user"
	"github.com/bastean/bingo/pkg/context/user/application/verify"
	"github.com/gin-gonic/gin"
)

type Param struct {
	Id string
}

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Redirect(http.StatusFound, "/")
			}
		}()

		verifyParam := Param{Id: c.Param("id")}

		command := verify.Command(verifyParam)

		user.UserVerifyHandler.Handle(&command)

		c.Redirect(http.StatusFound, "/")
	}
}
