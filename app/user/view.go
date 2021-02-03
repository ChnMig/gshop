package user

import (
	"gshop/tools"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @id userLogin
// @router /v1/user [get]
// @summary user login
// @description Access this api to request login, if successful, return the user's token
// @accept application/json
// @produce application/json
// @param object body model.nameLoginParam true "body"
// @securitydefinitions.apikey ApiKeyAuth
// @response 200 {object} model.nameLoginParam
func nameLoginHandler(c *gin.Context) {
	params := nameLoginParam{}
	if err := c.ShouldBindJSON(&params); err != nil {
		ei := "The request api parameter does not match the setting"
		tools.FormatError(c, 1001, ei)
		tools.Log.Warn(ei, zap.Error(err))
		return
	}
}
