package main

import (
	"task-test/router"
	"task-test/utils"
)

func main() {
	utils.CreateTable()
	route := router.InitRouter()
	route.Run()
}
