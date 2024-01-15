package main

import (
	"math"
)

// PackageInfo contains the package size and the quantity of defaultItems. It is used to return the result of the calculation.
type PackageInfo struct {
	PackageSize int
	Quantity    int
}

// Calculate uses a brute force approach to calculate the number of packages needed to fulfill the order based on the
// knapsack problem, see article "[Knapsac Problem]". It is not used in the final solution, but it is kept here for reference. The solution is not optimal
// and it is very slow for large numbers of items. ex: 23, 31, 53 with itemsOrdered = 5000000. It takes a huge amount of time
// to calculate the result.
// [Knapsac Problem]: https://en.wikipedia.org/wiki/Knapsack_problem
func Calculate(xi []int, W int) []PackageInfo {
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
