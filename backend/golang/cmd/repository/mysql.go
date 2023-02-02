package repository

import (
	"log"

	"github.com/getground/tech-tasks/backend/cmd/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	GuestRepo domain.GuestRepository
	TableRepo domain.TableRepository
}

func New() *Repository {
	//db, err := gorm.Open(mysql.Open("user:password@/getground?parseTime=true"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open("Izy:P@ssw0rd@/getground?parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Auto Migration
	db.AutoMigrate(&domain.Guest{}, &domain.Table{})

	// Return all Repository Implementations
	return &Repository{
		GuestRepo: NewGuestRepository(db),
		TableRepo: NewTableRepository(db),
	}
}
