package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irron21/book-management-system/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("127.0.0.1:3306", r))

}
