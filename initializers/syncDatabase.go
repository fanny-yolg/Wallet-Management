package initializers

import "wallet-tennet/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Wallet{})
	DB.AutoMigrate(&models.Asset{})
	DB.AutoMigrate(&models.Transaction{})
}
