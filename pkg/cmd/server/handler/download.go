package handler

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func Download() gin.HandlerFunc {
	return func(c *gin.Context) {
		pathToFile := c.Query("filepath")

		verify := regexp.MustCompile(`^temp\/build-[0-9]{1,10}\/build\/[\w-]{1,50}$`)

		if !verify.MatchString(pathToFile) {
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}

		_, err := os.Stat(pathToFile)

		if errors.Is(err, os.ErrNotExist) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		tempDir, filename := filepath.Split(pathToFile)

		c.FileAttachment(pathToFile, filename)

		folderToRemove := filepath.Dir(strings.TrimSuffix(tempDir, "/"))

		os.RemoveAll(folderToRemove)
	}
}
