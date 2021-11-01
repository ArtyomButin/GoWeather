package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/postgres", getConnection)
}

func getConnection(c *gin.Context){

}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome",
	})
	return
}
