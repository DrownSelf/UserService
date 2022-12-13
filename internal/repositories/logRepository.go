package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	configs "github.com/DrownSelf/UserService/internal/config"
	"github.com/DrownSelf/UserService/internal/entities"
)

type ILogRepo interface {
	ReportLog(ctx context.Context, log entities.Log) error
}

type LogRepo struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewLogRepo(ctx context.Context, config *configs.Config) (*LogRepo, error) {
	clientOptions := options.Client().ApplyURI(config.MongoDbConnSting)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	collection := client.Database("logDb").Collection("logs")
	return &LogRepo{collection, client}, nil
}

func (r *LogRepo) DestroyRepo(ctx context.Context) error {
	err := r.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *LogRepo) ReportLog(ctx context.Context, log entities.Log) error {
	_, err := r.collection.InsertOne(ctx, log)
	if err != nil {
		return err
	}
	return err
}
