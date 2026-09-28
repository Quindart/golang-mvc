package main

import (
	"example/golang-mvc/infra"
	"example/golang-mvc/router"
)

func main() {

	var appDB infra.MyDB
	appDB.InitDB()

	connection := appDB.Connector
	r := router.SetupRouter(connection)
	r.Run(":5001")
}
