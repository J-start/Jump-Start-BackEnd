FROM golang:1.23

# set working directory
WORKDIR /app

# Copy the source code
COPY . . 

# Build the Go app
RUN go build -o main cmd/main.go

# Run the executable
ENTRYPOINT ["/app/main"]