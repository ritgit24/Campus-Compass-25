package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SetupRoutes(r *gin.Engine) {
	// Protected profile routes
	profileGroup := r.Group("/api/profile")
	profileGroup.Use(AuthMiddleware())
	
	// Current user's profile
	profileGroup.GET("/", GetProfileHandler)
	
	// Specific profile by ID
	profileGroup.GET("/:id", GetProfileByIDHandler)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // Use your actual secret key
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store claims in context
		c.Set("claims", token.Claims)
		c.Next()
	}
}