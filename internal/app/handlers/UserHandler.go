package handlers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"InnoTaxi/internal/app/appErrors"
	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/middlewares"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/dto"
)

type Handler struct {
	userService services.IUserService
}

func New(service services.IUserService) Handler {
	return Handler{userService: service}
}

func (h *Handler) Register(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	id, err := h.userService.RegisterUser(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, id)
}

func (h *Handler) LogIn(ctx *gin.Context) {
	var user dto.LogInUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	response, err := h.userService.LogInUser(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var user dto.ChangeUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	if err := h.userService.UpdateUser(ctx, user); err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Updated successfully")
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	var id dto.DeleteUserRequest
	if err := ctx.ShouldBindJSON(&id); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	if err := h.userService.DeleteUser(ctx, id.Id); err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Deleted successfully")
}

func (h *Handler) GetUser(ctx *gin.Context) {
	var phoneNumber dto.GetUserInfoRequest
	if err := ctx.ShouldBindJSON(&phoneNumber); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	user, err := h.userService.GetUserByPhone(ctx, phoneNumber.PhoneNumber)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	userResponse := dto.GetUserResponse{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Rating:      user.Rating,
	}
	ctx.JSON(http.StatusOK, userResponse)
}

func (h *Handler) LogOut(ctx *gin.Context) {
	if err := h.userService.LogOutUser(ctx, ctx.GetHeader("Authorization")); err != nil {
		_ = ctx.Error(err)
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

	v1 := router.Group("/api/v1")
	{
		userGroup := v1.Group("/user")
		{
			userGroup.POST("/register", h.Register)
			userGroup.POST("/login", h.LogIn)
			authGroup := userGroup.Group("", middlewares.TokenDecoderMiddleware(dependencies.Forger, dependencies.CacheRepository))
			{
				authGroup.GET("/logout", h.LogOut)
				authGroup.GET("", h.GetUser)
				authGroup.PUT("", h.UpdateUser)
				authGroup.DELETE("", h.DeleteUser)
			}
		}
	}
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	pprof.Register(router)
	pprof.RouteRegister(v1, "pprof")
}
