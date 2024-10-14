package utils

import (
	"SEP/internal/models/dataModels"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type JwtTool struct{}

func (j *JwtTool) GenerateLoginToken(user *dataModels.User) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	claims := jwt.MapClaims{
		"UserId":  user.ID,
		"IsAdmin": user.UserIsAdmin,
		"exp":     expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("jwt.jwtSecret")))
	Log.WithFields(logrus.Fields{
		"jwtToken":  t,
		"jwtsecret": viper.GetString("jwt.jwtSecret"),
	}).Info("JWT Token in log in")
	if err != nil {
		return "", err
	}
	return t, nil
}
