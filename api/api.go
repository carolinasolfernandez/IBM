package api

import (
	_ "github.com/carolinasolfernandez/IBM/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//SetupRouter setea las rutas y sus handlers
func SetupRouter(r *gin.Engine) {

	v1 := r.Group("/api")
	{
		v1.POST("/submit", postEndpoint)
	}
	r.GET("/api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // ../api/doc/index.html
}
