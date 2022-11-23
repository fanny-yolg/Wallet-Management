package routes

import (
	"wallet-tennet/controllers"

	"github.com/gin-gonic/gin"
)

func WalletRoutes(g *gin.RouterGroup) {
	g.POST("/", controllers.CreateWallet)
	g.GET("/", controllers.GetWallet)
	g.GET("/:id", controllers.FindWallet)
	g.PUT("/:id", controllers.UpdateWallet)
	g.DELETE("/:id", controllers.DeleteWallet)

}
