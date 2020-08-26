package router

import (
	"MyfirstGO/router/middleware"
	"MyfirstGO/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(g *gin.Engine){
	//middllwares := []gin.HandlerFunc{}
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound,"wrong path")
	})
	router:=g.Group("/user")
	{
		router.POST("/addUser",service.AddUser)
		router.POST("/selectUser",service.SelectUser)
	}
}
