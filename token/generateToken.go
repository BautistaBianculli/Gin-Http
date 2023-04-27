package token

import (
	"GORUTINE/models"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(user string, password string) (string, error) {
	tokenLife, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"authorized": true,
		"user":       user,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenLife)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func ValidationToken(c *gin.Context) error {

	token := ExtraerToken(c)
	claims := new(models.JWTClaims)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	// if time.Now().After(time.Unix(int64(claims.ExpiresAt), 0)) {
	// 	return fmt.Errorf("ExpirdToken %v", time.Unix(claims.ExpiresAt, 0))
	// }
	if claims.User != "BBIAN" {
		return fmt.Errorf("UnhautorizedUser %s", claims.User)
	}
	return nil
}

func ExtraerToken(c *gin.Context) string {

	token := c.Request.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}
