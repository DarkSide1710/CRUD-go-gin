package main

import (
	"DarkSide1710/CRUD-go-gin/http-server/handlers"
	"DarkSide1710/CRUD-go-gin/repository"
	"DarkSide1710/CRUD-go-gin/service"
	"database/sql"
)

func main() {

	db, _ := sql.Open("postgres", "postgres", "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable")
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	m.Up()

	repo := repository.New()
	service := services.New(repo)

	router := handlers.NewRouter(service)

	router.Run(":8080")

}
