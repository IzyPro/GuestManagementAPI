package rest

import (
	"github.com/getground/tech-tasks/backend/cmd/delivery/rest/guest"
	"github.com/getground/tech-tasks/backend/cmd/delivery/rest/table"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(r *fiber.App, config Config) {
	guestRepo := config.GuestRepo
	tableRepo := config.TableRepo

	guestRouter := r.Group("/")
	tableRouter := r.Group("/tables")

	guest.New(guestRouter, tableRepo, guestRepo)
	table.New(tableRouter, tableRepo)

	// ping
	guestRouter.Get("ping", HandlerPing)
}

func HandlerPing(c *fiber.Ctx) error {
	return c.SendString("pong\n")
}
