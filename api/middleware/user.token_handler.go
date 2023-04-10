package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/kdsama/book_five/service"
)

type userContextKeyType string

const UserContextKey userContextKeyType = "user_id"

type UserTokenHandler struct {
	service service.UserTokenDI
}

func NewUserTokenHandler(service service.UserTokenDI) *UserTokenHandler {
	return &UserTokenHandler{service}
}

func (uth *UserTokenHandler) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString != "" {
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		}
		id, err := uth.service.ValidateUserTokenAndGetUserID(tokenString)

		// Check for errors and validate the token
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Add the user ID to the request context

		ctx := context.WithValue(r.Context(), UserContextKey, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
