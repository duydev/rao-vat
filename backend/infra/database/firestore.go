package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var db *firestore.Client

func GetDBInstance() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./service_account_private_key.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	db = client

}

// Firestore note, no need to call when program exit
func CloseDBConnection() {
	if db != nil {
		defer db.Close()
	}
}
