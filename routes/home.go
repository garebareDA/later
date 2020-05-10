package routes

import (
	"github.com/gin-gonic/gin"
	"later/database"
	"log"
	"net/http"
	"strconv"
)

type HomeGets struct {
	ID       string
	Title    string
	Content  string
	UserName string
}

//Home ホームの表示
func Home(c *gin.Context)  {
	c.HTML(http.StatusFound, "index.html", gin.H{})
}

//HomeGet ホームの記事の取得
func HomeGet(c *gin.Context) {
	get := c.Query("number")

	getNumber, err := strconv.Atoi(get)
	if err != nil {
		log.Println("number error")
		statusError(c, "number error", 402)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー", 500)
		return
	}
	defer db.Close()

	publics := []database.Public{}
	getPublic := []HomeGets{}

	err = db.Offset(getNumber - 10).Order("id desc").Limit(10).Find(&publics).Error
	if err != nil {
		log.Println("get number error")
		statusError(c, "データベースエラー", 500)
		return
	}

	if len(publics) == 0 {
		log.Println("empty")
		c.JSON(200, gin.H{"continue": false, "get": getPublic})
		c.Abort()
		return
	}

	for _, public := range publics {
		getPublic = append(getPublic, HomeGets{ID: public.UUID, Title: public.Title, Content: public.Content, UserName: public.UserName})
	}

	c.JSON(200, gin.H{"continue": true, "get": getPublic})
}
