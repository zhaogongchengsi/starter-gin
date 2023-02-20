package utils

import (
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expireseAs))),
			Issuer:    issuer,
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
	token, e := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SigningKey), nil
	})

	if claims, ok := token.Claims.(Claims[C]); ok && token.Valid {
		c = GetClaimsInfo(claims)
		return c, nil
	}

	return c, e

}
