package Controllers

import (
	"GoAssig/GoAssig/Models"
	"GoAssig/GoAssig/Service"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Infos struct {
	Email    string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	if err != nil {
		return err
	}

	return nil
}

func EmailReg(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func Signup(c *gin.Context) {
	var user Models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Insert your data again"})
		return
	}
	if !EmailReg(user.Email) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Your E-mail format is wrong"})
		return
	}

	user, CheckUser, err := Models.AskUser(user.Email)
	if err != nil {
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
			return
		}
	}
	if CheckUser {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This user does exist"})
		return
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Creation Failed"})
		return
	}
	user.Password = hash
	newID, err := Models.NewUser(user)

	if err != nil {
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Creation Failed"})
			return
		}
	}
	c.IndentedJSON(http.StatusCreated, newID)
	return
}

func SignIN(c *gin.Context) {
	var user Models.User
	var info Infos

	if err := c.BindJSON(&info); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Insert your data again"})
		return
	}
	user, CheckUser, err := Models.AskUser(info.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}
	if !CheckUser {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This user doesn't exist"})
		return
	}
	err = CheckPasswordHash(user.Password, info.Password)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "Signing Failed"})
		return
	}

	tokenString, err := Service.CreateToken(user)

	c.Header("Authorization", "Bearer "+tokenString)
	c.IndentedJSON(http.StatusOK, tokenString)

}
