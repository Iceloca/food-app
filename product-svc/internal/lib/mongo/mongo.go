package mongo

import (
	"context"
	"fmt"
	"github.com/r1nb0/food-app/product-svc/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

func InitClient(cfg *config.Config) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%s", cfg.MongoDB.Host, cfg.MongoDB.Port)

	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
