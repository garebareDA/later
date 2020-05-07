package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebase"
	"log"
)

type draftGet struct {
	ID      string `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}

type removesDraft struct {
	Token   string `json:"token" binding:"required"`
	UUID    string `json:"uuid"  binding:"required"`
}

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

func RemoveDraft(c *gin.Context) {
	var removesDraft removesDraft
	c.BindJSON(&removesDraft)
	token := removesDraft.Token
	uuid := removesDraft.UUID

	if token == "" {
		log.Println("user not login")
		statusError(c, "エラー")
	}

	_, err := firebase.FirebaseToken(token)
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

	db.Where("draft_id = ?", uuid).Delete(&database.Draft{})

	c.JSON(200, gin.H{
		"status": "削除しました",
	})
}

func DraftsIfinite(c *gin.Context) {
	get := c.Query("number")
	token := c.Query("token")

	user, getNumber := infiniteAuxiliary(c, get, token)

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー")
	}
	defer db.Close()

	drafts := []database.Draft{}
	getDrafts := []draftGet{}

	err = db.Where("user_id = ? AND id BETWEEN ? AND ?", user.UID, getNumber-10, getNumber-1).Find(&drafts).Error
	if err != nil {
		log.Println("get number error")
		statusError(c, "データベースエラー")
	}

	if len(drafts) == 0 {
		log.Println("arry is empty")
		c.JSON(200, "empty")
		c.Abort()
		return
	}

	for _, draft := range drafts {
		getDrafts = append(getDrafts, draftGet{ID: draft.DraftID, Title: draft.Title, Content: draft.Content})
	}

	last := database.Draft{}
	db.Last(&last)
	if getNumber > last.ID {
		log.Println("empty")
		c.JSON(200, gin.H{"continue":false, "get":getDrafts})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"continue":true, "get":getDrafts})
}