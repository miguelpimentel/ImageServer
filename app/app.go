package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Model

type Photo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var photos []Photo

// HTTP Methods

func getPhotos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(photos)
}

func createPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var photo Photo
	_ = json.NewDecoder(r.Body).Decode(&photo)
	photo.ID = strconv.Itoa(rand.Intn(100000000))
	photos = append(photos, photo)
	json.NewEncoder(w).Encode(photo)
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is the RESTful api")
}

func main() {

	photos = append(photos, Photo{ID: "nd327dh39w4", Name: "photo324.jpg"})
	photos = append(photos, Photo{ID: "372834whefw", Name: "photo232.jpg"})

	// Router
	router := httprouter.New()
	router.GET("/", indexHandler)       // GET
	router.GET("/photos", getPhotos)    // GET
	router.POST("/photos", createPhoto) // POST

	// Trigger server
	env_config()
	http.ListenAndServe(":3003", router)
}

// Enviroment Setup

func env_config() {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}
}

// Test

func return10() int {
	return 10
}
