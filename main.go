package main

import (
	"github.com/Yadier01/golang-rest-api/db"
	"github.com/Yadier01/golang-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

  routes.RegisterRoutes(server)
	server.Run()
}
