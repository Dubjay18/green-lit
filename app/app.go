package app

import (
	"github.com/Dubjay18/green-lit/domain"
	"github.com/Dubjay18/green-lit/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func GetEnvVar() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

}

func SanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined...")
	}
}

func getDbClient() *sqlx.DB {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	constr := "user=" + dbUser + " dbname=" + dbName + " password=" + dbPass + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"

	db, err := sqlx.Open("postgres", constr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)

	}
	return db
}

func Start() {
	GetEnvVar()
	SanityCheck()
	dbClient := getDbClient()
	articleRepository := domain.NewArticleRepositoryDB(dbClient)
	articleService := service.NewArticleService(articleRepository)
	articleHandler := ArticleHandler{service: articleService}
	router := mux.NewRouter()
	router.HandleFunc("/articles", articleHandler.GetAllArticles).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADDRESS")+":"+os.Getenv("SERVER_PORT"), router))
}
