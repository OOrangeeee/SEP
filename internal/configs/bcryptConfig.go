package configs

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type encryptionTool struct {
}

func (eT *encryptionTool) encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "密码加密失败",
		}).Panic("密码加密失败")
	}
	return string(hashedPassword), nil
}

func (eT *encryptionTool) comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
