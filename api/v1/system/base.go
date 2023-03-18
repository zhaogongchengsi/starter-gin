package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/server-gin/common"
	"github.com/server-gin/global"
	"github.com/server-gin/module"
	"github.com/server-gin/service/system"
)

var store = base64Captcha.DefaultMemStore

type CaptchaResponse struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

func Captcha(ctx *gin.Context) {
	config := global.AppConfig.Captcha
	// qw := ctx.DefaultQuery("width", strconv.FormatInt(int64(config.Height), 10))
	// qh := ctx.DefaultQuery("height", strconv.FormatInt(int64(config.Width), 10))
	// width, err := strconv.ParseInt(qw, 10, 8)
	// if err != nil {
	// 	common.NewFailResponse().AddError(err, "参数(宽度)无法转换为数字").Send(ctx)
	// }
	// height, err := strconv.ParseInt(qh, 10, 8)

	// if err != nil {
	// 	common.NewFailResponse().AddError(err, "参数(长度)无法转换为数字").Send(ctx)
	// }

	// 使用 redis
	// var store = core.NewRedisStore(global.Redis)

	driver := base64Captcha.NewDriverDigit(config.Height, config.Width, config.Length, config.MaxSkew, config.DotCount)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()

	if err != nil {
		common.NewParamsError(ctx, err)
		return
	}

	common.NewResponseWithData(CaptchaResponse{id, b64s}).Send(ctx)
}

func UpLoad(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		common.NewParamsError(c, err)
		return
	}

	fileinfo := system.UpdateFileInfo{FileName: file.Filename, FileHeader: file}

	ce, err := fileinfo.SaveFile()

	if err != nil {
		common.NewFailResponse().AddError(err, "文件上传失败").Send(c)
		return
	}

	common.NewResponseWithData(ce).Send(c)

}

func UploadMult(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]

	succFile := []module.File{}
	failFile := []string{}

	for _, file := range files {
		fileinfo := system.UpdateFileInfo{FileName: file.Filename, FileHeader: file}
		ce, err := fileinfo.SaveFile()
		if err != nil {
			failFile = append(failFile, file.Filename)
		} else {
			succFile = append(succFile, ce)
		}

	}

	common.NewResponseWithData(gin.H{"fail": failFile, "success": succFile}).Send(c)
}
