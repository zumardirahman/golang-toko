package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
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

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", "localhost", "root", "", "gotokodb", "5432")
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed on connecting to the database server")
	}

	server.Router = mux.NewRouter()
	server.initializeRoutes() //untuk initialize di routes
}

//3. unutk run web server
func (server *Server) Run(addr string) {
	fmt.Printf("Listening to Port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erorr on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoToko")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9001")

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
