package main

import (
	"context"
	"fmt"
	"github.com/ArtyomButin/GoWeather/pkg/handlers"
	"github.com/ArtyomButin/GoWeather/pkg/repository"
	"github.com/ArtyomButin/GoWeather/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		ConnStr: fmt.Sprintf("%s://%s:%s@%s:%s/%s",
			viper.GetString("db.driver"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.docker_host"),
			viper.GetString("db.port"),
			viper.GetString("db.dbname"),
		)})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	r := repository.NewRepository(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("http.port")),
		Handler: h.InitRoutes(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	defer db.Close()
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("main")
	viper.AddConfigPath("./configs")
	return viper.ReadInConfig()
}
