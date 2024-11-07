# Start from the latest golang image
FROM golang:1.19

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main ./cmd

# Expose port
EXPOSE 8080

# Start the application
CMD ["./main"]
