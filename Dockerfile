# Use the official Golang image
FROM golang:1.19

# Install Redis
RUN apt-get update && apt-get install -y redis

# Set the working directory
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 5000

# Define the command to start the app
CMD ["./main"]
