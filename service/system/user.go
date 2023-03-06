package system

import (
	"errors"

	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
)

type User struct {
	Phone    string `json:"phone" validate:"required, len=11"`
	Password string `form:"password" json:"password" binding:"required" validate:"gte=6,lte=18"`
	NickName string `form:"nickName" json:"nickName"`
	Email    string `json:"email" validate:"e-mail"`
}

var ErrUserNotFound = errors.New("err:user does not exist, please check and try again(用户不存在)")
var ErrWrongPassword = errors.New("err:wrong password(密码错误)")
var ErrTokenSigningFailed = errors.New("err:Token signing failed(token 签发错误)")
var ErrUserExt = errors.New("err:User already exists(用户已存在)")

func (L *User) Login() (user *system.User, err error) {

	user = system.NewFindUser(L.Phone, L.Password)

	u, err := user.FirstByPhone(global.Db)

	if err != nil {
		return user, ErrUserNotFound
	}

	if !user.ComparePassword(u.Password) {
		return user, ErrWrongPassword
	}

	return u, nil
}

func (R *User) Register() (user *system.User, err error) {
	user = system.CreateUser(R.Password, R.Phone, R.NickName, R.Email)
	oldu, err := user.FirstByPhone(global.Db)
	if err == nil {
		return oldu, ErrUserExt
	}

	newUser, err := user.CreateUser(global.Db)

	if err == nil {
		return user, err
	}

	return newUser, nil

}
