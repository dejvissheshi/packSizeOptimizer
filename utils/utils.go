package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertStringToArray(input string) ([]int, error) {
	// Remove spaces from the input string
	input = strings.ReplaceAll(input, " ", "")

	// Split the input string by commas
	strNumbers := strings.Split(input, ",")

	// Initialize an array to store the converted numbers
	var numbers []int

	// Iterate through the split strings and convert each to an integer
	for _, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to number: %v", err)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
