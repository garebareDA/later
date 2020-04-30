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

	user, err := firebase.FirebaseUser(token)
	if err != nil {
		statusError(c, "ログインしていません")
		log.Println("user not login")
	}

	db, err := database.ConnectDB()
	if err != nil {
		statusError(c, "データベースエラー")
		log.Println("database is closed")
	}
	defer db.Close()

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
