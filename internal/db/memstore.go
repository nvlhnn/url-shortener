package db

import (
	"context"

	"github.com/gofiber/storage/redis/v3"
	"github.com/nvlhnn/url-shortener/internal/config"
)

var Ctx = context.Background()

// redis instance for rate limiting
func CreateClient(config *config.MemStore) *redis.Storage {

	store := redis.New(redis.Config{
		Host:      config.Host,
		Port:      config.Port,
		Password:  config.Password,
		Database:  config.Database,
	})
	
	return store
}