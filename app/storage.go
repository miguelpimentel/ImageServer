package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"cloud.google.com/go/storage"
)

// Google Cloud Storage

// Read

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

func readImage(filename string) ([]byte, error) {

	fmt.Println("dasdasd")

	bucket := "nexte-profile-images"
	context := context.Background()
	client, err := storage.NewClient(context)

	// [START download_file]
	rc, err := client.Bucket(bucket).Object(filename).NewReader(context)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Write

func upload(filename string, r *http.Request) {

	bucket := "nexte-profile-images"
	context := context.Background()
	client, err := storage.NewClient(context)

	if err != nil {
		log.Fatal(err)
	}

	if err := write(client, bucket, filename, r); err != nil {
		log.Fatalf("Cannot write object: %v", err)
	}
}

func write(client *storage.Client, bucket, filename string, r *http.Request) error {

	ctx := context.Background()

	file, _, err := r.FormFile("file") // Key able to be used in request
	if err != nil {
		return err
	}
	defer file.Close()

	wc := client.Bucket(bucket).Object(filename).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	// [END upload_file]
	return nil
}
