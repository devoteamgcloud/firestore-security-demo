package v1

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/article"
	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/articles"
	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/database"
)

func Router(router gin.IRouter) {

	datasource := database.New()
	group := router.Group("api/v1/articles")
	var contPost = article.PostArticle{Database: datasource}
	group.POST("/article", contPost.View)
	var contList = articles.GetArticles{Database: datasource}
	group.GET("/articles", contList.View)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "127.0.0.1:8080"
		},
		MaxAge: 12 * time.Hour,
	}))
}
