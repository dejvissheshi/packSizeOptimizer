Pack Size Optimizer

This project is the work corresponding to the RE coding challenge.

The project holds the logic of pack optimisation, as well as a Rest API to provide 
the calculation of the packages and customization of delivery packages.

How to run the application locally:


Run the docker comands to build and run the application locally.
1. docker build -t packsizeoptimizer .
2. docker run -p 8080:8080 packsizeoptimizer

The application will be running on localhost:8080

Attached to the Dockerfile is also a docker-compose file to run the application locally. 
In order to run the application using docker-compose, run the following command:
1. docker-compose up -> This will build and run the application together with the database. In the docker-compose file,
 the database is exposed on port 3306, so make sure that the port is not used by any other application. Also the application
 image is mounted to the volume of the container, so any changes made to the application will be reflected in the container.
   (For simplicity purpose, the deployed application is using a file storage instead of a database.)

API Documentation:

 - /calculate/:itemsOrdered -> GET method to calculate the packages for the given number of orders.

 - /add/:package -> GET method to add a new package.
 - /remove/:package -> GET method to remove a package.
 - /read -> GET method to read the packages.
 - /rollback -> GET method to rollback the changes made to the packages to default values.

Added functionality to serve a static html page to test the calculate API.

 - /visual/calculator -> GET method to serve the static html page and calculate the packages.
Receives an array of packageSizes and itemsOrdered.

Base Url to access the API: https://packsizeoptimizer-weathered-river-9585.fly.dev/

FYI - The application is deployed on Fly using containerised approach . 
The application is deployed on a free tier, so it might take some time to load the application.