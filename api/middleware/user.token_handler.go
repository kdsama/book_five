package middleware

import (
	"context"
	"net/http"

	"github.com/kdsama/book_five/service"
)

type userContextKeyType string

const userContextKey userContextKeyType = "user_id"

type UserTokenHandler struct {
	service service.UserTokenDI
}

func NewUserTokenHandler(service service.UserTokenDI) *UserTokenHandler {
	return &UserTokenHandler{service}
}

func (uth *UserTokenHandler) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len("Bearer "):]
		id, err := uth.service.ValidateUserTokenAndGetUserID(tokenString)

		// Check for errors and validate the token
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Add the user ID to the request context
		ctx := context.WithValue(r.Context(), userContextKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
