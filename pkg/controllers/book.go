package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"go-crud/pkg/utils"
	"go-crud/pkg/modles"
)

var NewBook models.Book

func GetBooks(w *http.ResponseWritter,r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _:=json.Marsha(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
} 

func GetBookById(w *http.ResponseWritter, r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if(err!=nil){
		fmt.Println("error while parsing")
	}
	bookDetails, _:=models.GetBookById(ID)
	res, _:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}

func CreateBook(w *http.ResponseWritter, r *http.Request){
	CreateBook:=&models.Book{}
	utils.ParseBody(r,CreateBook)
	b:= CreateBook.CreateBook()
	res, _:=json.Marshal(b)
	w.Header.Set("Content-Type","pkglication")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}

func DeleteBookById(w *http.ResponseWritter, r *http.Request){
	vars :=mux.Vars()
	bookid:=vars["bookId"]
	ID, err:=strconv.ParseInt(bookid,0,0)
	if(err!=nil){
		fmt.Println("error while parsing")
	}
	book:=models.DeleteBookById(ID)
	res, _:=json.Marshal(b)
	w.Header.Set("Content-Type","pkglication")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}

func UpdateBook(w *http.ResponseWritter, r *http.Request){
	var updateBook=&models.Book{}
	utils.Parse(r,updateBook)
	vars:=mux.Vars(r)
	bookId=vars["bookId"]
	ID, err:=strconv.ParseInd(bookId,0,0)
	if(err!=nil){
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if(updateBook.Name!=""){
		bookDetails.Name=updateBook.Name
	}
	if(updateBook.Author!=""){
		bookDetails.Author=updateBook.Author
	}
	if(updateBook.Publication!=""){
		bookDetails.Publication=updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _:=json.Marshal(bookDetails)
	w.Header.Set("Content-Type","pkglication")
	w.WriteHeader(http.StatusOk)
	w.Write(res);
}