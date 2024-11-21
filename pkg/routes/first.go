package routes
import (
	"go-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoute=function(router *mux.Router){
	router.HandleFunc("/book/",controllers.createBook).methods("POST")
	router.HandleFunc("/book/",controllers.getBooks).methods("GET")
	router.HandleFunc("/book/{bookId}/",controllers.getBook).methods("GET")
}