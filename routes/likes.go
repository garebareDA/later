package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebase"
	"log"
)

type likePosts struct {
	Token   string `json:"token" binding:"required"`
	UUID string `json:"uuid"  binding:"required"`
}

func LikePost(c *gin.Context){
	var likePosts likePosts
	c.BindJSON(&likePosts)

	token := likePosts.Token
	uuid := likePosts.UUID

	if token == "" {
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

	like := database.Like{}
	if !db.Where("uuid = ? AND user_id = ?", uuid, user.UID).First(&like).RecordNotFound() {
		log.Println("like found")
		statusError(c, "すでにいいねしてあります")
	}

	like.UUID = uuid
	like.UserID = user.UID
	db.Create(&like)

	c.JSON(200, gin.H{"status": "ok"})
}

func LikeRemove(c *gin.Context) {
	var likePosts likePosts
	c.BindJSON(&likePosts)

	token := likePosts.Token
	uuid := likePosts.UUID

	if token == "" {
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

	err = db.Where("uuid = ? AND user_id = ?", uuid, user.UID).Delete(&database.Like{}).Error
	if err != nil{
		log.Println("like delete error")
		statusError(c, "いいねを取り消せません")
	}

	c.JSON(200, gin.H{"status":"OK"})
}

func LikeGet(c *gin.Context){
	id := c.Query("uuid")
	token := c.Query("token")

	if token == ""{
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

	like := database.Like{}
	if db.Where("uuid = ? AND user_id = ?", id, user.UID).First(&like).RecordNotFound() {
		c.JSON(200, gin.H{"status":"not like"})
		c.Abort()
	}

	c.JSON(200, gin.H{"status":"like"})
}