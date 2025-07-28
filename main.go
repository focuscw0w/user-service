package main

import (
	"fmt"
	"github.com/focuscw0w/microservices/internal/config"
	middleware2 "github.com/focuscw0w/microservices/internal/middleware"
	"github.com/focuscw0w/microservices/internal/user/security"
	"log"
	"net/http"

	"github.com/focuscw0w/microservices/internal/db"
	"github.com/focuscw0w/microservices/internal/user/handler"
	"github.com/focuscw0w/microservices/internal/user/repository"
	user "github.com/focuscw0w/microservices/internal/user/service"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	handler *handler.Handler
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	security.InitJWT(cfg.SecretKey)

	db, err := db.InitDB("app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	userService := user.NewService(repo)
	apiHandler := handler.NewHandler(userService)

	app := application{handler: apiHandler}

	router := http.NewServeMux()

	router.HandleFunc("POST /sign-up", app.handler.HandleSignUp)
	router.HandleFunc("POST /sign-in", app.handler.HandleSignIn)
	router.HandleFunc("POST /sign-out", app.handler.HandleSignOut)

	router.Handle("PUT /users/update/{id}", middleware2.Authorize(middleware2.CheckPermission(http.HandlerFunc(app.handler.HandleUpdateUser))))
	router.HandleFunc("GET /users", app.handler.HandleGetUsers)
	router.Handle("GET /users/{id}", middleware2.Authorize(middleware2.CheckPermission(http.HandlerFunc(app.handler.HandleGetUser))))
	router.Handle("DELETE /users/{id}", middleware2.Authorize(middleware2.CheckPermission(http.HandlerFunc(app.handler.HandleDeleteUser))))

	stack := middleware2.CreateStack(
		middleware2.Logging,
	)

	addr := fmt.Sprintf(":%s", cfg.Port)
	server := http.Server{
		Addr:    addr,
		Handler: stack(router),
	}

	log.Println("Server running on port: 8080")

	err = server.ListenAndServe()
	if err != nil {
		log.Println("Error while listening to port 8080.")
	}
}
