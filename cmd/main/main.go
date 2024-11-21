package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialect/mysql"
	"go-crud/pkg/routes"
)

func main(){
	r :=mux.NewRoute()
	routes.RegisterBookStoreRoute(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}