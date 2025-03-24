package httpmlog

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// process request
		c.Next()

		if len(c.Errors) > 0 {
			requestId, _ := c.Get("RequestID")
			for _, err := range c.Errors {
				log.Printf("[ERROR] id: %v | method: %v | uri: %v | error: %s",
					requestId, c.Request.Method, c.Request.RequestURI, strings.ReplaceAll(err.Error(), "\n", "; "))
			}
		}
	}
}
