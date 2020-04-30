package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebase"
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

	}

	db, err := database.ConnectDB()
	if err != nil {

	}
	defer db.Close()

	draft := database.Draft{}
	if db.Where("DraftID = ?", uuid).First(&draft).RecordNotFound() {
		draft.Content = content
		draft.UserID = user.UID
		draft.Title = title
		db.Create(&draft)
	} else {
		db.Model(&draft).Where("DraftID = ?", uuid).Updates(map[string]interface{}{"Title": title, "Content": content})
	}

	c.JSON(200, gin.H{
		"status": "posted",
	})
}
