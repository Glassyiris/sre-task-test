package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/favicon.svg", "./favicon.svg")
	r.Static("/static", "./static")
	user := r.Group("/user")
	{
		user.GET("/login", userLogin)
		user.POST("/register", userRegister)
		user.POST("/update", useProfileUpdate)
	}

	index := r.Group("/")
	{
		index.GET("", Index)
		index.GET("login", indexLogin)
		index.GET("register", indexRegister)
	}

	return r
}
