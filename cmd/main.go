package main

import (
	"context"
	"fmt"
	"github.com/ArtyomButin/GoWeather/configs"
	"github.com/ArtyomButin/GoWeather/pkg/handlers"
	"github.com/ArtyomButin/GoWeather/pkg/repository"
	"github.com/ArtyomButin/GoWeather/pkg/services"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config := configs.GetConfig()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	db, err := repository.NewPostgresDB(repository.Config{
		ConnStr: fmt.Sprintf("%s://%s:%s@%s:%s/%s",
			config.Database.Driver,
			config.Database.Username,
			config.Database.Password,
			config.Database.DockerHost,
			config.Database.Port,
			config.Database.DbName,
		)})
	defer db.Close()
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	r := repository.NewRepository(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Server.Port),
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
}
