package dto

import "time"

type Category struct {
	ID           uint      `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Posts        []*Post   `json:"posts,omitempty"`
}
