FROM golang:1.23

WORKDIR /app

COPY . . 

RUN go build -o main cmd/main.go

ENTRYPOINT ["/app/main"]