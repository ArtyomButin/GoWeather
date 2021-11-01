package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/postgres", getConnection)
}

func getConnection(c *gin.Context){
	pool, err := pgxpool.Connect(c, "postgres://golang_user:golang_pass@go-db:5432/weather")
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer pool.Close()
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome",
		"pool": pool,
	})
	return
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome",
	})
	return
}
