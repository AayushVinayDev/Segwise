# Use the official Golang image to build the application
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code from the `backend-api` directory
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend-api ./cmd/main.go

# Use a minimal base image for the final build
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /backend-api .

# Expose the port the API runs on
EXPOSE 8080

# Command to run the executable
CMD ["./backend-api"]
