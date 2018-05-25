package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"

	"github.com/julienschmidt/httprouter"
)

func TestReturn10(t *testing.T) {

	var number = return10()

	if number != 10 {
		t.Error("Expected 10, got ", number)
	}
}

func testJsonResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	jsonResponse(rr, http.StatusOK, "The test has passed")
	//testMessage := "The test has passed"

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong message")
	}

}

func TestGetImage(t *testing.T) {

	router := httprouter.New()
	router.GET("/images", getImage)

	req, _ := http.NewRequest("GET", "/images", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}

func TestRedirectPostImage(t *testing.T) {
	testRouter := httprouter.New()
	testRouter.POST("/upload", postImage)

	req, _ := http.NewRequest("GET", "/upload", nil)
	rr := httptest.NewRecorder()
	t.Log(rr.Code)

	testRouter.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Wrong status")
	}
}

func TestPostImageWithoutImage(t *testing.T) {

	router := httprouter.New()
	router.POST("/upload", postImage)

	req, _ := http.NewRequest("POST", "/upload", nil)
	rr := httptest.NewRecorder()
	t.Log(rr.Code)

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}



func TestPostImageWithImage(t *testing.T) {

	router := httprouter.New()
	router.POST("/upload", postImage)

	image, _ := os.Open("a.png")
	req, _ := http.NewRequest("POST", "/upload", image)
	rr := httptest.NewRecorder()
	t.Log(rr.Code)

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}


}

func testSaveFile(t *testing.T) {

	
	
}


