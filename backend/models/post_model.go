package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID      int    `json:"id" gorm:"primaryKey;unique;AUTO_INCREMENT"`
	Title   string `json:"title"`
	Content string `json:"desc"`
	// Owner of post is a user, OwnerID = foreign key
	OwnerID string
	Owner   User
}
