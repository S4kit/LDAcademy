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
	router.POST("/shops", Middleware.VerifyAuth(), Controllers.CreateShop)
	router.PUT("/shops/:id", Middleware.VerifyAuth(), Controllers.UpdateShop)
	router.DELETE("/shops/:id", Middleware.VerifyAuth(), Controllers.DeleteShop)

	router.GET("/products", Controllers.GetAllProducts)
	router.GET("/products/:id", Controllers.GetProductByID)
	router.POST("/products", Middleware.VerifyAuth(), Controllers.CreateProduct)
	router.PUT("/products/:id", Middleware.VerifyAuth(), Controllers.UpdateProduct)
	router.DELETE("/products/:id", Middleware.VerifyAuth(), Controllers.DeleteProduct)

	router.GET("/categories", Controllers.GetAllCategories)

	router.Run("localhost:8080")
}
