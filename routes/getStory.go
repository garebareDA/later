package routes
import (
	"github.com/gin-gonic/gin"
	"later/database"
	"log"
)

type storyJson struct{
	Content string `json: "content"`
	Title string `json: "title"`
	UserName string `json: "userName"`
}

func GetStory(c *gin.Context) {
	id := c.Query("uuid")

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database is closed")
		statusError(c, "データベースエラー")
	}
	defer db.Close()

	public := database.Public{}
	if db.Where("uuid = ?", id).First(&public).RecordNotFound() {
		log.Println("recoad not found")
		statusError(c, "記事がみつかりませんでした")
	}

	json := storyJson{Content: public.Content, Title: public.Title, UserName: public.UserName}
	c.JSON(200, json)
}