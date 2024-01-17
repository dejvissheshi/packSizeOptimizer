package routers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"packSizeOptimizer/file"
	"strconv"
	"testing"
)

func TestAddNewPackages_Success(t *testing.T) {
	// Create a request with a sample ID
	req, err := http.NewRequest("GET", "/add/42", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	File := file.PackFiles{
		Filename: "fixtures/data.csv",
	}
	handler := HttpHandler{
		File,
	}

	handler.AddNewPackages(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `[42]`
	if body := rr.Body.String(); body != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", body, expected)
	}

	RemovePackage(t, 42)
}

func TestAddNewPackages_Failure_BadRequest(t *testing.T) {
	// Create a request with a sample ID
	req, err := http.NewRequest("GET", "/add/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	File := file.PackFiles{
		Filename: "fixtures/data.csv",
	}
	handler := HttpHandler{
		File,
	}

	handler.AddNewPackages(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
	}
}

func TestRemovePackages_Failure_BadRequest(t *testing.T) {

	CreateNewPackage(t, 42)

	File := file.PackFiles{
		Filename: "fixtures/data.csv",
	}
	handler := HttpHandler{
		File,
	}

	req, err := http.NewRequest("GET", "/remove/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler.RemovePackages(rr, req)
	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
	}
}

func TestReadPackages_Success(t *testing.T) {

	CreateNewPackage(t, 42)
	CreateNewPackage(t, 24)

	// Create a request with a sample ID
	req, err := http.NewRequest("GET", "/read", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	File := file.PackFiles{
		Filename: "fixtures/data.csv",
	}
	handler := HttpHandler{
		File,
	}

	handler.ReadPackages(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `[24,42]`
	if body := rr.Body.String(); body != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", body, expected)
	}

	RemovePackage(t, 42)
	RemovePackage(t, 24)
}

func TestCalculateData_WrongMethod(t *testing.T) {

	CreateNewPackage(t, 42)
	CreateNewPackage(t, 24)

	// Create a request with a sample ID
	req, err := http.NewRequest("GET", "form/calculate", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	CalculateData(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusMethodNotAllowed)
	}

	RemovePackage(t, 42)
	RemovePackage(t, 24)
}

func TestCalculateData_Success(t *testing.T) {

	CreateNewPackage(t, 42)
	CreateNewPackage(t, 24)

	body := SubmitRequest{
		Items:     42,
		PackSizes: []int{42, 24},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request with a sample ID
	req, err := http.NewRequest("POST", "/form/calculate/", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	CalculateData(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `[{"PackageSize":42,"Quantity":1}]`
	if body := rr.Body.String(); body != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", body, expected)
	}

	RemovePackage(t, 42)
	RemovePackage(t, 24)
}

func CreateNewPackage(t *testing.T, pack int) {
	// Create a request with a sample ID
	req, err := http.NewRequest("GET", "/add/"+strconv.Itoa(pack), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	File := file.PackFiles{
		Filename: "fixtures/data.csv",
	}
	handler := HttpHandler{
		File,
	}

	handler.AddNewPackages(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func RemovePackage(t *testing.T, pack int) {
	// Create a request with a sample ID
	req, err := http.NewRequest("GET", "/remove/"+strconv.Itoa(pack), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	File := file.PackFiles{
		Filename: "fixtures/data.csv",
	}
	handler := HttpHandler{
		File,
	}

	handler.RemovePackages(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}
