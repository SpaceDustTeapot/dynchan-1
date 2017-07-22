package user

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/users", List)
	router.GET("/users/:id", Read)
}
