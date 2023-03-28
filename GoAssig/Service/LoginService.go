package Service

import (
	"GoAssig/GoAssig/Models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte(os.Getenv("SECRET_TOKEN"))
)

func CreateToken(user Models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
