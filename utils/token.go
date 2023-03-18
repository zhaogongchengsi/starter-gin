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
	CustomClaims
}

type CustomClaims struct {
	//BufferTime time.Time
	jwt.RegisteredClaims
}

var (
	ErrTokenIsNotValidPeriod = errors.New("token is not valid period") // 令牌不在有效期
)

func CreateClaims[I any](loadInfo I, issuedAt, express time.Time, issuer string) Claims[I] {
	return Claims[I]{
		Info: loadInfo,
		CustomClaims: CustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				// 过期时间
				ExpiresAt: jwt.NewNumericDate(express),
				// 发布时间
				IssuedAt: jwt.NewNumericDate(issuedAt),
				// 签发人
				Issuer: issuer,
			},
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

func ParseToken[C any](tokenString string, SigningKey string) (Claims[C], error) {

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
			return cla, ErrTokenIsNotValidPeriod
		} else {
			return cla, e
		}
	}

	if !token.Valid {
		return cla, fmt.Errorf("token is valid")
	}

	return cla, nil

}

var key = "claims"

func ShouldBindUserWith[C any](c *gin.Context, claims any) {
	cla, ok := claims.(Claims[C])
	if ok {
		info := GetClaimsInfo(cla)
		c.Set(key, info)
	}
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
