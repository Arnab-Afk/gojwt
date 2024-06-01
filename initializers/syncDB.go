package initializers

import (
	"gopro/models"
)

func SyncDB() {
	// Sync database here
	DB.AutoMigrate(&models.User{})
}
