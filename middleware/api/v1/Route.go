package v1

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/database"
	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/gossip"
	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/gossips"
)

func Router(router gin.IRouter) {

	datasource := database.New()
	group := router.Group("api/v1/gossips")
	var contPost = gossip.PostGossip{Database: datasource}
	group.POST("/gossip", contPost.View)
	var contList = gossips.GetGossips{Database: datasource}
	group.GET("/gossips", contList.View)
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
