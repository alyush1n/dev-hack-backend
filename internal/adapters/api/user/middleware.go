package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	userCtx             = "user_id"
	authorizationHeader = "Authorization"
	headerEmptyError    = "header is empty"
	invalidHeaderError  = "invalid auth header"
	tokenEmptyError     = "token is empty"
)

func (h *handler) userIdentity(c *gin.Context) {
	id, err := h.parseAuthHandler(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Set(userCtx, id)
}

func (h *handler) parseAuthHandler(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", fmt.Errorf(headerEmptyError)
	}

	headerSplit := strings.Split(header, " ")
	if len(headerSplit) != 2 || headerSplit[0] != "Bearer" {
		return "", fmt.Errorf(invalidHeaderError)
	}

	if len(headerSplit[1]) == 0 {
		return "", fmt.Errorf(tokenEmptyError)
	}
	return h.userService.ParseToken(headerSplit[1])
}
