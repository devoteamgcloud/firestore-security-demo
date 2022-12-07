package usecase

import (
	"github.com/kaan-devoteam/firestore-security-demo/core/entity"
	"github.com/kaan-devoteam/firestore-security-demo/modules/data"
)

func CreateArticle(token, title, content string) (entity.Article, error) {
	return data.Article{}.Create(token, title, content)
}
