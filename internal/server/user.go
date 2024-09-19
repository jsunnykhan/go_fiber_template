package server

import "github.com/gofiber/fiber/v2"

func (s *FiberServer) GetUsers(ctx *fiber.Ctx) error {
	users := s.db.FindUsers()

	s.cache.Set("user", "This is set all route")

	return ctx.JSON(users)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}
