# User Service
***
The service provides functionality to work with user requests. It works with order(by grpc) 
and analytics(by kafka) service. All deployed in docker.

# External requirements
***
    Go 1.18
    Docker
    Docker-compose
    Goose
    Godotenv
    pq
    client-golang/prometheus
    mongo-driver/mongo
    go-redis/redis
    go-redis/cache
    swaggo/swag
    github.com/DrownSelf/OrderService
    github.com/DrownSelf/AnalyticsService

    


## Configuration

The service could be configured by providing environment variables.

| Name               | Meaning                      | Example                                                             |
|--------------------|------------------------------|---------------------------------------------------------------------|
| PG_SOURCE          | Postrgres connection string  | postgresql://postgres:password@postgres:5432/userDb?sslmode=disable |
| MONGODB_CONNSTRING | Mongo connections string     | mongodb://mongo:passsword@mongo:27017/?authSource=logDb             |
| SERVERPORT         | server port                  | 8080                                                                |
| SECRET             | secret for JWT Token         | YourSecret                                                          |
| REDIS_HOST         | Host for redis               | redis:6379                                                          |
| REDIS_PASSWORD     | Password for redis           | password                                                            |
| EXPTIME            | Expiration time of JWT Token | 1h                                                                  |
| GRPC_CLIENT        | Connection to grpc server    | 127.0.0.1:8082                                                      |
| KAFKA_TOPIC        | Topic for kafka              | allnamewhichyouwant                                                 |
| KAFKA_CONNECTION   | Connection for kafka         | localhost:9092                                                      |
