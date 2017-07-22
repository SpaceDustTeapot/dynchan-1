package boards

import (
	"github.com/gin-gonic/gin"
	jsonschema "github.com/mcuadros/go-jsonschema-generator"
)

var SchemaDefinition *jsonschema.Document

func Schema(c *gin.Context) {
	c.JSON(200, SchemaDefinition)
}
