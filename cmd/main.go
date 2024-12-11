package main

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/karthik11135/golang-backend/db"
	"github.com/karthik11135/golang-backend/routes"
)

func main() {
	// res := db.ConnectDb()
	// if !res {
	// 	fmt.Println("Could not is iso connect to database")
	// 	return
	// }

	r := mux.NewRouter()

	routes.HandleAllRoutes(r)

	log.Fatal(http.ListenAndServe(":3000", r))
}