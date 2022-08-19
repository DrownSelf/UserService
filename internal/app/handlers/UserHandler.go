package handlers

import (
	"InnoTaxi/internal/app/errors"
	"InnoTaxi/internal/app/middlewares"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/DTO"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	service   services.IUserService
	validator *validator.Validate
}

func New(service services.IUserService, validator *validator.Validate) Handler {
	return Handler{service: service, validator: validator}
}

func (h *Handler) Register(ctx *gin.Context) {
	var user DTO.CreateUserRequest
	if err := ctx.Bind(&user); err != nil {
		ctx.Error(errors.ErrInvalidData)
		return
	}

	if err := h.validator.Struct(user); err != nil {
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
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(errors.ErrInvalidData)
		return
	}

	if err := h.validator.Struct(user); err != nil {
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

func (h *Handler) InitRoutes(router *gin.Engine, repo *repositories.LogRepo) {
	router.Use(cors.Default())
	router.Use(gin.Recovery(), middlewares.Logger(repo))
	router.Use(errors.HandleErr)
	router.POST("/user", h.Register)
	router.GET("/user", h.LogIn)
}
