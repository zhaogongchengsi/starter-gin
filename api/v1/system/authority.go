package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/common"
	systemService "github.com/zhaogongchengsi/starter-gin/service/system"
	"github.com/zhaogongchengsi/starter-gin/utils"
)

// GetAuthorities 获取所有权限
func GetAuthorities(c *gin.Context) {
	var page common.Page
	err := utils.QueryBindStruct(c, &page)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}

	auth := systemService.NewAuthorityService()

	list, msg, err := auth.GetAuths(page)
	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}

	common.NewResponseWithData(list).SendAfterChangeMessage(msg, c)
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
