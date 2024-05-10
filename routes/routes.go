package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	var port = 8080
	router := mux.NewRouter()

	log.Println("Application running on port", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
