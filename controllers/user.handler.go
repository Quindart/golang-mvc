package controllers

import (
	"database/sql"
	"example/golang-mvc/models"
	_ "example/golang-mvc/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

func (userController *UserController) GetAllUsers(context *gin.Context) {

	query := "select id, fullName, age from users"
	rows, err := userController.DB.Query(query)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch user",
		})
		return
	}
	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Age,
		)
		if err != nil {
			fmt.Println("Error when parse user")
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to fetch user",
			})
		}
		users = append(users, user)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Get all users successfully !",
		"users":   users,
	})

}
