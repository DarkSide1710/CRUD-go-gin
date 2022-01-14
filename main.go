package main

import (
	"DarkSide1710/CRUD-go-gin/handlers"
	repository "DarkSide1710/CRUD-go-gin/repository"
	services "DarkSide1710/CRUD-go-gin/service"
)

func main() {

	repo := repository.Contact()
	service := services.New(repo)

	router := handlers.NewRouter(service)

	router.Run(":8080")

}
