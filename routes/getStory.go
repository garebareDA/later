package routes
import (
	"github.com/gin-gonic/gin"
	"later/database"
	"log"
)

type storyJSON struct{
	Content string
	Title string
	UserName string
}

//GetStory 投稿一つの取得
func GetStory(c *gin.Context) {
	id := c.Query("uuid")

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー")
	}
	defer db.Close()

	public := database.Public{}
	if db.Where("uuid = ?", id).First(&public).Error != nil {
		log.Println("recoad not found")
		statusError(c, "記事がみつかりませんでした")
	}

	json := storyJSON{Content: public.Content, Title: public.Title, UserName: public.UserName}
	c.JSON(200, json)
}