package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/common"
	systemService "github.com/zhaogongchengsi/starter-gin/service/system"
	"github.com/zhaogongchengsi/starter-gin/utils"
)

func GetAuthority(c *gin.Context) {
	uc, ok := utils.GetUserWith(c)
	if !ok {
		common.NewFailResponse().SendAfterChangeMessage("用户未登录请重试", c)
		return
	}

	auth := systemService.NewAuthorityService()

	list, msg, err := auth.GetAuths(uc.UUID)
	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}
	common.NewResponseWithData(list).Send(c)
}

func AddAuthority(c *gin.Context) {
	var auth systemService.AuthorityService
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}
	msg, err := auth.CreateAuth()
	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}

	common.NewOkResponse().SendAfterChangeMessage(msg, c)
}

// DeleteAuthority todo: 删除权限
func DeleteAuthority(c *gin.Context) {

}
