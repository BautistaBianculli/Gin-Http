package middleware

import (
	"GORUTINE/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.ValidationToken(c) 
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized: " + err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
