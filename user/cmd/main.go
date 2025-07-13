package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/focuscw0w/microservices/user/internal/handler"
	"github.com/focuscw0w/microservices/user/internal/service"
	"github.com/focuscw0w/microservices/user/internal/store"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	handler *handler.Handler
}

func main() {
	// database
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
		return
	}
	defer db.Close()

	// storage
	storage := store.NewSqlStorage(db)

	// service
	userService := service.NewUserService(storage)

	// handler
	handler := handler.NewHandler(userService)
	app := application{handler: handler}

	// router
	router := http.NewServeMux()

	router.HandleFunc("GET /", app.handler.HandleHome)
	router.HandleFunc("POST /register", app.handler.Register)

	// server init
	server := http.Server {
		Addr: ":8080",
		Handler: router,
	}

	log.Println("Server running on port: 8080")

	err = server.ListenAndServe()
	if err != nil {
		log.Println("Error while listening to port 8080.")
	}
}
