package boards

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	if SchemaDefinition == nil {
		//SchemaDefinition = &jsonschema.Document{}
		//SchemaDefinition.Read(&models.Board{})
	}

	router.GET("/schema:boards", Schema)

	router.GET("/boards", List)
	router.POST("/boards", Create)

	router.GET("/boards/:id", Read)
	router.PUT("/boards/:id", Update)
	router.DELETE("/boards/:id", Delete)
}
