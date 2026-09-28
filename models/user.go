package models

type User struct {
	ID       int64  `json:"user_id"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
}
