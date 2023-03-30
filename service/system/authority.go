package system

import (
	"errors"
	"fmt"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/module"
)

type AuthorityService struct {
	AuthorityId   int    `json:"authorityId" validate:"required"`
	AuthorityName string `json:"authorityName" validate:"required"`
	ParentId      int    `json:"parentId"`
}

// ErrBeingUsed 权限正在被使用
var ErrBeingUsed = errors.New("err: permissions are being used")

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

func (a *AuthorityService) DeleteAuths(id int) (string, error) {
	uas := new(module.UserAuthority)
	ok := uas.WhetherByAuthId(id, global.Db)
	if ok {
		return fmt.Sprintf("%v 用户有角色正在使用，拒绝删除！", id), ErrBeingUsed
	}

	auth := module.NewAuthority(id)

	err := auth.FindAuthority(global.Db)
	if err != nil {
		return "删除失败, 角色不存在", err
	}

	err = auth.DeleteAuthority(id, global.Db)

	if err != nil {
		return "删除失败", err
	}

	return "删除成功", nil
}
