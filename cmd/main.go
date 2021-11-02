package main

import (
	"github.com/ArtyomButin/GoWeather/pkg/repository"
	"github.com/ArtyomButin/GoWeather/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("config initialization failed: %s", err.Error())
	}
	dbConn := repository.Config{ConnStr: viper.GetString("db.url")}
	db, err := repository.NewPostgresDB(dbConn)
	if err != nil {
		log.Fatalf(err.Error())
	}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	routes.Routes(router)
	err = router.Run(":" + viper.GetString("http.port"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("main")
	viper.AddConfigPath("./configs")
	return viper.ReadInConfig()
}
