package infrastructure

import (
	"crypto/rand"
	"encoding/hex"

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
