package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebase"
	"log"
	"strconv"
)

type draftGet struct {
	ID      string `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}

func DraftsIfinite(c *gin.Context) {
	get := c.Query("number")
	token := c.Query("token")

	if get == "" || token == "" {
		log.Println("user not login")
		statusError(c, "エラー")
	}

	user, err := firebase.FirebaseToken(token)
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

	getNumber, err := strconv.Atoi(get)
	if err != nil {
		log.Println("number error")
		statusError(c, "number error")
	}

	drafts := []database.Draft{}
	getDrafts := []draftGet{}

	err = db.Where("user_id = ? AND id BETWEEN ? AND ?", user.UID, getNumber-10, getNumber-1).Find(&drafts).Error
	if err != nil {
		log.Println("get number error")
		statusError(c, "データベースエラー")
	}

	if len(drafts) == 0 {
		log.Println("arry is empty")
		c.JSON(404, "配列が0です")
		c.Abort()
	}

	for _, draft := range drafts {
		getDrafts = append(getDrafts, draftGet{ID: draft.DraftID, Title: draft.Title, Content: draft.Content})
	}

	c.JSON(200, getDrafts)
}
