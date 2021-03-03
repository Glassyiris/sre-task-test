package main

import (
	"task-test/router"
)

func main() {
	route := router.InitRouter()
	route.Run()
}
