package handlers

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gugi97/learn-go-redis/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Get data from redis
	key := "book-" + strconv.Itoa(id)
	redisResult, err := h.getDataFromRedis(key)
	if err != redis.Nil {
		log.Printf("Error retrieving value for key '%s'\n", key)
		return
	}

	if err == redis.Nil {
		// Get data from db
		if result := h.DB.First(&book, id); result.Error != nil {
			log.Println(result.Error)
			return
		}
	} else {
		// Unmarshal the JSON string into the struct
		if err = json.Unmarshal([]byte(redisResult), &book); err != nil {
			log.Println("Error:", err)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
