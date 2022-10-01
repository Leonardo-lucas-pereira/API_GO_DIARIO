package middlewares

import (
	"api/services"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Beare_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Beare_schema):]

		if !services.NewJWTServices().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
