package usecase

import (
	"cloud.google.com/go/firestore"
	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/modules/data"
)

func GetUser(db *firestore.Client, id string) (err error, user entity.User) {
	user = data.User{}.Get(db, id)
	if user == nil {
		err := entity.UserNotFoundError{}.New(id)
		return err, nil
	}
	return nil, user
}
