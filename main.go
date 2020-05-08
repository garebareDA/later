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
	db.CreateTable(&database.Public{})
	db.CreateTable(&database.Like{})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static", "./static/")

	//routesの中身を全体的にリファクタリング
	//URLをまとめる
	router.GET("/", routes.Home)
	router.GET("/draft",routes.DraftsIfinite)
	router.GET("/public", routes.PublicInfnite)
	router.GET("/likes", routes.LikeInfiniteGet)
	router.GET("/like", routes.LikeGet)
	router.GET("/story", routes.GetStory)

	router.POST("/draft", routes.DraftPost)
	router.POST("/public", routes.PublicPost)
	router.POST("/like", routes.LikePost)

	router.DELETE("/draft", routes.RemoveDraft)
	router.DELETE("/public", routes.RemovePublic)
	router.DELETE("/like", routes.LikeRemove)

	router.Run(":8000")
}
