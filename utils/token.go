package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/server-gin/modules/system"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims[I any] struct {
	Info I
	jwt.RegisteredClaims
}

var (
	ErrTokenIsNotValidPeriod = errors.New("token is not valid period") // 令牌不在有效期
)

func CreateClaims[I any](loadInfo I, issuedAt, express time.Time, issuer string) Claims[I] {
	return Claims[I]{
		Info: loadInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(express),
			// 发布时间
			IssuedAt: jwt.NewNumericDate(issuedAt),
			// 签发人
			Issuer: issuer,
		},
	}
}

func GetClaimsInfo[I any](cla Claims[I]) I {
	return cla.Info
}

func CreateToken[I any](loadInfo I, SigningKey string, issuedAt, express time.Time, issuer string) (string, error) {
	claims := CreateClaims(loadInfo, issuedAt, express, issuer)
	// Sign and get the complete encoded token as a string using the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SigningKey))
}

func ParseToken[C any](tokenString string, SigningKey string) (c C, e error) {

	cla := Claims[C]{}

	token, e := jwt.ParseWithClaims(tokenString, &cla, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SigningKey), nil
	})

	if e != nil {
		if errors.Is(e, jwt.ErrTokenExpired) || errors.Is(e, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet
			return c, ErrTokenIsNotValidPeriod
		} else {
			return c, e
		}
	}

	if !token.Valid {
		return c, fmt.Errorf("token is valid")
	}

	claims, ok := token.Claims.(*Claims[C])

	if ok {
		c = GetClaimsInfo(*claims)
		return c, nil
	}

	return c, e

}

var key = "claims"

func ShouldBindUserWith(c *gin.Context, claims system.User) {
	c.Set(key, claims)
}

func GetUserWith(c *gin.Context) (system.User, bool) {
	userClaims, ok := c.Get(key)
	if !ok {
		return system.User{}, false
	}
	user, is := userClaims.(system.User)
	if !is {
		return system.User{}, false
	}
	return user, true
}
