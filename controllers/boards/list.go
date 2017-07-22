package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/jlettman/dynchan/globals"
	"github.com/jlettman/dynchan/models"
)

func List(c *gin.Context) {
	var users []models.Board

	if err := globals.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, users)
	}
}
