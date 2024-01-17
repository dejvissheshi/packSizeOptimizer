package routers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"packSizeOptimizer/helpers"
	"packSizeOptimizer/service"
	"strconv"
	"strings"

	"packSizeOptimizer/utils"
)

var defaultItems = []int{250, 500, 1000, 2000, 5000}

func NewRouter(httpHandler HttpHandler) *mux.Router {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/rollback", httpHandler.RollbackPackageChanges)
	myRouter.HandleFunc("/add/{packages}", httpHandler.AddNewPackages)
	myRouter.HandleFunc("/remove/{packages}", httpHandler.RemovePackages)
	myRouter.HandleFunc("/read", httpHandler.ReadPackages)

	myRouter.HandleFunc("/calculate/{items}", httpHandler.CalculatePackages)
	myRouter.HandleFunc("/form/calculate", CalculateData).Methods("POST")

	myRouter.HandleFunc("/", Index)
	myRouter.HandleFunc("/visual/calculate/", CalculateTemplate)
	myRouter.HandleFunc("/submit", SubmitHandler)

	return myRouter
}

// HttpHandler is the builder for the handler functions
type HttpHandler struct {
	PackService service.PackService
}

// RollbackPackageChanges is a handler for the rollback endpoint
func (h HttpHandler) RollbackPackageChanges(w http.ResponseWriter, r *http.Request) {
	err := h.PackService.Rollback(defaultItems)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	numbers, err := h.PackService.Read()
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
	// Extract the "packages" parameter from the URL
	packages := strings.TrimPrefix(r.URL.Path, "/add/")
	packages = strings.TrimSuffix(packages, "/")

	newPackageSize, err := strconv.Atoi(packages)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = h.PackService.Add([]int{newPackageSize})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	numbers, err := h.PackService.Read()
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
	// Extract the "items" parameter from the URL
	packages := strings.TrimPrefix(r.URL.Path, "/remove/")
	packages = strings.TrimSuffix(packages, "/")

	newPackageSize, err := strconv.Atoi(packages)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = h.PackService.Delete([]int{newPackageSize})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	numbers, err := h.PackService.Read()
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
	numbers, err := h.PackService.Read()
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
func (h HttpHandler) CalculatePackages(w http.ResponseWriter, r *http.Request) {
	// Extract the "items" parameter from the URL
	items := strings.TrimPrefix(r.URL.Path, "/calculate/")
	items = strings.TrimSuffix(items, "/")

	itemsOrdered, err := strconv.Atoi(items)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Read packages
	packageSizes, err := h.PackService.Read()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	packages := helpers.Calculate(packageSizes, itemsOrdered)
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

	result := helpers.Calculate(requestData.PackSizes, requestData.Items)
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

// CalculateTemplate is a handler for the visual Calculate endpoint.
// It takes packetSizes and items as input and returns
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

// SubmitRequest is the request body for the submit endpoint
type SubmitRequest struct {
	PackSizes []int `json:"packSizes"`
	Items     int   `json:"items"`
}

// SubmitHandler is a handler for the submit endpoint
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

	var packagesInfo []helpers.PackageInfo
	err = json.Unmarshal(body, &packagesInfo)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	data := struct {
		Packages []helpers.PackageInfo
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

// Index is a handler that renders the index page
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing HTML template", http.StatusInternalServerError)
		return
	}
}
