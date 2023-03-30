package system

import (
	"errors"
	"github.com/zhaogongchengsi/starter-gin/common"
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

func (a *AuthorityService) GetAuths(page common.Page) ([]module.Authority, string, error) {
	auth := new(module.Authority)
	as, err := auth.GetAuthorities(page, global.Db)
	if err != nil {
		return as, "获取失败", err
	}
	return as, "获取成功", nil
}

func (a *AuthorityService) CreateAuth() (string, error) {
	auth := module.NewFullAuthority(a.AuthorityId, a.ParentId, a.AuthorityName)
	err := auth.CreateAuth(global.Db)
	if err != nil {
		if errors.Is(err, module.ErrAuthExist) {
			return "创建角色失败, 角色已存在", err
		}
		return "创建角色失败", err
	}
	return "创建成功", nil
}

func (a *AuthorityService) DeleteAuths(ids []int) (string, error) {
	auth := new(module.Authority)
	err := auth.DeleteAuthority(ids, global.Db)
	if err != nil {
		return "删除失败", err
	}
	return "删除成功", nil
}
