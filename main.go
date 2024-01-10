package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"packSizeOptimizer/db/mysql"
)

// Calculator defines the interface for a calculator.
type Calculator interface {
	Calculate(a []int, b int) []PackageInfo
}

// SingletonCalculator is a singleton that holds an instance of a calculator.
type SingletonCalculator struct {
	calculator Calculator
	once       sync.Once
}

// GetInstance returns the singleton instance of the calculator.
func (s *SingletonCalculator) GetInstance() Calculator {
	s.once.Do(func() {
		// You can choose between InitialCalculator, OptimizedCalculator and AdvancedCalculator here
		s.calculator = &AdvancedCalculator{}
	})
	return s.calculator
}

func main() {
	//
	//	// TODO: for better configuration, we could use a config structure
	//	// that reads from the environment variables the DB connection configuration.
	//	// For the sake of simplicity, we will hardcode configuration of db in the main code.
	//	//Initialise DB
	dbPersister := &mysql.MySQLPersister{}

	packageRepository := dbPersister
	packageInterface := &HttpHandler{
		Repository: packageRepository,
		UseFile:    false,
	}

	if packageInterface.UseFile {
		err := dbPersister.Init()
		if err != nil {
			log.Fatal(err)
		}
		defer dbPersister.Close()

		err = dbPersister.Migrate()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Endpoints
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/calculate/", CalculatePackages)
	myRouter.HandleFunc("/rollback", packageInterface.RollbackPackageChanges)
	myRouter.HandleFunc("/add/", packageInterface.AddNewPackages)
	myRouter.HandleFunc("/remove/", packageInterface.RemovePackages)
	myRouter.HandleFunc("/read", packageInterface.ReadPackages)
	myRouter.HandleFunc("/form/calculate", CalculateData).Methods("POST")

	myRouter.HandleFunc("/visual/calculate/", CalculateTemplate)
	myRouter.HandleFunc("/submit", SubmitHandler)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), myRouter)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
