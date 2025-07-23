package main

import (
	"github.com/focuscw0w/microservices/api"
	"log"
	"net/http"

	"github.com/focuscw0w/microservices/internal/db"
	"github.com/focuscw0w/microservices/internal/user/handler"
	"github.com/focuscw0w/microservices/internal/user/repository"
	"github.com/focuscw0w/microservices/internal/user/service"
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

	repo := repository.NewRepository(db)

	userService := service.NewService(repo)

	apiHandler := handler.NewHandler(userService)

	app := application{handler: apiHandler}

	router := http.NewServeMux()

	router.HandleFunc("POST /sign-up", app.handler.HandleSignUp)
	router.HandleFunc("POST /sign-in", app.handler.HandleSignIn)
	router.HandleFunc("POST /sign-out", app.handler.HandleSignOut)
	router.HandleFunc("GET /users", app.handler.HandleGetUsers)
	router.HandleFunc("GET /users/", app.handler.HandleGetUser)
	router.HandleFunc("DELETE /users/", app.handler.HandleDeleteUser)

	server := http.Server{
		Addr:    ":8080",
		Handler: api.Logging(router),
	}

	log.Println("Server running on port: 8080")

	err = server.ListenAndServe()
	if err != nil {
		log.Println("Error while listening to port 8080.")
	}
}
