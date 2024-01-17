package file

import (
	"reflect"
	"sort"
	"testing"
)

func Test_ReadNumbersFromCSV(t *testing.T) {
	files := PackFiles{
		Filename: "fixtures/file_1.csv",
	}
	// Test case 1: Valid CSV file with numbers
	numbers, err := files.Read()
	if err != nil {
		t.Errorf("Error reading numbers from CSV file: %v", err)
	}
	expectedNumbers := []int{250, 400, 500, 600}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected numbers %v, but got %v", expectedNumbers, numbers)
	}

	files.Filename = "fixtures/file_2_invalid_data.csv"
	// Test case 2: Invalid CSV file with non-numeric values
	_, err = files.Read()
	if err == nil {
		t.Error("Expected error reading numbers from invalid CSV file, but got none")
	}
}

func Test_AddNumbersToCSV(t *testing.T) {
	files := PackFiles{
		Filename: "fixtures/file_3.csv",
	}
	// Test case 1: Valid CSV file with numbers
	err := files.Add([]int{700, 800})
	if err != nil {
		t.Errorf("Error adding numbers to CSV file: %v", err)
	}

	numbers, err := files.Read()
	sort.Ints(numbers)
	if err != nil {
		t.Errorf("Error reading numbers from CSV file: %v", err)
	}

	expectedNumbers := []int{250, 400, 500, 600, 700, 800}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected numbers %v, but got %v", expectedNumbers, numbers)
	}

	err = files.Rollback([]int{250, 400, 500, 600})
	if err != nil {
		t.Errorf("Error rolling back test file to initial state: %v", err)
	}
}

func Test_DeleteNumbersFromCSV(t *testing.T) {
	files := PackFiles{
		Filename: "fixtures/file_4.csv",
	}
	// Test case 1: Valid CSV file with numbers
	err := files.Delete([]int{250, 400})
	if err != nil {
		t.Errorf("Error deleting numbers from CSV file: %v", err)
	}

	numbers, err := files.Read()
	sort.Ints(numbers)
	if err != nil {
		t.Errorf("Error reading numbers from CSV file: %v", err)
	}

	expectedNumbers := []int{500, 600}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected numbers %v, but got %v", expectedNumbers, numbers)
	}

	err = files.Rollback([]int{250, 400, 500, 600})
	if err != nil {
		t.Errorf("Error rolling back test file to initial state: %v", err)
	}
}
