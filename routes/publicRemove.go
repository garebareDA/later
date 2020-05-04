package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebase"
	"log"
)

func RemovePublic(c *gin.Context) {
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

	db.Where("uuid = ?", uuid).Delete(&database.Public{})

	c.JSON(200, gin.H{
		"status": "削除しました",
	})
}
