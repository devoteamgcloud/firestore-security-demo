package models

import "github.com/kaan-devoteam/one-click-deploy-demo/core/entity"

type CreateGossipRequestModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	//User    string `json:"user"`
}

type CreateGossipResponseModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	//User    string `json:"user"`
}

func (g *CreateGossipResponseModel) FromEntity(entity entity.Gossip) *CreateGossipResponseModel {
	g.Content = entity.Content()
	g.Title = entity.Title()
	return g
}
