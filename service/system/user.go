package system

import (
	"errors"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/module"
)

type User struct {
	Phone    string `json:"phone" validate:"required, len=11"`
	Password string `form:"password" json:"password" binding:"required" validate:"gte=6,lte=18"`
	NickName string `form:"nickName" json:"nickName"`
	Email    string `json:"email" validate:"e-mail"`
}

var (
	ErrWrongPassword = errors.New("err: wrong password")
	ErrUserExt       = errors.New("err: User already exists")
	ErrUserNotFound  = errors.New("err: user does not exist, please check and try again")
)

func (u *User) Login() (*module.User, string, error) {

	user := module.NewFindUser(u.Phone, u.Password) // 准备登陆的账号

	ut, err := user.FirstByPhone(global.Db) // 从数据库内查出来的用户

	if err != nil {
		return user, "账号不存在", ErrUserNotFound
	}
	// 用 登陆的明文密码 和 数据库内加密后的密码进行对比
	if !user.ComparePassword(ut.Password) {
		return user, "密码错误", ErrWrongPassword
	}

	return ut, "登录成功", nil
}

func (u *User) Register() (*module.User, string, error) {
	user := module.CreateUser(u.Password, u.Phone, u.NickName, u.Email)
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

func (u *User) ChangePassword(newPwd string) (*module.User, error) {
	us, _, err := u.Login()

	if err != nil {
		return us, err
	}

	ut, e := us.UpdatePwd(global.Db, newPwd)

	if e != nil {
		return ut, e
	}

	return ut, nil
}

func (u *User) DeletedUser() error {
	user := module.User{
		Phone: u.Phone,
		Email: u.Email,
	}

	return user.UsePhoneDeleted(global.Db)
}

func (u *User) GetAuths() ([]module.Authority, string, error) {
	var list []module.Authority
	user := module.NewFindUser(u.Phone, u.Password)
	list, err := user.GetAuthoritysByPhone(global.Db)
	if err != nil {
		return list, "获取失败", err
	}

	return list, "获取成功", nil
}

func (u *User) GetUserRouters() ([]module.RouterRecord, string, error) {
	user := module.NewFindUser(u.Phone, u.Password)
	list, err := user.GetAuthoritysByPhone(global.Db)
	if err != nil {
		return []module.RouterRecord{}, "获取路由失败", err
	}
	routers := []module.RouterRecord{}

	for _, authority := range list {
		routers = append(routers, authority.RouterRecords...)
	}

	return routers, "获取成功", nil
}

func (u *User) AddAuthority(uid int, authid int) (string, error) {

	auth := module.NewAuthority(authid)

	err := auth.FindAuthority(global.Db)

	if err != nil {
		if errors.Is(err, module.ErrAuthNotExist) {
			return "权限不存在", err
		}
		return "未知错误", err
	}

	user := module.User{BaseMode: module.BaseMode{ID: uint(uid)}}
	_, err = user.FirstById(global.Db)
	if err != nil {
		if errors.Is(err, module.ErrUserNotExist) {
			return "用户不存在", err
		}
		return "未知错误", err
	}

	ua := module.NewUserAuthority(uid, authid)
	err = ua.CreateUserAuth(global.Db)

	if err != nil {
		if errors.Is(err, module.ErrUserAuthExists) {
			return "已添加，请勿重复添加", err
		}
		return "添加失败", err
	}

	return "添加成功", err
}
