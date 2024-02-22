package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/redis/v3"
	"github.com/nvlhnn/url-shortener/internal/config"
	"github.com/nvlhnn/url-shortener/internal/db"
)

type LimiterMiddleware struct {
	config *config.Limiter
	redis  *redis.Storage
}

func NewLimiterMiddleware(configLimit *config.Limiter, configRedis *config.MemStore) *LimiterMiddleware {
	return &LimiterMiddleware{
		config: configLimit,
		redis:  db.CreateClient(configRedis),
	}
}

func (l *LimiterMiddleware) LimiterMiddleware() func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:          l.config.Max,
		Expiration:   time.Duration(l.config.Expiration) * time.Second,
		LimitReached: l.limitClient,
		Storage: l.redis,
	})
}

func (l *LimiterMiddleware) limitClient(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"message": "Too many requests",
	})
}