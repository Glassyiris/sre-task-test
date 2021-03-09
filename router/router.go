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
	store, err := redis.NewStore(100, "tcp", fmt.Sprintf("%s:%d", utils.Configs.Redis.Addr, utils.Configs.Redis.Port), "", []byte(secret))

	if err != nil {
		logger.Error("Redis connect field")
	}
	r.Use(sessions.Sessions("session", store))
	//r.Use(middleware.Auth())
	//load static resource
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/favicon.svg", "./favicon.svg")
	r.Static("/static", "./static")
	r.StaticFS("/avatar", http.Dir("./avatar"))

	//add the router for user's operation
	user := r.Group("/user")
	{
		user.POST("/login", CreateJwt)
		user.POST("/register", userRegister)
		user.POST("/update", middleware.Auth(), useProfileUpdate)
		user.GET("/profile", middleware.Auth(), profile)
		user.POST("/logout", userLogout)
	}

	index := r.Group("/")
	{
		index.GET("", Index)
		index.GET("register", indexRegister)
		index.GET("test", test)
	}

	return r
}
