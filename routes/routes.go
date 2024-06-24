package routes

import (
	"github.com/gin-gonic/gin"
)



func RegisterRoutes(server *gin.Engine){
	server.GET("/event",  ShowEvents )
	server.GET("/event/:id", ShowEvent)
	server.POST("/event", CreateEvent)
  server.PUT("event/:id", updateEvent)
}
