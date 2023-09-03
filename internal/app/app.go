package app

import (
	"log"

	"github.com/zarasfara/pet-adoption-platoform/internal/config"
	"github.com/zarasfara/pet-adoption-platoform/internal/handler"
	"github.com/zarasfara/pet-adoption-platoform/internal/repository"
	"github.com/zarasfara/pet-adoption-platoform/internal/server"
	"github.com/zarasfara/pet-adoption-platoform/internal/service"
)

func Run() {
	cfg := config.Init()

	db, err := repository.NewPostgresDB(*cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer(cfg, handlers.InitRoutes())

	if err := srv.Run(); err != nil  {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}