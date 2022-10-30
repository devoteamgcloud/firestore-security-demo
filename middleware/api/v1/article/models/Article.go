package models

import "github.com/kaan-devoteam/one-click-deploy-demo/core/entity"

type CreateArticleRequestModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	//User    string `json:"user"`
}

type CreateArticleResponseModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	//User    string `json:"user"`
}

func (g *CreateArticleResponseModel) FromEntity(entity entity.Gossip) *CreateArticleResponseModel {
	g.Content = entity.Content()
	g.Title = entity.Title()
	return g
}
