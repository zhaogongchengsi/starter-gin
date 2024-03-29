package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/module"
	systemService "github.com/zhaogongchengsi/starter-gin/service/system"
	"github.com/zhaogongchengsi/starter-gin/utils"
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
	User  module.User `json:"user"`
	Token Token       `json:"authorization"`
}

type UserAndAuthority struct {
	Userid int `json:"user_id"`
	AuthId int `json:"auth_id"`
}

// Login 登录
func Login(c *gin.Context) {
	var loginRes LoginRes
	err := c.ShouldBindJSON(&loginRes)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}

	if isOk := global.CaptchaStore.Verify(loginRes.Captcha.Id, loginRes.Captcha.Value, true); !isOk {
		common.NewFailResponse().SendAfterChangeMessage("验证码验证失败", c)
		return
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
	user.Phone = ""

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

// Register 注册用户
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

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var changeling ChangePasswordReq
	err := c.ShouldBindJSON(&changeling)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}

	user := systemService.User{
		Phone:    changeling.Phone,
		Password: changeling.OldPassword,
	}

	us, err := user.ChangePassword(changeling.NewPassword)

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

// DeleteUser 删除用户
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

// GetUserAuthorities 获取用户权限列表
func GetUserAuthorities(c *gin.Context) {
	uc, ok := utils.GetUserWith(c)
	if !ok {
		common.NewFailResponse().SendAfterChangeMessage("用户未登录请重试", c)
		return
	}
	user := systemService.NewUser()
	list, msg, err := user.GetAuths(uc.UUID)
	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}

	common.NewResponseWithData(list).SendAfterChangeMessage(msg, c)
}

// SetUserAuthority 给一个用户设置权限
func SetUserAuthority(c *gin.Context) {
	var ua UserAndAuthority
	err := c.ShouldBindJSON(&ua)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}

	user := systemService.NewUser()
	msg, err := user.AddAuthority(ua.Userid, ua.AuthId)

	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}

	common.NewOkResponse().SendAfterChangeMessage(msg, c)
}

// DeleteUserAuthority 删除一个用户的权限
func DeleteUserAuthority(c *gin.Context) {
	var ua UserAndAuthority
	err := c.ShouldBindJSON(&ua)
	if err != nil {
		common.NewParamsError(c, err)
		return
	}

	user := systemService.User{}
	msg, err := user.DeleteAuthority(ua.Userid, ua.AuthId)

	if err != nil {
		common.NewFailResponse().AddError(err, msg).Send(c)
		return
	}

	common.NewOkResponse().SendAfterChangeMessage(msg, c)
}

// GetUserRouters 获取用户所有的路由
func GetUserRouters(c *gin.Context) {
	uc, ok := utils.GetUserWith(c)

	if !ok {
		common.NewFailResponse().SendAfterChangeMessage("用户未登录请重试", c)
		return
	}
	user := new(systemService.User)

	routers, s, err := user.GetUserRouters(uc.UUID)
	if err != nil {
		common.NewFailResponse().AddError(err, s).Send(c)
		return
	}

	common.NewResponseWithData(routers).SendAfterChangeMessage(s, c)
}
