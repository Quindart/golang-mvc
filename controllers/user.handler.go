package controllers

import (
	"example/golang-mvc/models"
	"example/golang-mvc/repositories"
	"example/golang-mvc/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo *repositories.UserRepository
}

func NewUserController(repo *repositories.UserRepository) *UserController {
	return &UserController{
		Repo: repo,
	}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.Repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all users successfully !",
		"users":   users,
	})
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	if err := c.Repo.Create(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create user successfull!",
		"user":    user,
	})
}
func (c *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user := c.Repo.GetByID(id)
 	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user successfull!",
		"user":    user,
	})
}

func (c *UserController) GetToken(ctx *gin.Context) {
	token, err := utils.GenerateToken(1, "Quang")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to generate token",
		})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Generate token success!",
		"token":   token,
	})

}
