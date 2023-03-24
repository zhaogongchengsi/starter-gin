package system

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/zhaogongchengsi/starter-gin/module"

	"github.com/zhaogongchengsi/starter-gin/global"
)

type User struct {
	Phone    string `json:"phone" validate:"required, len=11"`
	Password string `form:"password" json:"password" binding:"required" validate:"gte=6,lte=18"`
	NickName string `form:"nickName" json:"nickName"`
	Email    string `json:"email" validate:"e-mail"`
}

var (
	ErrWrongPassword      = errors.New("err: wrong password")
	ErrTokenSigningFailed = errors.New("err: Token signing failed")
	ErrUserExt            = errors.New("err: User already exists")
	ErrUserNotFound       = errors.New("err: user does not exist, please check and try again")
	ErrUuidInvalid        = errors.New("err: uuid is invalid")
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

func (u *User) AddAuthority(uid string, authid int) (string, error) {

	id, err := uuid.FromString(uid)
	if err != nil {
		return "uuid 无效", ErrUuidInvalid
	}

	user := module.User{
		UUID:       id,
		Authoritys: []module.Authority{{AuthorityId: authid}},
	}

	err = user.AddAssociation(global.Db)

	if err != nil {
		return "添加失败", err
	}

	return "添加成功", err
}
