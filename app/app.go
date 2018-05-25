package main

import (

	// Formatting and Setting up

	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"mime/multipart"

	// Http and request

	"net/http"
	"os"
	"strconv"

	// External frameroks
	"github.com/julienschmidt/httprouter"
)

// Main function

func main() {

	// Router
	router := httprouter.New()
	router.GET("/image", handler)      // GET image from server PATH
	router.POST("/upload", UploadFile) // POST image using multipart/form-data

	router

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

	existingImageFile, err := os.Open("a.png")
	if err != nil {
	}

	defer existingImageFile.Close()

	imageData, imageType, err := image.Decode(existingImageFile)

	if err != nil {
	}

	fmt.Println(imageData)
	fmt.Println(imageType)

	existingImageFile.Seek(0, 0)

	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
	}

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

// UploadFile uploads a file to the server
func UploadFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		fmt.Println("Here1")
		return
	}

	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)

		return
	}
	jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

// Testing

func return10() int {
	return 10
}
