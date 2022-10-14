package main

import "Assignment3/router"

func main() {
	startRoute := router.StartApp()
	startRoute.Run(":8080")
}
