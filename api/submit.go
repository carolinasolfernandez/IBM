package api

import (
	"github.com/carolinasolfernandez/IBM/models"
	"github.com/carolinasolfernandez/IBM/services"
	"github.com/gin-gonic/gin"
)

// postEndpoint godoc
// @Summary Endpoint for IBM
// @Accept json
// @Produce json
// @Success 202
// @Failure 400
// @Failure 500
// @Router /submit [post]
func postEndpoint(ctx *gin.Context) {
	var request models.Submit
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, "Request error")
	}
	status := services.CallSubmit(request)

	ctx.JSON(status, gin.H{})
	return
}
