package user

import (
	api "dev-hack-backend/internal/adapters/api"
	"dev-hack-backend/pkg/apperror"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user_id"
)

func (h *handler) userIdentity(c *gin.Context) {
	id, err := h.parseAuthHandler(c)
	if err != nil {
		api.NewAbortResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, id)
	c.Next()
}

func (h *handler) parseAuthHandler(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", apperror.HeaderEmptyError
	}

	headerSplit := strings.Split(header, " ")
	if len(headerSplit) != 2 || headerSplit[0] != "Bearer" {
		return "", apperror.InvalidHeaderError
	}

	if len(headerSplit[1]) == 0 {
		return "", apperror.TokenEmptyError
	}

	return h.userService.ParseToken(headerSplit[1])
}

func (h *handler) tokenValid(c *gin.Context) {

}
