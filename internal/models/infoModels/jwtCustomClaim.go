package infoModels

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaim struct {
	UserId  uint `json:"userId"`
	IsAdmin bool `json:"isAdmin"`
	jwt.StandardClaims
}
