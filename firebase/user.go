package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"log"
	"os"
)

func FirebaseToken(tokens string) (*auth.Token, error) {
	ctx := context.Background()
	credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_CONFIG")))
	if err != nil {
		log.Println(err)
	}

	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(ctx, nil, opt)
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

func FirebaseUser(tokens string) (*auth.UserRecord, error) {
	ctx := context.Background()
	var err error

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
