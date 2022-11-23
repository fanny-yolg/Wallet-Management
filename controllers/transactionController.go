package controllers

import (
	"wallet-tennet/initializers"
	"wallet-tennet/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var body struct {
		Src_wallet_id  int
		Src_asset_id   int
		Dest_wallet_id int
		Dest_asset_id  int
		Amount         float32
		Gas_fee        float32
		Total          float32
	}
	c.Bind(&body)

	// Create a transaction
	transaction := models.Transaction{Src_wallet_id: body.Src_wallet_id, Src_asset_id: body.Src_asset_id, Dest_wallet_id: body.Dest_wallet_id, Dest_asset_id: body.Dest_asset_id, Amount: body.Amount, Gas_fee: body.Gas_fee, Total: body.Total}
	result := initializers.DB.Create(&transaction)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create asset",
		})
		return
	}

	//Update assets amount
	//Find the source asset and decrease the balance amount
	var asset []models.Asset
	initializers.DB.Where(&models.Asset{ID: body.Src_asset_id, Wallet_id: body.Src_wallet_id}).First(&asset)
	initializers.DB.Model(&asset).Updates(models.Asset{
		Balance: asset[0].Balance - body.Total,
	})

	//Find the destination asset and increase the balance amount
	var asset_dest []models.Asset
	initializers.DB.Where(&models.Asset{ID: body.Dest_asset_id, Wallet_id: body.Dest_wallet_id}).First(&asset_dest)
	initializers.DB.Model(&asset_dest).Updates(models.Asset{
		Balance: asset_dest[0].Balance + body.Amount,
	})

	c.JSON(200, gin.H{
		"message":      "Successfully create transaction between your wallets!",
		"transaction":  transaction,
		"source_asset": asset,
		"dest_asset":   asset_dest,
	})
}
