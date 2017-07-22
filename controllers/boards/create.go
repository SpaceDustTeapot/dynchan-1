package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/jlettman/dynchan/globals"
	"github.com/jlettman/dynchan/models"
)

func Create(c *gin.Context) {
	var board models.Board

	c.BindJSON(&board)
	globals.DB.Create(&board)
	c.JSON(200, board)
}
