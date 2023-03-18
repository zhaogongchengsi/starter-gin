package system

import (
	"github.com/gin-gonic/gin"
	"github.com/server-gin/common"
	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
	systemService "github.com/server-gin/service/system"
	"github.com/server-gin/utils"
	"time"
)

type captcha struct {
	Id    string `binding:"required" json:"id"`
	Value string `binding:"required" json:"value"`
}

type LoginRes struct {
	systemService.User
	Captcha captcha `binding:"required" json:"captcha"`
}

type Token struct {
	IssuedAt  time.Time `json:"issued_at"`
	ExpressAt time.Time `json:"express_at"`
	Token     string    `json:"token"`
}
type LoginReq struct {
	User  system.User `json:"user"`
	Token Token       `json:"authorization"`
}

// 登录
func Login(c *gin.Context) {
	var loginRes LoginRes
	err := c.ShouldBindJSON(&loginRes)
	if err != nil {
		common.NewParamsError(c, err)
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

	user, msg, err := login.Login()

	if err != nil {
		common.NewFailResponse().ChangeCode(common.AuthFailed).AddError(err, msg).Send(c)
		return
	}

	jwtConf := global.AppConfig.Jwt
	// 删除隐私信息
	user.Password = ""
	token := ""
	it := time.Now()
	et := it.Add(time.Duration(jwtConf.ExpiresAt) * time.Minute)
	token, err = utils.CreateToken(user, jwtConf.SigningKey, it, et, jwtConf.Issuer)

	if err != nil {
		common.NewFailResponse().ChangeCode(common.AccreditFail).AddError(err, "token 签发失败").Send(c)
		return
	}

	common.NewResponse(common.Ok, LoginReq{*user, Token{
		ExpressAt: et,
		IssuedAt:  it,
		Token:     token,
	}}, msg).Send(c)
}

// 注册用户
func Register(c *gin.Context) {
	var regUser systemService.User
	err := c.ShouldBindJSON(&regUser)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}

	user, msg, err := regUser.Register()

	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}

	common.NewResponse(200, user, msg).Send(c)
}

type ChangePasswordReq struct {
	Phone       string `json:"phone" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

// 修改密码
func ChangePassword(c *gin.Context) {
	var changeinfo ChangePasswordReq
	err := c.ShouldBindJSON(&changeinfo)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}

	user := systemService.User{
		Phone:    changeinfo.Phone,
		Password: changeinfo.OldPassword,
	}

	us, err := user.ChangePassword(changeinfo.NewPassword)

	if err != nil {
		common.NewFailResponse().AddError(err, "修改失败").Send(c)
		return
	}

	common.NewResponse(200, us, "修改成功").Send(c)
}

type DeleteUserInfo struct {
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email"`
}

// 删除用户
func DeleteUser(c *gin.Context) {
	var di DeleteUserInfo
	err := c.ShouldBindJSON(&di)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}

	regUser := systemService.User{
		Phone: di.Phone,
		Email: di.Email,
	}

	err = regUser.DeletedUser()

	if err != nil {
		common.NewFailResponse().AddError(err, "删除失败").Send(c)
		return
	}

	common.NewOkResponse().SendAfterChangeMessage("删除成功", c)
}

func GetUsers(c *gin.Context) {
	uc, ok := utils.GetUserWith(c)
	if !ok {
		common.NewFailResponse().SendAfterChangeMessage("获取失败", c)
		return
	}

	//c.Header("new-authorization", "asdjasiodhaoisdhaiosdh")
	//c.Header("new-expires-at", "asdjasiodhaoisdhaiosdh")
	//c.Header("new-issued-at", "asdjasiodhaoisdhaiosdh")
	common.NewResponseWithData(uc).SendAfterChangeMessage("获取成功", c)
}
