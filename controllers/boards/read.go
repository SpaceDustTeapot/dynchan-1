package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/jlettman/dynchan/globals"
	"github.com/jlettman/dynchan/models"
)

// Read handles requests for a single Board.
func Read(c *gin.Context) {
	var board models.Board
	id := c.Params.ByName("id")

	if err := globals.DB.Where("id = ?", id).First(&board).Error; err != nil {
		c.AbortWithStatus(404)
		log.Error("read error", err)
	} else {
		c.JSON(200, board)
	}
}
