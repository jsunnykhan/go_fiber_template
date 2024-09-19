package server

import (
	"github.com/gofiber/fiber/v2"
	"go_fiber_template/internal/redis"

	"go_fiber_template/internal/database"
)

type FiberServer struct {
	*fiber.App
	db    database.Service
	cache redis.RedisClient
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go_fiber_template",
			AppName:      "go_fiber_template",
		}),

		db:    database.New(),
		cache: redis.NewRedisClient(),
	}

	return server
}
