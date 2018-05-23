package main

import (

	// Formatting and Setting up

	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"

	// Image Handle

	// Http and request
	"math/rand"
	"net/http"
	"os"
	"strconv"

	// External frameroks
	"github.com/julienschmidt/httprouter"
)

type Photo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var photos []Photo

// HTTP Handlers

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

// Main function

func main() {

	photos = append(photos, Photo{ID: "nd327dh39w4", Name: "photo324.jpg"})
	photos = append(photos, Photo{ID: "372834whefw", Name: "photo232.jpg"})

	// Router
	router := httprouter.New()
	router.GET("/", indexHandler)       // GET
	router.GET("/photos", getPhotos)    // GET
	router.POST("/photos", createPhoto) // POST
	router.GET("/image", handler)       // GET image from server PATH

	// Working with files

	router.ServeFiles("/src/*filepath", http.Dir("/src"))

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

// Get Image from PATH

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	img := getImage()
	writeImage(w, &img)
}

func getImage() image.Image {

	// Read image from file that already exists
	existingImageFile, err := os.Open("a.png")

	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, imageType, err := image.Decode(existingImageFile)
	if err != nil {
	}

	fmt.Println(imageData)
	fmt.Println(imageType)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handler error
	}

	fmt.Println(loadedImage)
	return loadedImage
}

func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer) // Define a buffer to store an image to requests

	// Error Handler
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")                        // Header to set content type an image
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes()))) // Header for file size

	// Error Handler
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

// Testing

func return10() int {
	return 10
}
