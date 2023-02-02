package test

import (
	"testing"

	"github.com/getground/tech-tasks/backend/cmd/helpers"
)

func TestContains(t *testing.T) {
	actual := helpers.Contains([]string{"application/json", "test"}, "test")
	expected := true

	if actual != expected {
		t.Errorf("\n-TEST FAILED-\nActual: %t\nExpected:%t", actual, expected)
	}
}
