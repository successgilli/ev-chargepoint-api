package request

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewGetByIDRequest(c *gin.Context) (string, error) {
	id := strings.TrimSpace(c.Param("id"))

	if id == "" {
		return "", errors.New("id cannot be empty")
	}

	return id, nil
}
