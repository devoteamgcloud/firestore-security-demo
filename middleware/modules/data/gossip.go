package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
	"github.com/kaan-devoteam/one-click-deploy-demo/modules/data/models"
)

type Gossip struct {
	model models.GossipModel
}

func (g Gossip) Create(token, title, content string) (entity.Gossip, error) {
	urlApi := "https://firestore.googleapis.com/v1/"
	parent := "projects/kaan-sandbox/databases/(default)/documents/gossips"
	key := "?key=AIzaSyC2gQdpHk-rSNgRfvMtNIUccJQ8dy5kMGs"
	urlFinal := fmt.Sprintf("%s%s%s", urlApi, parent, key)
	body := getBodyFromRestModel(title, content)
	response, err := RequestPostWithToken(urlFinal, token, body)
	if err != nil {
		return nil, err
	}
	var createdGossip models.RestDoc
	errUn := json.Unmarshal(response, &createdGossip)
	if errUn != nil {
		log.Error(errUn.Error())
		return nil, errUn
	}
	var data = models.GossipModel{Title: createdGossip.Fields.Title.Value, Content: createdGossip.Fields.Gossip.Value, User: token}
	g.model = data
	return &g, nil
}

func getBodyFromRestModel(title, content string) io.Reader {
	data := models.RestDoc{Fields: models.Rest{Gossip: models.RestValue{Value: content}, Title: models.RestValue{Value: title}}}
	bodyBytes, _ := json.Marshal(data)
	body := bytes.NewBuffer(bodyBytes)
	return body
}

func (g Gossip) Content() string {
	return g.model.Content
}

func (g Gossip) Title() string {
	return g.model.Title
}

func (g Gossip) User() string {
	return g.model.User
}

func (g Gossip) Policy() entity.PolicyCode {
	return entity.PolicyCodeFromString(g.model.Policy)
}
