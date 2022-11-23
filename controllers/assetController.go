package controllers

import (
	"wallet-tennet/initializers"
	"wallet-tennet/models"

	"github.com/gin-gonic/gin"
)

func CreateAsset(c *gin.Context) {
	//Get data off req body
	var body struct {
		Name      string
		Symbol    string
		Network   string
		Address   string
		Balance   float32
		Wallet_id int
	}
	c.Bind(&body)

	// Create an asset
	asset := models.Asset{
		Name:      body.Name,
		Symbol:    body.Symbol,
		Network:   body.Network,
		Address:   body.Address,
		Balance:   body.Balance,
		Wallet_id: body.Wallet_id,
	}
	result := initializers.DB.Create(&asset)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create asset",
		})
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"message": "Succesfully create asset!",
		"data":    asset,
	})
}

func GetAsset(c *gin.Context) {
	// Get the assets
	var assets []models.Asset
	initializers.DB.Find(&assets)

	//Respond with them
	c.JSON(200, gin.H{
		"message": "Succesfully get assets",
		"data":    assets,
	})
}

func FindAsset(c *gin.Context) {
	//Get ID from URL
	id := c.Param("id")

	var asset []models.Asset
	find := initializers.DB.First(&asset, id)

	if find.Error != nil {
		c.JSON(404, gin.H{
			"message": "We cannot find the data you were looking",
		})
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"message": "Succesfully get asset",
		"data":    asset,
	})
}

func UpdateAsset(c *gin.Context) {
	//Get the id from the url
	id := c.Param("id")

	//Get the data off the body
	var body struct {
		Name      string
		Symbol    string
		Network   string
		Address   string
		Balance   float32
		Wallet_id int
	}
	c.Bind(&body)

	//Find the asset we're updating
	var asset []models.Asset
	find := initializers.DB.First(&asset, id)

	if find.Error != nil {
		c.JSON(404, gin.H{
			"message": "We cannot find the data you were looking",
		})
		return
	}
	//Update it
	initializers.DB.Model(&asset).Updates(models.Asset{
		Name:      body.Name,
		Symbol:    body.Symbol,
		Network:   body.Network,
		Address:   body.Address,
		Balance:   body.Balance,
		Wallet_id: body.Wallet_id,
	})

	//Return it
	c.JSON(200, gin.H{
		"message": "Succesfully udpate asset",
		"data":    asset,
	})

}

func DeleteAsset(c *gin.Context) {
	//Get the id from URL
	id := c.Param("id")

	// Delete the posts
	find := initializers.DB.First(&models.Asset{}, id)
	if find.Error != nil {
		c.JSON(400, gin.H{
			"message": "Cannot find the data you were looking",
		})
		return
	}
	initializers.DB.Delete(&models.Asset{}, id)

	//Respond
	c.JSON(200, gin.H{
		"message": "Deleted successfully",
	})
}
