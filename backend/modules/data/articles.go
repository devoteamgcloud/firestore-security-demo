package data

import (
	"encoding/json"
	"fmt"

	"github.com/kaan-devoteam/firestore-security-demo/core/entity"
	"github.com/kaan-devoteam/firestore-security-demo/log"
	"github.com/kaan-devoteam/firestore-security-demo/modules/data/models"
	"github.com/kaan-devoteam/firestore-security-demo/settings"
)

type Articles struct {
	User entity.User
}

func (a Articles) GetAllArticles(token string) ([]entity.Article, error) {
	urlApi := settings.FirestoreRestUrl
	parent := fmt.Sprintf("projects/%s/databases/(default)/documents/articles", settings.DatabaseProjectID)
	key := fmt.Sprintf("?%s", settings.ApiKey)
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
