package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", index)
}

func index(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index page",
		"author": "John Smith",
	})
	return
}
