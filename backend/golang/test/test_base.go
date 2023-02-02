package test

import (
	"net/http"
	"testing"

	"github.com/getground/tech-tasks/backend/cmd/delivery/rest"
	"github.com/getground/tech-tasks/backend/cmd/repository"
	"github.com/gofiber/fiber/v2"
)

func setUp() *fiber.App {
	repo := repository.New()

	// DI COnfiguration
	httpConfig := rest.Config{
		GuestRepo: repo.GuestRepo,
		TableRepo: repo.TableRepo,
	}

	// Congigure Routes
	app := rest.RunHttpServer(httpConfig)
	return app
}

func ExecuteRequest(req *http.Request) *http.Response {
	a := setUp()
	req.Header.Set("Content-Type", "application/json")
	resp, _ := a.Test(req, -1)

	//fmt.Println(a.GetRoutes())
	return resp
}

func CheckResponseCode(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("\n-TEST FAILED-\nActual: %d\nExpected:%d", actual, expected)
	}
}
