package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey;unique;AUTO_INCREMENT"`
	Username string `json:"username" gorm:"type:varchar(28);unqiue; not null"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
	Email    string `json:"email"`
}
