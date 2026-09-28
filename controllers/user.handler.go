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
			"message": "Failed to fetch user: " + err.Error(),
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
				"message": "Failed to fetch user: " + err.Error(),
			})
		}
		users = append(users, user)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Get all users successfully !",
		"users":   users,
	})

}

func (userController *UserController) CreateUser(context *gin.Context) {

	var user models.User
	context.ShouldBindJSON(&user)

	query := `
	insert into users (fullName, age) values (?, ?)
	`
	stmt, err := userController.DB.Prepare(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create user:" + err.Error(),
		})
		return
	}

	result, err := stmt.Exec(user.FullName, user.Age)

	id, err := result.LastInsertId()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create user:" + err.Error(),
		})
		return
	}
	user.ID = id
	context.JSON(http.StatusCreated, gin.H{
		"message": "Create user successfull!",
		"user":    user,
	})
}
