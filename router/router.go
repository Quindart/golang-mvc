package router

import (
	"database/sql"
	"example/golang-mvc/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(ginServer *gin.Engine, db *sql.DB) {

	ginServer.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok "})
	})

	//User layer
	userController := controllers.NewUserController(db)
	userRouter := NewUserRouter(ginServer, userController)

	userRouter.getAllUser()
}
