package models

import{
	"github.com/jinzhu/gorm"
	"go-crud/pkg/config"
}
var db *gorm.DB

type Book struct{
	gorm.Modle
	Name string `gorm:"json":"name"`
	Author string `gorm:"json":"author"`
	Publication string `gorm:"json":"publication"`
}

function init(){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book())
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b);
	db.Create(b);
	return b;
}

function GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id *int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.where("ID?=",Id).Find(&getBook)
	return &getBook,db
} 

func DeleteBookById(Id *int64) Book{
	var book Book
	db.where("ID?=",Id).Delete(Book)
	return book
}