package gossips

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/gossip/models"
	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/core/usecase"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
)

type GetGossipsResponse struct {
	Gossips []models.CreateGossipResponseModel `json:"gossips"`
}

func (g *GetGossipsResponse) fromEntity(entities []entity.Gossip) *GetGossipsResponse {
	for _, gossip := range entities {
		gossipFetched := models.CreateGossipResponseModel{}
		g.Gossips = append(g.Gossips, *gossipFetched.FromEntity(gossip))
	}
	return g
}

type GetGossips struct {
	Database *firestore.Client
}

func (controller *GetGossips) View(c *gin.Context) {
	header := c.Request.Header["Authorization"]
	if len(header) == 0 {
		c.IndentedJSON(http.StatusUnauthorized, "A valid Firebase token must be provided")
		return
	}
	token := header[0]
	gossips, err := usecase.GetGossips(token)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}
	var gossipsResponse GetGossipsResponse
	gossipsResponse.fromEntity(gossips)
	c.IndentedJSON(http.StatusOK, gossipsResponse)
}

func badRequestIfError(c *gin.Context, err error) {
	log.Error(err.Error())
	c.IndentedJSON(http.StatusBadRequest, err.Error())
}
