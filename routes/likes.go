package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebases"
	"log"
)

type likePosts struct {
	Token   string `json:"token" binding:"required"`
	UUID string `json:"uuid"  binding:"required"`
}

//LikePost いいねの投稿
func LikePost(c *gin.Context){
	var likePosts likePosts
	c.BindJSON(&likePosts)

	token := likePosts.Token
	uuid := likePosts.UUID

	if token == "" {
		log.Println("user not login")
		statusError(c, "ログインしていません", 403)
		return
	}

	user, err := firebases.FirebaseToken(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません", 403)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー", 500)
		return
	}
	defer db.Close()

	like := database.Like{}
	if !db.Where("uuid = ? AND user_id = ?", uuid, user.UID).First(&like).RecordNotFound() {
		log.Println("like found")
		statusError(c, "すでにいいねしてあります", 500)
		return
	}

	like.UUID = uuid
	like.UserID = user.UID
	db.Create(&like)

	c.JSON(201, gin.H{"status": "ok"})
}

//LikeGet いいねの取得
func LikeGet(c *gin.Context){
	id := c.Query("uuid")
	token := c.Query("token")

	if token == ""{
		log.Println("user not login")
		statusError(c, "ログインしていません", 204)
	}

	user, err := firebases.FirebaseToken(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません", 204)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー", 500)
	}
	defer db.Close()

	like := database.Like{}
	if db.Where("uuid = ? AND user_id = ?", id, user.UID).First(&like).RecordNotFound() {
		c.JSON(200, gin.H{"status":"not like"})
		c.Abort()
	}

	c.JSON(201, gin.H{"status":"like"})
}

//LikeRemove いいねの削除
func LikeRemove(c *gin.Context) {
	var likePosts likePosts
	c.BindJSON(&likePosts)

	token := likePosts.Token
	uuid := likePosts.UUID

	if token == "" {
		log.Println("user not login")
		statusError(c, "ログインしていません", 204)
		return
	}

	user, err := firebases.FirebaseToken(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません", 204)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー", 500)
		return
	}
	defer db.Close()

	err = db.Where("uuid = ? AND user_id = ?", uuid, user.UID).Delete(&database.Like{}).Error
	if err != nil{
		log.Println("like delete error")
		statusError(c, "いいねを取り消せません", 500)
		return
	}

	c.JSON(200, gin.H{"status":"OK"})
}

//LikeInfiniteGet ユーザのいいねした一覧
func LikeInfiniteGet(c *gin.Context){
	get := c.Query("number")
	token := c.Query("token")

	user, getNumber ,err := infiniteAuxiliary(get, token)
	if err != nil{
		statusError(c, "ユーザー認証エラー", 402)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー", 500)
		return
	}
	defer db.Close()

	likes := []database.Like{}
	likesGet := []draftGet{}

	err = db.Offset(getNumber - 10).Where("user_id = ?", user.UID,).Order("id desc").Limit(10).Find(&likes).Error
	if err != nil {
		log.Println("get number error")
		statusError(c, "データベースエラー", 500)
		return
	}

	if len(likes) == 0 {
		log.Println("empty")
		c.JSON(200, gin.H{"continue":false, "get":likesGet})
		c.Abort()
		return
	}

	for _, like := range likes {
		public := database.Public{}
		if db.Where("uuid = ?", like.UUID).First(&public).Error != nil {
			log.Println("recoad not found")
			if db.Where("uuid = ?", like.UUID).Delete(&database.Like{}).Error != nil{
				log.Println("likes delete error")
				statusError(c, "データベースエラー", 500)
				return
			}
			log.Println("like delete")
			continue
		}
		likesGet = append(likesGet, draftGet{ID:public.UUID, Title:public.Title , Content: public.Content})
	}

	c.JSON(200, gin.H{"continue":true, "get":likesGet})
}