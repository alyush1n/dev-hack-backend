package user

import (
	"context"
	"dev-hack-backend/internal/adapters/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userURL  = "/:user_id"
	userPost = "/"
)

type handler struct {
	userService Service
}

func NewHandler(userService Service) api.Handler {
	return &handler{
		userService: userService,
	}
}

func (h *handler) Register(router *gin.Engine) {
	auth := router.Group("auth")
	auth.POST(userPost, h.SignIn)
	auth.POST(userPost, h.SignUp)

	user := router.Group("user", h.userIdentity)
	user.GET(userURL, h.GetUser)
	user.PUT(userURL, h.PutUser)
	user.PATCH(userURL, h.PathUser)
	user.DELETE(userURL, h.DeleteUser)
}

func (h *handler) SignIn(c *gin.Context) {
	ctx := context.Background()

}

func (h *handler) SignUp(c *gin.Context) {
	ctx := context.Background()
	var userDTO CreateUserDTO

	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		api.NewResponse(c, http.StatusBadRequest, "not all parameters are specified "+err.Error())
		return
	}
	user, err := h.userService.InsertUser(ctx, &userDTO)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}

func (h *handler) GetUser(c *gin.Context) {

}

func (h *handler) PutUser(c *gin.Context) {
	ctx := context.Background()

}

func (h *handler) PathUser(c *gin.Context) {
	ctx := context.Background()

}

func (h *handler) DeleteUser(c *gin.Context) {
	ctx := context.Background()

}
