package handlers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/errors"
	"InnoTaxi/internal/app/middlewares"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/DTO"
)

type Handler struct {
	service services.IUserService
}

func New(service services.IUserService) Handler {
	return Handler{service: service}
}

func (h *Handler) Register(ctx *gin.Context) {
	var user DTO.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(errors.ErrInvalidData)
		return
	}

	id, err := h.service.RegisterUser(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, id)
}

func (h *Handler) LogIn(ctx *gin.Context) {
	var user DTO.LogInUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(errors.ErrInvalidData)
		return
	}

	response, err := h.service.LogInUser(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) ChangePassword(ctx *gin.Context) {
	var user DTO.ChangeUserPassword
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(errors.ErrInvalidData)
		return
	}

	if err := h.service.ChangeUserPassword(ctx, user); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Updated successfully")
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	var id DTO.DeleteUserRequest
	if err := ctx.ShouldBindJSON(&id); err != nil {
		ctx.Error(errors.ErrInvalidData)
		return
	}

	if err := h.service.DeleteUser(ctx, id.Id); err != nil {
		ctx.Error(errors.ErrInvalidData)
		return
	}
	ctx.JSON(http.StatusOK, "Deleted successfully")
}

func (h *Handler) InitRoutes(router *gin.Engine, repo *repositories.LogRepo, forger auth.TokenForger) {
	router.Use(cors.Default())
	router.Use(gin.Recovery(), middlewares.Logger(repo))
	router.Use(errors.HandleErr)

	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", h.Register)
		userGroup.POST("/login", h.LogIn)
		authGroup := userGroup.Group("/", middlewares.AuthMiddleware(forger))
		{
			authGroup.PUT("", h.ChangePassword)
			authGroup.DELETE("", h.DeleteUser)
		}
	}
}
