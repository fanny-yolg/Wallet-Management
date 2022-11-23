package controllers

import (
	"wallet-tennet/initializers"
	"wallet-tennet/models"

	"github.com/gin-gonic/gin"
)

func CreateWallet(c *gin.Context) {
	var body struct {
		Name string
	}
	c.Bind(&body)

	// Create a wallet
	wallet := models.Wallet{Name: body.Name}
	result := initializers.DB.Create(&wallet)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create asset",
		})
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"message": "Succesfully create Wallet",
		"data":    wallet,
	})
}

func GetWallet(c *gin.Context) {
	// Get the wallets
	var wallets []models.Wallet
	initializers.DB.Find(&wallets)

	//Respond with them
	c.JSON(200, gin.H{
		"message": "Succesfully get wallets",
		"data":    wallets,
	})
}

func FindWallet(c *gin.Context) {
	//Get ID from URL
	id := c.Param("id")

	var wallet []models.Wallet
	find := initializers.DB.First(&wallet, id)

	if find.Error != nil {
		c.JSON(404, gin.H{
			"message": "We cannot find the data you were looking",
		})
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"message": "Succesfully get wallet",
		"data":    wallet,
	})
}

func UpdateWallet(c *gin.Context) {
	//Get the id from the url
	id := c.Param("id")

	//Get the data off the body
	var body struct {
		Name string
	}
	c.Bind(&body)

	//Find the post we're updating
	var wallet []models.Wallet
	find := initializers.DB.First(&wallet, id)

	if find.Error != nil {
		c.JSON(400, gin.H{
			"message": "Cannot find the data you were looking",
		})
		return
	}
	//Update it
	initializers.DB.Model(&wallet).Updates(models.Wallet{
		Name: body.Name,
	})

	//Return it
	c.JSON(200, gin.H{
		"message": "Succesfully udpate Wallet",
		"data":    wallet,
	})
}

func DeleteWallet(c *gin.Context) {
	//Get the id from URL
	id := c.Param("id")

	// Delete the posts
	find := initializers.DB.First(&models.Wallet{}, id)
	if find.Error != nil {
		c.JSON(400, gin.H{
			"message": "Cannot find the data you were looking",
		})
		return
	}
	initializers.DB.Delete(&models.Wallet{}, id)

	//Respond
	c.JSON(200, gin.H{
		"message": "Deleted successfully",
	})
}
