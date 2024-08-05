# Use the official Golang image to create a build artifact.
# This is the first stage of a multi-stage build.
FROM golang:1.22.5 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Verify the contents of the /app directory (optional, for debugging purposes)
RUN ls -la /app

# Change to the directory containing main.go
WORKDIR /app/cmd/api

# Verify the contents of the /app/cmd/api directory (optional, for debugging purposes)
RUN ls -la /app/cmd/api

# Build the Go app with CGO disabled for compatibility with Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Start a new stage from scratch
FROM alpine:latest  

# Install necessary packages (optional, for debugging purposes)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Verify the contents of the /root directory (optional, for debugging purposes)
RUN ls -la /root

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]