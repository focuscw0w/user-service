package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/focuscw0w/microservices/user/internal/service"
	"github.com/focuscw0w/microservices/user/internal/store"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	userService *internal.UserService
}

func (app *application) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("create user handler"))
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
	
	// create service
	userService := internal.NewUserService(storage)

	// create handler
	app := application{userService: userService}

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Homepage!")
	})
	http.HandleFunc("POST /user", app.handleCreateUser)


	// server init
	log.Println("Server starting on port: 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error while listening to port 8080.")
	}
}