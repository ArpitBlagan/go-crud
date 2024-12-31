package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

	type Data struct{
		Name string `json:"name"`
		Age int `json:"age"`
		Address string `json:"address"`
	}
	type User struct{
		ID uint `json:"id" gorm:"primaryKey"`
		Name string `json:"name"`
		Age int `json:"age"`
		Address string `json:"address"`
		Email string `json:"email"`
	}
	type Mess struct{
		Message string `json:"message"`
	}
	var db *gorm.DB

	func initDB(){
		var err error
		dsn := "host=localhost user=postgres password=yourpassword dbname=testdb port=5432 sslmode=disable"
		db,err =gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if(err!=nil){
			fmt.Println("Error while connecting to DB",err)
			log.Fatal("Not able to connect to DB")
			return
		}
		db.AutoMigrate(&User{})
		fmt.Println("Connect to DB sucessfully :)")
	}

	

	func getInfo(w http.ResponseWriter,r * http.Request){
		// w.Header().Set("Content-Type","pkglication")
		// w.WriteHeader(http.StatusOK)
		// w.Write(res)
		data :=Data{Name:"Arpit Blagan",
		Age:23,
		Address:"Lohana hill, Palampur, Kangra, Himachal Pradesh"}
		res,err:=json.Marshal(data)
		if(err!=nil){
			fmt.Println("Not able parse the body into json format :(")
			return
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	}
	func createUser(w http.ResponseWriter,r *http.Request){
		var user User
		if err:=json.NewDecoder(r.Body).Decode(&user); err!=nil{
			fmt.Println("Error while taking data from the body to create a new user")
			http.Error(w,"Invalid user input",http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		// db.Create(&user)
		if err := db.Create(&user).Error; err != nil {
			// If database insertion fails, return an error
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			fmt.Println("Error while creating user in the database:", err)
			return
		}
		data,err:=json.Marshal(user)
		if(err!=nil){
			http.Error(w,"Error while marshaling the user",http.StatusInternalServerError)
			fmt.Println("Not able to convert the data into json format")
			return
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
	func deleteUser(w http.ResponseWriter,r *http.Request){
		params:=mux.Vars(r)
		var user User
		if err:=db.First(&user,params["id"]).Error;err!=nil{
			http.Error(w,"User not found",http.StatusInternalServerError)
			return
		}
		mess:=Mess {
			Message:"user with given id was deleted sucessfully",
		}
		res,err:=json.Marshal(mess)
		if(err!=nil){
			http.Error(w,"Error while marshaling",http.StatusInternalServerError)
			return
		}
		db.Delete(&user)
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	func updateUser(w http.ResponseWriter,r *http.Request){
		params:=mux.Vars(r)
		var user User
		if err:=db.First(&user,params["id"]).Error;err!=nil{
			http.Error(w,"User not found",http.StatusInternalServerError)
			return
		}
		var updatedUser User
		err:=json.NewDecoder(r.Body).Decode(&updatedUser)
		if(err!=nil){
			http.Error(w,"Not able to decode user form the body",http.StatusInternalServerError)
			return
		}
		//update logic
		user.Name=updatedUser.Name
		//save the db after the update
		db.Save(&user)
		res,err:=json.Marshal(user)
		if(err!=nil){
			http.Error(w,"Not able to marshal the data",http.StatusInternalServerError)
		}
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	func getAllUser(w http.ResponseWriter,r *http.Request){
		var users []User
		if(db==nil){
			http.Error(w, "Error in DB connectivity", http.StatusInternalServerError)
			fmt.Println("DB is not there to run any type of operations")
			return
		}
		if err:=db.Find(&users).Error; err!=nil{
			http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		fmt.Println("Error while retrieving users:", err)
		return
		}
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusOK)
		data,err:=json.Marshal(users)
		if(err!=nil){
			http.Error(w, "Error retrieving users", http.StatusInternalServerError)
			fmt.Println("Error while marshaling the users into json",err)
			return;
		}
		w.Write(data)
	}
	func getUser(w http.ResponseWriter,r *http.Request){
		fmt.Println("Hit")
		var user User
		params:=mux.Vars(r)
		id:=params["id"]
		if err:=db.First(&user,id).Error; err!=nil{
			http.Error(w, "Error retrieving users", http.StatusInternalServerError)
			fmt.Println("Error while get user of pariticular id:",id,err)
			return
		}
		
		res,err:=json.Marshal(user)
		if(err!=nil){
			http.Error(w, "Error while marshaling the reterive data", http.StatusInternalServerError)
			fmt.Println("Error while marshaling the user data")
			return
		}
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	func main(){
		// lets create a simple crud app using Go
		fmt.Println("Hey from the simple crud app :)")
		// Let's first connect to DB...
		initDB()
		r:=mux.NewRouter()
		r.HandleFunc("/create",createUser).Methods("POST")
		r.HandleFunc("/delete/{id}",deleteUser).Methods("DELETE")
		r.HandleFunc("/update/{id}",updateUser).Methods("PUT")
		r.HandleFunc("/getUsers",getAllUser).Methods("GET")
		r.HandleFunc("/getUser/{id}",getUser).Methods("GET")
		err:=http.ListenAndServe(":4000",r)
		if(err!=nil){
			fmt.Println("Error while starting the Go server please check the error",err)
			return;
		}
		fmt.Println("Server is running")
	}

	// func main(){
	// 	fmt.Println("Hello from Golang")
	// 	//1.
	// 	//Another and correct way of doing routing is to use third party lib which other features on top
	// 	//User gorilla/mux
	// 	r:=mux.NewRouter()
	// 	r.HandleFunc("/getInfo",getInfo).Methods("GET")
	// 	err:=http.ListenAndServe(":4000",r)
	// 	if(err!=nil){
	// 		fmt.Println("Getting error while starting the server",err)
	// 		return
	// 	}
	// 	//2.
	// 	// http.HandleFunc("/getInfo",getInfo)
	// 	// err:=http.ListenAndServe(":4000",nil);
	// 	// if(err!=nil){	
	// 	// 	fmt.Println("Error while starting the server at some port")
	// 	// 	return
	// 	// }
	// 	fmt.Println("Server is running...")
	// }