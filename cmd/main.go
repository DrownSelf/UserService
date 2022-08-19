package main

import (
	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/handlers"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/configs"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	config, err := configs.LoadConnectionConfig()
	if err != nil {
		log.Fatalf("error during reading config: %v", err)
	}

	userRepo, err := repositories.NewUserRepository(config)
	defer userRepo.DestroyRepository()
	if err != nil {
		log.Fatalf("error during connect DB: %v", err)
	}

	logRepo, err := repositories.NewLogRepo(context.Background(), config)
	if err != nil {
		log.Fatalf("error during connect DB: %v", err)
	}
	router := gin.New()

	var service services.IUserService = services.New(userRepo, auth.NewJwt(config.Secret), &auth.Hasher{})
	handler := handlers.New(service, validator.New())
	handler.InitRoutes(router, logRepo)
	err = router.Run(":" + config.ServerPort)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
