package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gugi97/learn-go-redis/pkg/models"
	"log"
	"net/http"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	// Use the SCAN command to iterate through keys without a limit
	ctx := context.Background()
	var cursor uint64
	keys := []string{}

	for {
		// Execute the SCAN command
		var err error
		keys, cursor, err = h.Redis.Scan(ctx, cursor, "book-*", 0).Result()
		if err != nil {
			log.Fatalf("Error scanning keys: %v", err)
			return
		}

		// Retrieve the values for each key
		for _, key := range keys {
			var tempBook models.Book
			var value string

			// Get data from Redis
			value, err = h.getDataFromRedis(key)
			if err != redis.Nil {
				log.Printf("Error retrieving value for key '%s'\n", key)
				return
			}

			// Unmarshal the JSON string into the struct
			if err = json.Unmarshal([]byte(value), &tempBook); err != nil {
				log.Println("Error:", err)
				return
			}
			books = append(books, tempBook)
		}

		if cursor == 0 { // no more keys
			break
		}
	}

	// Get data from DB
	if len(books) == 0 {
		if result := h.DB.Find(&books); result.Error != nil {
			log.Println(result.Error)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
