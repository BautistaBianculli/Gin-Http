package models

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	User     string `json:"user"`
	Password string `json:"password"`
	jwt.StandardClaims
}
