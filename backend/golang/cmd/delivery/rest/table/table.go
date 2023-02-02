package table

import (
	"context"

	"github.com/getground/tech-tasks/backend/cmd/domain"
	"github.com/getground/tech-tasks/backend/cmd/helpers"
	"github.com/gofiber/fiber/v2"
)

type TableHandler struct {
	TableRepo domain.TableRepository
}

func (h *TableHandler) Create(c *fiber.Ctx) error {
	res, err := helpers.ParseAndValidate(c, new(domain.CreateTableRequest))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	model := res.(*domain.CreateTableRequest)
	var table domain.Table = domain.Table{
		Capacity: model.Capacity,
	}
	result, err := h.TableRepo.Create(context.TODO(), &table)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"id":       result.ID,
		"capacity": result.Capacity,
	})
}

func New(tableRouter fiber.Router, repo domain.TableRepository) {
	handler := &TableHandler{
		TableRepo: repo,
	}

	tableRouter.Post("/", handler.Create)
}
