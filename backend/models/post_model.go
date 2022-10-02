package models

import (
	"time"

	"gorm.io/gorm"
)

// Raw Struct for defining database schema
type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	// Owner of post is a user, OwnerID = foreign key
	OwnerID uint
	Owner   User
}

// Lightweight struct omitting confidential info for JSON responses
type SanitizedPost struct {
	ID        uint          `json:"id"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Owner     SanitizedUser `json:"owner"`
}

func (p *Post) SanitizePost() SanitizedPost {
	return SanitizedPost{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Title:     p.Title,
		Content:   p.Content,
		Owner:     p.Owner.SanitizeUser(),
	}
}
