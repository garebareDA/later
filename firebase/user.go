package firebase

import(
	"context"
	"log"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func FirebaseUser(tokens string) (*auth.Token, error) {
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

	return token, err
}