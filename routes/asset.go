package routes

import (
	"wallet-tennet/controllers"

	"github.com/gin-gonic/gin"
)

func AssetRoutes(g *gin.RouterGroup) {
	g.POST("/", controllers.CreateAsset)
	g.GET("/", controllers.GetAsset)
	g.GET("/:id", controllers.FindAsset)
	g.PUT("/:id", controllers.UpdateAsset)
	g.DELETE("/:id", controllers.DeleteAsset)
}
