package main

import (
	"DarkSide1710/CRUD-go-gin/http-server/handlers"
	"DarkSide1710/CRUD-go-gin/repository"
	"DarkSide1710/CRUD-go-gin/service"
)

func main() {

	repo := repository.New()
	service := services.New(repo)

	router := handlers.NewRouter(service)

	router.Run(":8080")

}
