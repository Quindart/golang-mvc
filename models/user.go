package models

type User struct {
    ID       int64  `json:"user_id"`
    FullName string `json:"full_name"`
    Age      int    `json:"age"`
    UserName string `json:"user_name" binding:"required,email"`    // Đã sửa (bỏ khoảng trắng)
    Password string `json:"password" binding:"required,min=6"`     // Đã sửa (bỏ khoảng trắng)
}