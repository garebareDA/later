package routes

import (
	"github.com/gin-gonic/gin"
	"later/firebase"
	"firebase.google.com/go/auth"
	"log"
	"strconv"
)

func infiniteAuxiliary(c *gin.Context, get string, token string)(*auth.Token, int) {
	if get == "" || token == "" {
		log.Println("user not login")
		statusError(c, "エラー")
	}

	user, err := firebase.FirebaseToken(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません")
	}

	getNumber, err := strconv.Atoi(get)
	if err != nil {
		log.Println("number error")
		statusError(c, "number error")
	}

	return user, getNumber
}
