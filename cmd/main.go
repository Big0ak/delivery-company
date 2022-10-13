package main

import (
	"log"
	"os"
	api "github.com/Big0ak/DeliveryCompany"
	handler "github.com/Big0ak/DeliveryCompany/pkg/hendler"
	"github.com/Big0ak/DeliveryCompany/pkg/repository"
	"github.com/Big0ak/DeliveryCompany/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil{
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.GetDB(repository.Config{
		Server:   viper.GetString("db.server"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:  	  viper.GetString("db.port"),
		Database: viper.GetString("db.database"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	//handlers := new(handler.Handler)

	srv := new(api.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
