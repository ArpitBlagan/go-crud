package routes
import (
	"go-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoute=func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}/",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}/",controllers.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/book/{bookId}/",controllers.UpdateBook).Methods("PUT")
}