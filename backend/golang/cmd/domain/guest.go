package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	Name               string `json:"name"`
	AccompanyingGuests int    `json:"accompanying_guests"`
	TableID            uint
	Table              Table     `json:"table"`
	CheckedIn          bool      `json:"checked_in"`
	TimeArrived        time.Time `json:"time_arrived"`
}

type AddGuestRequest struct {
	Table              int `json:"table" validate:"required,numeric,gt=0"`
	AccompanyingGuests int `json:"accompanying_guests" validate:"required,numeric,gt=0"`
}

type GuestCheckInRequest struct {
	AccompanyingGuests int `json:"accompanying_guests" validate:"required,numeric,gt=0"`
}

type GuestCheckInResponse struct {
	Name               string    `json:"name"`
	AccompanyingGuests int       `json:"accompanying_guests"`
	TimeArrived        time.Time `json:"time_arrived"`
}

type GuestListResponse struct {
	Name               string `json:"name"`
	Table              int    `json:"table"`
	AccompanyingGuests int    `json:"accompanying_guests"`
}

// Guest Repository Interface
type GuestRepository interface {
	GetAll(ctx context.Context) (*[]Guest, error)
	GetCheckedIn(ctx context.Context) (*[]Guest, error)
	GetByName(ctx context.Context, name string) (*Guest, error)
	Create(ctx context.Context, guest *Guest) (*Guest, error)
	Update(ctx context.Context, guest *Guest) (*Guest, error)
	Delete(ctx context.Context, name string) (*Guest, error)
}
