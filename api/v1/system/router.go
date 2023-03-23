package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/service/system"
	"github.com/zhaogongchengsi/starter-gin/utils"
)

func GetRouters(c *gin.Context) {
	uc, ok := utils.GetUserWith(c)
	if !ok {
		common.NewFailResponse().SendAfterChangeMessage("用户未登录请重试", c)
		return
	}
	r := system.NewSysRouter()
	list, msg, err := r.GetUserRouters(uc.Phone, uc.Password)
	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}

	common.NewResponseWithData(list).Send(c)
}
