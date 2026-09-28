package models

type User struct {
	ID       string `json:"user_id"`
	FullName string `json:"full_name"`
	Age      int
}
