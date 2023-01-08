package initializers

import "github.com/vermaexe/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Post{})
}
