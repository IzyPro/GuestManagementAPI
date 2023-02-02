package test

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber"
)

func TestPing(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	res := ExecuteRequest(req)

	CheckResponseCode(t, res.StatusCode, fiber.StatusOK)
}
