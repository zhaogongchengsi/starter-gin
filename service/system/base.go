package system

import (
	"fmt"

	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
	"github.com/server-gin/utils"
)

type Login struct {
	Phone    string `json:"phone" validate:"required, len=11"`
	Password string `form:"password" json:"password" binding:"required" validate:"gte=6,lte=18"`
	NickName string `form:"nickName" json:"nickName"`
	Email    string `json:"email" validate:"email"`
}

func (L *Login) Login() (user *system.User, token string, err error) {

	// 假装查询

	user = system.NewFindUser(L.Phone, L.Password)

	user, err = user.FirstByPhone(global.Db)

	if err != nil {
		return user, "", fmt.Errorf("user does not exist, please check and try again")
	}

	jwtConf := global.AppConfig.Jwt
	// 删除隐私信息
	user.Password = ""
	user.Phone = ""

	token, err = utils.CreateToken(user, jwtConf.SigningKey, jwtConf.ExpiresAt, jwtConf.Issuer)
	if err != nil {
		return user, "", err
	}

	return user, token, nil
}
