package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/service/system"
)

// GetRouters todo: 获取所有路由
func GetRouters(c *gin.Context) {
	r := system.NewSysRouter()
	rs, msg, err := r.GetUserRouters()
	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}
	common.NewResponseWithData(rs).Send(c)
}
