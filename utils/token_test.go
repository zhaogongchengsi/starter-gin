package utils_test

import (
	"errors"
	"testing"
	"time"

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
	it := time.Now()
	et := it.Add(expireseAs + time.Minute)
	claims := utils.CreateClaims(info, it, et, issuer)
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
	it := time.Now()
	et := it.Add(expireseAs + time.Minute)

	token, err := utils.CreateToken(info, SigningKey, it, et, issuer)

	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Errorf("token is empty")
	}

}

func TestParseToken(t *testing.T) {

	it := time.Now()
	et := it.Add(expireseAs * time.Minute)
	token, err := utils.CreateToken(info, SigningKey, it, et, issuer)

	if err != nil {
		t.Error(err)
	}

	c, err := utils.ParseToken[Info](token, SigningKey)

	if err != nil {
		t.Error(err)
		return
	}

	if c.Username != info.Username {
		t.Error("claims info not equal")
	}

	if c.Nickname != info.Nickname {
		t.Error("claims info not equal")
	}

	it = time.Now()
	et = it.Add(expireseAs * time.Microsecond)
	token, err = utils.CreateToken(info, SigningKey, it, et, issuer)

	if err != nil {
		t.Error(err)
	}

	time.Sleep(expireseAs * time.Microsecond)

	c, err = utils.ParseToken[Info](token, SigningKey)

	if err != nil {
		if !errors.Is(err, utils.ErrTokenIsNotValidPeriod) {
			t.Error(err)
		}
		return
	}

}
