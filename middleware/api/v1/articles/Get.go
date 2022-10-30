package articles

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/article/models"
	"github.com/kaan-devoteam/one-click-deploy-demo/core/entity"
	"github.com/kaan-devoteam/one-click-deploy-demo/core/usecase"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
)

type GetArticlesResponse struct {
	Articles []models.CreateArticleResponseModel `json:"articles"`
}

func (g *GetArticlesResponse) fromEntity(entities []entity.Article) *GetArticlesResponse {
	for _, article := range entities {
		articleFetched := models.CreateArticleResponseModel{}
		g.Articles = append(g.Articles, *articleFetched.FromEntity(article))
	}
	return g
}

type GetArticles struct {
	Database *firestore.Client
}

func (controller *GetArticles) View(c *gin.Context) {
	header := c.Request.Header["Authorization"]
	if len(header) == 0 {
		c.IndentedJSON(http.StatusUnauthorized, "A valid Firebase token must be provided")
		return
	}
	token := header[0]
	articles, err := usecase.GetArticles(token)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}
	var articlesResponse GetArticlesResponse
	articlesResponse.fromEntity(articles)
	c.IndentedJSON(http.StatusOK, articlesResponse)
}

func badRequestIfError(c *gin.Context, err error) {
	log.Error(err.Error())
	c.IndentedJSON(http.StatusBadRequest, err.Error())
}
