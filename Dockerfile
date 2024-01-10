# Use the official Go image as the base image
FROM  --platform=linux/amd64 golang:latest

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y python3 python3-pip

RUN apt install python3.11-venv -y

RUN python3 -m venv .venv

RUN . .venv/bin/activate && pip3 install --upgrade pip

RUN pip3 install pulp --break-system-packages

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

#
## Use an official Go runtime as a parent image
#FROM golang:latest
#
## Set the working directory to /go/src/app
#WORKDIR /app
#
## Install Python 3 and create a virtual environment
#RUN apt-get update && \
#    apt-get install -y python3 python3-pip && \
#    python3 -m venv venv
#
## Activate the virtual environment and install pulp
#RUN . venv/bin/activate && \
#    pip install --upgrade pip && \
#    pip install pulp
#
## Add the current directory contents to the container at /go/src/app
#COPY . .
#
## Build the Go application
#RUN go get -d -v ./...
#RUN go install -v ./...
#
## Define environment variables
#ENV PATH="/go/bin:${PATH}"
#
#RUN go build -o main .
#
#EXPOSE 8080
#
#CMD ["./main"]
