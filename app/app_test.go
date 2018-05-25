package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestReturn10(t *testing.T) {

	var number = return10()

	if number != 10 {
		t.Error("Expected 10, got ", number)
	}
}

func testJsonResponse(t *testing.T) {

}

func TestHandler(t *testing.T) {

	router := httprouter.New()
	router.GET("/images", handler)

	req, _ := http.NewRequest("GET", "/images", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}
