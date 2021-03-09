package main

import (
	"task-test/cache"
	"task-test/logger"
	"task-test/router"
	"task-test/utils"
)

func main() {
	utils.InitDB()
	err := cache.Setup()
	if err != nil {
		logger.Error(err.Error())
	}
	utils.InitPath()
	route := router.InitRouter()
	route.Run()
}
