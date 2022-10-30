package article

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"

	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/article/models"
	"github.com/kaan-devoteam/one-click-deploy-demo/core/usecase"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
)

type PostArticle struct {
	Database *firestore.Client
}

func (controller *PostArticle) View(c *gin.Context) {
	var newArticle models.CreateArticleRequestModel
	var err error
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	header := c.Request.Header["Authorization"]
	if len(header) == 0 {
		c.IndentedJSON(http.StatusUnauthorized, "A valid Firebase token must be provided")
		return
	}
	token := header[0]
	//log.Info(string(jsonData))
	//log.Info(c.Request.Header["Authorization"][0])
	//log.Info(c.Request.Header["Token"][0])
	if err != nil {
		badRequestIfError(c, err)
	}
	err = json.Unmarshal(jsonData, &newArticle)
	if err != nil {
		badRequestIfError(c, err)
	}
	var articleCreated models.CreateArticleResponseModel
	article, err := usecase.CreateGossip(token, newArticle.Title, newArticle.Content)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}
	articleCreated.FromEntity(article)
	c.IndentedJSON(http.StatusCreated, articleCreated)
}

func badRequestIfError(c *gin.Context, err error) {
	log.Error(err.Error())
	c.IndentedJSON(http.StatusBadRequest, err.Error())
}
