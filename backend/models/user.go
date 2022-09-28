package models

type User struct {
	ID           int    `json:"id" gorm:"primaryKey;unique;AUTO_INCREMENT"`
	Username     string `json:"username" gorm:"type:varchar(28);unqiue; not null"`
	PasswordHash string `json:"password" gorm:"type:varchar(100);not null"`
	Email        string `json:"email"`
	Role         int
}
