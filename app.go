package  app

import (
	u "apiExample/utils"
	"net/http"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"fmt"
)

// JWT Bearer auth
var JwtAuthentication =  func (next http.Handler) http.Handler {
	/*
	*	@params w : Http.ResponseWriter
	*	@params r : Http.Request
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// List of endpoints that doesn't require auth
		notAuth := []string{"/api/user/new", "/api/user/login"}
		// Current request path
		requestPath := r.URL.Path
		// Check if req doesn't need auth & serve it
		for _, value := range notAuth {
			if value == requestPath {
				// w -> http.Writer r -> Http.Request
				next.ServeHTTP(w, r)
				return
			}
		}

		// Init response map
		response := make(map[string]interface{})
		// Get token
		tokenHeader := r.Header.Get("Authorization")

		// If token is empty
		if tokenHeader == "" {
			// Set response usable obj
			response = u.Message(false, "Not a valid token")

			// Set status 403 - Forbidden
			w.WriteHeader(http.StatusForbidden)
			// Set content-type in headers
			w.Header().Add("Content-Type", "application/json")
			// Send response w utils func
			u.Respond(w, response)
			return
		}

		// Remove 'Bearer' keyword from token w split func
		splitted := strings.Split(tokenHeader, " ")
		// If token doesnt contains Bearer
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// Contains token
		tokenPart := splitted[1]
		tk := &models.Token{}

		// Check server signing
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nill
		})

		if err != nil {
			response = u.Message(false, "Malformed token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response )
			return
		}

		if !token.Valid {
			response = u.Message(false, "Token not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// 200 OK
		fmt.Sprintf("User %", tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		// Pass middleware
			// w -> res, r -> req
		next.ServeHTTP(w, r)
	})
}