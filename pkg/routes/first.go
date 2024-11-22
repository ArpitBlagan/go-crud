package routes
import (
	"go-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoute=function(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).methods("POST")
	router.HandleFunc("/book/",controllers.GetBooks).methods("GET")
	router.HandleFunc("/book/{bookId}/",controllers.GetBookById).methods("GET")
	router.HandleFunc("/book/{bookId}/",controllers.DeleteBookById).methods("DELETE")
	router.HandleFunc("/book/{bookId}/",controllers.UpdateBook).methods("PUT")
}