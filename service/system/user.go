package system

import (
	"errors"
	"github.com/server-gin/module"

	"github.com/server-gin/global"
)

type User struct {
	Phone    string `json:"phone" validate:"required, len=11"`
	Password string `form:"password" json:"password" binding:"required" validate:"gte=6,lte=18"`
	NickName string `form:"nickName" json:"nickName"`
	Email    string `json:"email" validate:"e-mail"`
}

var ErrUserNotFound = errors.New("err: user does not exist, please check and try again")
var ErrWrongPassword = errors.New("err: wrong password")
var ErrTokenSigningFailed = errors.New("err: Token signing failed")
var ErrUserExt = errors.New("err: User already exists")

func (L *User) Login() (user *module.User, msg string, err error) {

	user = module.NewFindUser(L.Phone, L.Password)

	u, err := user.FirstByPhone(global.Db)

	if err != nil {
		return user, "账号不存在", ErrUserNotFound
	}

	if !user.ComparePassword(u.Password) {
		return user, "密码错误", ErrWrongPassword
	}

	return u, "登录成功", nil
}

func (R *User) Register() (*module.User, string, error) {
	user := module.CreateUser(R.Password, R.Phone, R.NickName, R.Email)
	oldu, err := user.FirstByPhone(global.Db)

	if err == nil {
		return oldu, "用户已存在", ErrUserExt
	}

	newUser, err := user.CreateUser(global.Db)

	if err != nil {
		return user, "创建失败", err
	}

	return newUser, "注册成功", nil

}

func (R *User) ChangePassword(newPwd string) (*module.User, error) {
	us, _, err := R.Login()

	if err != nil {
		return us, err
	}

	u, e := us.UpdatePwd(global.Db, newPwd)

	if e != nil {
		return u, e
	}

	return u, nil
}

func (R *User) DeletedUser() error {
	user := module.User{
		Phone: R.Phone,
		Email: R.Email,
	}

	return user.UsePhoneDeleted(global.Db)
}
