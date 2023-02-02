package repository

import (
	"context"
	"fmt"

	"github.com/getground/tech-tasks/backend/cmd/domain"
	"gorm.io/gorm"
)

type guestRepository struct {
	Db *gorm.DB
}

func NewGuestRepository(db *gorm.DB) domain.GuestRepository {
	return guestRepository{
		Db: db,
	}
}

func (repo guestRepository) GetAll(ctx context.Context) (*[]domain.Guest, error) {
	var guests []domain.Guest
	//.Model(&Guest{}) before preload
	if result := repo.Db.Preload("Table").Find(&guests); result.Error != nil {
		return nil, result.Error
	}
	return &guests, nil
}

func (repo guestRepository) GetCheckedIn(ctx context.Context) (*[]domain.Guest, error) {
	var guests []domain.Guest
	if result := repo.Db.Preload("Table").Where("checked_in = ?", true).Find(&guests); result.Error != nil {
		return nil, result.Error
	}
	return &guests, nil
}

func (repo guestRepository) Create(ctx context.Context, guest *domain.Guest) (*domain.Guest, error) {
	if _, err := repo.GetByName(ctx, guest.Name); err == nil {
		return nil, fmt.Errorf("guest %s already exists", guest.Name)
	}
	if result := repo.Db.Create(&guest); result.Error != nil {
		return nil, result.Error
	}
	return guest, nil
}

func (repo guestRepository) GetByName(ctx context.Context, name string) (*domain.Guest, error) {
	var guest domain.Guest
	if result := repo.Db.Preload("Table").Where("name = ?", name).First(&guest); result.Error != nil {
		return nil, result.Error
	}
	return &guest, nil
}

func (repo guestRepository) Update(ctx context.Context, guest *domain.Guest) (*domain.Guest, error) {
	if result := repo.Db.Save(&guest); result.Error != nil {
		return nil, result.Error
	}
	return guest, nil
}

func (repo guestRepository) Delete(ctx context.Context, name string) (*domain.Guest, error) {
	var guests domain.Guest
	if result := repo.Db.Where("name = ?", name).Delete(&guests); result.Error != nil {
		return nil, result.Error
	}
	return &guests, nil
}
