package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Endpoints
	http.HandleFunc("/test", testHandler)

	http.HandleFunc("/calculate/", CalculatePackages)

	http.HandleFunc("/rollback", RollbackPackageChanges)
	http.HandleFunc("/add/", AddNewPackages)
	http.HandleFunc("/remove/", RemovePackages)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
