package repository

import (
	"context"

	"github.com/getground/tech-tasks/backend/cmd/domain"
	"gorm.io/gorm"
)

type tableRepository struct {
	Db *gorm.DB
}

func NewTableRepository(db *gorm.DB) domain.TableRepository {
	return tableRepository{
		Db: db,
	}
}

func (repo tableRepository) Create(ctx context.Context, table *domain.Table) (*domain.Table, error) {
	if result := repo.Db.Create(&table); result.Error != nil {
		return nil, result.Error
	}
	return table, nil
}

func (repo tableRepository) Get(ctx context.Context, id int) (*domain.Table, error) {
	var table *domain.Table
	if result := repo.Db.First(&table, id); result.Error != nil {
		return nil, result.Error
	}
	return table, nil
}

func (repo tableRepository) GetAll(ctx context.Context) (*[]domain.Table, error) {
	var tables *[]domain.Table
	if result := repo.Db.Find(&tables); result.Error != nil {
		return nil, result.Error
	}
	return tables, nil
}
