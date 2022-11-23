package routes

import (
	"wallet-tennet/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(g *gin.RouterGroup) {
	g.POST("/", controllers.CreateTransaction)
}
