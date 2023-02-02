package rest

import (
	"github.com/getground/tech-tasks/backend/cmd/domain"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	GuestRepo domain.GuestRepository
	TableRepo domain.TableRepository
}

func RunHttpServer(config Config) *fiber.App {
	app := fiber.New()
	InitRouter(app, config)
	return app
}
