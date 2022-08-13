package handlers

import (
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/DTO"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/spf13/viper"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	err := h.validator.Struct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	id, err := h.service.RegisterUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id, "name": user.Name, "phoneNumber": user.PhoneNumber})
}

func (h *Handler) LogIn(ctx *gin.Context) {
	var user DTO.LogInUserRequest
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		ctx.Abort()
		return
	}

	err := h.validator.Struct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	response, err := h.service.LogInUser(ctx, user)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": response.Id, "token": response.Token})
}

func (h *Handler) InitRoutes(router *gin.Engine) {

	router.POST("/user", h.Register)
	router.GET("/user", h.LogIn)
}
