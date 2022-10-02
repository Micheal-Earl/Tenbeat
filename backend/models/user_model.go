package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Raw Struct for defining database schema
type User struct {
	gorm.Model
	// ID           int    `json:"id" gorm:"primaryKey;unique;AUTO_INCREMENT"`
	Username     string `json:"username" gorm:"type:varchar(28);unique;not null"`
	PasswordHash string `json:"password" gorm:"type:varchar(100);not null"`
	Email        string `json:"email" gorm:"unique"`
	Role         int
}

// Lightweight struct omitting confidential info for JSON responses
type SanitizedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u *User) SanitizeUser() SanitizedUser {
	return SanitizedUser{
		Username: u.Username,
		Email:    u.Email,
	}
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	u.PasswordHash = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
