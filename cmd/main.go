package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/handlers"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/configs"
)

func main() {
	config, err := configs.LoadConnectionConfig()
	if err != nil {
		log.Fatalf("error during reading config: %v", err)
	}

	userRepo, err := repositories.NewUserRepo(config)
	if err != nil {
		log.Fatalf("error during connect DB: %v", err)
	}

	logRepo, err := repositories.NewLogRepo(context.Background(), config)
	if err != nil {
		log.Fatalf("error during connect DB: %v", err)
	}

	cacheRepo := repositories.NewCacheRepo(*config)

	router := gin.New()
	metricsRepo := repositories.NewMetricsRepo(router)
	tokenForger := auth.NewJwt(config.Secret)
	service := services.NewUserService(userRepo, cacheRepo, tokenForger, &auth.Hasher{}, config)
	handler := handlers.New(service)
	handler.InitRoutes(router, handlers.MiddlewareDependencies{
		LogRepository:    logRepo,
		Forger:           tokenForger,
		CacheRepository:  cacheRepo,
		MetricRepository: metricsRepo,
	})

	srv := &http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error during creating server: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = userRepo.DestroyRepository(ctx); err != nil {
		log.Printf("Error during shutdown db: %s", err)
	}

	if err = logRepo.DestroyRepo(ctx); err != nil {
		log.Printf("Error during shutdown db: %s", err)
	}

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 sec")

	log.Println("Server exiting")
}
