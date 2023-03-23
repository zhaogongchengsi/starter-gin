package system

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/module"
)

type AuthorityService struct {
	AuthorityId   int    `json:"authorityId" validate:"required"`
	AuthorityName string `json:"authorityName" validate:"required"`
	ParentId      int    `json:"parentId"`
}

func NewAuthorityService() *AuthorityService {
	return &AuthorityService{}
}

func (a *AuthorityService) GetAuths(uid uuid.UUID) ([]module.Authority, string, error) {
	auth := module.NewAuthority(a.AuthorityId, a.ParentId, a.AuthorityName)
	user, err := auth.GetUserAuths(uid.String(), global.Db)
	if err != nil {
		return user.Authoritys, "获取失败", err
	}
	return user.Authoritys, "获取成功", nil
}

func (a *AuthorityService) CreateAuth() (string, error) {
	auth := module.NewAuthority(a.AuthorityId, a.ParentId, a.AuthorityName)
	err := auth.CreateAuth(global.Db)
	if err != nil {
		if errors.Is(err, module.ErrAuthExist) {
			return "创建角色失败, 角色已存在", err
		}
		return "创建角色失败", err
	}
	return "创建成功", nil
}
