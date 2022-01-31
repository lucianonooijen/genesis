package middleware

import "github.com/gin-gonic/gin"

// EnsureKeysMap is Gin middleware to ensure c.Keys is not nil.
func EnsureKeysMap() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
	}
}
