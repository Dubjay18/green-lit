package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	r := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))

}
