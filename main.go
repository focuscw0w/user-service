package main

import (
	"log"
	"net/http"

	"github.com/focuscw0w/microservices/handlers"
	"github.com/focuscw0w/microservices/internal/db"
	"github.com/focuscw0w/microservices/repositories"
	"github.com/focuscw0w/microservices/services"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	handler *handler.Handler
}

func main() {
	db, err := db.InitDB("app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// storage
	storage := repository.NewSqlStorage(db)

	// service
	userService := service.NewUserService(storage)

	// handler
	handler := handler.NewHandler(userService)
	app := application{handler: handler}

	// router
	router := http.NewServeMux()

	router.HandleFunc("GET /", app.handler.HandleHome)
	router.HandleFunc("POST /register", app.handler.Register)
	router.HandleFunc("GET /users", app.handler.GetUsers)

	// server init
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server running on port: 8080")

	err = server.ListenAndServe()
	if err != nil {
		log.Println("Error while listening to port 8080.")
	}
}
