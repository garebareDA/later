package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebases"
	"log"
)

type publicJSON struct {
	Token   string `json:"token" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"  binding:"required"`
	UUID    string `json:"uuid"  binding:"required"`
}

type publicGet struct {
	ID      string
	Title   string
	Content string
}

//PublicPost 記事の公開
func PublicPost(c *gin.Context) {
	var publicJSON publicJSON
	c.BindJSON(&publicJSON)
	token := publicJSON.Token
	title := publicJSON.Title
	content := publicJSON.Content
	uuid := publicJSON.UUID

	if content == "" {
		log.Println("content is empty")
		statusError(c, "記事の中身がからです", 500)
		return
	}

	if title == "" {
		log.Println("title is empty")
		statusError(c, "タイトルがからです", 500)
		return
	}

	user, err := firebases.FirebaseToken(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません", 402)
		return
	}

	profile, err := firebases.FirebaseUser(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません", 402)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー", 500)
		return
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
		c.JSON(201, gin.H{
			"status": "公開しました",
		})
	} else {
		db.Model(&public).Where("uuid = ?", uuid).Updates(map[string]interface{}{"Title": title, "Content": content})
		c.JSON(201, gin.H{
			"status": "更新しました",
		})
	}
}

//RemovePublic 公開記事の削除
func RemovePublic(c *gin.Context) {
	var removesDraft removesDraft
	c.BindJSON(&removesDraft)
	token := removesDraft.Token
	uuid := removesDraft.UUID

	if token == "" {
		log.Println("user not login")
		statusError(c, "エラー", 402)
		return
	}

	_, err := firebases.FirebaseToken(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません", 402)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー", 500)
		return
	}
	defer db.Close()

	if db.Where("uuid = ?", uuid).Delete(&database.Public{}).Error != nil {
		statusError(c, "削除できませんでした", 500)
	}

	c.JSON(200, gin.H{
		"status": "削除しました",
	})
}

//PublicInfnite ユーザの公開記事の一覧
func PublicInfnite(c *gin.Context) {
	get := c.Query("number")
	token := c.Query("token")

	user, getNumber := infiniteAuxiliary(c, get, token)

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー", 500)
		return
	}
	defer db.Close()

	publics := []database.Public{}
	getPublic := []publicGet{}

	err = db.Offset(getNumber - 10).Where("user_id = ?", user.UID).Order("id desc").Limit(10).Find(&publics).Error
	if err != nil {
		log.Println("get number error")
		statusError(c, "データベースエラー", 500)
		return
	}

	if len(publics) == 0 {
		log.Println("empty")
		c.JSON(200, gin.H{"continue":false, "get":getPublic})
		c.Abort()
		return
	}

	for _, public := range publics {
		getPublic = append(getPublic, publicGet{ID: public.UUID, Title: public.Title, Content: public.Content})
	}

	c.JSON(200, gin.H{"continue":true, "get":getPublic})
}