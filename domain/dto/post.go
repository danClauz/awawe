package dto

import "time"

type (
	Post struct {
		ID         uint        `json:"id"`
		UserID     uint        `json:"user_id,omitempty"`
		Title      string      `json:"title"`
		Content    string      `json:"content"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		User       *User       `json:"author,omitempty"`
		Categories []*Category `json:"categories,omitempty"`
		Tags       []*Tag      `json:"tags,omitempty"`
		Comments   []*Comment  `json:"comments,omitempty"`
	}
)
