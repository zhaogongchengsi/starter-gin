package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/service/system"
	"github.com/zhaogongchengsi/starter-gin/utils"
)

func GetAuthority(c *gin.Context) {
	user, is := utils.GetUserWith(c)
	if !is {
		common.NewParamsError(c, errors.New("user not logged in"))
		return
	}

	global.Logger.Info("日志写入测试")

	authservice := system.NewAuthorityService(user)
	authservice.GetAuths()

	common.NewResponseWithData(user).Send(c)

}
