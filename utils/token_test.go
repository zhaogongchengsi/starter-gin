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

const (
	SigningKey = "test"
	expireseAs = 3
	issuer     = "test"
)

func TestCreateClaims(t *testing.T) {
	claims := utils.CreateClaims(info, expireseAs, issuer)
	if claims.Info != info {
		t.Error("claims info not equal")
	}

	if claims.Issuer != issuer {
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

func TestCreateToken(t *testing.T) {

	token, err := utils.CreateToken(info, SigningKey, expireseAs, issuer)

	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Errorf("token is empty")
	}

}

func TestParseToken(t *testing.T) {
	token, err := utils.CreateToken(info, SigningKey, expireseAs, issuer)

	if err != nil {
		t.Error(err)
	}

	c, err := utils.ParseToken[Info](token, SigningKey)

	if err != nil {
		t.Error(err)
	}

	if c.Username != info.Username {
		t.Error("claims info not equal")
	}

	if c.Nickname != info.Nickname {
		t.Error("claims info not equal")
	}

}
