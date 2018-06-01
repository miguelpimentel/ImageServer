package main

import (

	// OS and File

	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"mime/multipart"

	// HTTP

	"net/http"
	"os"
	"strconv"

	// Router
	"github.com/julienschmidt/httprouter"
)

func main() {

	// Router
	router := httprouter.New()
	router.GET("/image", getImage)    // GET image from server PATH
	router.POST("/upload", postImage) // POST image using multipart/form-data

	// Server
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

func getImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	img := imageHandle()
	writeImage(w, &img)
}

func imageHandle() image.Image {

	existingImageFile, err := os.Open("a.png")
	errorHandler(err)

	defer existingImageFile.Close()

	_, imageType, err := image.Decode(existingImageFile)
	errorHandler(err)

	fmt.Println(imageType)
	existingImageFile.Seek(0, 0)

	loadedImage, err := png.Decode(existingImageFile)
	errorHandler(err)

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

// POST image

func postImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")

	switch mimeType {
	case "image/jpeg":
		saveFile(w, file, handle)
	case "image/png":
		saveFile(w, file, handle)
	default:
		jsonResponse(w, http.StatusBadRequest, "The format file is not valid.")
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
	errorHandler(err)
	jsonResponse(w, http.StatusCreated, "File uploaded successfully!")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

// Error Handler

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
