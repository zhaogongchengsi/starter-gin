package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/module"
	"github.com/zhaogongchengsi/starter-gin/service/system"
	"github.com/zhaogongchengsi/starter-gin/utils"
)

type CaptchaResponse struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

// Captcha 获取验证码
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

	driver := base64Captcha.NewDriverDigit(config.Height, config.Width, config.Length, config.MaxSkew, config.DotCount)
	cp := base64Captcha.NewCaptcha(driver, global.CaptchaStore)
	id, b64s, err := cp.Generate()
	if err != nil {
		common.NewParamsError(ctx, err)
		return
	}
	common.NewResponseWithData(CaptchaResponse{id, b64s}).Send(ctx)
}

// AddCaptchaBlacklist 拉黑验证码
func AddCaptchaBlacklist(c *gin.Context) {
	id := c.Query("captcha_id")
	global.CaptchaStore.Delete(id)
	common.NewOkResponse().Send(c)
}

// UpLoad 文件上传
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

// UploadMult 多个文件上传
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

func Health(c *gin.Context) {

	info := c.Query("info")

	if len(info) < 1 {
		common.NewOkResponse().SendAfterChangeMessage("ok", c)
		return
	}

	var health utils.Health
	os := utils.NewOs()
	health.Os = *os
	cpu, err := utils.NewCpu()
	if err != nil {
		common.NewFailResponse().AddError(err, "获取cpu 状态失败").Send(c)
		return
	}
	health.Cpu = cpu
	ram, err := utils.NewRAM()
	if err != nil {
		common.NewFailResponse().AddError(err, "获取ram 状态失败").Send(c)
		return
	}
	health.Ram = ram

	disk, err := utils.NewDisk()
	if err != nil {
		common.NewFailResponse().AddError(err, "获取disk 状态失败").Send(c)
		return
	}
	health.Disk = disk

	common.NewResponseWithData(health).Send(c)
}
