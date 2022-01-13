package main

import (
	"example/web-service-gin/handlers"
	"example/web-service-gin/service"
)

func main() {

	services := services.NewServices()

	router := handlers.NewRouter(services)

	router.Run(":8080")

}
