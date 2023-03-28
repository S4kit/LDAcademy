package main

import (
	"GoAssig/GoAssig/Controllers"
	"GoAssig/GoAssig/Initializers"
	"GoAssig/GoAssig/Middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	Initializers.ConnectToDb()

	router := gin.Default()

	router.POST("/signup", Controllers.Signup)
	router.POST("/login", Controllers.SignIN)

	router.GET("/shops", Controllers.GetAllShops)
	router.GET("/shops/:id", Controllers.GetShopByID)
	router.POST("/shops", Middleware.Auth(), Controllers.CreateShop)
	router.PUT("/shops/:id", Middleware.Auth(), Controllers.UpdateShop)
	router.DELETE("/shops/:id", Middleware.Auth(), Controllers.DeleteShop)

	router.GET("/products", Controllers.GetAllProducts)
	router.GET("/products/:id", Controllers.GetProductByID)
	router.POST("/products", Middleware.Auth(), Controllers.CreateProduct)
	router.PUT("/products/:id", Middleware.Auth(), Controllers.UpdateProduct)
	router.DELETE("/products/:id", Middleware.Auth(), Controllers.DeleteProduct)

	router.GET("/categories", Controllers.GetAllCategories)

	router.Run("localhost:8080")
}
