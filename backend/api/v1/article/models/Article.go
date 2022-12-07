package models

import "github.com/kaan-devoteam/firestore-security-demo/core/entity"

type CreateArticleRequestModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateArticleResponseModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (g *CreateArticleResponseModel) FromEntity(entity entity.Article) *CreateArticleResponseModel {
	g.Content = entity.Content()
	g.Title = entity.Title()
	return g
}
