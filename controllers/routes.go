package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jlettman/dynchan/controllers/boards"
	"github.com/jlettman/dynchan/controllers/user"
)

func Routes(router *gin.Engine) {
	router.Group("/api/v1")
	{
		boards.Routes(router)
		user.Routes(router)
	}
}
