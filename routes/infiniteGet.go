package routes

import (
	"later/firebases"
	"firebase.google.com/go/auth"
	"log"
	"strconv"
	"errors"
)

func infiniteAuxiliary(get string, token string)(*auth.Token, int, error) {
	var err error
	if get == "" || token == "" {
		log.Println("user not login")
		err = errors.New("ユーザーがログインしていません")
		return nil, 0, err
	}

	user, err := firebases.FirebaseToken(token)
	if err != nil {
		log.Println("user not login")
		err = errors.New("ユーザーがログインしていません")
		return nil, 0, err
	}

	getNumber, err := strconv.Atoi(get)
	if err != nil {
		log.Println("number error")
		err = errors.New("ナンバーエラー")
		return nil, 0, err
	}

	return user, getNumber, err
}
