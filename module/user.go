package module

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseMode
	UUID       uuid.UUID   `json:"uuid" gorm:"index;comment:用户UUID"`
	UserName   string      `json:"username" gorm:"index;comment:用户登录名"`
	NickName   string      `json:"nickname" gorm:"comment:用户昵称"`
	Email      string      `json:"email" gorm:"comment:用户邮箱"`
	Phone      string      `json:"phone" gorm:"comment:用户手机号"`
	Password   string      `json:"-" gorm:"comment:用户登录密码"`
	Mode       string      `json:"mode" gorm:"default:dark; comment:用户使用的主题  黑色(dark)或白色(light)"`
	AvatarUrl  string      `json:"avatarUrl" gorm:"comment:用户头像url"`
	Enable     int         `json:"enable" gorm:"default:1;comment:账号使用状态 1 正常 2 封禁"`
	Authoritys []Authority `json:"authoritys" gorm:"many2many:user_authoritys"`
}

func CreatePassworld(paw string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(paw), bcrypt.DefaultCost)
	return string(bytes)
}

func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUser(uname, pass, phone, nname, email string) *User {

	v4 := uuid.NewV4()

	return &User{
		UUID:     v4,
		UserName: uname,
		Password: CreatePassworld(pass),
		Phone:    phone,
		NickName: nname,
		Email:    email,
		Enable:   1,
	}
}

func CreateUser(pass, phone, nname, email string) *User {
	v4 := uuid.NewV4()
	return &User{
		UUID:     v4,
		Phone:    phone,
		Password: CreatePassworld(pass),
		NickName: nname,
		Email:    email,
	}
}

func NewFindUser(phone, pass string) *User {
	return &User{
		Password: pass,
		Phone:    phone,
	}
}

func (user *User) FirstUser(db *gorm.DB, query any, values ...any) (*User, error) {
	var u User
	result := db.Where(query, values...).First(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &User{}, gorm.ErrRecordNotFound
	}
	return &u, nil
}

func (user *User) FirstByEmail(db *gorm.DB) (*User, error) {
	res, err := user.FirstUser(db, "email = ?", user.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (user *User) FirstByPhone(db *gorm.DB) (*User, error) {
	res, err := user.FirstUser(db, "phone = ?", user.Phone)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

// 明文 = 密码
func (user *User) ComparePassword(pass string) bool {
	return BcryptCheck(user.Password, pass)
}

func (user *User) CreateUser(db *gorm.DB) (*User, error) {
	result := db.Create(user)
	if result.Error != nil {
		return &User{}, result.Error
	}
	return user, nil
}

func (user *User) Conditions(db *gorm.DB, query any, args ...any) *gorm.DB {
	return db.Model(user).Where(query, args...)
}

func (user *User) UpdatePwd(db *gorm.DB, newPwd string) (*User, error) {
	result := user.Conditions(db, "phone = ? AND password = ?", user.Phone, user.Password).Update("password", CreatePassworld(newPwd))
	if result.Error != nil {
		return &User{}, result.Error
	}
	return user, nil
}

func (user *User) UsePhoneDeleted(db *gorm.DB) error {
	result := user.Conditions(db, "phone = ?", user.Phone).Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
