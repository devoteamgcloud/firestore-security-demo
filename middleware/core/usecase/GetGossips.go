package usecase

import (
	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/modules/data"
)

func GetGossips(token string) ([]entity.Gossip, error) {
	gossips, err := data.Gossips{}.GetAllGossips(token)
	return gossips, err
}
