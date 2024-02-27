package models

import "time"

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Hash      string `json:"hash"`
	AvatarURL string
	CreatedAt time.Time
	DeletedAt time.Time
}

type Topic struct {
	ID          int    `json:"id"`
	Name 		string `json:"name"`
	Description string `json:"description"`
	AvatarURL   string
	CreatedAt   time.Time
	DeletedAt   time.Time
}

type Post struct {
	ID        int
	Title     string
	Content   string
	Author    int // User.ID
	Topic     int // Topic.ID
	CreatedAt time.Time
	DeletedAt time.Time
}

type Comment struct {
	ID              int
	Content         string
	ParentCommentID int // Comment.ID
	CreatedAt       time.Time
	DeletedAt       time.Time
}

type Vote struct {
	ID       int
	Positive bool
	UserID   int // User.ID
}
