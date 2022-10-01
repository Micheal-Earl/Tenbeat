package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID           int    `json:"id" gorm:"primaryKey;unique;AUTO_INCREMENT"`
	Username     string `json:"username" gorm:"type:varchar(28);unique; not null"`
	PasswordHash string `json:"password" gorm:"type:varchar(100);not null"`
	Email        string `json:"email" gorm:"unique"`
	Role         int
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.PasswordHash = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
