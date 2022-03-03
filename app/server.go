package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

//1. bikin struct
type Server struct {
	DB     *gorm.DB    //db nya
	Router *mux.Router //routernya
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

//2. membuat metod baru
func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Welcome to " + appConfig.AppName) //menampilkajn pesan

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
	var appConfig = AppConfig{}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erorr on loading .env file")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
