package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// @Author KHighness
// @Update 2022-02-13

// 用户模型
type User struct {
	gorm.Model
	UserName        string  `gorm:unique`
	PasswordDigest  string
}

const (
	PasswordLevel = 12 // 密码加密级别
)

// 设置密码
func (user *User) setPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordLevel)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// 校验密码
func (user *User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
