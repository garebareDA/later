package main

import(
	"later/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static", "./static/")

	router.GET("/", routes.Home)
	router.Run(":8000");
}