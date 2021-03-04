package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/favicon.svg", "./favicon.svg")
	r.Static("/static", "./static")
	r.StaticFS("/avatar", http.Dir("./avatar"))
	user := r.Group("/user")
	{
		user.POST("/login", CreateJwt)
		user.POST("/register", userRegister)
		user.POST("/update", useProfileUpdate)
		user.GET("/logout", userLogout)
	}

	index := r.Group("/")
	{
		index.GET("", Index)
		index.GET("login", indexLogin)
		index.GET("register", indexRegister)
	}

	return r
}
