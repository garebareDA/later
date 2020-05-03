package routes
import (
	"github.com/gin-gonic/gin"
	"later/database"
	"later/firebase"
	"log"
	"strconv"
)

type draftGet struct {
	Id   string    `json: "id"`
	Title string `json: "title"`
}

func DraftsIfinite(c *gin.Context) {
	get := c.Query("number")
	token := c.Query("token")

	if get == "" || token == ""{
		statusError(c, "エラー")
		log.Println("user not login")
	}

	user, err := firebase.FirebaseToken(token)
	if err != nil {
		statusError(c, "ログインしていません")
		log.Println("user not login")
	}

	db, err := database.ConnectDB()
	if err != nil {
		statusError(c, "データベースエラー")
		log.Println("database is closed")
	}
	defer db.Close()

	getNumber, err := strconv.Atoi(get)
	if err != nil {
		statusError(c, "number error")
	}

	drafts := []database.Draft{}
	getDrafts := []draftGet{}
	db.Where("user_id = ? AND id BETWEEN ? AND ?", user.UID, getNumber -10, get).Find(&drafts)
	for _, draft := range drafts{
		getDrafts = append(getDrafts,draftGet{Id:draft.DraftID, Title:draft.Title})
	}
	c.JSON(200, getDrafts)
}