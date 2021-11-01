package main

import (
	"github.com/ArtyomButin/GoWeather/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("config initialization failed: %s", err.Error())
	}
	router := gin.Default()
	routes.Routes(router)
	err := router.Run(":"+viper.GetString("http.port"))
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config.yml")
	return viper.ReadInConfig()
}
