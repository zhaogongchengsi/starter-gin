package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/server-gin/common"
	"github.com/server-gin/service/system"
	"github.com/server-gin/utils"
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
