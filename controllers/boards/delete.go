package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/jlettman/dynchan/globals"
	"github.com/jlettman/dynchan/models"
)

func Delete(c *gin.Context) {
	var board models.Board
	id := c.Params.ByName("id")

	globals.DB.Where("id = ?", id).Delete(&board)
	c.AbortWithStatus(200)
}
