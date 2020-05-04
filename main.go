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
	db.CreateTable(&database.Public{});

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static", "./static/")

	router.GET("/", routes.Home)
	router.GET("/drafts",routes.DraftsIfinite)
	router.GET("/publics", routes.PublicInfnite)
	router.POST("/draft", routes.DraftPost)
	router.POST("/public", routes.PublicPost)
	router.DELETE("/draft/remove", routes.RemoveDraft)
	router.DELETE("/public/remove", routes.RemovePublic)
	router.Run(":8000")
}
