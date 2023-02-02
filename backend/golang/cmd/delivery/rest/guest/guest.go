package guest

import (
	"context"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/getground/tech-tasks/backend/cmd/domain"
	"github.com/getground/tech-tasks/backend/cmd/helpers"
	"github.com/gofiber/fiber/v2"
)

type GuestHandler struct {
	GuestRepo domain.GuestRepository
	TableRepo domain.TableRepository
}

func (h *GuestHandler) GetAllGuests(c *fiber.Ctx) error {
	guests, err := h.GuestRepo.GetAll(context.TODO())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var response []domain.GuestListResponse
	linq.From(*guests).Select(func(item interface{}) interface{} {
		return domain.GuestListResponse{
			Name:               item.(domain.Guest).Name,
			Table:              int(item.(domain.Guest).Table.ID),
			AccompanyingGuests: item.(domain.Guest).AccompanyingGuests,
		}
	}).ToSlice(&response)
	return c.JSON(response)
}

func (h *GuestHandler) AddGuest(c *fiber.Ctx) error {
	guestName := c.Params("name")
	res, err := helpers.ParseAndValidate(c, new(domain.AddGuestRequest))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	model := res.(*domain.AddGuestRequest)
	table, err := h.TableRepo.Get(context.TODO(), model.Table)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Table ID",
		})
	}
	var guest domain.Guest = domain.Guest{
		Name:               guestName,
		AccompanyingGuests: model.AccompanyingGuests,
		Table:              *table,
		TableID:            table.ID,
	}
	result, err := h.GuestRepo.Create(context.TODO(), &guest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"name": result.Name,
	})
}

func (h *GuestHandler) CheckIn(c *fiber.Ctx) error {
	guestName := c.Params("name")
	res, err := helpers.ParseAndValidate(c, new(domain.GuestCheckInRequest))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	model := res.(*domain.GuestCheckInRequest)

	count, err := h.calculateAvailability(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else if count < model.AccompanyingGuests+1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Not enough room for guests",
		})
	}

	guest, err := h.GuestRepo.GetByName(context.TODO(), guestName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else if guest.Table.Capacity <= model.AccompanyingGuests {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Number of Accompanying Guest is above Table Capacity",
		})
	}
	guest.AccompanyingGuests = model.AccompanyingGuests
	guest.CheckedIn = true
	guest.TimeArrived = time.Now()
	result, err := h.GuestRepo.Update(context.TODO(), guest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"name": result.Name,
	})
}

func (h *GuestHandler) GetCheckedIn(c *fiber.Ctx) error {
	result, err := h.GuestRepo.GetCheckedIn(context.TODO())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var response []domain.GuestCheckInResponse
	linq.From(*result).Select(func(item interface{}) interface{} {
		return domain.GuestCheckInResponse{
			Name:               item.(domain.Guest).Name,
			AccompanyingGuests: item.(domain.Guest).AccompanyingGuests,
			TimeArrived:        item.(domain.Guest).TimeArrived,
		}
	}).ToSlice(&response)
	return c.JSON(response)
}

func (h *GuestHandler) CheckOut(c *fiber.Ctx) error {
	guestName := c.Params("name")
	guest, err := h.GuestRepo.GetByName(context.TODO(), guestName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else if !guest.CheckedIn {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot check out a guest that was never checked in",
		})
	}
	if _, err := h.GuestRepo.Delete(context.TODO(), guestName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(nil)
}

func (h *GuestHandler) EmptySeats(c *fiber.Ctx) error {
	count, err := h.calculateAvailability(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"seats_empty": count,
	})
}

func (h *GuestHandler) calculateAvailability(c *fiber.Ctx) (int, error) {
	var noOfTables, noOfOccupiedTables int
	tables, e := h.TableRepo.GetAll(context.TODO())
	if e != nil {
		return 0, e
	}
	guests, err := h.GuestRepo.GetCheckedIn(context.TODO())
	if err != nil {
		return 0, err
	}

	for _, val := range *tables {
		noOfTables += val.Capacity
	}
	for _, val := range *guests {
		noOfOccupiedTables += val.AccompanyingGuests + 1
	}
	return noOfTables - noOfOccupiedTables, nil
}

func New(guestRouter fiber.Router, tableRepo domain.TableRepository, repo domain.GuestRepository) {
	handler := &GuestHandler{
		GuestRepo: repo,
		TableRepo: tableRepo,
	}

	guestRouter.Get("guest_list", handler.GetAllGuests)
	guestRouter.Get("guests", handler.GetCheckedIn)
	guestRouter.Get("seats_empty", handler.EmptySeats)
	guestRouter.Post("guest_list/:name", handler.AddGuest)
	guestRouter.Put("guests/:name", handler.CheckIn)
	guestRouter.Delete("guests/:name", handler.CheckOut)
}
