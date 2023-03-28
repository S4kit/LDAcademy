package Controllers

import (
	"GoAssig/GoAssig/Models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllShops(c *gin.Context) {

	AllShops, err := Models.ShowAllShops()
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, AllShops)
	return
}
func GetShopByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ShopID, err := Models.ShowShop(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, ShopID)
	return
}
func CreateShop(c *gin.Context) {
	var shops Models.Shop
	if err := c.BindJSON(&shops); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	result, err := Models.AddShop(shops)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.IndentedJSON(http.StatusOK, result)
}
func UpdateShop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var shops Models.Shop
	if err := c.BindJSON(&shops); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := Models.EditShop(id, shops)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
func DeleteShop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := Models.DeleteShop(id)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
