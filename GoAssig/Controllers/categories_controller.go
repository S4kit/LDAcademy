package Controllers

import (
	"GoAssig/GoAssig/Models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {

	AllCat, err := Models.ShowAllCat()
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, AllCat)
	return
}
