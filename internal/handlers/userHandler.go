package handlers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/DrownSelf/UserService/internal/appErrors"
	"github.com/DrownSelf/UserService/internal/auth"
	"github.com/DrownSelf/UserService/internal/entities"
	"github.com/DrownSelf/UserService/internal/middlewares"
	"github.com/DrownSelf/UserService/internal/repositories"
	"github.com/DrownSelf/UserService/internal/services"
)

type Handler struct {
	userService services.IUserService
}

func NewHandler(service services.IUserService) Handler {
	return Handler{userService: service}
}

func (h *Handler) Register(ctx *gin.Context) {
	var user entities.User
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
	var user LogInUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	response, err := h.userService.LogInUser(ctx, user.PhoneNumber, user.Password)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var changedUser ChangeUserRequest
	if err := ctx.ShouldBindJSON(&changedUser); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		_ = ctx.Error(appErrors.ErrInvalidToken)
		return
	}
	phoneNumber := user.(auth.TokenClaims).PhoneNumber

	if err := h.userService.UpdateUser(ctx, entities.User{
		Name:        changedUser.NewName,
		Email:       changedUser.NewEmail,
		PhoneNumber: changedUser.NewPhoneNumber,
	}, phoneNumber); err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Updated successfully")
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		_ = ctx.Error(appErrors.ErrInvalidToken)
		return
	}
	id := user.(auth.TokenClaims).Id

	if err := h.userService.DeleteUser(ctx, id); err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Deleted successfully")
}

func (h *Handler) GetUser(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		_ = ctx.Error(appErrors.ErrInvalidToken)
		return
	}
	phoneNumber := user.(auth.TokenClaims).PhoneNumber

	response, err := h.userService.GetUser(ctx, phoneNumber)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	userResponse := GetUserResponse{
		Name:        response.Name,
		PhoneNumber: response.PhoneNumber,
		Email:       response.Email,
		Rating:      response.Rating,
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

func (h *Handler) OrderTaxi(ctx *gin.Context) {
	var orderRequest MakeOrderRequest
	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		_ = ctx.Error(appErrors.ErrInvalidToken)
		return
	}
	phoneNumber := user.(auth.TokenClaims).PhoneNumber

	gottenUser, err := h.userService.GetUser(ctx, phoneNumber)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.userService.MakeOrder(ctx, gottenUser, orderRequest.From, orderRequest.To, orderRequest.TaxiType)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, UserRideResponse{
		Id:                response.OrderId,
		DriverName:        response.Driver.Name,
		DriverPhoneNumber: response.Driver.PhoneNumber,
		TaxiType:          response.Driver.TaxiType,
		From:              response.From,
		To:                response.To,
	})
}

func (h *Handler) RateUserFromOrder(ctx *gin.Context) {
	var rateFromOrder RateUserFromOrderRequest
	if err := ctx.ShouldBindJSON(&rateFromOrder); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	if err := h.userService.UpdateUserRating(ctx, rateFromOrder.PhoneNumber, float64(rateFromOrder.Rating)); err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Rating added")
}

func (h *Handler) RateRide(ctx *gin.Context) {
	var rideRequest RateRideRequest
	if err := ctx.ShouldBindJSON(&rideRequest); err != nil {
		_ = ctx.Error(appErrors.ErrInvalidData)
		return
	}

	err := h.userService.RateRideFromUser(ctx, rideRequest.Rating, rideRequest.Id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, "Rating added successfully")
}

type MiddlewareDependencies struct {
	LogRepository    repositories.ILogRepo
	Forger           auth.TokenAuth
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
		orderGroup := v1.Group("/order")
		{
			orderGroup.POST("/rateUser", h.RateUserFromOrder)
		}
		userGroup := v1.Group("/user")
		{
			userGroup.POST("/register", h.Register)
			userGroup.POST("/login", h.LogIn)
			authGroup := userGroup.Group("", middlewares.TokenDecoderMiddleware(dependencies.Forger, dependencies.CacheRepository))
			{
				authGroup.POST("/orderTaxi", h.OrderTaxi)
				authGroup.POST("/rateOrder", h.RateRide)
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
