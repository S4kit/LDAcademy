package Controllers

import (
	"GoAssig/GoAssig/Models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {

	AllProducts, err := Models.ShowAllProducts()
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, AllProducts)
	return
}
func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ProdID, err := Models.ShowProduct(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, ProdID)
	return
}
func CreateProduct(c *gin.Context) {
	var prd Models.Product
	if err := c.BindJSON(&prd); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	result, err := Models.AddProduct(prd)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.IndentedJSON(http.StatusOK, result)
}
func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var prd Models.Product
	if err := c.BindJSON(&prd); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := Models.EditProduct(id, prd)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := Models.DeleteProduct(id)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
