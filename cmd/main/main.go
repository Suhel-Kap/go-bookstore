package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/suhel-kap/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.BookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
