package utils

import (
	"dyd/entity"
	"dyd/log"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// 定义过期时间
const tokenExpireDuration = time.Hour * 5

// 定义密钥
var secret = []byte("waterIB")

// 定义签发人
const issuer = "didyoudead"

// 定义UserToken生成
func EnruptUserToken(username string, password string) (string, error) {
	userClaim := entity.UserClaim{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	return token.SignedString(secret)
}

// 定义UserToken解析
func DeruptUserToken(tokenString string) (*entity.UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &entity.UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		log.Warning.Println("token解析失败，token："+tokenString, err)
	}

	if claims, ok := token.Claims.(*entity.UserClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的token")
}
