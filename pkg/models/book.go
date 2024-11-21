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