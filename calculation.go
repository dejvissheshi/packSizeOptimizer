package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"sort"
)

// PackageInfo contains the package size and the quantity of defaultItems. It is used to return the result of the calculation.
type PackageInfo struct {
	PackageSize int
	Quantity    int
}

// InitialCalculator is the initial implementation of the calculator.
type InitialCalculator struct{}

// OptimizedCalculator is the optimized implementation of the calculator.
type OptimizedCalculator struct{}

// AdvancedCalculator is the advanced implementation of the calculator.
type AdvancedCalculator struct{}

// Calculate calculates the number of packages needed to fulfill the order
// Not correct for all cases
func (i InitialCalculator) Calculate(items []int, itemsOrdered int) []PackageInfo {
	sort.Ints(items)
	packages := make([]PackageInfo, 0)
	leftOver := itemsOrdered

	mapOfItems := make(map[int]struct{})
	for _, item := range items {
		mapOfItems[item] = struct{}{}
	}

	for i := len(items) - 1; i >= 0; i-- {
		packageSize := items[i]
		needed := leftOver/packageSize > 0
		isLatest := i == 0

		if needed {
			quantityNeeded := leftOver / packageSize
			leftOver = leftOver % packageSize
			packages = append(packages, PackageInfo{PackageSize: packageSize, Quantity: quantityNeeded})
		}

		// case when the last package together with the additional remainig package is more than the next biggest package
		if i == 1 && leftOver > 0 {
			if leftOver > items[0] {
				lastPackageQuantity := leftOver / items[0]
				lastPackageItems := items[0] * (lastPackageQuantity + 1)
				if items[1] < lastPackageItems {
					packages = append(packages, PackageInfo{PackageSize: items[1], Quantity: 1})
					break
				}
			}
		}

		if isLatest && leftOver > 0 {
			packages = append(packages, PackageInfo{PackageSize: packageSize, Quantity: 1})
		}

	}
	//return packages
	return optimizePackages(packages, mapOfItems)
}

func optimizePackages(packages []PackageInfo, mapOfItems map[int]struct{}) []PackageInfo {
	optimizedPackages := make([]PackageInfo, 0)
	fmt.Println("packages", packages)
	mapOfPackages := make(map[int]int)
	for i := 0; i < len(packages); i++ {
		fmt.Println(" optimizePackages info", packages[i])
		if _, ok := mapOfPackages[packages[i].PackageSize]; !ok {
			mapOfPackages[packages[i].PackageSize] = packages[i].Quantity
		} else {
			quantity := mapOfPackages[packages[i].PackageSize]
			mapOfPackages[packages[i].PackageSize] = quantity + packages[i].Quantity
		}
	}

	fmt.Println("mapOfPackages", mapOfPackages)
	for i := 0; i < len(mapOfItems); i++ {
		mapOfPackages = removeUnnecessaryPackages(mapOfPackages, mapOfItems)
	}
	fmt.Println("mapOfPackagesUpdated", mapOfPackages)

	for k, v := range mapOfPackages {
		if _, ok := mapOfItems[k]; ok {
			optimizedPackages = append(optimizedPackages, PackageInfo{PackageSize: k, Quantity: v})
		}
	}

	return optimizedPackages
}

func removeUnnecessaryPackages(mapOfPackages map[int]int, mapOfItems map[int]struct{}) map[int]int {
	optimizedPackages := make(map[int]int)
	for k, v := range mapOfPackages {
		if _, ok := mapOfItems[k*v]; ok {
			if q, ok := optimizedPackages[k*v]; !ok {
				optimizedPackages[k*v] = 1
			} else {
				optimizedPackages[k*v] = 1 + q
			}
			continue
		}

		if q, ok := optimizedPackages[k]; !ok {
			optimizedPackages[k] = v
		} else {
			optimizedPackages[k] = v + q
		}
	}
	return optimizedPackages
}

// Calculate uses a brute force approach to calculate the number of packages needed to fulfill the order based on the
// knapsack problem, see article "[Knapsac Problem]". It is not used in the final solution, but it is kept here for reference. The solution is not optimal
// and it is very slow for large numbers of items. ex: 23, 31, 53 with itemsOrdered = 5000000. It takes a huge amount of time
// to calculate the result.
// [Knapsac Problem]: https://en.wikipedia.org/wiki/Knapsack_problem
func (o OptimizedCalculator) Calculate(xi []int, W int) []PackageInfo {
	minWeights := make([]int, len(xi))
	for i := range minWeights {
		minWeights[i] = math.MaxInt32
	}
	minSum := math.MaxInt32
	minPackets := math.MaxInt32
	var minPackItems []PackageInfo

	var solve func(index int, weights []int, currentSum int, currentPackets int, currentPackItems []PackageInfo)
	solve = func(index int, weights []int, currentSum int, currentPackets int, currentPackItems []PackageInfo) {
		// Base case: if we have processed all packets
		if index == len(xi) {
			if currentSum >= W {
				if currentSum < minSum || (currentSum == minSum && currentPackets < minPackets) {
					minSum = currentSum
					minPackets = currentPackets
					copy(minWeights, weights)
					minPackItems = append([]PackageInfo(nil), currentPackItems...)
				}
			}
			return
		}

		// Recursive case: try adding the current packet 0 or more times
		for count := 0; currentSum+xi[index]*count <= minSum; count++ {
			newWeights := append([]int(nil), weights...)
			newWeights[index] += count

			newSum := currentSum + xi[index]*count
			newPackets := currentPackets + count

			newPackItem := PackageInfo{PackageSize: xi[index], Quantity: count}
			newPackItems := append([]PackageInfo(nil), currentPackItems...)
			newPackItems = append(newPackItems, newPackItem)

			solve(index+1, newWeights, newSum, newPackets, newPackItems)
		}
	}

	solve(0, make([]int, len(xi)), 0, 0, nil)

	if minSum == math.MaxInt32 {
		return nil // No solution found
	}

	var packageInfoWithoutEmpty []PackageInfo
	for _, val := range minPackItems {
		if val.Quantity != 0 {
			packageInfoWithoutEmpty = append(packageInfoWithoutEmpty, val)
		}
	}

	return packageInfoWithoutEmpty
}

// Calculate for the advanced calculator calculates the number of packages needed to fulfill the order by using a Python script that uses a
// linear programming solver. The algorithm is based on the knapsack problem. This solutions uses the PuLP library for Python.
// Gonum library for Go could be used as well, but it is not as easy to use as PuLP. ChatGPT for reference on the Python script
func (a AdvancedCalculator) Calculate(packages []int, items int) []PackageInfo {
	sort.Ints(packages)
	var args []string
	for _, i := range packages {
		fmt.Println(i)
		args = append(args, fmt.Sprint(i))
	}
	args = append(args, fmt.Sprint(items))

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Advanced Calculator, Error getting current working directory:", err)
		return nil
	}

	filepath := dir + "/script.py"
	cmd := exec.Command("python3", filepath)
	cmd.Args = append(cmd.Args, args...)

	output, _ := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Advanced Calculator, Error executing Python script:", err)
		return nil
	}

	var result []int
	// Unmarshal the JSON data into the result slice
	err = json.Unmarshal(output, &result)
	if err != nil {
		fmt.Println("Advanced Calculator, Error unmarshalling JSON data:", err)
		return nil
	}

	var packageInfo []PackageInfo
	for i, val := range packages {
		p := PackageInfo{PackageSize: val, Quantity: result[i]}
		packageInfo = append(packageInfo, p)
	}

	var packageInfoWithoutEmpty []PackageInfo
	for _, val := range packageInfo {
		if val.Quantity != 0 {
			packageInfoWithoutEmpty = append(packageInfoWithoutEmpty, val)
		}
	}

	// Write the output of the Python script to the HTTP response
	return packageInfoWithoutEmpty
}
