package main

import (
	"example/golang-mvc/infra"
	"example/golang-mvc/router"
	"github.com/gin-gonic/gin"
)

func main() {

	var appDB infra.MyDB

	appDB.InitDB()

	connection := appDB.Connector
	server := gin.Default()

	router.RegisterRouter(server, connection)

	server.Run(":5001")
}
