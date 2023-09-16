package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gugi97/learn-go-redis/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Get data from redis
	key := "book-" + strconv.Itoa(id)
	_, err := h.getDataFromRedis(key)
	if err != redis.Nil {
		log.Printf("Error retrieving value for key '%s'\n", key)
		return
	}

	// Delete the key from redis
	if err != redis.Nil {
		err = h.Redis.Del(context.Background(), key).Err()
		if err != nil {
			log.Printf("Error deleting key '%s': %v", key, err)
			return
		}
	}

	// Find the book by ID from db
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		log.Println(result.Error)
		return
	}

	// Delete that book from db
	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
