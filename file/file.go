package file

import (
	"encoding/csv"
	"os"
	"strconv"
)

// ReadNumbersFromCSV Function to read numbers from a CSV file and return an array of integers
func ReadNumbersFromCSV(filename string) ([]int, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Initialize an array to store numbers
	var numbers []int

	for _, record := range records {
		for _, value := range record {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}
	}

	return numbers, nil
}

// AddNumbersToCSV Function to add numbers to a CSV file removing duplicates
func AddNumbersToCSV(filename string, newData []int) error {

	existingData, err := ReadNumbersFromCSV(filename)
	if err != nil {
		return err
	}

	listOfData := make([]int, 0)
	listOfData = append(listOfData, existingData...)
	listOfData = append(listOfData, newData...)

	mapOfNumbersToAdd := make(map[int]struct{})
	for _, num := range listOfData {
		if _, ok := mapOfNumbersToAdd[num]; !ok {
			mapOfNumbersToAdd[num] = struct{}{}
		}
	}

	dataToWrite := make([]string, 0)
	for num := range mapOfNumbersToAdd {
		dataToWrite = append(dataToWrite, strconv.Itoa(num))
	}

	// Open the CSV file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	// Erase data from the file by truncating it to size 0
	if err := file.Truncate(0); err != nil {
		return err
	}

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Write newData to CSV file
	err = writer.Write(dataToWrite)
	if err != nil {
		return err
	}

	// Flush pending operations to the file
	writer.Flush()

	return nil
}

// RollbackFileToInitialState Function to rollback a test file to its initial state
func RollbackFileToInitialState(filename string, initialStateValues []int) error {
	// Open the CSV file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	// Erase data from the file by truncating it to size 0
	if err := file.Truncate(0); err != nil {
		return err
	}

	// Create a CSV writer
	writer := csv.NewWriter(file)

	dataToWrite := make([]string, 0)
	for _, num := range initialStateValues {
		dataToWrite = append(dataToWrite, strconv.Itoa(num))
	}

	// Write newData to CSV file
	err = writer.Write(dataToWrite)
	if err != nil {
		return err
	}

	// Flush pending operations to the file
	writer.Flush()

	return nil
}

// DeleteNumbersFromCSV Function to delete numbers from a CSV file
func DeleteNumbersFromCSV(filename string, numbersToDelete []int) error {
	existingData, err := ReadNumbersFromCSV(filename)
	if err != nil {
		return err
	}

	mapOfNumbersToDelete := make(map[int]struct{})
	for _, num := range numbersToDelete {
		if _, ok := mapOfNumbersToDelete[num]; !ok {
			mapOfNumbersToDelete[num] = struct{}{}
		}
	}

	dataToWrite := make([]string, 0)
	for _, num := range existingData {
		if _, ok := mapOfNumbersToDelete[num]; !ok {
			dataToWrite = append(dataToWrite, strconv.Itoa(num))
		}
	}

	// Open the CSV file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	// Erase data from the file by truncating it to size 0
	if err := file.Truncate(0); err != nil {
		return err
	}

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Write newData to CSV file
	err = writer.Write(dataToWrite)
	if err != nil {
		return err
	}

	// Flush pending operations to the file
	writer.Flush()

	return nil
}
