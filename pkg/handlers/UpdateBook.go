package handlers

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gugi97/learn-go-redis/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook models.Book
	var book models.Book

	defer r.Body.Close()

	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	json.Unmarshal(body, &updatedBook)

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
		if err = json.Unmarshal([]byte(redisResult), &book); err != nil {
			log.Println("Error:", err)
			return
		}
	}

	// set value for update data
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc

	// update to DB
	h.DB.Save(&book)

	// Marshal the struct into a JSON string
	jsonData, err := json.Marshal(book)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	// update to Redis
	err = h.setDataToRedis(key, string(jsonData))
	if err != nil {
		log.Printf("Error update value for key '%s'\n", key)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
