package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	// Endpoints
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/calculate/{id}", CalculatePackages)
	myRouter.HandleFunc("/rollback", RollbackPackageChanges)
	myRouter.HandleFunc("/add/{id}", AddNewPackages)
	myRouter.HandleFunc("/remove/{id}", RemovePackages)
	myRouter.HandleFunc("/read", ReadPackages)
	myRouter.HandleFunc("/form/calculate", CalculateData).Methods("POST")

	myRouter.HandleFunc("/", Index)
	myRouter.HandleFunc("/visual/calculate/", CalculateTemplate)
	myRouter.HandleFunc("/submit", SubmitHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
