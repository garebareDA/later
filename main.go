package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"later/routes"
	"later/database"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.CreateTable(&database.Draft{})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static", "./static/")

	router.GET("/", routes.Home)
	router.POST("/draft", routes.DraftPost)
	router.Run(":8000")
}
