package mongo

import (
	"context"
	"fmt"
	"github.com/r1nb0/food-app/product-svc/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func InitClient(cfg *config.Config) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%s", cfg.MongoDB.Host, cfg.MongoDB.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	defer cancel()
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
