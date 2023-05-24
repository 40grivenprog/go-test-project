package infrastructure

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
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
	}
}
