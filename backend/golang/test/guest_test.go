package test

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber"
)

func TestGetAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/guest_list", nil)
	res := ExecuteRequest(req)

	CheckResponseCode(t, res.StatusCode, fiber.StatusOK)
}

func TestGetCheckedIn(t *testing.T) {
	req, _ := http.NewRequest("GET", "/guests", nil)
	res := ExecuteRequest(req)

	CheckResponseCode(t, res.StatusCode, fiber.StatusOK)
}

func TestDeleteWithInvalidGuest(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/guests/albert", nil)
	res := ExecuteRequest(req)

	CheckResponseCode(t, res.StatusCode, fiber.StatusBadRequest)
}

func TestGetEmptySeats(t *testing.T) {
	req, _ := http.NewRequest("GET", "/seats_empty", nil)
	res := ExecuteRequest(req)

	CheckResponseCode(t, res.StatusCode, fiber.StatusOK)
}
