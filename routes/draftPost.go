package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebase"
	"log"
)

type draftJson struct {
	Token   string `json:"token" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"  binding:"required"`
	DraftID string `json:"draftID"  binding:"required"`
}

func DraftPost(c *gin.Context) {
	var draftPosted draftJson
	c.BindJSON(&draftPosted)
	token := draftPosted.Token
	title := draftPosted.Title
	content := draftPosted.Content
	uuid := draftPosted.DraftID

	if content == "" {
		log.Println("content is empty")
		statusError(c, "記事の中身がからです")
	}

	if title == "" {
		log.Println("title is empty")
		statusError(c, "タイトルがからです")
	}

	user, err := firebase.FirebaseUser(token)
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

	public := database.Public{}
	if !db.Where("uuid = ?", uuid).First(&public).RecordNotFound() {
		log.Println("was public")
		statusError(c, "公開してあるため下書きには保存できません")
	}

	draft := database.Draft{}
	if db.Where("draft_id = ?", uuid).First(&draft).RecordNotFound() {
		draft.Content = content
		draft.UserID = user.UID
		draft.Title = title
		draft.DraftID = uuid
		db.Create(&draft)
	} else {
		db.Model(&draft).Where("draft_id = ?", uuid).Updates(map[string]interface{}{"Title": title, "Content": content})
	}

	c.JSON(200, gin.H{
		"status": "保存しました",
	})
}
