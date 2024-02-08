package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"utkarsh/Fetch/api/handlers"
	"utkarsh/Fetch/api/models"

	"github.com/gin-gonic/gin"
)

// Test if API is online
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

// Send a JSON get ID then send ID and get points. Check if points are correct.
func TestJSON1(t *testing.T) {
	// Set up a test router with the GetPing Handler function.
	r := gin.Default()
	r.POST("/post", handlers.GetID)
	r.GET("/reciepts/:id/points", handlers.GetPoints)

	file, err := os.Open("test/test1.json")
	if err != nil {
		t.Fatal(err)
		return
	}
	defer file.Close()

	// Decode the JSON data into a struct
	var reciept1 models.RecieptModel
	err = json.NewDecoder(file).Decode(&reciept1)
	if err != nil {
		t.Fatal(err)
		return
	}
	// Convert struct to JSON
	jsonData, err := json.Marshal(reciept1)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP request with some JSON data.
	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

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

	id := onlineReturn.ID

	// Create a mock HTTP request. send the recieved id to get points
	req1, err1 := http.NewRequest("GET", "/reciepts/"+id+"/points", nil)
	if err1 != nil {
		t.Fatal(err1)
	}

	// Create a mock HTTP response recorder.
	w2 := httptest.NewRecorder()

	// Perform the request.
	r.ServeHTTP(w2, req1)

	// Check the status code.
	if w2.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, w.Code)
	}

	// Unmarshal JSON data into a PointsModel struct
	var pointsModel models.PointsModel
	err = json.Unmarshal([]byte(w2.Body.String()), &pointsModel)

	if err != nil {
		t.Errorf("Invalid return type")
		return
	}
	// Fail if points not equal to expected val
	if pointsModel.Points != 28 {
		t.Errorf("Invalid return value")
		return
	}

}
