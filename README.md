# Order-service
***
The service provides functionality to work with orders.

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
