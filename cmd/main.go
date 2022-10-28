package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/DrownSelf/OrderService/pkg/grpc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"

	_ "github.com/DrownSelf/UserService/cmd/docs"
	"github.com/DrownSelf/UserService/internal/auth"
	config "github.com/DrownSelf/UserService/internal/config"
	"github.com/DrownSelf/UserService/internal/handlers"
	"github.com/DrownSelf/UserService/internal/repositories"
	"github.com/DrownSelf/UserService/internal/services"
)

func main() {
	connectionConfig, err := config.LoadConnectionConfig()
	if err != nil {
		log.Fatalf("error during reading connectionConfig: %v", err)
	}

	userRepo, err := repositories.NewUserRepo(connectionConfig)
	if err != nil {
		log.Fatalf("error during connect DB: %v", err)
	}

	logRepo, err := repositories.NewLogRepo(context.Background(), connectionConfig)
	if err != nil {
		log.Fatalf("error during connect DB: %v", err)
	}

	cacheRepo := repositories.NewCacheRepo(*connectionConfig)
	router := gin.New()
	metricsRepo := repositories.NewMetricsRepo(router)
	tokenForger := auth.NewJwt(connectionConfig.Secret)
	service := services.NewUserService(userRepo, cacheRepo, tokenForger, &auth.Hasher{}, connectionConfig)
	grpcConnection, err := grpc.Dial(connectionConfig.GrpcClient, grpc.WithInsecure(), grpc.WithBlock())
	client := pb.NewOrderServiceClient(grpcConnection)
	if err != nil {
		log.Fatalf("error during setup GRPC: %v", err)
	}

	handler := handlers.New(service, client)
	handler.InitRoutes(router, handlers.MiddlewareDependencies{
		LogRepository:    logRepo,
		Forger:           tokenForger,
		CacheRepository:  cacheRepo,
		MetricRepository: metricsRepo,
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	srv := &http.Server{
		Addr:    ":" + connectionConfig.ServerPort,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error during creating server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
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
