package dto

type Tag struct {
	ID      uint    `json:"id"`
	TagName string  `json:"tag_name"`
	Posts   []*Post `json:"post,omitempty"`
}
