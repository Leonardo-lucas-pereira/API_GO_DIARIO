package migrations

import (
	"api/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Login{})
}
