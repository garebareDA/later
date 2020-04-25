package main

import (
	"github.com/gin-gonic/gin"
	"later/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static", "./static/")

	router.GET("/", routes.Home)
	router.Run(":8000")
}
