package system

import "github.com/zhaogongchengsi/starter-gin/module"

type AuthorityService struct {
	UserInfo module.User
}

func NewAuthorityService(userInfo module.User) *AuthorityService {
	return &AuthorityService{UserInfo: userInfo}
}

func (a *AuthorityService) GetAuths() ([]module.Authority, string, error) {
	var authlist []module.Authority

	return authlist, "获取成功", nil
}
