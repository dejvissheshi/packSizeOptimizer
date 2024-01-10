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