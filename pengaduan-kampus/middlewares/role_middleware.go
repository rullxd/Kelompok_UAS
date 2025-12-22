package middlewares

import "github.com/gin-gonic/gin"

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != "admin" {
			c.AbortWithStatus(403)
			return
		}
		c.Next()
	}
}
