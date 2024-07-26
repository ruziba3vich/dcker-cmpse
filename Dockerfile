# Use the official Golang image
FROM golang:1.22.5-alpine

# Set the working directory
WORKDIR /app

# Copy all source files into the working directory
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Command to run the executable
CMD ["./main"]
