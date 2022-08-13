package main

import (
	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/handlers"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/configs"
	"github.com/gin-contrib/cors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	config, err := configs.LoadConnectionConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	router := gin.Default()
	repo, err := repositories.NewUserRepository(config)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var service services.IUserService = services.New(repo, auth.NewJwt(config.Secret), &auth.Hasher{})

	handler := handlers.New(service, validator.New())
	handler.InitRoutes(router)
	router.Use(cors.Default())
	err = router.Run(":" + config.ServerPort)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
