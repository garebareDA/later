package firebase

import(
	"context"
	"log"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func firebaseUser(tokens string) (*auth.Token, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./key/later-d4187-firebase-adminsdk-7n7ps-4dd39b0207.json")
	var err error

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