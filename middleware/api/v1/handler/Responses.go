package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaan-devoteam/one-click-deploy-demo/api/v1/models"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
)

func UnprocessableEntityResponse(err error, ctx *gin.Context) {
	log.Debug("couldn't parse request with error %s", err)
	ctx.JSON(
		http.StatusBadRequest, &models.Response[models.ErrorResponse]{
			Code: "UnprocessableEntity",
			Content: models.ErrorResponse{
				Field: "query",
				Issue: err.Error(),
			},
		},
	)
}
