# Start from the official Golang image.
FROM golang:1.22.1 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy migration files

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application, disable CGO to create a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o demo-app ./cmd/abr-plus

# Use a smaller image to run the app
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN ls -la

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/demo-app .
COPY --from=builder /app/pkg/abr-plus/migrations ./migrations

# Command to run the executable
CMD ["./demo-app"]
