package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	v1 "github.com/DrownSelf/UserService/internal/api/v1"
	"github.com/DrownSelf/UserService/internal/appErrors"
	"github.com/DrownSelf/UserService/internal/handlers"
	"github.com/DrownSelf/UserService/internal/middlewares"
	"github.com/DrownSelf/UserService/internal/repositories"
)

type ApiGroup struct {
	handler     *handlers.Handler
	ApiV1       *v1.ApiV1
	logRepo     repositories.ILogRepo
	metricsRepo *repositories.PrometheusRepository
}

func NewApiGroup(handler *handlers.Handler, apiV1 *v1.ApiV1, logRepo repositories.ILogRepo, metricsRepo *repositories.PrometheusRepository) *ApiGroup {
	return &ApiGroup{handler: handler, ApiV1: apiV1, logRepo: logRepo, metricsRepo: metricsRepo}
}

func (a *ApiGroup) InitRouterGroups(router *gin.Engine) {
	router.Use(cors.Default())
	router.Use(gin.Recovery(), middlewares.Logger(a.logRepo))
	router.Use(appErrors.HandleErr)
	router.Use(middlewares.ObserveStats(a.metricsRepo))
	api := router.Group("/api")
	{
		a.ApiV1.InitApiV1Groups(api)
	}
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	pprof.Register(router)
	pprof.RouteRegister(api, "pprof")
}
