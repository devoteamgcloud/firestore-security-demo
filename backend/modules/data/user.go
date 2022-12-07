package data

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
	"github.com/kaan-devoteam/one-click-deploy-demo/modules/data/models"
)

type User struct {
	model    models.UserModel
	id       string
	database *firestore.Client
}

func (u User) Get(db *firestore.Client, id string) entity.User {
	ctx := context.TODO()
	collection := db.Collection("users")
	docRef := collection.Doc(id)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	var user models.UserModel
	err = docSnap.DataTo(&user)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	u.model = user
	u.id = docRef.ID
	u.database = db
	return &u
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Database() *firestore.Client {
	return u.database
}

func (u *User) Role() entity.RoleCode {
	return entity.RoleCodeFromString(u.model.Role)
}
