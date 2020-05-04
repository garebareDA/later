package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebase"
	"log"
	"strconv"
)

type publicGet struct {
	ID      string `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}


func PublicInfnite(c *gin.Context) {
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

	publics := []database.Public{}
	getPublic := []publicGet{}
	err = db.Where("user_id = ? AND id BETWEEN ? AND ?", user.UID, getNumber-10, getNumber-1).Find(&publics).Error

	if err != nil {
		log.Println("get number error")
		statusError(c, "データベースエラー")
	}

	if len(publics) == 0 {
		log.Println("arry is empty")
		c.JSON(404, "配列が0です")
		c.Abort()
	}

	for _, public := range publics {
		getPublic = append(getPublic, publicGet{ID: public.UUID, Title: public.Title, Content: public.Content})
	}
	c.JSON(200, getPublic)
}