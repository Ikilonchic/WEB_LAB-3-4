package models

// Post ...
type Post struct {
	ID                int    `json:"id"`
	UserID			  int	 `json:"user"`
	Title	          string `json:"title"`
	Message			  string `json:"message"`
}