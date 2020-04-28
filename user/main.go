package main

import (
	"user/router"
	"user/router/v1"
)

func main() {
	router.Init()
	v1.Init()
	router.GinApp.Run(":8888")
}
