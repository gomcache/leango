package controller


import (
	"LearnGoProject/httpServer"
	"github.com/gin-gonic/gin"
)


func Init() *gin.Engine {
	router := gin.Default()
	v2 := router.Group("/user")
	{
		v2.GET("/login", httpServer.Index)
		v2.GET("/submit", httpServer.Index)
	}

	return router
}

