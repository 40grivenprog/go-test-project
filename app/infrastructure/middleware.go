package infrastructure

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/utils/token"
	"github.com/gin-gonic/gin"
)

// CorrelationIDGenerator is middleware for generating correlation_id
func CorrelationIDGenerator() gin.HandlerFunc {
	return func(c *gin.Context) {
		length := 16

		randomBytes := make([]byte, length)
		_, err := rand.Read(randomBytes)
		if err != nil {
			return
		}

		randomHex := hex.EncodeToString(randomBytes)

		queryParams := c.Request.URL.Query()
		queryParams.Set("correlation_id", randomHex)
		c.Request.URL.RawQuery = queryParams.Encode()

		c.Next()
	}
}

// ErrorHandler is middleware for error handling
func ErrorHandler(logger usecases.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			logger.LogError(ginErr.Error())
		}

		if len(c.Errors) > 0 {
			c.JSON(c.Writer.Status(), c.Errors)
		}
	}
}

// JwtAuthMiddleware is middleware that checks JWT token
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

// SetCurrentUser is middleware that save current user
func SetCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := token.ExtractTokenID(c)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Set("userID", uid)
		c.Next()
	}
}
