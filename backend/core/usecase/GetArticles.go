package usecase

import (
	"github.com/kaan-devoteam/firestore-security-demo/core/entity"
	"github.com/kaan-devoteam/firestore-security-demo/modules/data"
)

func GetArticles(token string) ([]entity.Article, error) {
	articles, err := data.Articles{}.GetAllArticles(token)
	return articles, err
}
