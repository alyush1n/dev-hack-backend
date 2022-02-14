package user

import (
	"context"
	"dev-hack-backend/internal/adapters/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	auth := router.Group("/auth")
	auth.POST(userPost, h.SignIn)
	auth.POST(userPost, h.SignUp)

	user := router.Group("/user", h.userIdentity)
	user.GET(userURL, h.GetUser)
	user.PATCH(userURL, h.PathUser)
	user.DELETE(userURL, h.DeleteUser)
}

func (h *handler) SignIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*15)
	defer cancel()
	var userDTO SignInUserDTO

	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		api.NewResponse(c, http.StatusBadRequest, "not all parameters are specified "+err.Error())
		return
	}
	user, err := h.userService.Authorize(ctx, &userDTO)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	aToken, rToken, err := h.userService.CreateSession(ctx, user.Id.Hex())
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.ResponseWithTokens(c, http.StatusCreated, aToken, rToken)
}

func (h *handler) SignUp(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*15)
	defer cancel()
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

	aToken, rToken, err := h.userService.CreateSession(ctx, user.Id.Hex())
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.ResponseWithTokens(c, http.StatusCreated, aToken, rToken)
}

func (h *handler) GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*15)
	defer cancel()
	user, err := h.userService.GetUser(ctx)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	api.ResponseUser(c, http.StatusOK, user)
}

func (h *handler) PathUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*15)
	defer cancel()
	var userDTO UpdateUserDTO

	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		api.NewResponse(c, http.StatusBadRequest, "not all parameters are specified "+err.Error())
		return
	}
	user, err := h.userService.UpdateUser(ctx, &userDTO)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
	}
	api.ResponseUser(c, http.StatusOK, user)
}

func (h *handler) DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*15)
	defer cancel()
	err := h.userService.DeleteUser(ctx)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
	}
	api.NewResponse(c, http.StatusNoContent, "")
}
