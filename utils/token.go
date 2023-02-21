package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims[I any] struct {
	Info I
	jwt.RegisteredClaims
}

func CreateClaims[I any](loadInfo I, expireseAs int, issuer string) Claims[I] {
	return Claims[I]{
		Info: loadInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expireseAs))),
			// 发布时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 签发人
			Issuer: issuer,
		},
	}
}

func GetClaimsInfo[I any](cla Claims[I]) I {
	return cla.Info
}

func CreateToken[I any](loadInfo I, SigningKey string, expiresAt int, issuer string) (string, error) {
	claims := CreateClaims(loadInfo, expiresAt, issuer)
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
		if errors.Is(e, jwt.ErrTokenMalformed) {
			return c, fmt.Errorf("that's not even a token")
		} else if errors.Is(e, jwt.ErrTokenExpired) || errors.Is(e, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet
			return c, fmt.Errorf("timing is everything")
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
