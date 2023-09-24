package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sd/internal/config"
	"time"
)

func NewClient(ctx context.Context, cfg *config.MongoConfig) (*mongo.Database, error) {
	var mongoDBURL string
	var anonymous bool
	if cfg.Username == "" || cfg.Password == "" {
		anonymous = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
	} else {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoDBURL)
	if !anonymous {
		clientOptions.SetAuth(options.Credential{
			AuthSource:  cfg.AuthSource,
			Username:    cfg.Username,
			Password:    cfg.Password,
			PasswordSet: true,
		})
	}

	client, err := mongo.Connect(reqCtx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к mongodb: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к mongodb: %v", err)
	}

	return client.Database(cfg.Database), nil
}
