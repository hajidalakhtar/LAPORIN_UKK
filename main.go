package main

import (
	"fmt"
	"laporin_go/routes"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

func main() {

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.Handle("/", routes.SetRoutes())
	fmt.Println("Backend started on port : 9000")
	// log.Fatal(http.ListenAndServe(":9000", routes.SetRoutes()))
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(originsOk, headersOk, methodsOk)(routes.SetRoutes())))
}
