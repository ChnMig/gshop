package user

import (
	"fmt"
	tools "gshop/tool"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
func nickNameLoginHandler(c *gin.Context) {
	// get data
	params := nickNameLoginParam{}
	if err := c.ShouldBindJSON(&params); err != nil {
		ei := "The request api parameter does not match the setting"
		tools.FormatError(c, 1001, ei)
		tools.Log.Warn(ei, zap.Error(err))
		return
	}
	u, err := useNickNamegetUserFromDB(params.NickName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ei := "This user was not found"
			tools.FormatError(c, 1103, ei)
			return
		}
		ei := "An unknown error occurred during database query"
		tools.FormatError(c, 1501, ei)
		tools.Log.Warn(ei, zap.Error(err))
		return
	}
	fmt.Println(u)
}
