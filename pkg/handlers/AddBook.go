package handlers

import (
	"encoding/json"
	"github.com/gugi97/learn-go-redis/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var book models.Book
	json.Unmarshal(body, &book)

	// Insert to DB
	if result := h.DB.Create(&book); result.Error != nil {
		log.Println(result.Error)
		return
	}

	// Marshal the struct into a JSON string
	jsonData, err := json.Marshal(book)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	// Insert to Redis
	key := "book-" + strconv.Itoa(book.Id)
	err = h.setDataToRedis(key, string(jsonData))
	if err != nil {
		log.Printf("Error add value for key '%s'\n", key)
		return
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
