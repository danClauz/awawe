package model

import "time"

type (
	Post struct {
		ID         uint        `json:"id"`
		UserID     uint        `json:"user_id"`
		Title      string      `json:"title"`
		Content    string      `json:"content"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		User       *User       `json:"user"`
		Categories []*Category `json:"categories"`
		Tags       []*Tag      `json:"tags"`
		Comments   []*Comment  `json:"comments"`
	}

	PostCategory struct {
		PostID     uint `json:"post_id"`
		CategoryID uint `json:"category_id"`
	}

	PostTag struct {
		PostID uint `json:"post_id"`
		TagID  uint `json:"tag_id"`
	}
)

func (Post) TableName() string {
	return "posts"
}

func (PostCategory) TableName() string {
	return "post_category"
}

func (PostTag) TableName() string {
	return "post_tag"
}
