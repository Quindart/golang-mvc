package router

import (
	"example/golang-mvc/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	Engine  *gin.Engine
	Handler *controllers.UserController
}

func NewUserRouter(engine *gin.Engine,
	handler *controllers.UserController) *UserRouter {
	return &UserRouter{
		Engine:  engine,
		Handler: handler,
	}
}

func (userRouter *UserRouter) getAllUser() {
	userRouter.Engine.GET(
		"/users",
		func(c *gin.Context) { userRouter.Handler.GetAllUsers(c) },
	)
}

// create user
func (userRouter *UserRouter) createUser() {
	userRouter.Engine.POST(
		"/users",
		func(c *gin.Context) { userRouter.Handler.CreateUser(c) },
	)
}
