package routes

import (
	"net/http"

	"github.com/focuscw0w/microservices/user/internal/handler"
)

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/", handler.HandleHome)
}

