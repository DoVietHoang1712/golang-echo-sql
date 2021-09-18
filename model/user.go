package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null"`
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Image    *string
}

func (u *User) HashPassword(plainText string) (string, error) {
	if len(plainText) == 0 {
		return "", errors.New("passoword not empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plainText), 10)
	return string(h), err
}

func (u *User) CheckPassword(plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	return err == nil
}
