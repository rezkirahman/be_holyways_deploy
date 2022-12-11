package main

import (
	"os"
	"fmt"
	"net/http"
	"holyways/database"
	"holyways/pkg/mysql"
	"holyways/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code

	// Setup allowed Header, Method, and Origin for CORS on this below code ...
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})
	var DB_PORT = os.Getenv("DB_PORT")
	//var port = "5000"
	fmt.Println("server running:" + DB_PORT)

	// Embed the setup allowed in 2 parameter on this below code ...
	http.ListenAndServe(":"+DB_PORT, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
