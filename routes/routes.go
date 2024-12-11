package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/karthik11135/golang-backend/controllers"
	"github.com/karthik11135/golang-backend/middleware"
)

func HandleAllRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.SignupController).Methods("POST")
	r.HandleFunc("/login", controllers.LoginController).Methods("POST")
	r.HandleFunc("/logout", controllers.LogoutController).Methods("GET")
	
	myPlaylistGetHandler := http.HandlerFunc(controllers.GetMyPlaylists)
	myPlaylistPostHandler := http.HandlerFunc(controllers.PostMyPlaylist)
	r.Handle("/playlists", middleware.AuthMiddleware(myPlaylistGetHandler)).Methods("GET")
	r.Handle("/playlists", middleware.AuthMiddleware(myPlaylistPostHandler)).Methods("POST")
}

