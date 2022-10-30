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

type Article struct {
	model models.ArticleModel
}

func (a Article) Create(token, title, content string) (entity.Article, error) {
	urlApi := "https://firestore.googleapis.com/v1/"
	parent := "projects/kaan-sandbox/databases/(default)/documents/articles"
	key := "?key=AIzaSyC2gQdpHk-rSNgRfvMtNIUccJQ8dy5kMGs"
	urlFinal := fmt.Sprintf("%s%s%s", urlApi, parent, key)
	body := getBodyFromRestModel(title, content)
	response, err := RequestPostWithToken(urlFinal, token, body)
	if err != nil {
		return nil, err
	}
	var createdArticle models.RestDoc
	errUn := json.Unmarshal(response, &createdArticle)
	if errUn != nil {
		log.Error(errUn.Error())
		return nil, errUn
	}
	var data = models.ArticleModel{Title: createdArticle.Fields.Title.Value, Content: createdArticle.Fields.Content.Value, User: token}
	a.model = data
	return &a, nil
}

func getBodyFromRestModel(title, content string) io.Reader {
	data := models.RestDoc{Fields: models.Rest{Content: models.RestValue{Value: content}, Title: models.RestValue{Value: title}}}
	bodyBytes, _ := json.Marshal(data)
	body := bytes.NewBuffer(bodyBytes)
	return body
}

func (a Article) Content() string {
	return a.model.Content
}

func (a Article) Title() string {
	return a.model.Title
}

func (a Article) User() string {
	return a.model.User
}

func (a Article) Policy() entity.PolicyCode {
	return entity.PolicyCodeFromString(a.model.Policy)
}
