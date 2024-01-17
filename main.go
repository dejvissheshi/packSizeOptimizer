package main

import (
	"fmt"
	"net/http"
	"os"
	"packSizeOptimizer/file"

	"packSizeOptimizer/routers"
)

func main() {

	packService := file.PackFiles{
		Filename: "data.csv",
	}

	httpHandler := routers.HttpHandler{
		PackService: packService,
	}
	myRouter := routers.NewRouter(httpHandler)

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
