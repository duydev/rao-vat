package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App
var ctx context.Context
var db *firestore.Client

func GetDBInstance() (context.Context, *firestore.Client) {
	var err error

	if ctx == nil {
		ctx = context.Background()
	}

	if app == nil {
		sa := option.WithCredentialsFile("../service_account_private_key.json")
		app, err = firebase.NewApp(ctx, nil, sa)

		if err != nil {
			log.Fatalln(err)
		}
	}

	if db == nil {
		db, err = app.Firestore(ctx)

		if err != nil {
			log.Fatalln(err)
		}
	}

	return ctx, db
}

// Firestore note, no need to call when program exit
func CloseDBConnection() {
	if db != nil {
		defer db.Close()
	}
}
