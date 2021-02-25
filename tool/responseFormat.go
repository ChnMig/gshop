package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FormatOk ok
func FormatOk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "success",
	})
	// Return directly
	c.Abort()
}

// FormatError err
func FormatError(c *gin.Context, errorCode int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": errorCode,
		"data": message,
	})
	// Return directly
	c.Abort()
}

// FormatData data
func FormatData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
	// Return directly
	c.Abort()
}
