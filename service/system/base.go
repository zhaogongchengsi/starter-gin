package system

import (
	"github.com/server-gin/global"
	"github.com/server-gin/utils"
)

type Login struct {
	Phone    string `json:"phone" validate:"required, len=11"`
	Password string `form:"password" json:"password" binding:"required" validate:"gte=6,lte=18"`
	NickName string `form:"nickName" json:"nickName"`
	Email    string `json:"email" validate:"email"`
}

func (L *Login) Login() (token string, err error) {

	// 假装查询

	jwtConf := global.AppConfig.Jwt
	// 删除隐私信息
	L.Password = ""
	L.Phone = ""

	token, err = utils.CreateToken(L, jwtConf.SigningKey, jwtConf.ExpiresAt, jwtConf.Issuer)
	if err != nil {
		return "", err
	}

	return token, nil
}
