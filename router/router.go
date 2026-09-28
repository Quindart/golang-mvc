package router

import (
	"database/sql"
	"example/golang-mvc/controllers"
	"example/golang-mvc/middlewares"
	"example/golang-mvc/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// Health check chung
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	// Khởi tạo các tầng cho User (Dependency Injection)
	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	// Tạo một API Version group (ví dụ /api/v1)
	apiV1 := router.Group("/api/v1")
	users := apiV1.Group("/users")
	{
		users.GET("", middlewares.Authenticate, userController.GetAllUsers)
		users.GET("/token", userController.GetToken)
		users.POST("", userController.CreateUser)
	}

	return router
}
