package main

import (
	"fmt"
	"sort"
)

// PackageInfo contains the package size and the quantity of defaultItems. It is used to return the result of the calculation.
type PackageInfo struct {
	PackageSize int
	Quantity    int
}

// Calculate calculates the number of packages needed to fulfill the order
func Calculate(itemsOrdered int, items []int) []PackageInfo {
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
