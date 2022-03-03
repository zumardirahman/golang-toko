package app

import (
	"github.com/gorilla/mux"
	"github.com/zumardirahman/golang-toko/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET") //routing tes pertama
}
