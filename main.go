package main

import (
	"wallet-tennet/controllers"
	"wallet-tennet/initializers"
	"wallet-tennet/middleware"
	"wallet-tennet/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	//gin for route
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Login)

	v1 := r.Group("/v1")
	{
		v1.Use(middleware.RequireAuth)
		routes.WalletRoutes(v1.Group("/wallet"))
		routes.AssetRoutes(v1.Group("/asset"))
		routes.TransactionRoutes(v1.Group("/transaction"))
	}

	r.Run()
}
