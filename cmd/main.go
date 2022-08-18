package main

import (
	"boosters/api/config"
	"boosters/api/pkg/handler"
	"boosters/api/pkg/repository"
	"boosters/api/pkg/router"
	"boosters/api/pkg/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// @title MiniBlog
// @version 1.0
// @description API for Blog can make CRUD operation with posts

// @host localhost:8000
// @BasePath /

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(".env file dont load :", err)
	}

	c := config.GetConfig()

	db, err := repository.PostgresConnect(c)
	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	hand := handler.NewHandler(srv)

	router := router.NewRouter(*hand)
	router.StartRouter()
}
