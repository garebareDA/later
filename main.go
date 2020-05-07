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
	router.GET("/drafts",routes.DraftsIfinite)
	router.GET("/publics", routes.PublicInfnite)
	router.GET("/story", routes.GetStory)
	router.GET("/like", routes.LikeGet)
	router.GET("/likes", routes.LikeInfiniteGet)

	router.POST("/draft", routes.DraftPost)
	router.POST("/public", routes.PublicPost)
	router.POST("/like", routes.LikePost)
	router.DELETE("/draft/remove", routes.RemoveDraft)
	router.DELETE("/public/remove", routes.RemovePublic)
	router.DELETE("/like", routes.LikeRemove)
	router.Run(":8000")
}
