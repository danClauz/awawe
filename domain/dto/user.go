package dto

type User struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Password  string  `json:"password,omitempty"`
	Posts     []*Post `json:"posts"`
}
