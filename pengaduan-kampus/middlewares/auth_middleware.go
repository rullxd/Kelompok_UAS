package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"pengaduan-kampus/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
		token, _ := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("role", claims["role"])
		c.Next()
	}
}
