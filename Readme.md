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
/rollback -> GET method to rollback the changes made to the packages to default values.

Base Url to access the API: https://quiet-anchorage-41647-601b859a018b.herokuapp.com/

FYI - The application is deployed on Heroku. 
The application is deployed on a free dyno, so it might take some time to load the application.
The application deployed on Heroku is based on another repository where server is using gin-gonic framework.
This was necessary to deploy the application on Heroku since gin was a requirement for the deployment.

Reference link of repository deployed on Heroku: https://github.com/dejvissheshi/go-getting-started