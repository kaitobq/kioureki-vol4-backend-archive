package middleware

import (
	"backend/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(st *service.TokenService) gin.HandlerFunc{
	return func(c *gin.Context) {
		valid, err := st.TokenValid(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
