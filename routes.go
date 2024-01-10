package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"packSizeOptimizer/utils"
	"strconv"
	"strings"

	"packSizeOptimizer/db"
	"packSizeOptimizer/file"
)

var defaultItems = []int{250, 500, 1000, 2000, 5000}

// PackageInterface is a custom interface for HTTP handlers.
type PackageInterface interface {
	RollbackPackageChanges(w http.ResponseWriter, r *http.Request)
	AddNewPackages(w http.ResponseWriter, r *http.Request)
	RemovePackages(w http.ResponseWriter, r *http.Request)
}

// HttpHandler is a type that implements MyHandlerInterface.
type HttpHandler struct {
	Repository db.PackagesRepository
	UseFile    bool
}

// testHandler is a simple handler for the test endpoint
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is the test endpoint!")
}

// RollbackPackageChanges is a handler for the rollback endpoint
func (h HttpHandler) RollbackPackageChanges(w http.ResponseWriter, r *http.Request) {
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
func (h HttpHandler) AddNewPackages(w http.ResponseWriter, r *http.Request) {
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
func (h HttpHandler) RemovePackages(w http.ResponseWriter, r *http.Request) {
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

// ReadPackages is a handler for the read endpoint
func (h HttpHandler) ReadPackages(w http.ResponseWriter, r *http.Request) {
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
	id := strings.TrimPrefix(r.URL.Path, "/calculate/")
	id = strings.TrimSuffix(id, "/")

	itemsOrdered, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	calculator := SingletonCalculator{}
	packages := calculator.GetInstance().Calculate(defaultItems, itemsOrdered)
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

// CalculateData is a handler for the visual Calculate endpoint. It takes packetSizes and items as input and returns
// the combination as JSON response
func CalculateData(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData SubmitRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}

	calculator := SingletonCalculator{}
	result := calculator.GetInstance().Calculate(requestData.PackSizes, requestData.Items)
	responseJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error creating response JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the response JSON to the response writer
	w.Write(responseJSON)
}

func CalculateTemplate(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/calculate_form.html")

	if err != nil {
		http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
		return
	}

	// Execute the template with the data and write to the response
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing HTML template", http.StatusInternalServerError)
		return
	}
}

type SubmitRequest struct {
	PackSizes []int `json:"packSizes"`
	Items     int   `json:"items"`
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	packs := r.FormValue("packSizes")
	packs = strings.ReplaceAll(packs, " ", "")

	packSizes, err := utils.ConvertStringToArray(packs)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if len(packSizes) == 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	itemsStr := r.FormValue("items")
	items, err := strconv.Atoi(itemsStr)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	request := SubmitRequest{
		PackSizes: packSizes,
		Items:     items,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "Failed to marshall structure", http.StatusInternalServerError)
		return
	}

	response, err := http.Post("http://localhost:8080/form/calculate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer response.Body.Close() // Close the response body when done

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Submitted body: ", string(body))

	var packagesInfo []PackageInfo
	err = json.Unmarshal(body, &packagesInfo)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	data := struct {
		Packages []PackageInfo
	}{
		Packages: packagesInfo,
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/submit.html")
	if err != nil {
		http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
		return
	}

	// Execute the template with the data and write to the response
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing HTML template", http.StatusInternalServerError)
		return
	}
}
