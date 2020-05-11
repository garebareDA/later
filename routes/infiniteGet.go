package routes

import (
	"github.com/gin-gonic/gin"
	"later/firebases"
	"firebase.google.com/go/auth"
	"log"
	"strconv"
)

func infiniteAuxiliary(c *gin.Context, get string, token string)(*auth.Token, int) {
	if get == "" || token == "" {
		log.Println("user not login")
		statusError(c, "エラー", 403)
	}

	user, err := firebases.FirebaseToken(token)
	if err != nil {
		log.Println("user not login")
		statusError(c, "ログインしていません", 403)
	}

	getNumber, err := strconv.Atoi(get)
	if err != nil {
		log.Println("number error")
		statusError(c, "構文が無効です", 400)
	}

	return user, getNumber
}
