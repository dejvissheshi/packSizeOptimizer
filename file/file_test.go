package file

import (
	"reflect"
	"sort"
	"testing"
)

func Test_ReadNumbersFromCSV(t *testing.T) {
	// Test case 1: Valid CSV file with numbers
	numbers, err := ReadNumbersFromCSV("fixtures/file_1.csv")
	if err != nil {
		t.Errorf("Error reading numbers from CSV file: %v", err)
	}
	expectedNumbers := []int{250, 400, 500, 600}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected numbers %v, but got %v", expectedNumbers, numbers)
	}

	// Test case 2: Invalid CSV file with non-numeric values
	_, err = ReadNumbersFromCSV("fixtures/file_2_invalid_data.csv")
	if err == nil {
		t.Error("Expected error reading numbers from invalid CSV file, but got none")
	}
}

func Test_AddNumbersToCSV(t *testing.T) {
	// Test case 1: Valid CSV file with numbers
	err := AddNumbersToCSV("fixtures/file_3.csv", []int{700, 800})
	if err != nil {
		t.Errorf("Error adding numbers to CSV file: %v", err)
	}

	numbers, err := ReadNumbersFromCSV("fixtures/file_3.csv")
	sort.Ints(numbers)
	if err != nil {
		t.Errorf("Error reading numbers from CSV file: %v", err)
	}

	expectedNumbers := []int{250, 400, 500, 600, 700, 800}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected numbers %v, but got %v", expectedNumbers, numbers)
	}

	err = RollbackFileToInitialState("fixtures/file_3.csv", []int{250, 400, 500, 600})
	if err != nil {
		t.Errorf("Error rolling back test file to initial state: %v", err)
	}
}

func Test_DeleteNumbersFromCSV(t *testing.T) {
	// Test case 1: Valid CSV file with numbers
	err := DeleteNumbersFromCSV("fixtures/file_4.csv", []int{250, 400})
	if err != nil {
		t.Errorf("Error deleting numbers from CSV file: %v", err)
	}

	numbers, err := ReadNumbersFromCSV("fixtures/file_4.csv")
	sort.Ints(numbers)
	if err != nil {
		t.Errorf("Error reading numbers from CSV file: %v", err)
	}

	expectedNumbers := []int{500, 600}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected numbers %v, but got %v", expectedNumbers, numbers)
	}

	err = RollbackFileToInitialState("fixtures/file_4.csv", []int{250, 400, 500, 600})
	if err != nil {
		t.Errorf("Error rolling back test file to initial state: %v", err)
	}
}
