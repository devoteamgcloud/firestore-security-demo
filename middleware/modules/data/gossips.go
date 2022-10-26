package data

import (
	"encoding/json"
	"fmt"

	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
	"github.com/kaan-devoteam/one-click-deploy-demo/modules/data/models"
)

type Gossips struct {
	User entity.User
}

func (g Gossips) GetAllGossips(token string) ([]entity.Gossip, error) {
	urlApi := "https://firestore.googleapis.com/v1/"
	parent := "projects/kaan-sandbox/databases/(default)/documents/gossips"
	key := "?key=AIzaSyC2gQdpHk-rSNgRfvMtNIUccJQ8dy5kMGs"
	urlFinal := fmt.Sprintf("%s%s%s", urlApi, parent, key)
	response, err := RequestGetWithToken(urlFinal, token)

	if err != nil {
		return nil, err
	}
	var listGossips models.RestList
	errUn := json.Unmarshal(response, &listGossips)
	if errUn != nil {
		log.Error(errUn.Error())
		return nil, errUn
	}
	var gossipList []entity.Gossip
	for _, doc := range listGossips.Docs {
		gossi := Gossip{model: models.GossipModel{Title: doc.Fields.Title.Value, Content: doc.Fields.Gossip.Value}}
		gossipList = append(gossipList, gossi)
	}
	return gossipList, nil
}

/*
func (g Gossips) GetByUsers(users []string) (gossipList []entity.Gossip) {
	client := g.User.Database()
	gossips := client.Collection("gossips")
	var docs []*firestore.DocumentSnapshot
	for _, u := range users {
		docs = append(docs, getByUser(gossips, u)...)
		for _, doc := range docs {
			var g models.GossipModel
			err := doc.DataTo(&g)
			if err != nil {
				log.Error(err.Error())
			}
			gossipList = append(gossipList, Gossip{model: g})
		}
	}
	return
}

func getByUser(collection *firestore.CollectionRef, user string) []*firestore.DocumentSnapshot {
	ctx := context.TODO()
	docs, err := collection.Where("user", "==", user).Documents(ctx).GetAll()
	if err != nil {
		log.Error(err.Error())
	}
	return docs
}
*/
