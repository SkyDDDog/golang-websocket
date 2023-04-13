package repository

import (
	"demo04/pkg/constants"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"unique"`
	Password string
	Model
}

func (user *User) CheckUserIdExist() bool {
	if err := DB.Where("id = ?", user.Id).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (user *User) CheckUsernameExist() bool {
	if err := DB.Where("username = ?", user.Username).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (user *User) UserRegister() (*User, error) {
	user.Id = SF.NextVal()
	err := DB.Create(&user).Error
	return user, err
}

func (user *User) SetPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), constants.PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 密码校验
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) UserLogin() (*User, error) {
	plainPassword := user.Password
	DB.Where("username = ?", user.Username).First(&user)
	if user.CheckPassword(plainPassword) {
		return user, nil
	}
	return user, errors.New("密码校验失败")
}
