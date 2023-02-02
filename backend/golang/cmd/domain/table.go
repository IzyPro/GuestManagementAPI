package domain

import (
	"context"

	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Capacity int `json:"capacity"`
}

type CreateTableRequest struct {
	Capacity int `json:"capacity" validate:"required,numeric,gt=0"`
}

// Table Repository Interface
type TableRepository interface {
	Create(ctx context.Context, table *Table) (*Table, error)
	Get(ctx context.Context, id int) (*Table, error)
	GetAll(ctx context.Context) (*[]Table, error)
}
