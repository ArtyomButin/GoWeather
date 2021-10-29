package main

import (
	"github.com/ArtyomButin/GoWeather/pkg/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	routes.Routes(router)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
