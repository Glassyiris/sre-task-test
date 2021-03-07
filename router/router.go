package router

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-test/logger"
	"task-test/middleware"
	"task-test/utils"
)

var secret = "timeToExit"

func InitRouter() *gin.Engine {
	r := gin.Default()
	store, err := redis.NewStore(100, "tcp", fmt.Sprintf("%s:%s", utils.Configs.Redis.Addr, utils.Configs.Redis.Port), "", []byte(secret))

	if err != nil {
		logger.Error("Redis connect field")
	}
	r.Use(sessions.Sessions("session", store))

	//load static resource
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/favicon.svg", "./favicon.svg")
	r.Static("/static", "./static")
	r.StaticFS("/avatar", http.Dir("./avatar"))

	//add the router for user's operation
	user := r.Group("/user")
	user.Use(middleware.Auth())
	{
		user.POST("/login", CreateJwt)
		user.POST("/register", userRegister)
		user.POST("/update", useProfileUpdate)
		user.GET("/logout", userLogout)
	}

	index := r.Group("/")
	{
		index.GET("", Index)
		index.GET("register", indexRegister)
	}

	return r
}
