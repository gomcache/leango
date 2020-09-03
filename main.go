package main

import (
	"LearnGoProject/controller"
)

func main() {
	router := controller.Init()
	router.StaticFile("/favicon.ico", "./static/images/favicon.png")
	router.Run(":8080")
}
