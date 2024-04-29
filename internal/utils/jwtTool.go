package utils

import (
	"SEP/internal/models/dataModels"
	"SEP/internal/models/infoModels"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type JwtTool struct{}

func (j *JwtTool) GenerateLoginToken(user *dataModels.User) (string, error) {
	claims := infoModels.JwtCustomClaim{
		UserId:         user.ID,
		IsAdmin:        user.UserIsAdmin,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("jwt.jwtSecret")))
	if err != nil {
		return "", err
	}
	return t, nil
}
