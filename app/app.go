package app

import (
	"fmt"
	"github.com/Dubjay18/green-lit/domain"
	"github.com/Dubjay18/green-lit/service"
	"github.com/Dubjay18/green-lit/utils"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
	"net/url"
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
	constr := "user=" + dbUser + " dbname=" + dbName + " password=" + dbPass + " host=" + dbHost + " port=" + dbPort + " sslmode=require"
	serviceURI := os.Getenv("DATABASE_URL")
	conn, _ := url.Parse(serviceURI)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"
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
	authRepository := domain.NewAuthRepositoryDB(dbClient)
	articleService := service.NewArticleService(articleRepository)
	authService := service.NewLoginService(authRepository, domain.GetRolePermissions())
	userRepository := domain.NewUserRepositoryDB(dbClient)
	userService := service.NewUserService(userRepository)
	userHandler := UserHandler{service: userService}
	articleHandler := ArticleHandler{service: articleService}
	authHandler := AuthHandler{service: authService}
	router := mux.NewRouter()
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins (adjust for production!)
		// ... add other options like AllowedMethods, AllowedHeaders, etc.
	})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Green-Lit API")
		utils.WriteJson(w, http.StatusOK, "Welcome to the Green-Lit API")
	}).Methods(http.MethodGet)

	router.HandleFunc("/users-populate", userHandler.PopulateUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/articles/users/{id:[0-9]+}", articleHandler.GetArticlesByUser).Methods(http.MethodGet)
	router.HandleFunc("/users", userHandler.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/articles", articleHandler.GetAllArticles).Methods(http.MethodGet)
	router.HandleFunc("/articles/{id:[0-9]+}", articleHandler.GetArticle).Methods(http.MethodGet)
	router.HandleFunc("/articles", articleHandler.CreateArticle).Methods(http.MethodPost)
	router.HandleFunc("/auth-signIn", authHandler.Login).Methods(http.MethodPost)
	corsHandler := corsOptions.Handler(router)
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADDRESS")+":"+os.Getenv("SERVER_PORT"), corsHandler))
}
