package utils_test

import (
	"testing"

	"github.com/server-gin/utils"
)

type Info struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

var info = Info{
	Username: "test",
	Nickname: "test",
}

func TestCreateClaims(t *testing.T) {
	claims := utils.CreateClaims(info, 3, "test")
	if claims.Info != info {
		t.Error("claims info not equal")
	}

	if claims.Issuer != "test" {
		t.Error("claims issuer not equal")
	}

	i := utils.GetClaimsInfo(claims)

	if i.Nickname != info.Nickname {
		t.Error("claims info not equal")
	}

	if i.Username != info.Username {
		t.Error("claims info not equal")
	}

}
