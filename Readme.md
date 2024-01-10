Pack Size Optimizer

This project is the work corresponding to the RE coding challenge.

The project holds the logic of pack optimisation, as well as a Rest API to provide 
the calculation of the packages and customization of delivery packages.

How to run the application locally:

1. Install the repository locally through git install command.
2. It is mandatory to have go locally.
3. To run the application go to the root directory and run:
   1. go build packSizeOptimizer
   2. go run packSizeOptimizer

API Documentation:

/calculate/:itemsOrdered -> GET method to calculate the packages for the given number of orders.

/add/:package -> GET method to add a new package.
/remove/:package -> GET method to remove a package.
/read -> GET method to read the packages.
/rollback -> GET method to rollback the changes made to the packages to default values.

Added functionality to serve a static html page to test the calculate API.

/visual/calculator -> GET method to serve the static html page and calculate the packages.
Receives an array of packageSizes and itemsOrdered.

Base Url to access the API: https://packsizeoptimizer-weathered-river-9585.fly.dev/

FYI - The application is deployed on Fly using containerised approach . 
The application is deployed on a free tier, so it might take some time to load the application.