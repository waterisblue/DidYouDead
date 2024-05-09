package entity

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Username string
	Password string
	jwt.StandardClaims
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDetail struct {
	Username      string
	Password      string
	Administrator bool
}
