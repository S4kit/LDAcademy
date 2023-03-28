package Middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	secretKey = []byte(os.Getenv("SECRET_TOKEN"))
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		prefix := "Bearer "
		header := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(header, prefix)

		if header == "" || tokenString == header {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "You need to be connected."})
			return
		}

		return
	}
}
