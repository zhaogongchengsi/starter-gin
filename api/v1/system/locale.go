package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/service/system"
)

func Getlocale(c *gin.Context) {

	lan := system.Messages{}

	msgs, err := lan.GetMessages()

	if err != nil {
		common.NewOkResponse().AddError(err, "获取语言失败").Send(c)
		return
	}

	common.NewResponseWithData(msgs).Send(c)

}
