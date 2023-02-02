package main

import (
	"log"

	"github.com/getground/tech-tasks/backend/cmd/delivery/rest"
	"github.com/getground/tech-tasks/backend/cmd/repository"
)

func main() {
	// Initialize DB and Repo
	repo := repository.New()

	// DI COnfiguration
	httpConfig := rest.Config{
		GuestRepo: repo.GuestRepo,
		TableRepo: repo.TableRepo,
	}

	// Congigure Routes
	app := rest.RunHttpServer(httpConfig)

	// Run
	log.Fatal(app.Listen(":3000"))
}
