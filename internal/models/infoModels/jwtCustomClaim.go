package infoModels

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaim struct {
	UserId  uint `json:"userId"`
	IsAdmin bool `json:"isAdmin"`
	jwt.RegisteredClaims
}
