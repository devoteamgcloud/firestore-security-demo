package data

import (
	"encoding/json"
	"fmt"

	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
	"github.com/kaan-devoteam/one-click-deploy-demo/modules/data/models"
)

type Articles struct {
	User entity.User
}

func (a Articles) GetAllArticles(token string) ([]entity.Article, error) {
	urlApi := "https://firestore.googleapis.com/v1/"
	parent := "projects/kaan-sandbox/databases/(default)/documents/articles"
	key := "?key=AIzaSyC2gQdpHk-rSNgRfvMtNIUccJQ8dy5kMGs"
	urlFinal := fmt.Sprintf("%s%s%s", urlApi, parent, key)
	response, err := RequestGetWithToken(urlFinal, token)

	if err != nil {
		return nil, err
	}
	var listArticles models.RestList
	errUn := json.Unmarshal(response, &listArticles)
	if errUn != nil {
		log.Error(errUn.Error())
		return nil, errUn
	}
	var articleList []entity.Article
	for _, doc := range listArticles.Docs {
		article := Article{model: models.ArticleModel{Title: doc.Fields.Title.Value, Content: doc.Fields.Content.Value}}
		articleList = append(articleList, article)
	}
	return articleList, nil
}
