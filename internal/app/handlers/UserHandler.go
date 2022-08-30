package handlers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"InnoTaxi/internal/app/appErrors"
	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/middlewares"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/DTO"
)

type Handler struct {
	userService services.IUserService
}

func New(service services.IUserService) Handler {
	return Handler{userService: service}
}

func (h *Handler) Register(ctx *gin.Context) {
	var user DTO.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(appErrors.ErrInvalidData)
		return
	}

	id, err := h.userService.RegisterUser(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, id)
}

func (h *Handler) LogIn(ctx *gin.Context) {
	var user DTO.LogInUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(appErrors.ErrInvalidData)
		return
	}

	response, err := h.userService.LogInUser(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) ChangeUserInfo(ctx *gin.Context) {
	var user DTO.ChangeUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(appErrors.ErrInvalidData)
		return
	}

	if err := h.userService.ChangeUserPassword(ctx, user); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Updated successfully")
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	var id DTO.DeleteUserRequest
	if err := ctx.ShouldBindJSON(&id); err != nil {
		ctx.Error(appErrors.ErrInvalidData)
		return
	}

	if err := h.userService.DeleteUser(ctx, id.Id); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Deleted successfully")
}

func (h *Handler) GetUserInfo(ctx *gin.Context) {
	var phoneNumber DTO.GetUserInfoRequest
	if err := ctx.ShouldBindJSON(&phoneNumber); err != nil {
		ctx.Error(appErrors.ErrInvalidData)
		return
	}

	user, err := h.userService.GetUserByPhone(ctx, phoneNumber.PhoneNumber)
	if err != nil {
		ctx.Error(err)
		return
	}

	userResponse := DTO.GetUserResponse{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Rating:      user.Rating,
	}
	ctx.JSON(http.StatusOK, userResponse)
}

func (h *Handler) LogOut(ctx *gin.Context) {
	if err := h.userService.LogOutUser(ctx, ctx.GetHeader("Authorization")); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Success log out")
}

type MiddlewareDependencies struct {
	LogRepository    repositories.ILogRepo
	Forger           auth.TokenForger
	CacheRepository  repositories.ICacheRepository
	MetricRepository *repositories.PrometheusRepository
}

func (h *Handler) InitRoutes(router *gin.Engine, dependencies MiddlewareDependencies) {
	router.Use(cors.Default())
	router.Use(gin.Recovery(), middlewares.Logger(dependencies.LogRepository))
	router.Use(appErrors.HandleErr)
	router.Use(middlewares.ObserveStats(dependencies.MetricRepository))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", h.Register)
		userGroup.POST("/login", h.LogIn)
		authGroup := userGroup.Group("", middlewares.TokenDecoderMiddleware(dependencies.Forger, dependencies.CacheRepository))
		{
			authGroup.GET("/logout", h.LogOut)
			authGroup.GET("", h.GetUserInfo)
			authGroup.PUT("", h.ChangeUserInfo)
			authGroup.DELETE("", h.DeleteUser)
		}
	}
}
