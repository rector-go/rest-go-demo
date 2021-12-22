package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"rest-go-demo/Controllers"
	"rest-go-demo/database"
	"rest-go-demo/entity"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090...")
	router := mux.NewRouter().StrictSlash(true)
	initializeHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/create", Controllers.Create).Methods("POST")
	router.HandleFunc("/get", Controllers.GetAll).Methods("GET")
	router.HandleFunc("/get/{id}", Controllers.GetById).Methods("GET")
	router.HandleFunc("/update/{id}", Controllers.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", Controllers.Delete).Methods("DELETE")
}

// 初始化数据库
func initDB() {
	config := database.Config{
		ServerName: "192.168.7.22",
		Port:       "3306",
		User:       "root",
		Password:   "123456",
		DB:         "rest-go-demo",
	}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err)
	}
	database.Migrate(&entity.Person{})
}
