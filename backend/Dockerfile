# Build the app
FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o webapp .

EXPOSE 8080

CMD ["./webapp"]

