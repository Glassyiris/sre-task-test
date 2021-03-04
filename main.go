package main

import (
	"task-test/router"
	"task-test/utils"
)

func main() {
	utils.CreateTable()
	utils.InitPath()
	route := router.InitRouter()
	route.Run()
}
