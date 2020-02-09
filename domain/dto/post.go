package dto

type (
	Post struct {
		ID         uint        `json:"id"`
		UserID     uint        `json:"user_id"`
		Title      string      `json:"title"`
		Content    string      `json:"content"`
		User       *User       `json:"user"`
		Categories []*Category `json:"categories"`
		Tags       []*Tag      `json:"tags"`
		Comments   []*Comment  `json:"comments"`
	}
)
