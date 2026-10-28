package jwtutil

import (
	"krm-backend/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSignKey = []byte(config.JwtSignKey)

type MyCustomerClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(username string) (string, error) {
	claims := MyCustomerClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtExpTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "krm-backend",
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSignKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ParseToken(tokenString string) (*MyCustomerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomerClaims{}, func(token *jwt.Token) (any, error) {
		return jwtSignKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 断言 token.Claims 为 *MyCustomerClaims 类型，并检查 token 是否有效
	if claims, ok := token.Claims.(*MyCustomerClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
