package main

import (
	library_app "library-app"
	"library-app/pkg/handler"
	"library-app/pkg/repository"
	"library-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(library_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
