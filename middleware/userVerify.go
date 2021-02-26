package middleware

import (
	"fmt"
	"gshop/db"
	"gshop/tool"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TokenVerify Get the token and verify its validity
func TokenVerify(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		tool.FormatError(c, 2003, "token expired or invalid")
		tool.Log.Warn(fmt.Sprintf("token invalid: %v", token))
		return
	}
	tokenID, err := tool.JWTDecrypt(token)
	if err != nil {
		tool.FormatError(c, 2003, "token expired or invalid")
		tool.Log.Warn(fmt.Sprintf("token invalid: %v", token), zap.Error(err))
		return
	}
	// get RDB token
	if val, err := db.RDB.Get(db.RDB.Context(), tokenID).Result(); err != nil || val != token {
		tool.FormatError(c, 2003, "token expired or invalid")
		tool.Log.Info(fmt.Sprintf("token expired: %v", token), zap.Error(err))
		return
	}
	// set tokenID to gin.Context
	c.Set("tokenID", tokenID)
	// Next
	c.Next()
}
