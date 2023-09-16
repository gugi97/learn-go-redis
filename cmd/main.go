package main

import (
	"github.com/gorilla/mux"
	"github.com/gugi97/learn-go-redis/pkg/db"
	"github.com/gugi97/learn-go-redis/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	DB := db.InitPostgres()
	redis := db.InitRedisClient()
	h := handlers.New(DB, redis)
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running in port 4000...")
	http.ListenAndServe(":4000", router)
}
