package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"packSizeOptimizer/file"
)

var defaultItems = []int{250, 500, 1000, 2000, 5000}

// testHandler is a simple handler for the test endpoint
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is the test endpoint!")
}

// RollbackPackageChanges is a handler for the rollback endpoint
func RollbackPackageChanges(w http.ResponseWriter, r *http.Request) {
	err := file.RollbackFileToInitialState("data.csv", defaultItems)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	numbers, err := file.ReadNumbersFromCSV("data.csv")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the response to JSON
	responseJSON, err := json.Marshal(numbers)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.Write(responseJSON)
}

// AddNewPackages is a handler for the add endpoint
func AddNewPackages(w http.ResponseWriter, r *http.Request) {
	// Extract the "id" parameter from the URL
	id := strings.TrimPrefix(r.URL.Path, "/add/")
	id = strings.TrimSuffix(id, "/")

	newPackageSize, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = file.AddNumbersToCSV("data.csv", []int{newPackageSize})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	numbers, err := file.ReadNumbersFromCSV("data.csv")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the response to JSON
	responseJSON, err := json.Marshal(numbers)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.Write(responseJSON)
}

// RemovePackages is a handler for the delete endpoint
func RemovePackages(w http.ResponseWriter, r *http.Request) {
	// Extract the "id" parameter from the URL
	id := strings.TrimPrefix(r.URL.Path, "/remove/")
	id = strings.TrimSuffix(id, "/")

	newPackageSize, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = file.DeleteNumbersFromCSV("data.csv", []int{newPackageSize})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	numbers, err := file.ReadNumbersFromCSV("data.csv")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the response to JSON
	responseJSON, err := json.Marshal(numbers)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.Write(responseJSON)
}

// CalculatePackages is a handler for the Calculate endpoint
func CalculatePackages(w http.ResponseWriter, r *http.Request) {
	// Extract the "id" parameter from the URL
	id := strings.TrimPrefix(r.URL.Path, "/Calculate/")
	id = strings.TrimSuffix(id, "/")

	itemsOrdered, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	packages := Calculate(itemsOrdered, defaultItems)
	// Create a response JSON
	response := packages

	// Convert the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.Write(responseJSON)
}
