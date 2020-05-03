package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebase"
	"log"
)

type publicJson struct {
	Token   string `json:"token" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"  binding:"required"`
	UUID    string `json:"uuid"  binding:"required"`
}

func PublicPost(c *gin.Context) {
	var publicJson publicJson
	c.BindJSON(&publicJson)
	token := publicJson.Token
	title := publicJson.Title
	content := publicJson.Content
	uuid := publicJson.UUID

	if content == "" {
		log.Println("content is empty")
		statusError(c, "記事の中身がからです")
	}

	if title == "" {
		log.Println("title is empty")
		statusError(c, "タイトルがからです")
	}

	user, err := firebase.FirebaseToken(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません")

	}

	profile, err := firebase.FirebaseUser(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません")

	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー")

	}
	defer db.Close()

	draft := database.Draft{}
	if !db.Where("draft_id = ?", uuid).First(&draft).RecordNotFound() {
		db.Where("draft_id = ?", uuid).Delete(&draft)
	}

	public := database.Public{}
	if db.Where("uuid = ?", uuid).First(&public).RecordNotFound() {
		public.UserID = user.UID
		public.UserName = profile.DisplayName
		public.Title = title
		public.Content = content
		public.Like = 0
		public.UUID = uuid
		db.Create(&public)
		c.JSON(200, gin.H{
			"status": "公開しました",
		})
	} else {
		db.Model(&public).Where("uuid = ?", uuid).Updates(map[string]interface{}{"Title": title, "Content": content})
		c.JSON(200, gin.H{
			"status": "更新しました",
		})
	}
}
