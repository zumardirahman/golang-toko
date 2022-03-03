package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//1. bikin struct
type Server struct {
	DB     *gorm.DB    //db nya
	Router *mux.Router //routernya
}

//2. membuat metod baru
func (server *Server) Initialize() {
	fmt.Println("Welcome to Golang Toko") //menampilkajn pesan

	server.Router = mux.NewRouter()
	server.initializeRoutes() //untuk initialize di routes
}

//3. unutk run web server
func (server *Server) Run(addr string) {
	fmt.Printf("Listening to Port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}

	server.Initialize()
	server.Run(":9001")
}
