package api

import (
	"database/sql"
	"net/http"
	"log"

	"github.com/focuscw0w/microservices/user/internal/routes"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	mainRouter := http.NewServeMux()
	apiRouter := http.NewServeMux()

	routes.RegisterRoutes(apiRouter)

	mainRouter.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))

	log.Println("Server is running on :8080")
	return http.ListenAndServe(s.addr, mainRouter)
}