package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"utkarsh/Fetch/api/handlers"
	"utkarsh/Fetch/api/models"

	"github.com/gin-gonic/gin"
)

func TestPingRequest(t *testing.T) {
	// Set up a test router with the GetPing Handler function.
	r := gin.Default()
	r.GET("/get", handlers.GetPing)

	// Create a mock HTTP request.
	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder.
	w := httptest.NewRecorder()

	// Perform the request.
	r.ServeHTTP(w, req)

	// Check the status code.
	if w.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, w.Code)
	}

	// Unmarshal JSON data into a IDModel struct
	var onlineReturn models.IDModel
	err = json.Unmarshal([]byte(w.Body.String()), &onlineReturn)

	if err != nil {
		t.Errorf("Invalid return type")
		return
	}
	// Fail if the JSON does not contain Online as ID
	if onlineReturn.ID != "Online" {
		t.Errorf("Invalid return type")
		return
	}
}
