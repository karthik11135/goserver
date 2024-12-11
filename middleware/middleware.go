package middleware

import (
	"fmt"

	// "github.com/karthik11135/golang-backend/middleware"

	// "github.com/joho/godotenv"
	"encoding/json"

	"net/http"

	"github.com/gorilla/context"
	"github.com/karthik11135/golang-backend/helpers"
)


func AuthMiddleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, errCookie := r.Cookie("token")

		if errCookie != nil {
			json.NewEncoder(w).Encode("Noo value in the cookie")
			return
		}

		_, err := helpers.VerifyToken(cookie.Value)

		if err != nil {
			json.NewEncoder(w).Encode("Could not verify")
			return 
		}

		fmt.Println("aodfasdfads", cookie.Value, "cookieval")

		decodedVal, _ := helpers.DecodeToken(cookie.Value)

		fmt.Println("Decoded value", decodedVal)

		if decodedVal == "" {
			// json.NewEncoder(w).Encode("could not decode the token")
			return 
		}

		context.Set(r, "email", decodedVal)

		originalHandler.ServeHTTP(w, r)
	})
}