package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/jlettman/dynchan/globals"
	"github.com/jlettman/dynchan/models"
)

func Update(c *gin.Context) {
	var board models.Board
	id := c.Params.ByName("id")

	if err := globals.DB.Where("id = ?", id).First(&board).Error; err != nil {
		c.AbortWithStatus(404)
		log.Error("update error", err)
	}

	c.BindJSON(&board)
	globals.DB.Save(&board)
	c.JSON(200, board)
}
