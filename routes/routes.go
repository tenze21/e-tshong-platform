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
	router.HandleFunc("/sellerdetail/{phonenumber}", controller.GetSellerDetails).Methods("GET")
	router.HandleFunc("/profile/{phonenumber}", controller.UpdateProfile).Methods("PUT")
	router.HandleFunc("/seller/{phonenumber}", controller.UpdateSellerDetails).Methods("PUT")
	router.HandleFunc("/logout", controller.Logout)

	// product routes
	router.HandleFunc("/product/{phonenumber}", controller.AddProduct).Methods("POST")
	router.HandleFunc("/products", controller.GetAllProducts)
	router.HandleFunc("/products/{phonenumber}", controller.GetProducts).Methods("GET")
	router.HandleFunc("/product/{pid}", controller.GetProduct).Methods("GET")
	router.HandleFunc("/product/{pid}", controller.DeleteProduct).Methods("DELETE")

	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application running on port", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
