package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func addBizFirestore(bizDoc map[string]interface{}) {
	// Initialize firestore setting with serviceKey from config
	ctx := context.Background()
	conf := option.WithCredentialsFile("./config/ServiceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, conf)

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// pass the bizDoc JSON to be added inside Firestore
	_, _, err = client.Collection("restaurants").Add(ctx, bizDoc)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	defer client.Close()
}
