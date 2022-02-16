package user

import (
	api "dev-hack-backend/internal/adapters/api"
	"dev-hack-backend/internal/service/user"
	"dev-hack-backend/pkg/apperror"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const (
	userURL       = "/:user_id"
	pathParamUser = "user_id"
	userLogin     = "/login"
	userRegister  = "/register"
	//tokenRefresh  = "/refresh_token"
	contextTTL = time.Second * 15
)

type handler struct {
	userService user.Service
}

func NewHandler(userService user.Service) api.Handler {
	return &handler{
		userService: userService,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.POST(userLogin, h.SignIn)
	router.POST(userRegister, h.SignUp)

	//auth := router.Group("/auth", h.tokenValid)
	//auth.POST(tokenRefresh, h.RefreshToken)

	userGroup := router.Group("/user", h.userIdentity)
	userGroup.GET(userURL, h.GetUser)
	userGroup.PATCH(userURL, h.PathUser)
	userGroup.DELETE(userURL, h.DeleteUser)
}

func (h *handler) SignIn(c *gin.Context) {
	ctx, cancel := h.userService.ContextWithTimeout(c, contextTTL, pathParamUser)
	defer cancel()

	var userDTO SignInUserDTO
	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		api.NewResponse(c, http.StatusBadRequest, "not all parameters are specified "+err.Error())
		return
	}

	currentUser, err := h.userService.Authorize(ctx, userDTO.Username, userDTO.Password)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	aToken, rToken, err := h.userService.CreateSession(ctx, currentUser.Id.Hex())
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.ResponseWithTokens(c, http.StatusCreated, currentUser.Id.Hex(), aToken, rToken)
}

func (h *handler) SignUp(c *gin.Context) {
	ctx, cancel := h.userService.ContextWithTimeout(c, contextTTL, pathParamUser)
	defer cancel()

	var userDTO CreateUserDTO
	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		api.NewResponse(c, http.StatusBadRequest, "not all parameters are specified "+err.Error())
		return
	}

	currentUser := userDTO.toUser()
	err = h.userService.InsertUser(ctx, currentUser)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	aToken, rToken, err := h.userService.CreateSession(ctx, currentUser.Id.Hex())
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.ResponseWithTokens(c, http.StatusCreated, currentUser.Id.Hex(), aToken, rToken)
}

//func (h *handler) RefreshToken(c *gin.Context) {
//	ctx, cancel := h.userService.ContextWithTimeout(c, contextTTL, pathParamUser)
//	defer cancel()
//
//	var token TokenDTO
//
//}

func (h *handler) GetUser(c *gin.Context) {
	ctx, cancel := h.userService.ContextWithTimeout(c, contextTTL, pathParamUser)
	defer cancel()

	currentUser, err := h.userService.GetUser(ctx)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.ResponseUser(c, http.StatusOK, currentUser)
}

func (h *handler) PathUser(c *gin.Context) {
	ctx, cancel := h.userService.ContextWithTimeout(c, contextTTL, pathParamUser)
	defer cancel()

	currentUser, err := h.userService.GetUser(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = c.ShouldBindJSON(&currentUser)
	if err != nil {
		api.NewResponse(c, http.StatusBadRequest, apperror.BadData.Error()+err.Error())
		return
	}

	ok := h.userService.CompareID(ctx, c.Param(pathParamUser))
	if !ok {
		api.NewResponseStatus(c, http.StatusUnauthorized)
		return
	}

	err = h.userService.UpdateUser(ctx, currentUser)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.NewResponseStatus(c, http.StatusOK)
}

func (h *handler) DeleteUser(c *gin.Context) {
	ctx, cancel := h.userService.ContextWithTimeout(c, contextTTL, pathParamUser)
	defer cancel()

	ok := h.userService.CompareID(ctx, c.Param(pathParamUser))
	if !ok {
		api.NewResponseStatus(c, http.StatusUnauthorized)
		return
	}

	err := h.userService.DeleteUser(ctx)
	if err != nil {
		api.NewResponse(c, http.StatusInternalServerError, err.Error())
	}

	api.NewResponseStatus(c, http.StatusNoContent)
}
