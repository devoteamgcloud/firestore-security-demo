package database

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kaan-devoteam/firestore-security-demo/settings"
)

func New() *firestore.Client {
	client, err := firestore.NewClient(context.Background(), settings.DatabaseProjectID)
	panicIfError(err)
	return client
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
