package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/service/system"
	"github.com/zhaogongchengsi/starter-gin/utils"
)

func GetAuthority(c *gin.Context) {
	user, is := utils.GetUserWith(c)
	if !is {
		common.NewParamsError(c, errors.New("user not logged in"))
		return
	}

	authservice := system.NewAuthorityService(user)
	authservice.GetAuths()

	common.NewResponseWithData(user).Send(c)

}
