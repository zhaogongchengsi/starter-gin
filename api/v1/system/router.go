package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/service/system"
)

// GetRouters 获取所有路由
func GetRouters(c *gin.Context) {
	r := system.NewSysRouter()
	rs, msg, err := r.GetAllRouters()
	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}
	common.NewResponseWithData(rs).Send(c)
}

func CreateRouter(c *gin.Context) {
	var r system.SysRouter
	err := c.ShouldBindJSON(&r)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}
	msg, err := r.CreateRouter()
	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}
	common.NewOkResponse().SendAfterChangeMessage(msg, c)
}

func DeleteRouter(c *gin.Context) {
	var id common.ID[int]
	if err := c.ShouldBindJSON(&id); err != nil {
		common.NewParamsError(c, err)
		return
	}

	r := system.NewSysRouter()

	err := r.DeleteRouter(id.Id)

	if err != nil {
		common.NewFailResponse().AddError(err, "删除失败").Send(c)
		return
	}

	common.NewOkResponse().SendAfterChangeMessage("删除成功", c)
}
