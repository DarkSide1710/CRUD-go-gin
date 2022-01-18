package main

import (
	"DarkSide1710/CRUD-go-gin/http-server/handlers"
	"DarkSide1710/CRUD-go-gin/repository"
	"DarkSide1710/CRUD-go-gin/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	uri := fmt.Sprintf("host=localhost port=5532 user=postgres password=12345 dbname=crud sslmode=disable")
	conn, err := sqlx.Connect("postgres", uri)
	if err != nil {
		panic(err)
	}

	repo := repository.New(conn)
	service := services.New(repo)

	router := handlers.NewRouter(service)

	router.Run(":8080")

}
