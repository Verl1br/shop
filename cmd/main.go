package main

import (
	"github.com/dhevve/shop"
	"github.com/dhevve/shop/internal/handler"
	"github.com/dhevve/shop/internal/repository"
	"github.com/dhevve/shop/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := intiConfig(); err != nil {
		logrus.Fatalf("error config: %s", err.Error())
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(shop.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error server %s", err.Error())
	}
}

func intiConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
