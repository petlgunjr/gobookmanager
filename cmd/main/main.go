package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinhzu/gorm/dialects/mysql"
	"github.com/petlgunjr/gobookmanager/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
