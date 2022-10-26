package data

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
)

func GetAllUsers(client *firestore.Client) (users []string) {
	ctx := context.TODO()
	collection := client.Collection("users")
	docs := collection.DocumentRefs(ctx)
	refs, err := docs.GetAll()
	if err != nil {
		log.Error(err.Error())
	}
	for _, doc := range refs {
		users = append(users, doc.ID)
	}
	return
}
