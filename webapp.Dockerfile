FROM golang:latest

WORKDIR /app

COPY go.mod main.go ./

WORKDIR /app/web_resources

COPY webapp /app/web_resources/
WORKDIR /app
CMD go run main.go

