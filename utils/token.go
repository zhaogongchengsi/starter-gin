package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/server-gin/global"
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

func CreateToken[I any](loadInfo I, SigningKey string) (string, error) {

	claims := CreateClaims(loadInfo, 30, global.JwtConfig.Issuer)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(SigningKey))

}

func ParseToken(tokenString string, hmacSampleSecret string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

}
