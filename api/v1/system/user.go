package system

import (
	"github.com/gin-gonic/gin"
	"github.com/server-gin/common"
	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
	systemService "github.com/server-gin/service/system"
	"github.com/server-gin/utils"
)

type captcha struct {
	Id    string `binding:"required" json:"id"`
	Value string `binding:"required" json:"value"`
}

type LoginRes struct {
	systemService.User
	Captcha captcha `binding:"required" json:"captcha"`
}

type LoginReq struct {
	User  system.User `json:"user"`
	Token string      `json:"token"`
}

func Login(c *gin.Context) {
	var loginRes LoginRes
	err := c.ShouldBindJSON(&loginRes)
	if err != nil {
		common.NewFailResponse().ErrorToString(err).Send(c)
		return
	}

	// 使用 redis
	// var store = core.NewRedisStore(global.Redis)

	if gin.Mode() == gin.ReleaseMode {
		if isOk := store.Verify(loginRes.Captcha.Id, loginRes.Captcha.Value, true); isOk {
			common.NewFailResponse().SendAfterChangeMessage("验证码验证失败", c)
			return
		}
	}

	login := systemService.User{
		Phone:    loginRes.Phone,
		Password: loginRes.Password,
		NickName: loginRes.NickName,
		Email:    loginRes.Email,
	}

	user, err := login.Login()

	if err != nil {
		common.NewFailResponse().ChangeCode(401).ErrorToString(err).Send(c)
		return
	}

	jwtConf := global.AppConfig.Jwt
	// 删除隐私信息
	user.Password = ""
	token := ""
	token, err = utils.CreateToken(user, jwtConf.SigningKey, jwtConf.ExpiresAt, jwtConf.Issuer)

	if err != nil {
		common.NewFailResponse().ChangeCode(402).ErrorToString(err).Send(c)
		return
	}

	common.NewResponse(200, LoginReq{*user, token}, "登录成功").Send(c)
}

func Register(c *gin.Context) {
	var regUser systemService.User
	err := c.ShouldBindJSON(&regUser)
	if err != nil {
		common.NewFailResponse().ErrorToString(err).Send(c)
		return
	}

	common.NewResponse(200, regUser, "注册成功").Send(c)
}
