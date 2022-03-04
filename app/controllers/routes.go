package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", server.Home).Methods("GET") //routing tes pertama
	server.Router.HandleFunc("/products", server.Products).Methods("GET")

	//setup public directory
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public").Handler(staticFileHandler).Methods("GET")

}
