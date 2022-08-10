package handlers

import (
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/DTO"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/spf13/viper"
	"net/http"
)

type Handler struct {
	service services.IUserService
}

func New(service services.IUserService) Handler {
	return Handler{service: service}
}

func (handler *Handler) Register(ctx *gin.Context) {
	var userRequest DTO.CreateUserRequest
	if err := ctx.Bind(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	id, err := handler.service.RegisterUser(ctx, userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id, "name": userRequest.Name, "phoneNumber": userRequest.PhoneNumber})
}

func (handler *Handler) LogIn(ctx *gin.Context) {
	var userRequest DTO.LogInUserRequest
	if err := ctx.BindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		ctx.Abort()
		return
	}

	response, err := handler.service.LogInUser(ctx, userRequest)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": response.Id, "token": response.Token})
}

func (handler *Handler) InitRoutes(router *gin.Engine) {
	router.POST("/user", handler.Register)
	router.GET("/user", handler.LogIn)
}
