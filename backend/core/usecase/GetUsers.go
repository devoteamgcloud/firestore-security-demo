package usecase

import (
	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/modules/data"
)

func GetUsersByUser(user entity.User) []string {
	users := data.GetAllUsers(user.Database())
	return users
}
