package routes

import (
	"net/http"

	"github.com/focuscw0w/microservices/services/user/internal/handler"
)

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/", handler.HandleHome)
}

