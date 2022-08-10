package main

import (
	auth2 "InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/handlers"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/configs"
	"database/sql"
	"github.com/gin-contrib/cors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	connectionString, err := configs.LoadConnectionConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("%v", err)
	}

	secret, err := configs.LoadSecretConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	router := gin.Default()
	var repo repositories.IUserRepository = &repositories.UserRepository{db}
	var service services.IUserService = services.New(repo, auth2.NewJwt(secret), &auth2.Hasher{}, validator.New())

	handler := handlers.New(service)
	handler.InitRoutes(router)
	router.Use(cors.Default())
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
