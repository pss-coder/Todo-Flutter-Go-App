package middlewares

import (
	"todo-backend/utils"

	"github.com/gin-gonic/gin"
)

// func returns a gin.HandlerFunc, used to process http request and response in gin web app
func IsAuthorised() gin.HandlerFunc {

	// func retrieve token from cookie from incoming HTTP request
	// and tries to parse it
	// if missing or parsing fails, it returns 401 unauthorized
	// if successful, it sets the role in the context,
	// and calls the next handler to continue processing request in the next middleware or handler in middleware chain
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Next()
	}
}
