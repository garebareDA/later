package routes

import(
	"github.com/gin-gonic/gin"
	"later/firebase"
	"later/database"
)

type draftJson struct {
	Token string `json:"token" binding:"required"`
	Title string `json:"title" binding:"required"`
	Content string `json:"content"  binding:"required"`
}

func DraftPost(c *gin.Context) {
	var draftPosted draftJson
	c.BindJSON(&draftPosted)
	token := draftPosted.Token
	title := draftPosted.Title
	content := draftPosted.Content

	user, err := firebase.FirebaseUser(token)
	if err != nil {

	}

	db, err := database.ConnectDB()
	if err != nil {

	}
	defer db.Close()

	draft := database.Draft{}
	draft.Content = content
	draft.UserID = user.UID
	draft.Title = title

	db.Create(&draft)
	c.JSON(200, gin.H{
		"status": "posted",
	})
}