package dto

import "github.com/golang-jwt/jwt/v5"

type JwtPayload struct {
	UserId uint   `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Time   int64  `json:"time"`
	jwt.RegisteredClaims
}
