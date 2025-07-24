package middleware

import (
	"context"
	"github.com/focuscw0w/microservices/internal/user/security"
	"net/http"
	"strconv"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized - no token", http.StatusUnauthorized)
			return
		}

		userID, err := security.VerifyToken(c.Value)
		if err != nil {
			http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CheckPermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxUserID, ok := r.Context().Value("userID").(int)
		if !ok {
			http.Error(w, "Unauthorized - no user id found", http.StatusUnauthorized)
			return
		}

		pathUserID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Unauthorized - invalid user id found", http.StatusUnauthorized)
			return
		}

		if ctxUserID != pathUserID {
			http.Error(w, "Unauthorized - no permission", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
