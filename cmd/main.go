package main

import (
	library_app "library-app"
	"library-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(library_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
