# Use the official Golang image to build the application
FROM golang:1.22 as builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code, including static files and main.go in the cmd folder
COPY . .

# Build the application binary from main.go in the cmd directory
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend-api ./cmd/main.go

# Use a minimal base image for the final build
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary and static files
COPY --from=builder /backend-api .
COPY --from=builder /app/static ./static

# Expose the port the API runs on
EXPOSE 8080

# Command to run the executable
CMD ["./backend-api"]
