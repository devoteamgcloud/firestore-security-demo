package usecase

import (
	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/modules/data"
)

func CreateGossip(token, title, content string) (entity.Gossip, error) {
	return data.Gossip{}.Create(token, title, content)
}
