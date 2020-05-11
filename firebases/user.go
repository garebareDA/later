package firebases

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"log"
)

//FirebaseToken ユーザーのトークンを取得
func FirebaseToken(tokens string) (*auth.Token, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Println(err)
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		log.Println(err)
	}

	token, err := auth.VerifyIDToken(ctx, tokens)
	if err != nil {
		log.Println(err)
	}

	return token, err
}

//FirebaseUser ユーザーのレコード取得
func FirebaseUser(tokens string) (*auth.UserRecord, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Println(err)
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		log.Println(err)
	}

	token, err := auth.VerifyIDToken(ctx, tokens)
	if err != nil {
		log.Println(err)
	}

	user, err := auth.GetUser(ctx, token.UID)
	if err != nil {
		log.Println(err)
	}

	return user, err
}
