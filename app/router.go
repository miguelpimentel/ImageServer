package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Routes

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"images",
		"GET",
		"/images/{imageId}",
		GetImage,
	},
	Route{
		"images",
		"POST",
		"/images/{imageId}",
		PostImage,
	},
	Route{
		"images",
		"DELETE",
		"/images/{imageId}",
		DeleteImage,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

// Handlers

func PostImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	upload(vars["imageId"], r)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file, err := readImage(vars["imageId"])

	if err != nil {
		fmt.Println("AUxs")
	}

	buffer := new(bytes.Buffer)
	img, _, _ := image.Decode(bytes.NewReader(file))

	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")                        // Header to set content type an image
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes()))) // Header for file size
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO DO: Almost finished")
}
