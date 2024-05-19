package routes

import (
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	var port = 8080
	router := mux.NewRouter()
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login)
	router.HandleFunc("/profile", controller.AddProfile).Methods("POST")
	router.HandleFunc("/seller/{phonenumber}", controller.GetSeller).Methods("GET")

	// product routes
	router.HandleFunc("/product", controller.AddProduct).Methods("POST")
	router.HandleFunc("/products", controller.GetAllProducts)

	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application running on port", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
