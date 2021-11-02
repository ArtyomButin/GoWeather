package main

import (
	"github.com/ArtyomButin/GoWeather/pkg/routes"
	postgres "github.com/ArtyomButin/GoWeather/repository/postgres.go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("config initialization failed: %s", err.Error())
	}
	dbConn := new postgres.Config{ConnStr: viper.GetString("db.url")}
	db := postgres.NewPostgresDB()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	routes.Routes(router)
	err := router.Run(":"+viper.GetString("http.port"))
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("main")
	viper.AddConfigPath("./configs")
	return viper.ReadInConfig()
}

