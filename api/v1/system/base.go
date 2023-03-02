package system

import (
	"github.com/gin-gonic/gin"
	"github.com/server-gin/common"
)

func Login(ctx *gin.Context) {
	common.NewOkResponse().Send(ctx)
}

type img struct {
	Dase64 string
	Id     string
}

func Captcha(ctx *gin.Context) {

	i := img{
		Dase64: "123123",
		Id:     "123123",
	}

	common.NewResponseWithData(i).SendAfterChangeData(img{
		Dase64: "asidhaosidh",
		Id:     "auisdgaiusdgaiusdg",
	}, ctx)
}
