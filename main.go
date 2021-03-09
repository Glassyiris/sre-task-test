package main

import (
	"task-test/router"
	"task-test/utils"
)

func main() {
	utils.InitDB()
	utils.InitPath()
	route := router.InitRouter()
	route.Run()
}
