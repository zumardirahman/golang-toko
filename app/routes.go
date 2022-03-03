package app

import "github.com/zumardirahman/golang-toko/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET") //routing tes pertama
}
