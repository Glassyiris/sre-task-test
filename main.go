package main

import (
	"task-test/cache"
	"task-test/router"
	"task-test/utils"
)

func main() {
	utils.InitDB()
	cache.InitRedis()
	utils.InitPath()
	route := router.InitRouter()
	route.Run()
}
